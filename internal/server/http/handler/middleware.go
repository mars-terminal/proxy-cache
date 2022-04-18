package handler

import (
	"bytes"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) Middleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w = &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           new(bytes.Buffer),
		}

		uri := r.URL.Query().Get("q")

		data, err := h.Service.GetData(r.Context(), uri)
		if err == nil && len(data.Data) > 0 && len(data.ContentType) > 0 {
			w.Header().Set("Content-Type", data.ContentType)
			w.WriteHeader(data.ResponseCode)
			_, _ = w.Write(data.Data)

			return
		}

		next(w, r)

		go func() {
			if len(uri) == 0 {
				return
			}

			err := h.Service.SetData(context.Background(), uri, w.(*responseWriter).body.Bytes(), w.(*responseWriter).Header().Get("Content-Type"), w.(*responseWriter).statusCode)
			if err != nil {
				logrus.WithError(err).Error("failed to set response data")
			}
		}()
	}
}
