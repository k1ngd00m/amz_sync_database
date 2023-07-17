package repositorio

import "github.com/k1ngd00m/amz_sync_database/dominio/entidad"

type RepositorioProducto interface {
	RegistrarProducto(producto *entidad.Producto) error
}
