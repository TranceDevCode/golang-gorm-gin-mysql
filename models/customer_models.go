package models

type Customer struct {
	Id        uint   `json:"id"`
	Nombre    string `json:"nombre" validate:"required"`
	Direccion string `json:"direccion" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	Telefono  string `json:"telefono" validate:"required"`
	Region    string `json:"region"`
	Ciudad    string `json:"ciudad"`
}

// Esto tiene el arrego de customer, ya que retornamos muchos customers
type Customers []Customer
