package register_approval

import (
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
)

type ControllerInterface interface {
	GetAllPendingApproval() (dto.BaseResponse, error)
	Approve(request ApproveRequest, actor entities.Actor) (dto.BaseResponse, error)
}

func NewController(approvalUseCase RegisterApprovalUseCaseInterface) ControllerInterface {
	return controller{approvalUseCase: approvalUseCase}
}

type controller struct {
	approvalUseCase RegisterApprovalUseCaseInterface
}

func (c controller) GetAllPendingApproval() (dto.BaseResponse, error) {
	approvals, err := c.approvalUseCase.GetAllPendingApproval()
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	result := Representation{Username: []string{}}
	for _, approval := range approvals {
		result.Username = append(result.Username, approval.Admin.Username)
	}
	return dto.Success("Success retrieve approval", result), nil
}

func (c controller) Approve(request ApproveRequest, actor entities.Actor) (dto.BaseResponse, error) {
	approved, err := c.approvalUseCase.Approve(request.Approve, actor.ID)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	rejected, err := c.approvalUseCase.Rejected(request.Reject, actor.ID)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	data := ApproveResponse{
		Success: append(approved["success"], rejected["success"]...),
		Fail:    append(approved["failed"], rejected["failed"]...),
	}
	return dto.Success("Success update approval", data), nil
}
