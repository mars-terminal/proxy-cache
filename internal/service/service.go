package service

import (
	"context"

	"github.com/mars-terminal/proxy-cache/internal/entities"
)

type ProxyService interface {
	GetData(ctx context.Context, url string) (*entities.URLResponse, error)
	SetData(ctx context.Context, url string, data []byte, contentType string, responseCode int) error
}
