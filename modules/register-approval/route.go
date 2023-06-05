package register_approval

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

type Route struct {
	approvalRequestHandler RequestHandlerInterface
}

func NewRoute(actorRepository repositories.ActorRepositoryInterface,
	approvalRepository repositories.RegisterApprovalRepositoryInterface) Route {
	approvalUseCase := NewRegisterApprovalUseCase(approvalRepository, actorRepository)
	approvalController := NewController(approvalUseCase)
	approvalRequestHandler := NewRequestHandler(approvalController)
	return Route{approvalRequestHandler: approvalRequestHandler}
}

func (r Route) Handle(router *gin.Engine) {
	group := router.Group("/actors", middleware.Authenticate(), middleware.AuthorizationWithRole([]string{"super_admin"}))
	group.GET("/approve", r.approvalRequestHandler.GetAllPendingApproval)
	group.PUT("/approve", r.approvalRequestHandler.Approve)
}
