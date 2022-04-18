package handler

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) handler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query().Get("q")

	if _, err := url.Parse(uri); len(uri) == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad request"))

		return
	}

	response, err := http.Get(uri)
	if err != nil {
		return
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if err := response.Body.Close(); err != nil {
		return
	}

	responseContentType := response.Header.Get("Content-Type")

	w.WriteHeader(response.StatusCode)
	w.Header().Add("Content-Type", responseContentType)
	w.Header().Add("Content-Length", strconv.FormatInt(int64(len(bytes)), 10))
	_, _ = w.Write(bytes)

	return
}
