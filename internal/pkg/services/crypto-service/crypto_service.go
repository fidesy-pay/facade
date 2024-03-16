package cryptoservice

import (
	"context"
	"fmt"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	coingecko_api "github.com/fidesy-pay/facade/pkg/coingecko-api"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
)

type Service struct {
	cryptoServiceClient crypto_service.CryptoServiceClient
	coinGeckoAPIClient  coingecko_api.CoinGeckoAPIClient
}

func New(
	cryptoServiceClient crypto_service.CryptoServiceClient,
	coinGeckoAPIClient coingecko_api.CoinGeckoAPIClient,
) *Service {
	return &Service{
		cryptoServiceClient: cryptoServiceClient,
		coinGeckoAPIClient:  coinGeckoAPIClient,
	}
}

type GetBalanceParams struct {
	Address string
	Chain   string
	Token   string
}

func (s *Service) GetBalance(ctx context.Context, params GetBalanceParams) (*model.Balance, error) {
	balanceResp, err := s.cryptoServiceClient.GetBalance(ctx, &crypto_service.GetBalanceRequest{
		Address: params.Address,
		Chain:   params.Chain,
		Token:   params.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("cryptoServiceClient.GetBalance: %w", err)
	}

	tokenPrice, err := s.coinGeckoAPIClient.GetPrice(ctx, &coingecko_api.GetPriceRequest{
		Symbol: params.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("coinGeckoAPIClient.GetPrice: %w", err)
	}

	return &model.Balance{
		Token:      params.Token,
		Balance:    balanceResp.GetBalance(),
		UsdBalance: balanceResp.GetBalance() * tokenPrice.GetPriceUsd(),
	}, nil
}
