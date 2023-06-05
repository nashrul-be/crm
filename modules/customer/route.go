package customer

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

type Route struct {
	customerRequestHandler RequestHandlerInterface
}

func NewRoute(customerRepository repositories.CustomerRepositoryInterface) Route {
	useCase := NewUseCase(customerRepository)
	customerController := NewController(useCase)
	requestHandler := NewRequestHandler(customerController)
	return Route{customerRequestHandler: requestHandler}
}

func (r Route) Handle(router *gin.Engine) {
	customer := router.Group("/customers", middleware.Authenticate())
	customer.GET("/:id", r.customerRequestHandler.GetByID)
	customer.GET("", r.customerRequestHandler.GetAll)
	customer.POST("", r.customerRequestHandler.CreateCustomer)
	customer.PUT("/:id", r.customerRequestHandler.UpdateCustomer)
	customer.DELETE("/:id", r.customerRequestHandler.DeleteCustomer)
}
