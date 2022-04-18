package proxy

import (
	"context"
	"fmt"

	"github.com/mars-terminal/proxy-cache/internal/entities"
)

func (s *Store) GetData(ctx context.Context, url string) (*entities.URLResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("url can't be zero value")
	}
	return s.Store.Get(ctx, url)
}
