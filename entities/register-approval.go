package entities

type RegisterApproval struct {
	ID           uint
	AdminID      uint
	Admin        Actor
	SuperAdminID uint
	Status       string
}
