package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"time"
)

type CustomerRepositoryInterface interface {
	GetByID(id uint) (customer entities.Customer, err error)
	Create(customer *entities.Customer) (err error)
	Update(customer *entities.Customer) (err error)
	UpdateOrCreate(customer *entities.Customer) (err error)
	Delete(id uint) (err error)
	IsExist(id uint) (exist bool, err error)
	IsEmailExist(customer entities.Customer) (bool, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryInterface {
	return customerRepository{db: db}
}

type customerRepository struct {
	db *gorm.DB
}

func (r customerRepository) IsExist(id uint) (exist bool, err error) {
	var count int64
	result := r.db.Model(&entities.Customer{}).Where("id = ?", id).Count(&count)
	if result.Error != nil {
		return
	}
	exist = count > 0
	return
}

func (r customerRepository) IsEmailExist(customer entities.Customer) (exist bool, err error) {
	var count int64
	result := r.db.Model(&entities.Customer{}).Where("email = ?", customer.Email).Where("id != ?", customer.ID).Count(&count)
	if result.Error != nil {
		return
	}
	exist = count > 0
	return
}

func (r customerRepository) GetByID(id uint) (customer entities.Customer, err error) {
	err = r.db.First(&customer, id).Error
	return
}

func (r customerRepository) Create(customer *entities.Customer) (err error) {
	err = r.db.Create(customer).Error
	return
}

func (r customerRepository) Update(customer *entities.Customer) (err error) {
	err = r.db.Updates(customer).Error
	return
}

func (r customerRepository) UpdateOrCreate(customer *entities.Customer) (err error) {
	customer.CreatedAt = time.Now()
	err = r.db.Save(customer).Error
	return
}

func (r customerRepository) Delete(id uint) (err error) {
	err = r.db.Delete(&entities.Customer{}, id).Error
	return
}
