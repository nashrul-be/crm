package customer

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/repositories"
)

type CustomerRoute struct {
	customerRequestHandler RequestHandlerInterface
}

func NewCustomerRoute(customerRepository repositories.CustomerRepositoryInterface) CustomerRoute {
	useCase := NewUseCase(customerRepository)
	customerController := NewController(useCase)
	requestHandler := NewRequestHandler(customerController)
	return CustomerRoute{customerRequestHandler: requestHandler}
}

func (r CustomerRoute) Handle(router *gin.Engine) {
	customer := router.Group("/customer")
	customer.GET("/:id", r.customerRequestHandler.GetByID)
	customer.POST("", r.customerRequestHandler.CreateCustomer)
	customer.PUT("/:id", r.customerRequestHandler.UpdateOrCreateCustomer)
	customer.DELETE("/:id", r.customerRequestHandler.DeleteCustomer)
}
