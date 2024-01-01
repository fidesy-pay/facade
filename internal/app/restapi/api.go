package restapi

import (
	"encoding/json"
	invoicesservice "github.com/fidesy-pay/facade/internal/pkg/services/invoices-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	"github.com/go-chi/chi"
	"net/http"
)

const APIKey = "api-key"

type Service struct {
	clientsClient   clients_service.ClientsServiceClient
	invoicesService *invoicesservice.Service
}

func New(
	router *chi.Mux,
	clientsClient clients_service.ClientsServiceClient,
	invoicesService *invoicesservice.Service,
) *Service {
	s := &Service{
		clientsClient:   clientsClient,
		invoicesService: invoicesService,
	}

	router.Post("/api/invoice", s.createInvoice)

	return s
}

func (s *Service) createInvoice(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get(APIKey)
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client, err := s.clientsClient.GetClient(r.Context(), &clients_service.GetClientRequest{
		Filter: &clients_service.GetClientRequest_Filter{
			ApiKeyIn: []string{apiKey},
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	invoiceID, err := s.invoicesService.CreateInvoice(r.Context(), client.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	body, _ := json.Marshal(map[string]string{
		"invoice_id": invoiceID,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
