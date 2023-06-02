package entities

type RegisterApproval struct {
	ID           uint
	AdminID      uint
	Admin        Actor `gorm:"foreignKey:AdminID"`
	SuperAdminID uint
	Status       string
}
