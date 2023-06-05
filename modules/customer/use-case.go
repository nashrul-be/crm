package customer

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	ValidateCustomer(customer entities.Customer, validations ...validateFunc) (error, error)
	GetByID(id uint) (customer entities.Customer, err error)
	GetAll(limit, offset uint) (customers []entities.Customer, err error)
	GetAllByEmail(email string, limit, offset uint) (customers []entities.Customer, err error)
	GetAllByName(name string, limit, offset uint) (customers []entities.Customer, err error)
	CreateCustomer(customer *entities.Customer) (err error)
	UpdateCustomer(customer *entities.Customer) (err error)
	DeleteCustomer(id uint) (err error)
}

func NewUseCase(repositoryInterface repositories.CustomerRepositoryInterface) UseCaseInterface {
	return useCase{customerRepository: repositoryInterface}
}

type useCase struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func (uc useCase) ValidateCustomer(customer entities.Customer, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(customer, uc.customerRepository)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc useCase) GetByID(id uint) (customer entities.Customer, err error) {
	return uc.customerRepository.GetByID(id)
}

func (uc useCase) GetAll(limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAll(limit, offset)
}

func (uc useCase) GetAllByEmail(email string, limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAllByEmail(email, limit, offset)
}

func (uc useCase) GetAllByName(name string, limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAllByName(name, limit, offset)
}

func (uc useCase) CreateCustomer(customer *entities.Customer) (err error) {
	return uc.customerRepository.Create(customer)
}

func (uc useCase) UpdateCustomer(customer *entities.Customer) (err error) {
	err = uc.customerRepository.Update(customer)
	return
}

func (uc useCase) DeleteCustomer(id uint) (err error) {
	err = uc.customerRepository.Delete(id)
	return
}
