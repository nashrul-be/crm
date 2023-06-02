package actor

import (
	"nashrul-be/crm/entities"
)

func mapCreateRequestToActor(request CreateRequest) entities.Actor {
	return entities.Actor{
		Username: request.Username,
		Password: request.Password,
		RoleID:   2,
	}
}

func mapUpdateRequestToActor(request UpdateRequest) entities.Actor {
	return entities.Actor{
		Username: request.Username,
		Password: request.Password,
	}
}

func mapActorToResponse(actor entities.Actor) Representation {
	return Representation{
		Username: actor.Username,
		Role:     actor.Role.RoleName,
		Verified: actor.Verified,
		Active:   actor.Active,
	}
}
