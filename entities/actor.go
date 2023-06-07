package entities

import (
	"errors"
	"time"
)

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

func (a Actor) Verify() Actor {
	a.Verified = true
	return a
}

func (a Actor) Activate() Actor {
	a.Active = true
	return a
}

func (a Actor) Deactivate() Actor {
	a.Active = false
	return a
}

func (a Actor) CanLogin() error {
	if !a.Verified {
		return errors.New("your account not verified yet")
	}
	if !a.Active {
		return errors.New("your account has been deactivated")
	}
	return nil
}
