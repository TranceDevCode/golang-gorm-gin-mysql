package dto

type CustomerDto struct {
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Rut       string `json:"rut"`
	Telefono  string `json:"telefono"`
	Region    string `json:"region"`
	Ciudad    string `json:"ciudad"`
}
