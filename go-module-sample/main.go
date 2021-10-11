package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Index(_ http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	// NewProduction builds a sensible production Logger that writes InfoLevel and above logs to standard error as JSON
	logger, _ := zap.NewProduction() // Create a production HTTP logger listener using Zap
	logger.Info("Successfully performed http request")
	logger.Info(quote.Hello())
	logger.Info(quoteV3.HelloV3())
}

func HealthCheck(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("Service is running well")
	fmt.Println("Headers:")
	res.Header().Set("Content-Type", "application/json")
	fmt.Println(res.Header())
	res.WriteHeader(http.StatusAccepted)
	fmt.Println("Params:")
	fmt.Println(params)
}

func main() {
	CyanPrintCLI("Initializing HTTP server at port 8080... ⏳")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/health-check", HealthCheck)

	log.Fatal(http.ListenAndServe(":8080", router)) // Fatal is equivalent to Print() followed by a call to os.Exit(1).
}
