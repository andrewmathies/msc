package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrewmathies/msl/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	//env.Parse()

	l := log.New(os.Stdout, "erd-api ", log.LstdFlags)

	// building the handlers
	erdHandler := handlers.NewERDs(l)

	// create a serve multiplexer and register handlers
	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	putRouter := mux.Methods(http.MethodPut).Subrouter()
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	deleteRouter := mux.Methods(http.MethodDelete).Subrouter()

	putRouter.Use(erdHandler.MiddlewareValidateERD)
	postRouter.Use(erdHandler.MiddlewareValidateERD)

	getRouter.HandleFunc("/erds", erdHandler.GetERDs)
	postRouter.HandleFunc("/erds", erdHandler.AddERD)
	putRouter.HandleFunc("/erds/{id:[0-9]+}", erdHandler.UpdateERD)
	deleteRouter.HandleFunc("/erds/{id:[0-9]+}", erdHandler.DeleteERD)

	// display swagger documentation in a neat way
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docHandler := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", docHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create server with specific parameters. These should be tuned depending on usage of service.
	// i.e. whether this service is client facing or used by other services
	server := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	l.Println("started listening on ", server.Addr)
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
