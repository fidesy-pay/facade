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

func NewResolver(
	clientsClient clients_service.ClientsServiceClient,
	cryptoServiceClient crypto_service.CryptoServiceClient,
	authService *authservice.Service,
	invoicesService *invoicesservice.Service,
) *Resolver {
	return &Resolver{
		clientsClient:       clientsClient,
		cryptoServiceClient: cryptoServiceClient,
		authService:         authService,
		invoicesService:     invoicesService,
	}
}
