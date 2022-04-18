package handler

import (
	"bytes"
	"net/http"

	"github.com/mars-terminal/proxy-cache/internal/service"
)

type Handler struct {
	Service service.ProxyService
}

func NewHandler(service service.ProxyService) *Handler {
	return &Handler{Service: service}
}

func (h Handler) SetupRouter(r *http.ServeMux) {
	r.HandleFunc("/", h.Middleware(h.handler))
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (w *responseWriter) Write(bytes []byte) (int, error) {
	_, _ = w.body.Write(bytes)
	return w.ResponseWriter.Write(bytes)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
