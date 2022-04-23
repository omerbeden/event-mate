package main

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/repositories"
)

func main() {
	// test code
	//TODO refeactor later
	database.InitMigration(&core.UserProfile{}, &core.UserAdress{})

	repo := repositories.NewUserRepo()

	user := &core.UserProfile{
		Name:     "omer",
		LastName: "beden",
		About:    "about about",
		Adress: core.UserAdress{
			City: "Sakarya",
		},
	}
	fmt.Printf("%v", user.Model)

	repo.InsertUser(user)

	users, _ := repo.GetUsers()
	fmt.Printf("%+v", users)

}
