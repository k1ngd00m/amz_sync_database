package servicio

import (
	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/repositorio"
)

type ServicioRegistrarProducto struct {
	repositorioProducto repositorio.RepositorioProducto
}

func NewServicioRegistrarProducto(repositorioProducto *repositorio.RepositorioProducto) ServicioRegistrarProducto {
	servicioRegistrarProducto := &ServicioRegistrarProducto{
		repositorioProducto: *repositorioProducto,
	}

	return *servicioRegistrarProducto
}

func (m *ServicioRegistrarProducto) Ejecutar(producto *entidad.Producto) error {
	return m.repositorioProducto.Registrar(producto)
}
