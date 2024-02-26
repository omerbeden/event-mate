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
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/redisadapter"
	activityRepo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	activityService "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
)

const applicationPort = ":3000"

func main() {
	dbPool := postgres.NewConn(&postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10}})

	var redisOption = redisadapter.RedisOption()
	redisClient := redis.NewClient(redisOption)
	activityRepository := activityRepo.NewActivityRepo(dbPool)
	activityRulesRepository := activityRepo.NewActivityRulesRepo(dbPool)
	activityFlowRepository := activityRepo.NewActivityFlowRepo(dbPool)
	locationRepository := activityRepo.NewLocationRepo(dbPool)

	activityService := activityService.NewService(activityRepository, activityRulesRepository, activityFlowRepository, locationRepository, *redisClient)

	userRepository := repo.NewUserProfileRepo(dbPool)
	userService := entrypoints.NewService(userRepository, *cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	}))

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
