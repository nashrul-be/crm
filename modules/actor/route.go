package actor

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

type ActorRoute struct {
	actorRequestHandler RequestHandlerInterface
}

func NewActorRoute(actorRepository repositories.ActorRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
	approvalRepository repositories.RegisterApprovalRepositoryInterface,
) ActorRoute {
	useCase := NewUseCase(actorRepository, roleRepository, approvalRepository)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return ActorRoute{actorRequestHandler: requestHandler}
}

func (r ActorRoute) Handle(router *gin.Engine) {
	actor := router.Group("/actors", middleware.Authenticate())
	actor.GET("/:id", r.actorRequestHandler.GetByID)
	actor.POST("", r.actorRequestHandler.CreateActor)
	actor.PATCH("/active", middleware.AuthorizationWithRole([]string{"super_admin"}), r.actorRequestHandler.ChangeActiveActor)
	actor.PATCH("/:id", r.actorRequestHandler.UpdateActor)
	actor.DELETE("/:id", middleware.AuthorizationWithRole([]string{"super_admin"}), r.actorRequestHandler.DeleteActor)
}
