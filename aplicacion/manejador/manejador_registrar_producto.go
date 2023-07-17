package manejador

import (
	"github.com/k1ngd00m/amz_sync_database/aplicacion/dto"
	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/servicio"
)

type ManejadorRegistrarProducto struct {
	servicioRegistrarProducto servicio.ServicioRegistrarProducto
	servicioBuscarCategoria   servicio.ServicioBuscarCategoria
}

func NewManejadorRegistrarProducto(servicioRegistrarProducto *servicio.ServicioRegistrarProducto,
	servicioBuscarCategoria *servicio.ServicioBuscarCategoria) ManejadorRegistrarProducto {

	manejadorRegistrarProducto := &ManejadorRegistrarProducto{
		servicioRegistrarProducto: *servicioRegistrarProducto,
		servicioBuscarCategoria:   *servicioBuscarCategoria,
	}

	return *manejadorRegistrarProducto
}

func (m *ManejadorRegistrarProducto) Ejecutar(dtoProducto dto.DtoRegistrarProducto) error {

	categoria, err := m.servicioBuscarCategoria.Ejecutar(dtoProducto.IdCategoria)

	if err != nil {
		return err
	}

	entidadCategoria := entidad.NewCategoria(categoria.ID, categoria.Nombre)

	entidadProducto := entidad.NewProducto(dtoProducto.ID,
		dtoProducto.Nombre,
		dtoProducto.Descripcion,
		dtoProducto.Stock,
		dtoProducto.Estado, &entidadCategoria)

	return m.servicioRegistrarProducto.Ejecutar(&entidadProducto)
}
