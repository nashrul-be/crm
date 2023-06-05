package customer

import "nashrul-be/crm/dto"

func customerNotFound() dto.BaseResponse {
	return dto.ErrorNotFound("customer")
}

type CreateRequest struct {
	FirstName string `binding:"required,alpha" json:"first_name"`
	LastName  string `binding:"required,alpha" json:"last_name"`
	Email     string `binding:"required,email" json:"email"`
	Avatar    string `json:"avatar"`
}

type UpdateRequest struct {
	ID        uint   `uri:"id" binding:"required,numeric"`
	FirstName string `binding:"omitempty,alpha" json:"first_name"`
	LastName  string `binding:"omitempty,alpha" json:"last_name"`
	Email     string `binding:"omitempty,email" json:"email"`
	Avatar    string `binding:"omitempty,uri" json:"avatar"`
}

type Representation struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type PaginationRequest struct {
	PerPage uint   `form:"perpage" binding:"required,numeric,gt=0"`
	Page    uint   `form:"page" binding:"required,numeric,gt=0"`
	Email   string `form:"email" json:"email,omitempty"`
	Name    string `form:"name" json:"name,omitempty"`
}

type ThirdPartyJSON struct {
	Data []CreateRequest `json:"data"`
}
