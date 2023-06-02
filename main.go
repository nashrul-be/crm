package main

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/modules/customer"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/db"
)

func main() {
	engine := gin.Default()

	dbConn, err := db.Connect(db.Config{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "mini",
	})

	if err != nil {
		panic(err.Error())
	}

	userRepo := repositories.NewCustomerRepository(dbConn)
	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)

	userRoute := customer.NewCustomerRoute(userRepo)
	userRoute.Handle(engine)

	actorRoute := actor.NewActorRoute(actorRepo, roleRepo)
	actorRoute.Handle(engine)

	if err := engine.Run(); err != nil {
		panic(err.Error())
	}
}
