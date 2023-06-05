package actor

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type validateFunc func(actor entities.Actor, actorRepo repositories.ActorRepositoryInterface) (error, error)

func validateUsername(actor entities.Actor, actorRepo repositories.ActorRepositoryInterface) (error, error) {
	exist, err := actorRepo.IsUsernameExist(actor)
	if err != nil {
		return nil, err
	}
	if exist {
		return errors.New("username already taken"), nil
	}
	return nil, nil
}

func validateId(actor entities.Actor, actorRepo repositories.ActorRepositoryInterface) (error, error) {
	exist, err := actorRepo.IsExist(actor.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return errors.New("actor doesn't exist"), nil
	}
	return nil, nil
}
