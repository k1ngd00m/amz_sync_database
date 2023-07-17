package entidad

type Producto struct {
	ID          string
	Nombre      string
	Descripcion string
	Categoria   *Categoria
	Stock       int
	Estado      int
}

func NewProducto(id string, nombre string, descripcion string, stock int, estado int, categoria *Categoria) Producto {

	producto := &Producto{
		ID:          id,
		Nombre:      nombre,
		Descripcion: descripcion,
		Stock:       stock,
		Estado:      estado,
		Categoria:   categoria,
	}

	return *producto

}
