package actor

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/hash"
)

type UseCaseInterface interface {
	GetByID(id uint) (actor entities.Actor, err error)
	CreateActor(actor *entities.Actor) (err error)
	UpdateActor(actor *entities.Actor) (err error)
	DeleteActor(id uint) (err error)
	IsUsernameExist(actor entities.Actor) (exist bool, err error)
	IsExist(id uint) (exist bool, err error)
}

func NewUseCase(repositoryInterface repositories.ActorRepositoryInterface, roleRepositoryInterface repositories.RoleRepositoryInterface) UseCaseInterface {
	return actorUseCase{
		actorRepository: repositoryInterface,
		roleRepository:  roleRepositoryInterface,
	}
}

type actorUseCase struct {
	actorRepository repositories.ActorRepositoryInterface
	roleRepository  repositories.RoleRepositoryInterface
}

func (uc actorUseCase) IsExist(id uint) (exist bool, err error) {
	exist, err = uc.actorRepository.IsExist(id)
	return
}

func (uc actorUseCase) IsUsernameExist(actor entities.Actor) (exist bool, err error) {
	exist, err = uc.actorRepository.IsUsernameExist(actor)
	return
}

func (uc actorUseCase) GetByID(id uint) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByID(id)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc actorUseCase) CreateActor(actor *entities.Actor) (err error) {
	actor.Password, err = hash.Hash(actor.Password)
	if err != nil {
		return
	}
	err = uc.actorRepository.Create(actor)
	if err != nil {
		return
	}
	*actor, err = uc.GetByID(actor.ID)
	return
}

func (uc actorUseCase) UpdateActor(actor *entities.Actor) (err error) {
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

func (uc actorUseCase) DeleteActor(id uint) (err error) {
	err = uc.actorRepository.Delete(id)
	return
}
