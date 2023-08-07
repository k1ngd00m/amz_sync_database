package servicio

import (
	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/repositorio"
)

type ServicioActualizarProducto struct {
	repostorioProducto repositorio.RepositorioProducto
}

func NewServicioActualizarProducto(repostorioProducto *repositorio.RepositorioProducto) ServicioActualizarProducto {

	return *&ServicioActualizarProducto{
		repostorioProducto: *repostorioProducto,
	}
}

func (m *ServicioActualizarProducto) Ejecutar(producto *entidad.Producto) error {
	return m.repostorioProducto.Actualizar(producto)
}
