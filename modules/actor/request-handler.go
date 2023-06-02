package actor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
	"strconv"
)

type RequestHandlerInterface interface {
	GetByID(c *gin.Context)
	CreateActor(c *gin.Context)
	UpdateActor(c *gin.Context)
	DeleteActor(c *gin.Context)
}

func NewRequestHandler(controllerInterface ControllerInterface) RequestHandlerInterface {
	return actorRequestHandler{actorController: controllerInterface}
}

type actorRequestHandler struct {
	actorController ControllerInterface
}

func (h actorRequestHandler) GetByID(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorNotFound(fmt.Sprintf("Actor %d not found", id)))
		return
	}
	response, err := h.actorController.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorNotFound(fmt.Sprintf("Actor %d not found", id)))
		return
	}
	c.JSON(response.Code, response)
}

func (h actorRequestHandler) CreateActor(c *gin.Context) {
	var request CreateRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.actorController.CreateActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h actorRequestHandler) UpdateActor(c *gin.Context) {
	var request UpdateRequest
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.actorController.UpdateActor(uint(id), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h actorRequestHandler) DeleteActor(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	err = h.actorController.DeleteActor(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(http.StatusNoContent, nil)
}
