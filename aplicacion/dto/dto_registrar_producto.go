package dto

type DtoRegistrarProducto struct {
	ID          string
	Nombre      string
	Descripcion string
	IdCategoria string
	Stock       int
	Estado      int
}
