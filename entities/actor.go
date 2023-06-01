package entities

import "time"

type Actor struct {
	ID        uint
	Username  string
	Password  string
	RoleID    uint
	Role      Role
	Verified  bool
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
