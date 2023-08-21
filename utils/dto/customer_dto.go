package dto

type CustomerDto struct {
	Nombre    string `json:"nombre" validate:"required"`
	Direccion string `json:"direccion" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	Telefono  string `json:"telefono" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Ciudad    string `json:"ciudad" validate:"required"`
}
