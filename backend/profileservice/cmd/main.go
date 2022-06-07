package main

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
	mgrpc "github.com/omerbeden/event-mate/backend/profileservice/infrastructure/grpc"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/repositories"
)

func main() {

	// test code
	//TODO refeactor later
	database.InitMigration(&core.UserProfile{}, &core.UserProfileAdress{})

	repo := repositories.NewUserRepo()

	user := &core.UserProfile{
		Name:     "omer",
		LastName: "beden",
		About:    "about about",
		Adress: core.UserProfileAdress{
			City: "Sakarya",
		},
	}
	fmt.Printf("%v", user.Model)

	repo.InsertUser(user)

	users, _ := repo.GetUsers()
	fmt.Printf("%+v", users)

	//GRPC SERVER
	mgrpc.StartGRPCServer()

}
