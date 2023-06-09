package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"nashrul-be/crm/migrations"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/customer"
	register_approval "nashrul-be/crm/modules/register-approval"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/translate"
	"reflect"
	"strings"
	"time"
)

func registerTranslator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		if err := en.RegisterDefaultTranslations(v, translate.DefaultTranslator()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	if err := registerTranslator(); err != nil {
		panic(err.Error())
	}

	engine := gin.Default()

	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = db.DefaultConnection()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err.Error())
	}

	migrations.Migrate(dbConn)

	userRepo := repositories.NewCustomerRepository(dbConn)
	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)
	approvalRepo := repositories.NewRegisterApprovalRepository(dbConn)

	migrations.Seed(dbConn, actorRepo, roleRepo, approvalRepo)

	userRoute := customer.NewRoute(userRepo)
	userRoute.Handle(engine)

	actorRoute := actor.NewRoute(actorRepo, roleRepo, approvalRepo)
	actorRoute.Handle(engine)

	approveRoute := register_approval.NewRoute(actorRepo, approvalRepo)
	approveRoute.Handle(engine)

	actorUseCase := actor.NewUseCase(actorRepo, roleRepo, approvalRepo)
	authRoute := authentication.NewRoute(actorUseCase)
	authRoute.Handle(engine)

	if err := engine.Run(); err != nil {
		panic(err.Error())
	}
}
