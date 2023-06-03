package actor

import (
	"fmt"
	"nashrul-be/crm/dto"
	"net/http"
)

type ControllerInterface interface {
	GetByID(id uint) (dto.BaseResponse, error)
	GetAllByUsername(req PaginationRequest) (dto.BaseResponse, error)
	CreateActor(req CreateRequest) (dto.BaseResponse, error)
	ChangeActiveActor(request ChangeActiveRequest) (dto.BaseResponse, error)
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

func (c controller) GetAllByUsername(req PaginationRequest) (dto.BaseResponse, error) {
	offset := (req.Page - 1) * req.PerPage
	actors, err := c.actorUseCase.GetAllByUsername(req.Username+"%", req.PerPage, offset)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	actorResponse := make([]Representation, 0)
	for _, actor := range actors {
		actorResponse = append(actorResponse, mapActorToResponse(actor))
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success retrieve actor",
		Data:    actorResponse,
	}, err
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

func (c controller) ChangeActiveActor(request ChangeActiveRequest) (dto.BaseResponse, error) {
	result := map[string][]string{
		"success": {},
		"failed":  {},
	}
	activate, err := c.actorUseCase.ActivateActor(request.Activate)
	if err != nil {
		result["failed"] = append(result["failed"], request.Activate...)
	} else {
		result["success"] = append(result["success"], activate["success"]...)
		result["failed"] = append(result["failed"], activate["failed"]...)
	}
	deactivate, err := c.actorUseCase.DeactivateActor(request.Deactivate)
	if err != nil {
		result["failed"] = append(result["failed"], request.Deactivate...)
	} else {
		result["success"] = append(result["success"], deactivate["success"]...)
		result["failed"] = append(result["failed"], deactivate["failed"]...)
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Activate/Deactivate success",
		Data:    result,
	}, nil
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
