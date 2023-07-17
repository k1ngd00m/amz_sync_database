package repositorioinfra

type DocumentoCategoria struct {
	ID     string `bson:"_id"`
	Nombre string `bson:"nombre"`
}

type DocumentoProducto struct {
	ID          string              `bson:"_id"`
	Nombre      string              `bson:"nombre"`
	Descripcion string              `bson:"descripcion"`
	Categoria   *DocumentoCategoria `bson:"categoria"`
	Stock       int                 `bson:"stock"`
	Estado      int                 `bson:"estado"`
}
