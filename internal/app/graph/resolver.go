package graph

import (
	authservice "github.com/fidesy-pay/facade/internal/pkg/services/auth-service"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService     *authservice.Service
	invoicesService *invoicesservice.Service
}

func NewResolver(
	authService *authservice.Service,
	invoicesService *invoicesservice.Service,
) *Resolver {
	return &Resolver{
		authService:     authService,
		invoicesService: invoicesService,
	}
}
