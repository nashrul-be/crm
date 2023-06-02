package actor

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/repositories"
)

type ActorRoute struct {
	actorRequestHandler RequestHandlerInterface
}

func NewActorRoute(actorRepository repositories.ActorRepositoryInterface, roleRepository repositories.RoleRepositoryInterface) ActorRoute {
	useCase := NewUseCase(actorRepository, roleRepository)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return ActorRoute{actorRequestHandler: requestHandler}
}

func (r ActorRoute) Handle(router *gin.Engine) {
	actor := router.Group("/actors")
	actor.GET("/:id", r.actorRequestHandler.GetByID)
	actor.POST("", r.actorRequestHandler.CreateActor)
	actor.PATCH("/:id", r.actorRequestHandler.UpdateActor)
	actor.DELETE("/:id", r.actorRequestHandler.DeleteActor)
}
