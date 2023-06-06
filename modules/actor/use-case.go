package actor

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/hash"
)

type UseCaseInterface interface {
	validateActor(actor entities.Actor, validations ...validateFunc) (error, error)
	GetByID(id uint) (actor entities.Actor, err error)
	GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error)
	GetByUsername(username string) (actor entities.Actor, err error)
	CreateActor(actor *entities.Actor) (err error)
	ActivateActor(usernames string) error
	DeactivateActor(usernames string) error
	UpdateActor(actor *entities.Actor) (err error)
	DeleteActor(id uint) (err error)
}

func NewUseCase(
	repositoryInterface repositories.ActorRepositoryInterface,
	roleRepositoryInterface repositories.RoleRepositoryInterface,
	registerApprovalRepositoryInterface repositories.RegisterApprovalRepositoryInterface,
) UseCaseInterface {
	return useCase{
		actorRepository:            repositoryInterface,
		roleRepository:             roleRepositoryInterface,
		registerApprovalRepository: registerApprovalRepositoryInterface,
	}
}

type useCase struct {
	actorRepository            repositories.ActorRepositoryInterface
	roleRepository             repositories.RoleRepositoryInterface
	registerApprovalRepository repositories.RegisterApprovalRepositoryInterface
}

func (uc useCase) validateActor(actor entities.Actor, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(actor, uc.actorRepository)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc useCase) GetByID(id uint) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByID(id)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc useCase) GetByUsername(username string) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByUsername(username)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc useCase) GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error) {
	actors, err := uc.actorRepository.GetAllByUsername(username, limit, offset)
	return actors, err
}

func (uc useCase) CreateActor(actor *entities.Actor) (err error) {
	actor.Password, err = hash.Hash(actor.Password)
	if err != nil {
		return
	}
	tx, err := uc.actorRepository.InitTransaction()
	if err != nil {
		return
	}
	actorTx := uc.actorRepository.Begin(tx)
	registerApprovalTx := uc.registerApprovalRepository.Begin(tx)
	if err = actorTx.Create(actor); err != nil {
		tx.Rollback()
		return
	}
	approval := &entities.RegisterApproval{AdminID: actor.ID, Status: "pending"}
	if err = registerApprovalTx.Create(approval); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	*actor, err = uc.GetByID(actor.ID)
	return
}

func (uc useCase) UpdateActor(actor *entities.Actor) (err error) {
	if actor.Password != "" {
		actor.Password, err = hash.Hash(actor.Password)
		if err != nil {
			return
		}
	}
	err = uc.actorRepository.Update(actor)
	if err != nil {
		return
	}
	*actor, err = uc.GetByID(actor.ID)
	return
}

func (uc useCase) changeActiveActor(username string, value bool) error {
	actor, err := uc.actorRepository.GetByUsername(username)
	if err != nil {
		return err
	}
	if !actor.Verified {
		return errors.New("actor not verified yet")
	}
	actor.Active = value
	return uc.actorRepository.Save(&actor)
}

func (uc useCase) ActivateActor(username string) error {
	return uc.changeActiveActor(username, true)
}

func (uc useCase) DeactivateActor(username string) error {
	return uc.changeActiveActor(username, false)
}

func (uc useCase) DeleteActor(id uint) (err error) {
	tx, err := uc.actorRepository.InitTransaction()
	if err != nil {
		return
	}
	actorTx := uc.actorRepository.Begin(tx)
	registerApprovalTx := uc.registerApprovalRepository.Begin(tx)
	if err = registerApprovalTx.DeleteByAdminId(id); err != nil {
		tx.Rollback()
		return
	}
	if err = actorTx.Delete(id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
