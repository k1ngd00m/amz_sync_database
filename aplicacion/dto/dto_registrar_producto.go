package dto

type DtoProducto struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	IdCategoria string `json:"idCategoria"`
	Stock       int    `json:"stock"`
	Estado      int    `json:"estado"`
}
