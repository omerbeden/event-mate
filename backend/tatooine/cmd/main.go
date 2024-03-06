package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/routes"
	activityRepoAdapter "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter"
	activityServiceEntryPoints "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	profileRepoAdapter "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
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
	activityRepository := activityRepoAdapter.NewActivityRepo(activityRepoAdapter.NewPgxAdapter(dbPool))
	activityRulesRepository := activityRepoAdapter.NewActivityRulesRepo(activityRepoAdapter.NewPgxAdapter(dbPool))
	activityFlowRepository := activityRepoAdapter.NewActivityFlowRepo(activityRepoAdapter.NewPgxAdapter(dbPool))
	locationRepository := activityRepoAdapter.NewLocationRepo(activityRepoAdapter.NewPgxAdapter(dbPool))

	activityService := activityServiceEntryPoints.NewService(activityRepository, activityRulesRepository, activityFlowRepository, locationRepository, *redisClient)

	userRepository := profileRepoAdapter.NewUserProfileRepo(profileRepoAdapter.NewPgxAdapter(dbPool))
	userAddressRepo := profileRepoAdapter.NewUserProfileAddressRepo(profileRepoAdapter.NewPgxAdapter(dbPool))
	userStatRepo := profileRepoAdapter.NewUserProfileStatRepo(profileRepoAdapter.NewPgxAdapter(dbPool))
	userService := entrypoints.NewService(userRepository, userStatRepo, userAddressRepo, *redisClient)

	app := fiber.New()
	api := app.Group("/api")
	routes.ActivityRouter(api, *activityService)
	routes.ProfileRouter(api, *userService)

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

	fmt.Println("Fiber was successful shutdown.")

}
