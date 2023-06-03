package actor

//TODO: create custom binding rule for username and password

type CreateRequest struct {
	Username string `json:"username" binding:"required,printascii"`
	Password string `json:"password" binding:"required,printascii"`
}

type UpdateRequest struct {
	ID       uint
	Username string `json:"username" binding:"printascii"`
	Password string `json:"password" binding:"printascii"`
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
