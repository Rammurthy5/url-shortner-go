package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rammurthy5/url-shortner-go/config"
	"github.com/Rammurthy5/url-shortner-go/internal/controllers"
	urls_mapping "github.com/Rammurthy5/url-shortner-go/internal/db/sqlc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	cfg, logger, db, midlware, _ := config.InitDependencies()
	logger.Info("Application starts..")
	dbInst := urls_mapping.New(db)
	defer config.ShutdownLogger()
	defer config.ShutDownDB()
	defer config.ShutDownCache()
	defer config.ShutdownTracer()

	baseController := controllers.BaseController{Cfg: cfg, Log: logger, Db: dbInst}

	homeController := controllers.NewHomeController(&baseController)
	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(homeController.ServeHandle), "home"))

	http.HandleFunc("/", homeController.ServeHandle)

	shortenController := controllers.NewShortenController(&baseController)
	http.Handle("/shorten", otelhttp.NewHandler(
		http.HandlerFunc(midlware.CheckIdempotency(shortenController.ServeHandle)),
		"shorten",
	))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPConfig.Port), nil))
}
