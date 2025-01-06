package main

import (
	"fmt"
	"github.com/Rammurthy5/url-shortner-go/config"
	"github.com/Rammurthy5/url-shortner-go/internal/controllers"
	urls_mapping "github.com/Rammurthy5/url-shortner-go/internal/db/sqlc"
	"log"
	"net/http"
)

func main() {
	cfg, logger, db := config.InitDependencies()
	logger.Info("Application starts..")
	dbInst := urls_mapping.New(db)
	baseController := controllers.BaseController{Cfg: cfg, Log: logger, Db: dbInst}
	defer config.CloseDB() // gracefully close the db connection

	homeController := controllers.NewHomeController(&baseController)
	http.HandleFunc("/", homeController.ServeHandle)

	shortenController := controllers.NewShortenController(&baseController)
	http.HandleFunc("/shorten", shortenController.ServeHandle)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPConfig.Port), nil))
}
