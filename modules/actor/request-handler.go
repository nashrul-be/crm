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
	GetAllByUsername(c *gin.Context)
	CreateActor(c *gin.Context)
	ChangeActiveActor(c *gin.Context)
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

func (h actorRequestHandler) GetAllByUsername(c *gin.Context) {
	perPage, _ := strconv.Atoi(c.Query("perpage"))
	page, _ := strconv.Atoi(c.Query("page"))
	username := c.Query("username")
	if perPage < 1 || page < 1 {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("perPage or page must be more than zero"))
		return
	}
	request := PaginationRequest{
		PerPage:  uint(perPage),
		Page:     uint(page),
		Username: username,
	}
	response, err := h.actorController.GetAllByUsername(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
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

func (h actorRequestHandler) ChangeActiveActor(c *gin.Context) {
	var request ChangeActiveRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	response, err := h.actorController.ChangeActiveActor(request)
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
	request.ID = uint(id)
	response, err := h.actorController.UpdateActor(request)
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
