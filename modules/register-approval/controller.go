package register_approval

import (
	"nashrul-be/crm/dto"
	"net/http"
)

type ControllerInterface interface {
	GetAllPendingApproval() (dto.BaseResponse, error)
	Approve(request ApproveRequest) (dto.BaseResponse, error)
}

func NewRegisterController(approvalUseCase RegisterApprovalUseCaseInterface) ControllerInterface {
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
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success retrieve approval",
		Data:    result,
	}, nil
}

func (c controller) Approve(request ApproveRequest) (dto.BaseResponse, error) {
	approved, err := c.approvalUseCase.Approve(request.Approve, 1)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	rejected, err := c.approvalUseCase.Rejected(request.Reject, 1)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success update approval",
		Data: ApproveResponse{
			Success: append(approved["success"], rejected["success"]...),
			Fail:    append(approved["failed"], rejected["failed"]...),
		},
	}, nil
}
