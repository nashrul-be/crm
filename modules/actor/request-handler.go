package actor

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
	"strconv"
)

type RequestHandlerInterface interface {
	GetByID(c *gin.Context)
	GetAllByUsername(c *gin.Context)
	CreateActor(c *gin.Context)
	ChangeActiveActor(c *gin.Context)
	UpdateActor(c *gin.Context)
	DeleteActor(c *gin.Context)
}

func NewRequestHandler(controllerInterface ControllerInterface) RequestHandlerInterface {
	return requestHandler{actorController: controllerInterface}
}

type requestHandler struct {
	actorController ControllerInterface
}

func (h requestHandler) GetByID(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusNotFound, actorNotFound())
		return
	}
	response, err := h.actorController.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, actorNotFound())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) GetAllByUsername(c *gin.Context) {
	var request PaginationRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		request = PaginationRequest{Page: 1, PerPage: 10}
	}
	response, err := h.actorController.GetAllByUsername(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) CreateActor(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.actorController.CreateActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) ChangeActiveActor(c *gin.Context) {
	var request ChangeActiveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.actorController.ChangeActiveActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateActor(c *gin.Context) {
	var request UpdateRequest
	actor, exist := c.Get("actor")
	if !exist {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	request.ID = actor.(entities.Actor).ID
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.actorController.UpdateActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) DeleteActor(c *gin.Context) {
	uriParam := c.Param("id")
	id, err := strconv.Atoi(uriParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, actorNotFound())
		return
	}
	if err := h.actorController.DeleteActor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(http.StatusNoContent, nil)
}
