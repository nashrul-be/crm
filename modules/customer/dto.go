package customer

type CreateRequest struct {
	FirstName string `binding:"required,alpha" json:"first_name"`
	LastName  string `binding:"required,alpha" json:"last_name"`
	Email     string `binding:"required,email" json:"email"`
	Avatar    string `json:"avatar"`
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
	Email   string `form:"email"`
	Name    string `form:"name"`
}

type ThirdPartyJSON struct {
	Data []CreateRequest `json:"data"`
}
