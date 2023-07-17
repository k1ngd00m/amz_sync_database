package adaptador

import (
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/dao"
	"github.com/k1ngd00m/amz_sync_database/dominio/puerto/repositorio"
	daoinfra "github.com/k1ngd00m/amz_sync_database/infra/adaptador/dao"
	repositorioinfra "github.com/k1ngd00m/amz_sync_database/infra/adaptador/repositorio"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adaptador struct {
	RepositorioProducto repositorio.RepositorioProducto
	DaoCategoria        dao.DaoCategoria
}

func NewAdaptador(cliente *mongo.Client) Adaptador {

	repositorioProducto := repositorioinfra.NewMongoRepositorioProducto(cliente)
	daoCategoria := daoinfra.NewCacheDaoCategoria()

	adapt := &Adaptador{
		RepositorioProducto: repositorioProducto,
		DaoCategoria:        daoCategoria,
	}

	return *adapt

}
