package repositorioinfra

import (
	"context"

	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/repositorio"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepositorioProducto struct {
	cliente *mongo.Client
}

func NewMongoRepositorioProducto(cliente *mongo.Client) repositorio.RepositorioProducto {
	repo := &MongoRepositorioProducto{
		cliente: cliente,
	}

	return repo
}

func (m *MongoRepositorioProducto) RegistrarProducto(producto *entidad.Producto) error {

	documento := &DocumentoProducto{
		ID:          producto.ID,
		Nombre:      producto.Nombre,
		Descripcion: producto.Descripcion,
		Stock:       producto.Stock,
		Estado:      producto.Estado,
		Categoria: &DocumentoCategoria{
			ID:     producto.Categoria.ID,
			Nombre: producto.Categoria.Nombre,
		},
	}

	ctx := context.TODO()

	coleccion := m.cliente.Database("catalogo").Collection("productos")
	_, err := coleccion.InsertOne(ctx, documento)

	if err != nil {
		return err
	}

	return nil

}
