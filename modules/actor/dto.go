package actor

import "nashrul-be/crm/dto"

//TODO: create custom binding rule for username and password

func actorNotFound() dto.BaseResponse {
	return dto.ErrorNotFound("Actor")
}

type CreateRequest struct {
	Username string `json:"username" binding:"required,printascii"`
	Password string `json:"password" binding:"required,printascii"`
}

type UpdateRequest struct {
	ID       uint   `uri:"id" binding:"required,numeric"`
	Username string `json:"username" binding:"omitempty,printascii"`
	Password string `json:"password" binding:"omitempty,printascii"`
}

type ChangeActiveRequest struct {
	Activate   []string `json:"activate"`
	Deactivate []string `json:"deactivate"`
}

type Representation struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
	Active   bool   `json:"active"`
}

type PaginationRequest struct {
	PerPage  uint   `form:"perpage" binding:"numeric,gt=0"`
	Page     uint   `form:"page" binding:"numeric,gt=0"`
	Username string `form:"username"`
}
