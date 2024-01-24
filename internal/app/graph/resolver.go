package graph

import (
	authservice "github.com/fidesy-pay/facade/internal/pkg/services/auth-service"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	clientsClient       clients_service.ClientsServiceClient
	cryptoServiceClient crypto_service.CryptoServiceClient
	authService         *authservice.Service
	invoicesService     *invoicesservice.Service
}

type ResolverOption func(r *Resolver)

func WithClientsClient(clientsClient clients_service.ClientsServiceClient) ResolverOption {
	return func(r *Resolver) {
		r.clientsClient = clientsClient
	}
}

func WithCryptoServiceClient(cryptoServiceClient crypto_service.CryptoServiceClient) ResolverOption {
	return func(r *Resolver) {
		r.cryptoServiceClient = cryptoServiceClient
	}
}

func WithAuthService(authService *authservice.Service) ResolverOption {
	return func(r *Resolver) {
		r.authService = authService
	}
}

func WithInvoicesService(invoicesService *invoicesservice.Service) ResolverOption {
	return func(r *Resolver) {
		r.invoicesService = invoicesService
	}
}

func NewResolver(options ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range options {
		opt(r)
	}

	return r
}
