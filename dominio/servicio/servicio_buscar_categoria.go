package servicio

import (
	"github.com/k1ngd00m/amz_sync_database/aplicacion/dto"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/dao"
)

type ServicioBuscarCategoria struct {
	daoCategoria dao.DaoCategoria
}

func NewServicioBuscarCategoria(daoCategoria *dao.DaoCategoria) ServicioBuscarCategoria {
	servicioBuscarCategoria := &ServicioBuscarCategoria{
		daoCategoria: *daoCategoria,
	}

	return *servicioBuscarCategoria
}

func (m *ServicioBuscarCategoria) Ejecutar(idCategoria string) (*dto.DtoCategoria, error) {
	return m.daoCategoria.ObtenerPorId(idCategoria)
}
