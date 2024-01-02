package invoicesservice

import (
	"context"
	"fmt"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	desc "github.com/fidesy-pay/facade/pkg/invoices-service"
)

type Service struct {
	invoicesClient desc.InvoicesServiceClient
}

func New(
	invoicesClient desc.InvoicesServiceClient,
) *Service {
	return &Service{
		invoicesClient: invoicesClient,
	}
}

func (s *Service) CreateInvoice(ctx context.Context, clientID string, usdAmount float64) (string, error) {
	invoiceResp, err := s.invoicesClient.CreateInvoice(ctx, &desc.CreateInvoiceRequest{
		ClientId:  clientID,
		UsdAmount: usdAmount,
	})
	if err != nil {
		return "", fmt.Errorf("invoicesClient.CreateInvoice: %w", err)
	}

	return invoiceResp.GetId(), nil
}

func (s *Service) UpdateInvoice(ctx context.Context, input model.UpdateInvoiceInput) (*desc.Invoice, error) {
	invoiceResp, err := s.invoicesClient.UpdateInvoice(ctx, &desc.UpdateInvoiceRequest{
		Id:    input.ID,
		Chain: input.Chain,
		Token: input.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("invoicesClient.UpdateInvoice: %w", err)
	}

	return invoiceResp.GetInvoice(), nil
}

func (s *Service) CheckInvoice(ctx context.Context, input model.CheckInvoiceInput) (*desc.Invoice, error) {
	invoiceResp, err := s.invoicesClient.CheckInvoice(ctx, &desc.CheckInvoiceRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("invoicesClient.CheckInvoice: %w", err)
	}

	return invoiceResp.GetInvoice(), nil
}

type ListInvoicesFilter struct {
	IDIn       []string
	ClientIDIn []string
}

func (s *Service) ListInvoices(ctx context.Context, filter ListInvoicesFilter) ([]*desc.Invoice, error) {
	reqFilter := &desc.ListInvoicesRequest_Filter{}
	if len(filter.IDIn) > 0 {
		reqFilter.IdIn = filter.IDIn
	}

	if len(filter.ClientIDIn) > 0 {
		reqFilter.ClientIdIn = filter.ClientIDIn
	}

	invoicesResp, err := s.invoicesClient.ListInvoices(ctx, &desc.ListInvoicesRequest{
		Filter: reqFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("invoicesClient.ListInvoices: %w", err)
	}

	return invoicesResp.GetInvoices(), nil
}
