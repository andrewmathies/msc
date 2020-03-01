package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrewmathies/msl/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	//env.Parse()

	l := log.New(os.Stdout, "erd-api", log.LstdFlags)

	// building the handler
	erdHandler := handlers.NewERDs(l)

	// create a serve multiplexer and register handlers
	mux := http.NewServeMux()
	mux.Handle("/", erdHandler)

	// build server with specific parameters. These can be tuned depending on usage of service.
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
