package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// Initialize the Service Weaver application.
	ctx := context.Background()
	root := weaver.Init(ctx)

	logger := root.Logger()

	// Get a client to the Reverser component.
	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	// Get a client to the Shuffler component.
	shuffler, err := weaver.Get[Shuffler](root)
	if err != nil {
		log.Fatal(err)
	}

	// Get a network listener on address "0.0.0.0:12345".
	opts := weaver.ListenerOptions{LocalAddress: "0.0.0.0:12345"}
	listener, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", listener)

	// Health Check
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "OK")
	})

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")
		fmt.Fprintf(writer, "Hello, %s!\n", name)

		logger.Debug("Hello!!")
	})

	// Serve the /shuffle endpoint.
	http.HandleFunc("/shuffle", func(writer http.ResponseWriter, request *http.Request) {
		shuffled, err := shuffler.Shuffle(request.Context(), request.URL.Query().Get("text"))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(writer, "%s", shuffled)
	})

	// Serve the /reverse endpoint.
	http.HandleFunc("/reverse", func(writer http.ResponseWriter, request *http.Request) {
		reversed, err := reverser.Reverse(request.Context(), request.URL.Query().Get("text"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(writer, "%s", reversed)
	})

	otelHandler := otelhttp.NewHandler(http.DefaultServeMux, "http")
	http.Serve(listener, otelHandler)
}
