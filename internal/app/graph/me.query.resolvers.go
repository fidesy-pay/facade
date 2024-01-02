package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/fidesy-pay/facade/internal/pkg/middleware/auth"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*clients_service.Client, error) {
	username := ctx.Value(auth.UsernameCtxName)

	client, err := r.clientsClient.GetClient(ctx, &clients_service.GetClientRequest{
		Filter: &clients_service.GetClientRequest_Filter{
			UsernameIn: []string{username.(string)},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("clientsClient.GetClient: %w", err)
	}

	return client, nil
}