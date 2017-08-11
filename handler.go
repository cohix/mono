package mono

import (
	"fmt"
	"net/http"

	"github.com/cohix/mono/log"
)

type Handler struct {
	Method      string
	Path        string
	handlerFunc http.HandlerFunc
	logger      *log.Logger
}

func InitHandler(method string, path string, handlerFunc http.HandlerFunc, l *log.Logger) *Handler {
	return &Handler{
		Method:      method,
		Path:        path,
		logger:      l,
		handlerFunc: handlerFunc,
	}
}

func (h *Handler) HandlerFunc() http.HandlerFunc {
	return h.logWrappedHandler()
}

func (h *Handler) logWrappedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.LogInfo(fmt.Sprintf("Serving %s", h.Path))

		h.handlerFunc(w, r)
	}
}
