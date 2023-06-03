package authentication

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/modules/actor"
)

type AuthRoute struct {
	authRequestHandler RequestHandlerInterface
}

func NewAuthRoute(actorUseCase actor.UseCaseInterface) AuthRoute {
	controller := NewAuthController(actorUseCase)
	requestHandler := NewRequestHandler(controller)
	return AuthRoute{
		authRequestHandler: requestHandler,
	}
}

func (r AuthRoute) Handle(router *gin.Engine) {
	router.POST("/login", r.authRequestHandler.Login)
}
