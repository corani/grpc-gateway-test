package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/corani/grpc-gateway-test/docs/api"
	"github.com/corani/grpc-gateway-test/internal/gateway"
	"github.com/corani/grpc-gateway-test/internal/server"
	"github.com/corani/grpc-gateway-test/static"
)

func main() {
	s := server.NewServer(8989)
	s.Start()

	r := mux.NewRouter()

	gateway.RegisterGateway(r, s, "/v1")
	api.RegisterOpenAPIHandler(r, "/openapi3.json")
	static.RegisterStaticContent(r, "/static/")

	srv := &http.Server{
		Addr:    ":8990",
		Handler: r,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8990")
	log.Fatalln(srv.ListenAndServe())
}
