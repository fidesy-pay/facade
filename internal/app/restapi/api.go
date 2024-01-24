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
	router.Get("/api/invoice/{invoice_id}", s.getInvoice)

	return s
}

type createInvoiceBody struct {
	USDAmount float64 `json:"usd_amount"`
}

func (s *Service) createInvoice(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get(APIKey)
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var invoiceBody *createInvoiceBody
	err := json.NewDecoder(r.Body).Decode(&invoiceBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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

	invoiceID, err := s.invoicesService.CreateInvoice(r.Context(), client.Id, invoiceBody.USDAmount)
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

func (s *Service) getInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID := chi.URLParam(r, "invoice_id")
	if invoiceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	invoices, err := s.invoicesService.ListInvoices(r.Context(), invoicesservice.ListInvoicesFilter{
		IDIn:       []string{invoiceID},
		ClientIDIn: []string{client.Id},
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	body, _ := json.Marshal(invoices)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
