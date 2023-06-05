package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
	"strconv"
)

type RequestHandlerInterface interface {
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func NewRequestHandler(controllerInterface ControllerInterface) RequestHandlerInterface {
	return requestHandler{customerController: controllerInterface}
}

type requestHandler struct {
	customerController ControllerInterface
}

func (h requestHandler) GetByID(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorNotFound(fmt.Sprintf("Customer %d not found", id)))
		return
	}
	response, err := h.customerController.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorNotFound(fmt.Sprintf("Customer %d not found", id)))
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) GetAll(c *gin.Context) {
	var request PaginationRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.customerController.GetAll(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) CreateCustomer(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.customerController.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateCustomer(c *gin.Context) {
	var request UpdateRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.customerController.UpdateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) DeleteCustomer(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	err = h.customerController.DeleteCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(http.StatusNoContent, nil)
}
