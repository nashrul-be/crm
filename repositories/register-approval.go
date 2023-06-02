package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type RegisterApprovalRepositoryInterface interface {
	Create(approval *entities.RegisterApproval) (err error)
	GetAllPendingApproval() (approvals []entities.RegisterApproval, err error)
	GetByAdminID(id uint) (approval entities.RegisterApproval, err error)
	GetByAdminIdBatch(id []uint) (approvals []entities.RegisterApproval, err error)
	Update(approval *entities.RegisterApproval) (err error)
	InitTransaction() (*gorm.DB, error)
	Begin(db *gorm.DB) RegisterApprovalRepositoryInterface
}

func NewRegisterApprovalRepository(db *gorm.DB) RegisterApprovalRepositoryInterface {
	return registerApprovalRepository{db: db}
}

type registerApprovalRepository struct {
	db *gorm.DB
}

func (r registerApprovalRepository) InitTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r registerApprovalRepository) Begin(db *gorm.DB) RegisterApprovalRepositoryInterface {
	return registerApprovalRepository{db: db}
}

func (r registerApprovalRepository) Create(approval *entities.RegisterApproval) (err error) {
	err = r.db.Omit("SuperAdminID").Create(approval).Error
	return
}

func (r registerApprovalRepository) GetAllPendingApproval() (approvals []entities.RegisterApproval, err error) {
	err = r.db.Preload("Admin").Find(&approvals, "status = ?", "pending").Error
	return
}

func (r registerApprovalRepository) GetByAdminID(id uint) (approval entities.RegisterApproval, err error) {
	err = r.db.First(&approval, "admin_id = ?", id).Error
	return
}

func (r registerApprovalRepository) GetByAdminIdBatch(id []uint) (approvals []entities.RegisterApproval, err error) {
	err = r.db.Find(&approvals, "admin_id IN ?", id).Error
	return
}

func (r registerApprovalRepository) Update(approval *entities.RegisterApproval) (err error) {
	err = r.db.Updates(approval).Error
	return
}
