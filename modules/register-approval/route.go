package register_approval

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

type ApprovalRoute struct {
	approvalRequestHandler RequestHandlerInterface
}

func NewApprovalRoute(actorRepository repositories.ActorRepositoryInterface,
	approvalRepository repositories.RegisterApprovalRepositoryInterface) ApprovalRoute {
	approvalUseCase := NewRegisterApprovalUseCase(approvalRepository, actorRepository)
	approvalController := NewRegisterController(approvalUseCase)
	approvalRequestHandler := NewRequestHandler(approvalController)
	return ApprovalRoute{approvalRequestHandler: approvalRequestHandler}
}

func (r ApprovalRoute) Handle(router *gin.Engine) {
	group := router.Group("/actors", middleware.Authenticate(), middleware.AuthorizationWithRole([]string{"super_admin"}))
	group.GET("/approve", r.approvalRequestHandler.GetAllPendingApproval)
	group.PUT("/approve", r.approvalRequestHandler.Approve)
}
