package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"stock-service/config"
	"stock-service/fetcher"
	"stock-service/handler"
	"stock-service/stockstore"
	"time"

	"github.com/gorilla/mux"
)

const (
	maxRouteStocks = 40
)

var (
	apiKeyPath = os.Getenv("IEX_CLOUD_KEY")
	devMode    = os.Getenv("DEV_MODE")
	apiKey     string
)

func init() {
	if apiKeyPath == "" {
		log.Fatal("No environment variable named IEX_CLOUD_KEY present")
	} else {
		content, err := ioutil.ReadFile(apiKeyPath)
		if err != nil {
			log.Fatal(err)
		}

		apiKey = string(content)
	}
}

func main() {
	ss, err := stockstore.New(maxRouteStocks)
	if err != nil {
		log.Fatal(err)
	}

	f, err := fetcher.New(apiKey, config.MaxRouteStocks)
	if err != nil {
		log.Fatal(err)
	}

	env := &config.Env{StockStore: ss, StockClient: f}

	go env.StockClient.FetchAndSave(context.Background(), env, time.Minute*15)

	r := mux.NewRouter()

	api := r.PathPrefix("/api/stocks/opt").Subrouter()

	api.HandleFunc("/gainers", handler.HandleGainers(env)).Methods("GET")
	api.HandleFunc("/losers", handler.HandleLosers(env)).Methods("GET")
	api.HandleFunc("/most-active", handler.HandleMostActive(env)).Methods("GET")

	fmt.Println("The stock service is ready")
	log.Fatal(http.ListenAndServe(":8082", r))
}
