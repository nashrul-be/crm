package register_approval

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type RegisterApprovalUseCaseInterface interface {
	GetAllPendingApproval() (approvals []entities.RegisterApproval, err error)
	Approve(username []string, superAdminID uint) (result map[string][]string, err error)
	Rejected(username []string, superAdminID uint) (result map[string][]string, err error)
}

func NewRegisterApprovalUseCase(
	approvalRepository repositories.RegisterApprovalRepositoryInterface,
	actorRepository repositories.ActorRepositoryInterface,
) RegisterApprovalUseCaseInterface {
	return registerApprovalUseCase{
		registerRepository: approvalRepository,
		actorRepository:    actorRepository,
	}
}

type registerApprovalUseCase struct {
	registerRepository repositories.RegisterApprovalRepositoryInterface
	actorRepository    repositories.ActorRepositoryInterface
}

func (uc registerApprovalUseCase) GetAllPendingApproval() (approvals []entities.RegisterApproval, err error) {
	approvals, err = uc.registerRepository.GetAllPendingApproval()
	return
}

func (uc registerApprovalUseCase) Approve(username []string, superAdminID uint) (result map[string][]string, err error) {
	actors, err := uc.actorRepository.GetByUsernameBatch(username)
	if err != nil {
		return
	}
	var ids []uint
	for _, actor := range actors {
		ids = append(ids, actor.ID)
	}
	approvals, err := uc.registerRepository.GetByAdminIdBatch(ids)
	if err != nil {
		return
	}
	result = map[string][]string{
		"success": {},
		"failed":  {},
	}
	for index, _ := range actors {
		err = uc.approved(&actors[index], &approvals[index], superAdminID)
		if err != nil {
			result["success"] = append(result["success"], actors[index].Username)
		} else {
			result["failed"] = append(result["failed"], actors[index].Username)
		}
	}
	return
}

func (uc registerApprovalUseCase) approved(
	actor *entities.Actor,
	approval *entities.RegisterApproval,
	superAdminID uint,
) (err error) {
	tx, err := uc.registerRepository.InitTransaction()
	if err != nil {
		return
	}
	actorTx := uc.actorRepository.Begin(tx)
	registerTx := uc.registerRepository.Begin(tx)
	approval.Status = "approved"
	approval.SuperAdminID = superAdminID
	if err = registerTx.Update(approval); err != nil {
		tx.Rollback()
		return
	}
	actor.Verified = true
	actor.Active = true
	if err = actorTx.Update(actor); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (uc registerApprovalUseCase) Rejected(username []string, superAdminID uint) (result map[string][]string, err error) {
	actors, err := uc.actorRepository.GetByUsernameBatch(username)
	if err != nil {
		return
	}
	var ids []uint
	for _, actor := range actors {
		ids = append(ids, actor.ID)
	}
	approvals, err := uc.registerRepository.GetByAdminIdBatch(ids)
	result = map[string][]string{
		"success": {},
		"failed":  {},
	}
	if err != nil {
		return
	}
	for index, _ := range actors {
		err = uc.approved(&actors[index], &approvals[index], superAdminID)
		if err != nil {
			result["success"] = append(result["success"], actors[index].Username)
		} else {
			result["failed"] = append(result["failed"], actors[index].Username)
		}
	}
	return
}

func (uc registerApprovalUseCase) rejected(approval *entities.RegisterApproval, superAdminID uint) (err error) {
	approval.Status = "rejected"
	approval.SuperAdminID = superAdminID
	if err = uc.registerRepository.Update(approval); err != nil {
		return
	}
	return
}
