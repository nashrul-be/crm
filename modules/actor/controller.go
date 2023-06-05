package actor

import (
	"nashrul-be/crm/dto"
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
		return actorNotFound(), err
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
	return dto.Success("Success retrieve actor", actorResponse), err
}

func (c controller) CreateActor(req CreateRequest) (dto.BaseResponse, error) {
	actor := mapCreateRequestToActor(req)
	validationErr, err := c.actorUseCase.validateActor(actor, validateUsername)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	if err := c.actorUseCase.CreateActor(&actor); err != nil {
		return dto.ErrorInternalServerError(), err
	}
	response := mapActorToResponse(actor)
	return dto.Created("Success create actor", response), nil
}

func (c controller) ChangeActiveActor(request ChangeActiveRequest) (dto.BaseResponse, error) {
	result := map[string][]string{
		"success": {},
		"failed":  {},
	}
	for _, username := range request.Activate {
		if err := c.actorUseCase.ActivateActor(username); err != nil {
			result["failed"] = append(result["failed"], username)
		} else {
			result["success"] = append(result["success"], username)
		}
	}
	for _, username := range request.Deactivate {
		if err := c.actorUseCase.DeactivateActor(username); err != nil {
			result["failed"] = append(result["failed"], username)
		} else {
			result["success"] = append(result["success"], username)
		}
	}
	return dto.Success("Activate/Deactivate success", result), nil
}

func (c controller) UpdateActor(req UpdateRequest) (dto.BaseResponse, error) {
	actor := mapUpdateRequestToActor(req)
	validationErr, err := c.actorUseCase.validateActor(actor, validateId, validateUsername)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	if err := c.actorUseCase.UpdateActor(&actor); err != nil {
		return dto.ErrorInternalServerError(), err
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
