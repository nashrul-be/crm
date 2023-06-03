package actor

import (
	"fmt"
	"nashrul-be/crm/dto"
	"net/http"
)

type ControllerInterface interface {
	GetByID(id uint) (dto.BaseResponse, error)
	CreateActor(req CreateRequest) (dto.BaseResponse, error)
	UpdateActor(req UpdateRequest) (dto.BaseResponse, error)
	DeleteActor(id uint) error
}

func NewController(useCaseInterface UseCaseInterface) ControllerInterface {
	return controller{
		actorUseCase: useCaseInterface,
	}
}

type controller struct {
	actorUseCase UseCaseInterface
}

func (c controller) GetByID(id uint) (dto.BaseResponse, error) {
	actor, err := c.actorUseCase.GetByID(id)
	if err != nil {
		return dto.ErrorNotFound(fmt.Sprintf("Actor %d doesn't exist", id)), err
	}
	actorRepresentation := mapActorToResponse(actor)
	return dto.Success("Success retrieve data", actorRepresentation), nil
}

func (c controller) CreateActor(req CreateRequest) (dto.BaseResponse, error) {
	actor := mapCreateRequestToActor(req)
	exist, err := c.actorUseCase.IsUsernameExist(actor)
	if err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create actor",
		}, err
	}
	if exist {
		return dto.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Username already exist",
		}, nil
	}
	if err := c.actorUseCase.CreateActor(&actor); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create actor",
		}, err
	}
	response := mapActorToResponse(actor)
	return dto.Created("Success create actor", response), nil
}

func (c controller) UpdateActor(req UpdateRequest) (dto.BaseResponse, error) {
	actor := mapUpdateRequestToActor(req)
	exist, err := c.actorUseCase.IsExist(actor.ID)
	if err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create actor",
		}, err
	}
	if !exist {
		return dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Actor not found",
		}, nil
	}
	if actor.Username != "" {
		exist, err := c.actorUseCase.IsUsernameExist(actor)
		if err != nil {
			return dto.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Can't create actor",
			}, err
		}
		if exist {
			return dto.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Username already exist",
			}, nil
		}
	}
	if err := c.actorUseCase.UpdateActor(&actor); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't update actor",
		}, err
	}
	response := mapActorToResponse(actor)
	return dto.Success("Success update actor", response), nil
}

func (c controller) DeleteActor(id uint) error {
	if err := c.actorUseCase.DeleteActor(id); err != nil {
		return err
	}
	return nil
}
