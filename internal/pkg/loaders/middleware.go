package loaders

import (
	"context"
	clientsservice "github.com/fidesy-pay/facade/internal/pkg/services/clients-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	"net/http"
)

// Middleware injects data loaders into the context
func Middleware(
	next http.Handler,
	clientsServiceClient clients_service.ClientsServiceClient,
) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(
			clientsservice.New(clientsServiceClient),
		)

		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
