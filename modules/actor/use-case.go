package actor

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/hash"
)

type UseCaseInterface interface {
	GetByID(id uint) (actor entities.Actor, err error)
	GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error)
	GetByUsername(username string) (actor entities.Actor, err error)
	CreateActor(actor *entities.Actor) (err error)
	ActivateActor(usernames []string) (map[string][]string, error)
	DeactivateActor(usernames []string) (map[string][]string, error)
	UpdateActor(actor *entities.Actor) (err error)
	DeleteActor(id uint) (err error)
	IsUsernameExist(actor entities.Actor) (exist bool, err error)
	IsExist(id uint) (exist bool, err error)
}

func NewUseCase(
	repositoryInterface repositories.ActorRepositoryInterface,
	roleRepositoryInterface repositories.RoleRepositoryInterface,
	registerApprovalRepositoryInterface repositories.RegisterApprovalRepositoryInterface,
) UseCaseInterface {
	return actorUseCase{
		actorRepository:            repositoryInterface,
		roleRepository:             roleRepositoryInterface,
		registerApprovalRepository: registerApprovalRepositoryInterface,
	}
}

type actorUseCase struct {
	actorRepository            repositories.ActorRepositoryInterface
	roleRepository             repositories.RoleRepositoryInterface
	registerApprovalRepository repositories.RegisterApprovalRepositoryInterface
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

func (uc actorUseCase) GetByUsername(username string) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByUsername(username)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc actorUseCase) GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error) {
	actors, err := uc.actorRepository.GetAllByUsername(username, limit, offset)
	return actors, err
}

func (uc actorUseCase) CreateActor(actor *entities.Actor) (err error) {
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

func notFound(queries []string, results []entities.Actor) []string {
	mapResult := make(map[string]bool)
	for _, result := range results {
		mapResult[result.Username] = true
	}
	diff := make([]string, 0)
	for _, query := range queries {
		if !mapResult[query] {
			diff = append(diff, query)
		}
	}
	return diff
}

func (uc actorUseCase) changeActiveActor(usernames []string, value bool) (map[string][]string, error) {
	actors, err := uc.actorRepository.GetByUsernameBatch(usernames)
	failed := notFound(usernames, actors)
	if err != nil {
		return nil, err
	}
	result := map[string][]string{
		"success": {},
		"failed":  failed,
	}
	for _, actor := range actors {
		if !actor.Verified {
			result["failed"] = append(result["failed"], actor.Username)
			continue
		}
		actor.Active = value
		if err := uc.actorRepository.UpdateOrCreate(&actor); err != nil {
			result["failed"] = append(result["failed"], actor.Username)
		} else {
			result["success"] = append(result["success"], actor.Username)
		}
	}
	return result, nil
}

func (uc actorUseCase) ActivateActor(usernames []string) (map[string][]string, error) {
	return uc.changeActiveActor(usernames, true)
}

func (uc actorUseCase) DeactivateActor(usernames []string) (map[string][]string, error) {
	return uc.changeActiveActor(usernames, false)
}

func (uc actorUseCase) DeleteActor(id uint) (err error) {
	err = uc.actorRepository.Delete(id)
	return
}
