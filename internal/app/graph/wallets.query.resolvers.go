package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/fidesy-pay/facade/internal/pkg/middleware/auth"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
)

// Wallets is the resolver for the wallets field.
func (r *queryResolver) Wallets(ctx context.Context, filter *model.WalletsFilter) (*model.WalletsPagination, error) {
	session := auth.GetSession(ctx)

	walletsResp, err := r.cryptoServiceClient.ListWallets(ctx, &crypto_service.ListWalletsRequest{
		Filter: &crypto_service.ListWalletsRequest_Filter{
			ClientIdIn: []string{session.ClientID},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cryptoServiceClient.ListWallets: %w", err)
	}

	return &model.WalletsPagination{
		Items: walletsResp.GetWallets(),
	}, nil
}
