package loaders

import (
	clientsservice "github.com/fidesy-pay/facade/internal/pkg/services/clients-service"
	clients_service "github.com/fidesy-pay/facade/pkg/clients-service"
	"github.com/vikstrous/dataloadgen"
	"time"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	ClientByIDLoader *dataloadgen.Loader[string, *clients_service.Client]
}

func NewLoaders(
	clientsService *clientsservice.Service,
) *Loaders {
	return &Loaders{
		ClientByIDLoader: dataloadgen.NewLoader(ClientsByID(clientsService), dataloadgen.WithWait(time.Millisecond)),
	}
}
