package main

import (
	"fmt"
	"github.com/Rammurthy5/url-shortner-go/config"
	"github.com/Rammurthy5/url-shortner-go/internal/controllers"
	"log"
	"net/http"
)

func main() {
	cfg, logger, db := config.InitDependencies()
	logger.Info("Application starts..")
	http.HandleFunc("/", controllers.ShowHomePage)
	http.HandleFunc("/shorten", controllers.ShortenHandle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpConfig.Port), nil))
}
