package authservice

import (
	"context"
	"fmt"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	auth_service "github.com/fidesy-pay/facade/pkg/auth-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
)

type Service struct {
	authClient    auth_service.AuthServiceClient
	clientsClient clients_service.ClientsServiceClient
}

func New(
	authClient auth_service.AuthServiceClient,
	clientsClient clients_service.ClientsServiceClient,
) *Service {
	return &Service{
		authClient:    authClient,
		clientsClient: clientsClient,
	}
}

func (s *Service) SignUp(ctx context.Context, input model.SignUpInput) (string, error) {
	client, err := s.clientsClient.CreateClient(ctx, &clients_service.CreateClientRequest{
		Username: input.Username,
	})
	if err != nil {
		return "", fmt.Errorf("clientsClient.CreateClient: %w", err)
	}

	signUpResp, err := s.authClient.SignUp(ctx, &auth_service.SignUpRequest{
		Username: input.Username,
		Password: input.Password,
		ClientId: client.Id,
	})
	if err != nil {
		return "", fmt.Errorf("authClient.SignUp: %w", err)
	}

	return signUpResp.Token, nil
}

func (s *Service) Login(ctx context.Context, input model.LoginInput) (string, error) {
	loginResp, err := s.authClient.Login(ctx, &auth_service.LoginRequest{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return "", fmt.Errorf("authClient.Login: %w", err)
	}

	return loginResp.Token, nil
}
