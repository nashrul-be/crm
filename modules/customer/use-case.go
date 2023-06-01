package customer

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetByID(id uint) (customer entities.Customer, err error)
	CreateCustomer(customer *entities.Customer) (err error)
	UpdateOrCreateCustomer(customer *entities.Customer) (err error)
	DeleteCustomer(id uint) (err error)
}

func NewUseCase(repositoryInterface repositories.CustomerRepositoryInterface) UseCaseInterface {
	return customerUseCase{customerRepository: repositoryInterface}
}

type customerUseCase struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func (uc customerUseCase) GetByID(id uint) (customer entities.Customer, err error) {
	customer, err = uc.customerRepository.GetByID(id)
	return
}

func (uc customerUseCase) CreateCustomer(customer *entities.Customer) (err error) {
	err = uc.customerRepository.Create(customer)
	return
}

func (uc customerUseCase) UpdateOrCreateCustomer(customer *entities.Customer) (err error) {
	err = uc.customerRepository.UpdateOrCreate(customer)
	return
}

func (uc customerUseCase) DeleteCustomer(id uint) (err error) {
	err = uc.customerRepository.Delete(id)
	return
}
