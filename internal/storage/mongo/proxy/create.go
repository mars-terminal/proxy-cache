package proxy

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) Set(ctx context.Context, url string, data []byte, contentType string, responseCode int) error {
	_, err := s.cache.UpdateOne(
		ctx,
		bson.M{"url": url},
		bson.M{"$set": bson.M{
			"data":          data,
			"content_type":  contentType,
			"response_code": responseCode,
		}},
		options.Update().SetUpsert(true),
	)

	return err
}
