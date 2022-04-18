package proxy

import (
	"context"
	"fmt"
)

func (s *Store) SetData(ctx context.Context, url string, data []byte, contentType string, responseCode int) error {
	if url == "" {
		return fmt.Errorf("url can't be zero value")
	}
	return s.Store.Set(ctx, url, data, contentType, responseCode)
}
