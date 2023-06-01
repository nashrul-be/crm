package customer

type CreateRequest struct {
	FirstName string `binding:"required,alpha"`
	LastName  string `binding:"required,alpha"`
	Email     string `binding:"required,email"`
	Avatar    string
}

type Representation struct {
	FirstName string
	LastName  string
	Email     string
	Avatar    string
}
