package proxy

import (
	"github.com/mars-terminal/proxy-cache/internal/storage"
)

type Store struct {
	Store storage.ProxyStorage
}

func NewStore(store storage.ProxyStorage) *Store {
	return &Store{Store: store}
}
