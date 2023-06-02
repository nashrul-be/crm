package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"time"
)

type ActorRepositoryInterface interface {
	GetByID(id uint) (customer entities.Actor, err error)
	IsUsernameExist(actor entities.Actor) (exist bool, err error)
	IsExist(id uint) (exist bool, err error)
	Create(customer *entities.Actor) (err error)
	Update(customer *entities.Actor) (err error)
	UpdateOrCreate(customer *entities.Actor) (err error)
	Delete(id uint) (err error)
}

func NewActorRepository(db *gorm.DB) ActorRepositoryInterface {
	return actorRepository{db: db}
}

type actorRepository struct {
	db *gorm.DB
}

func (r actorRepository) GetByID(id uint) (actor entities.Actor, err error) {
	err = r.db.First(&actor, id).Error
	return
}

func (r actorRepository) IsExist(id uint) (exist bool, err error) {
	var count int64
	result := r.db.Model(&entities.Actor{}).Where("id = ?", id).Count(&count)
	if result.Error != nil {
		return
	}
	exist = count > 0
	return
}

func (r actorRepository) IsUsernameExist(actor entities.Actor) (exist bool, err error) {
	var count int64
	result := r.db.Model(&entities.Actor{}).Where("username = ?", actor.Username).Where("id != ?", actor.ID).Count(&count)
	if result.Error != nil {
		return
	}
	exist = count > 0
	return
}

func (r actorRepository) Create(customer *entities.Actor) (err error) {
	err = r.db.Create(customer).Error
	return
}

func (r actorRepository) Update(customer *entities.Actor) (err error) {
	err = r.db.Updates(customer).Error
	return
}

func (r actorRepository) UpdateOrCreate(customer *entities.Actor) (err error) {
	customer.CreatedAt = time.Now()
	err = r.db.Save(customer).Error
	return
}

func (r actorRepository) Delete(id uint) (err error) {
	err = r.db.Delete(&entities.Actor{}, id).Error
	return
}
