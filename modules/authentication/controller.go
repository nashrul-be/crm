package authentication

import (
	"fmt"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/utils/hash"
	jwtUtil "nashrul-be/crm/utils/jwt"
	"net/http"
)

type ControllerInterface interface {
	Login(request LoginRequest) (dto.BaseResponse, error)
}

func NewAuthController(actorUseCase actor.UseCaseInterface) ControllerInterface {
	return authController{actorUseCase: actorUseCase}
}

type authController struct {
	actorUseCase actor.UseCaseInterface
}

func (c authController) Login(request LoginRequest) (dto.BaseResponse, error) {
	account, err := c.actorUseCase.GetByUsername(request.Username)
	if err != nil {
		return dto.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "Wrong Username/Password",
		}, err
	}
	if err := hash.Compare(request.Password, account.Password); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "Wrong Username/Password",
		}, err
	}
	if !account.Verified {
		return dto.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "Your account not verified yet",
		}, err
	}
	if !account.Active {
		return dto.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "Your account has been deactivated",
		}, err
	}
	token, err := jwtUtil.GenerateJWT(account)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Authenticated succes",
		Data:    LoginResponse{Token: fmt.Sprintf("Bearer %s", token)},
	}, nil
}
