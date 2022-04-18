package proxy

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mars-terminal/proxy-cache/internal/entities"
)

func (s *Storage) Get(ctx context.Context, url string) (*entities.URLResponse, error) {
	result := s.cache.FindOne(ctx, bson.M{"url": url})
	if err := result.Err(); err != nil {
		logrus.Info("this url doesnt exist")
		return nil, err
	}

	var data bson.M
	if err := result.Decode(&data); err != nil {
		return nil, err
	}
	if data == nil {
		return nil, mongo.ErrNoDocuments
	}

	bytes := make([]byte, len(data["data"].(primitive.Binary).Data))

	if d, ok := data["data"].(primitive.Binary); ok {
		copy(bytes[:], d.Data[:])
	}

	return &entities.URLResponse{
		Data:         bytes,
		ContentType:  data["content_type"].(string),
		ResponseCode: int(data["response_code"].(int32)),
	}, nil
}
