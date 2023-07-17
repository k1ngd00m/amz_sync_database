package entidad

type Categoria struct {
	ID     string
	Nombre string
}

func NewCategoria(id string, nombre string) Categoria {
	categoria := &Categoria{
		ID:     id,
		Nombre: nombre,
	}

	return *categoria
}
