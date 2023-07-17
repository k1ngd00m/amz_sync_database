package dao

import "github.com/k1ngd00m/amz_sync_database/aplicacion/dto"

type DaoCategoria interface {
	ObtenerPorId(id string) (*dto.DtoCategoria, error)
}
