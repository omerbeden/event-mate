package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/routes"
	activityRepoAdapter "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter"
	activityServiceEntryPoints "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	profileRepoAdapter "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
	ve "github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db/postgres"
)

const applicationPort = ":3000"

func main() {
	dbPool := postgres.NewConn(&postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10}})

	var redisOption = cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	}

	redisClient := cache.NewRedisClient(redisOption)
	pgxAdapter := postgres.NewPgxAdapter(dbPool)
	sugaredLogger := pkg.Logger()

	activityRepository := activityRepoAdapter.NewActivityRepo(pgxAdapter)
	activityRulesRepository := activityRepoAdapter.NewActivityRulesRepo(pgxAdapter)
	activityFlowRepository := activityRepoAdapter.NewActivityFlowRepo(pgxAdapter)
	locationRepository := activityRepoAdapter.NewLocationRepo(pgxAdapter)

	activityService := activityServiceEntryPoints.NewService(activityRepository, activityRulesRepository, activityFlowRepository, locationRepository, *redisClient, pgxAdapter)

	userRepository := profileRepoAdapter.NewUserProfileRepo(pgxAdapter)
	userAddressRepo := profileRepoAdapter.NewUserProfileAddressRepo(pgxAdapter)
	userStatRepo := profileRepoAdapter.NewUserProfileStatRepo(pgxAdapter)
	userBadgeRepo := profileRepoAdapter.NewBadgeRepo(pgxAdapter)
	userService := entrypoints.NewService(userRepository, userStatRepo, userAddressRepo, userBadgeRepo, *redisClient, pgxAdapter)
	validationService := ve.NewValidationService(sugaredLogger)

	app := fiber.New()
	app.Use(requestid.New())
	api := app.Group("/api")
	routes.ActivityRouter(api, *activityService)
	routes.ProfileRouter(api, *userService)
	routes.ValidationRouter(api, validationService)

	go func() {
		if err := app.Listen(applicationPort); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	redisClient.Close()
	dbPool.Close()
	sugaredLogger.Sync()

	fmt.Println("Fiber was successful shutdown.")

}
