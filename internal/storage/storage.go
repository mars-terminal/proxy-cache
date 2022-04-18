package storage

import (
	"context"

	"github.com/mars-terminal/proxy-cache/internal/entities"
)

type ProxyStorage interface {
	Get(ctx context.Context, url string) (*entities.URLResponse, error)
	Set(ctx context.Context, url string, data []byte, contentType string, responseCode int) error
}
