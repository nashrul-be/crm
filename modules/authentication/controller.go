package authentication

import (
	"fmt"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/utils/hash"
	jwtUtil "nashrul-be/crm/utils/jwt"
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
	defaultResponse := dto.ErrorUnauthorized("Wrong Username/Password")
	if err != nil {
		return defaultResponse, err
	}
	if err := hash.Compare(request.Password, account.Password); err != nil {
		return defaultResponse, err
	}
	if !account.Verified {
		return dto.ErrorUnauthorized("Your account not verified yet"), err
	}
	if !account.Active {
		return dto.ErrorUnauthorized("Your account has been deactivated"), err
	}
	token, err := jwtUtil.GenerateJWT(account)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	result := LoginResponse{Token: fmt.Sprintf("Bearer %s", token)}
	return dto.Success("Authenticated success", result), nil
}
