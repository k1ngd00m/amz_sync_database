package manejador

import (
	"github.com/k1ngd00m/amz_sync_database/aplicacion/dto"
	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/servicio"
)

type ManejadorActualizarProducto struct {
	servicioBuscarCategoria    servicio.ServicioBuscarCategoria
	servicioActualizarProducto servicio.ServicioActualizarProducto
}

func NewManejadorActualizarProducto(servicioActualizarProducto *servicio.ServicioActualizarProducto,
	servicioBuscarCategoria *servicio.ServicioBuscarCategoria) ManejadorActualizarProducto {

	return *&ManejadorActualizarProducto{
		servicioActualizarProducto: *servicioActualizarProducto,
		servicioBuscarCategoria:    *servicioBuscarCategoria,
	}
}

func (m *ManejadorActualizarProducto) Ejecutar(dtoProducto *dto.DtoProducto) error {

	dtoCategoria, err := m.servicioBuscarCategoria.Ejecutar(dtoProducto.IdCategoria)

	if err != nil {
		return err
	}

	categoria := entidad.NewCategoria(dtoCategoria.ID, dtoCategoria.Nombre)

	producto := entidad.NewProducto(dtoProducto.ID,
		dtoProducto.Nombre,
		dtoProducto.Descripcion,
		dtoProducto.Stock,
		dtoProducto.Estado,
		&categoria)

	return m.servicioActualizarProducto.Ejecutar(&producto)

}
