package app

import (
	elasticmid "airbnb-messaging-be/internal/app/middleware/elastic"
	smsevent "airbnb-messaging-be/internal/app/sms/api/event"
	"airbnb-messaging-be/internal/pkg/cache/auth"
	elastic "airbnb-messaging-be/internal/pkg/elasticsearch"
	"airbnb-messaging-be/internal/pkg/http/server"
	httprouter "airbnb-messaging-be/internal/pkg/http/server/router"
	kafkaconsumer "airbnb-messaging-be/internal/pkg/kafka/consumer"
	kafkaproducer "airbnb-messaging-be/internal/pkg/kafka/producer"
	"airbnb-messaging-be/internal/pkg/log"
	"context"
	"sync"

	"github.com/gin-gonic/gin"
)

var Instance = "App"

type Options struct {
	HttpServer    *server.Server
	EventListener *kafkaconsumer.Listener
	EventProducer *kafkaproducer.Producer

	SmsHandler *smsevent.Handler
}

type App struct {
	Options
}

// Run all the modules of the app.
func (a App) Run(ctx context.Context) {
	a.runModules(ctx)
	a.stopModules()
}

func (a App) runModules(ctx context.Context) {
	log.Event(Instance, "Starting...")

	// init app cache
	auth.InitAuthCache()

	// Init elasticsearch client
	elastic.InitElasticSearch()

	// Create required index in elastic
	elasticmid.CreateIndex()

	// recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// Register all routes
	a.registerHttpHandler()

	// Register kafka topic handlers
	a.registerEventHandler()

	go func() {
		err := a.HttpServer.Start()
		if err != nil {
			log.Fatal(Instance, "failed to start http server", err)
		}
	}()

	go func() {
		err := a.EventListener.Start(ctx)
		if err != nil {
			log.Fatal(Instance, "failed to start event listener", err)
		}
	}()

	<-ctx.Done()
}

func (a App) stopModules() {
	log.Event(Instance, "Stoping...")

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		err := a.EventProducer.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop event producer", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.EventListener.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop event listener", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.HttpServer.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop http server", err)
		}
	}()

	wg.Wait()
}
