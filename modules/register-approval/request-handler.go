package register_approval

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
)

type RequestHandlerInterface interface {
	GetAllPendingApproval(c *gin.Context)
	Approve(c *gin.Context)
}

func NewRequestHandler(approvalController ControllerInterface) RequestHandlerInterface {
	return requestHandler{approvalController: approvalController}
}

type requestHandler struct {
	approvalController ControllerInterface
}

func (h requestHandler) GetAllPendingApproval(c *gin.Context) {
	response, err := h.approvalController.GetAllPendingApproval()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) Approve(c *gin.Context) {
	var request ApproveRequest
	actorAny, _ := c.Get("actor")
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.approvalController.Approve(request, actorAny.(entities.Actor))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(response.Code, response)
}
