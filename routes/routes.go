package routes

import (
	"gotenancy/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	customerGroup := router.Group("/customers")

	// Ruta para clientes
	customerGroup.GET("/index", controllers.ListCustomers)
	customerGroup.GET("/:id", controllers.ListCustomersId)
	customerGroup.POST("/create", controllers.CreateCustomer)
	customerGroup.PUT("/update/:id", controllers.UpdateCustomer)
	customerGroup.DELETE("/delete/:id", controllers.DeleteCustomer)

	// Otras rutas para crear, actualizar, eliminar clientes, etc.
}
