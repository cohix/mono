package mono

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cohix/mono/log"
)

type Server struct {
	server   *http.Server
	services []*Service
	Logger   *log.Logger
}

func NewServer() *Server {
	return &Server{
		server:   &http.Server{},
		services: make([]*Service, 0),
		Logger:   log.New(),
	}
}

func (ms *Server) AddService(service *Service) {
	ms.services = append(ms.services, service)
}

func (ms *Server) ListenAndServe(port int) error {
	serviceName := os.Args[1]
	var service *Service
	for i, s := range ms.services {
		if strings.EqualFold(s.Name, serviceName) {
			service = ms.services[i]
			break
		}
	}

	if service == nil {
		return fmt.Errorf("no service with the name %s has been registered", serviceName)
	}

	ms.Logger.SetServiceName(serviceName)
	ms.Logger.LogInfo(fmt.Sprintf("Starting %s service on port %d", serviceName, service.Port))

	ms.server.Handler = service.Router()
	ms.server.Addr = fmt.Sprintf(":%d", service.Port)
	if err := ms.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
