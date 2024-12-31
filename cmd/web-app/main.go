package main

import (
	"fmt"
	"github.com/Rammurthy5/url-shortner-go/config"
	"github.com/Rammurthy5/url-shortner-go/internal/controllers"
	"log"
	"net/http"
)

func main() {
	configData, err := config.Load()
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", controllers.ShowHomePage)
	http.HandleFunc("/shorten", controllers.ShortenHandle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configData.HttpConfig.Port), nil))
}
