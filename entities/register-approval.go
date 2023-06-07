package entities

type RegisterApproval struct {
	ID           uint
	AdminID      uint
	Admin        Actor
	SuperAdminID uint
	Status       string
}

func (a RegisterApproval) Approve(superAdminID uint) RegisterApproval {
	a.Status = "approved"
	a.SuperAdminID = superAdminID
	return a
}

func (a RegisterApproval) Reject(superAdminID uint) RegisterApproval {
	a.Status = "rejected"
	a.SuperAdminID = superAdminID
	return a
}
