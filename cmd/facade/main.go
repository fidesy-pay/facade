package main

import (
	"context"
	graphqlServerHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/fidesy-pay/facade/internal/app/graph"
	"github.com/fidesy-pay/facade/internal/app/graph/generated"
	"github.com/fidesy-pay/facade/internal/pkg/middleware/auth"
	"github.com/fidesy-pay/facade/internal/pkg/sandbox"
	authservice "github.com/fidesy-pay/facade/internal/pkg/services/auth-service"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
	auth_service "github.com/fidesy-pay/facade/pkg/auth-service"
	invoices_service "github.com/fidesy-pay/facade/pkg/invoices-service"
	"github.com/fidesyx/platform/pkg/scratch"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// services
	authService := authservice.New(authClient)
	invoicesService := invoicesservice.New(invoicesClient)

	resolver := graph.NewResolver(
		authService,
		invoicesService,
	)

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	})

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		//AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		//Debug:            true,
	}).Handler)
	router.Use(auth.Auth(authClient))

	graphqlServer := graphqlServerHandler.NewDefaultServer(schema)
	//graphqlServer.AroundResponses(authMiddleware)

	go func() {
		//router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/", sandbox.ApolloSandboxHandler("GraphQL", "/query"))
		router.Handle("/query", graphqlServer)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()

	<-ctx.Done()
}