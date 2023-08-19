package controllers

import (
	"net/http"

	"gotenancy/config"
	"gotenancy/models"

	"gotenancy/utils/dto"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func ListCustomers(c *gin.Context) {
	customers := models.Customers{}
	config.Database.Order("id desc").Find(&customers)

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
	})
}

func ListCustomersId(c *gin.Context) {
	id := c.Param("id")
	customer := models.Customers{}
	if err := config.Database.First(&customer, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "Error",
			"mensaje": "Recurso no disponible",
		}
		c.JSON(http.StatusNotFound, respuesta)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"customer": customer,
		})
	}
}

func CreateCustomer(c *gin.Context) {
	var customer dto.CustomerDto
	//recibimos los parametros del front y validamos para guardar o mostrar el o los errores
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado": err.Error(),
		})
		return
	}
	// Usamos el validador para verificar si la estructura cumple con las reglas definidas
	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"message": err.Error(),
		})
		return
	}
	datos := models.Customer{Nombre: customer.Nombre, Direccion: customer.Direccion, Rut: customer.Rut, Telefono: customer.Telefono, Region: customer.Region, Ciudad: customer.Ciudad}
	config.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "OK",
		"mensaje": "Cliente creado con exito",
	}
	c.JSON(http.StatusCreated, gin.H{
		//"Create":   "Creating a New Customer",
		"respuesta": respuesta,
		//"customer":  customer,
	})
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer dto.CustomerDto
	//recibimos los parametros del front y validamos para guardar o mostrar el o los errores
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado": err.Error(),
		})
		return
	}
	// Usamos el validador para verificar si la estructura cumple con las reglas definidas
	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"message": err.Error(),
		})
		return
	}
	//ahora vamos a buscar el registro de la categoria
	datos := models.Customer{}
	if err := config.Database.First(&datos, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "El registro que intenta actualizar no existe",
		})
	} else { //Si es que encuentra el registro
		datos.Nombre = customer.Nombre
		datos.Direccion = customer.Direccion
		datos.Rut = customer.Rut
		datos.Telefono = customer.Telefono
		datos.Region = customer.Region
		datos.Ciudad = customer.Ciudad
		config.Database.Save(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "El registro se actualizo correctamente",
		}
		c.JSON(http.StatusOK, gin.H{
			"respuesta": respuesta,
		})
	}
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	//Buscamos el registro
	datos := models.Customer{}
	if err := config.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error, el registro que intenta eliminar no existe",
		}
		c.JSON(http.StatusNotFound, gin.H{
			"respuesta": respuesta,
		})
		return
	} else {
		config.Database.Delete(&datos)
		respuesta := map[string]string{
			"estado":  "OK",
			"mensaje": "El registro se elimino correctamente",
		}
		c.JSON(http.StatusOK, gin.H{
			"respuesta": respuesta,
		})
	}
}
