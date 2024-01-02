package main

import (
	"context"
	graphqlServerHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/fidesy-pay/facade/internal/app/graph"
	"github.com/fidesy-pay/facade/internal/app/graph/generated"
	"github.com/fidesy-pay/facade/internal/app/restapi"
	"github.com/fidesy-pay/facade/internal/config"
	"github.com/fidesy-pay/facade/internal/pkg/middleware/auth"
	"github.com/fidesy-pay/facade/internal/pkg/sandbox"
	authservice "github.com/fidesy-pay/facade/internal/pkg/services/auth-service"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
	auth_service "github.com/fidesy-pay/facade/pkg/auth-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
	invoices_service "github.com/fidesy-pay/facade/pkg/invoices-service"
	"github.com/fidesyx/platform/pkg/scratch"
	"github.com/go-chi/chi"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
	err    error
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT ENV is required")
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	defer cancel()

	tracer, closer, err = config.NewJaegerTracer()
	if err != nil {
		log.Fatalf("config.NewJaegerTracer: %v", err)
	}
	defer closer.Close()
	auth.Tracer = tracer

	// just init scratch
	_, err := scratch.New(ctx)
	if err != nil {
		log.Fatalf("scratch.New: %v", err)
	}

	// clients
	authClient, err := scratch.NewClient[auth_service.AuthServiceClient](
		ctx,
		auth_service.NewAuthServiceClient,
		"auth-service",
	)
	if err != nil {
		log.Fatalf("NewAuthClient: %v", err)
	}

	invoicesClient, err := scratch.NewClient[invoices_service.InvoicesServiceClient](
		ctx,
		invoices_service.NewInvoicesServiceClient,
		"invoices-service",
	)
	if err != nil {
		log.Fatalf("NewInvoicesClient: %v", err)
	}

	clientsClient, err := scratch.NewClient[clients_service.ClientsServiceClient](
		ctx,
		clients_service.NewClientsServiceClient,
		"clients-service",
	)
	if err != nil {
		log.Fatalf("NewClientsClient: %v", err)
	}

	cryptoServiceClient, err := scratch.NewClient[crypto_service.CryptoServiceClient](
		ctx,
		crypto_service.NewCryptoServiceClient,
		"crypto-service",
	)
	if err != nil {
		log.Fatalf("NewCryptoClient: %v", err)
	}

	// services
	authService := authservice.New(authClient, clientsClient)
	invoicesService := invoicesservice.New(invoicesClient)

	resolver := graph.NewResolver(
		clientsClient,
		cryptoServiceClient,
		authService,
		invoicesService,
	)

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	})
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		//Debug:            true,
	}).Handler)
	router.Use(auth.Auth(authClient))

	restapi.New(router, clientsClient, invoicesService)

	graphqlServer := graphqlServerHandler.NewDefaultServer(schema)
	//graphqlServer.AroundResponses(tracing.Tracing)

	go func() {
		//router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/", sandbox.ApolloSandboxHandler("GraphQL", "/query"))
		router.Handle("/query", graphqlServer)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()

	<-ctx.Done()
}
