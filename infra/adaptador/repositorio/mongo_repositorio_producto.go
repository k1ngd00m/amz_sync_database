package repositorioinfra

import (
	"context"
	"errors"

	"github.com/k1ngd00m/amz_sync_database/dominio/entidad"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/repositorio"
	"go.mongodb.org/mongo-driver/bson"
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

func (m *MongoRepositorioProducto) Registrar(producto *entidad.Producto) error {

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

func (m *MongoRepositorioProducto) Actualizar(producto *entidad.Producto) error {

	ctx := context.TODO()
	coleccion := m.cliente.Database("catalogo").Collection("productos")

	filter := bson.D{{Key: "_id", Value: producto.ID}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "nombre", Value: producto.Nombre},
		{Key: "descripcion", Value: producto.Descripcion},
		{Key: "stock", Value: producto.Stock},
		{Key: "estado", Value: producto.Estado},
		{Key: "categoria", Value: bson.D{
			{Key: "_id", Value: producto.Categoria.ID},
			{Key: "nombre", Value: producto.Categoria.Nombre},
		}},
	}}}

	result, err := coleccion.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("No se actualizo ningun registro")
	}

	return nil
}
