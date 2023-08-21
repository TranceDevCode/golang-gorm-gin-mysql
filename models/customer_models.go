package models

type Customer struct {
	Id        uint   `json:"id"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Rut       string `json:"rut"`
	Telefono  string `json:"telefono"`
	Region    string `json:"region"`
	Ciudad    string `json:"ciudad"`
}

// Esto tiene el arrego de customer, ya que retornamos muchos customers
type Customers []Customer
