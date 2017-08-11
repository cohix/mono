package mono

import (
	"github.com/gorilla/mux"
)

type Service struct {
	Name   string
	router *mux.Router
}

func InitService(name string) *Service {
	return &Service{
		Name:   name,
		router: mux.NewRouter(),
	}
}

func (ms *Service) Router() *mux.Router {
	return ms.router
}

func (ms *Service) AddRoute(handler *Handler) {
	ms.router.Methods(handler.Method).Path(handler.Path).HandlerFunc(handler.HandlerFunc())
}
