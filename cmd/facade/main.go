package main

import (
	"context"
	"fmt"
	graphqlServerHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/fidesy-pay/facade/internal/app/graph"
	"github.com/fidesy-pay/facade/internal/app/graph/generated"
	"github.com/fidesy-pay/facade/internal/app/restapi"
	"github.com/fidesy-pay/facade/internal/config"
	"github.com/fidesy-pay/facade/internal/pkg/loaders"
	"github.com/fidesy-pay/facade/internal/pkg/metrics"
	"github.com/fidesy-pay/facade/internal/pkg/middleware/auth"
	"github.com/fidesy-pay/facade/internal/pkg/sandbox"
	authservice "github.com/fidesy-pay/facade/internal/pkg/services/auth-service"
	cryptoservice "github.com/fidesy-pay/facade/internal/pkg/services/crypto-service"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
	auth_service "github.com/fidesy-pay/facade/pkg/auth-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	coingecko_api "github.com/fidesy-pay/facade/pkg/coingecko-api"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
	invoices_service "github.com/fidesy-pay/facade/pkg/invoices-service"
	"github.com/fidesy/sdk/common/grpc"
	"github.com/fidesy/sdk/common/logger"
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

	// just init service
	_, err = grpc.NewServer(
		grpc.WithDomainNameService(ctx, "domain-name-service:10000"),
		grpc.WithGraylog("graylog:5555"),
	)
	if err != nil {
		log.Fatalf("scratch.New: %v", err)
	}

	// clients

	authClient, err := grpc.NewClient[auth_service.AuthServiceClient](
		ctx,
		auth_service.NewAuthServiceClient,
		"rpc:///auth-service",
	)
	if err != nil {
		log.Fatalf("NewAuthClient: %v", err)
	}

	invoicesClient, err := grpc.NewClient[invoices_service.InvoicesServiceClient](
		ctx,
		invoices_service.NewInvoicesServiceClient,
		"rpc:///invoices-service",
	)
	if err != nil {
		log.Fatalf("NewInvoicesClient: %v", err)
	}

	clientsClient, err := grpc.NewClient[clients_service.ClientsServiceClient](
		ctx,
		clients_service.NewClientsServiceClient,
		"rpc:///clients-service",
	)
	if err != nil {
		log.Fatalf("NewClientsClient: %v", err)
	}

	cryptoServiceClient, err := grpc.NewClient[crypto_service.CryptoServiceClient](
		ctx,
		crypto_service.NewCryptoServiceClient,
		"rpc:///crypto-service",
	)
	if err != nil {
		log.Fatalf("NewCryptoClient: %v", err)
	}

	coinGeckoAPIClient, err := grpc.NewClient[coingecko_api.CoinGeckoAPIClient](
		ctx,
		coingecko_api.NewCoinGeckoAPIClient,
		"rpc:///external-api",
	)
	if err != nil {
		log.Fatalf("NewCoinGeckoAPIClient: %v", err)
	}

	// services
	authService := authservice.New(authClient, clientsClient)
	invoicesService := invoicesservice.New(invoicesClient)
	cryptoService := cryptoservice.New(cryptoServiceClient, coinGeckoAPIClient)

	resolver := graph.NewResolver(
		clientsClient,
		cryptoServiceClient,
		authService,
		invoicesService,
		cryptoService,
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

	serverWithDataloader := loaders.Middleware(
		graphqlServer,
		clientsClient,
	)

	go func() {
		//router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/", sandbox.ApolloSandboxHandler("GraphQL", "/query"))
		router.Handle("/query", serverWithDataloader)

		logger.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))

		if err = http.ListenAndServe(":"+port, router); err != nil {
			logger.Fatalf("%w", err)
		}
	}()

	go func() {
		metrics.Init()
		if err = metrics.Run(ctx, os.Getenv("METRICS_PORT")); err != nil {
			logger.Fatalf("metrics.Run: %w", err)
		}
	}()

	<-ctx.Done()
}
