package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type CustomerRepositoryInterface interface {
	GetByID(id uint) (customer entities.Customer, err error)
	Create(customer *entities.Customer) (err error)
	Update(customer *entities.Customer) (err error)
	UpdateOrCreate(customer *entities.Customer) (err error)
	Delete(id uint) (err error)
	IsExist(id uint) (bool, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryInterface {
	return customerRepository{db: db}
}

type customerRepository struct {
	db *gorm.DB
}

func (r customerRepository) IsExist(id uint) (exist bool, err error) {
	_, err = r.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}
		return
	}
	exist = true
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
	err = r.db.Save(customer).Error
	return
}

func (r customerRepository) Delete(id uint) (err error) {
	err = r.db.Delete(&entities.Customer{}, id).Error
	return
}
