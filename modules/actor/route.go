package actor

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

type Route struct {
	actorRequestHandler RequestHandlerInterface
}

func NewRoute(actorRepository repositories.ActorRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
	approvalRepository repositories.RegisterApprovalRepositoryInterface,
) Route {
	useCase := NewUseCase(actorRepository, roleRepository, approvalRepository)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return Route{actorRequestHandler: requestHandler}
}

func (r Route) Handle(router *gin.Engine) {
	router.POST("/register", r.actorRequestHandler.CreateActor) //too lazy to move it to authenticate package
	router.PATCH("/me", middleware.Authenticate(), r.actorRequestHandler.UpdateActor)
	actor := router.Group("/actors", middleware.Authenticate())
	actor.GET("/:id", r.actorRequestHandler.GetByID)
	actor.GET("", r.actorRequestHandler.GetAllByUsername)
	actor.POST("", r.actorRequestHandler.CreateActor)
	actor.PATCH("/active", middleware.AuthorizationWithRole([]string{"super_admin"}), r.actorRequestHandler.ChangeActiveActor)
	actor.DELETE("/:id", middleware.AuthorizationWithRole([]string{"super_admin"}), r.actorRequestHandler.DeleteActor)
}
