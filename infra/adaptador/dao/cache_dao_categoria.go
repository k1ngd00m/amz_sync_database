package daoinfra

import (
	"errors"
	"time"

	"github.com/k1ngd00m/amz_sync_database/aplicacion/dto"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/dao"
	"github.com/patrickmn/go-cache"
)

type CacheDaoCategoria struct {
	almacenDatos cache.Cache
}

func NewCacheDaoCategoria() dao.DaoCategoria {

	almacenDatos := cache.New(5*time.Minute, 10*time.Minute)

	key := "550e8400-e29b-41d4-a716-446655440000"
	dto := &dto.DtoCategoria{
		ID:     key,
		Nombre: "despensa",
	}

	almacenDatos.Add(key, dto, cache.NoExpiration)

	daoCategoria := &CacheDaoCategoria{
		almacenDatos: *almacenDatos,
	}

	return daoCategoria
}

func (m *CacheDaoCategoria) ObtenerPorId(id string) (*dto.DtoCategoria, error) {
	if x, found := m.almacenDatos.Get("foo"); found {
		categoria := x.(*dto.DtoCategoria)
		return categoria, nil
	}

	return nil, errors.New("no existe la categoria")
}
