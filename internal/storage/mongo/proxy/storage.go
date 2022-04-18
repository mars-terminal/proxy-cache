package proxy

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	cache *mongo.Collection
}

func NewStorage(cache *mongo.Collection) *Storage {
	return &Storage{cache: cache}
}
