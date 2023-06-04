package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/customer"
	register_approval "nashrul-be/crm/modules/register-approval"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	engine := gin.Default()

	dbConn, err := db.DefaultConnection()
	if err != nil {
		panic(err.Error())
	}

	userRepo := repositories.NewCustomerRepository(dbConn)
	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)
	approvalRepo := repositories.NewRegisterApprovalRepository(dbConn)

	userRoute := customer.NewCustomerRoute(userRepo)
	userRoute.Handle(engine)

	actorRoute := actor.NewActorRoute(actorRepo, roleRepo, approvalRepo)
	actorRoute.Handle(engine)

	approveRoute := register_approval.NewApprovalRoute(actorRepo, approvalRepo)
	approveRoute.Handle(engine)

	actorUseCase := actor.NewUseCase(actorRepo, roleRepo, approvalRepo)
	authRoute := authentication.NewAuthRoute(actorUseCase)
	authRoute.Handle(engine)

	if err := engine.Run(); err != nil {
		panic(err.Error())
	}
}
