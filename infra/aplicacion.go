package infra

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k1ngd00m/amz_sync_database/aplicacion/manejador"
	"github.com/k1ngd00m/amz_sync_database/dominio/servicio"
	"github.com/k1ngd00m/amz_sync_database/infra/adaptador"
	"github.com/k1ngd00m/amz_sync_database/infra/config"
	"github.com/k1ngd00m/amz_sync_database/infra/entrypoint/suscriptor"
)

func Start() {

	config.LoadEnvironment()

	log := config.GetLogger()

	// inicializar conexion a rdms read
	db := config.GetReadDBConexion()
	log.Info("conexion a read catalogo db existosa")

	// inicializar rabbitMQ
	rabbitMQ := config.GetRabbitMqConn()
	log.Info("conexion a rabbitMQ existosa")

	adapt := adaptador.NewAdaptador(db)

	servicioBuscarCategoria := servicio.NewServicioBuscarCategoria(&adapt.DaoCategoria)
	servicioRegistrarProducto := servicio.NewServicioRegistrarProducto(&adapt.RepositorioProducto)
	servicioActualizarProducto := servicio.NewServicioActualizarProducto(&adapt.RepositorioProducto)

	manejadorRegistrarProducto := manejador.NewManejadorRegistrarProducto(&servicioRegistrarProducto, &servicioBuscarCategoria)
	manejadorActualizarProducto := manejador.NewManejadorActualizarProducto(&servicioActualizarProducto, &servicioBuscarCategoria)

	// entrypoint message
	consumidorConfig := config.NewRabbitMQConsumer(rabbitMQ)

	suscriptor.NewSuscriptorRegistrarProducto(&consumidorConfig, &manejadorRegistrarProducto, &manejadorActualizarProducto)

	app := chi.NewRouter()

	// Default middleware config
	app.Use(middleware.RequestID)

	http.ListenAndServe(":3334", app)

}
