package storage

import (
	"context"

	"github.com/TropicalDog17/alert-checker/internal/model"
	"google.golang.org/protobuf/proto"
)

func (s *Storage) GetAlert(ctx context.Context, id string) (*model.Alert, error) {
	alertBytes, err := s.DB.HGet(ctx, "alerts", id).Result()
	if err != nil {
		return &model.Alert{}, err
	}
	alert := &model.Alert{}
	err = proto.Unmarshal([]byte(alertBytes), alert)
	if err != nil {
		return &model.Alert{}, err
	}
	return alert, nil
}

func (s *Storage) GetAlerts(ctx context.Context) ([]*model.Alert, error) {
	alerts, err := s.DB.HVals(ctx, "alerts").Result()
	if err != nil {
		return nil, err
	}

	var result []*model.Alert
	for _, alertBytes := range alerts {
		alert := &model.Alert{}
		err := proto.Unmarshal([]byte(alertBytes), alert)
		if err != nil {
			return nil, err
		}
		result = append(result, alert)
	}

	return result, nil
}

func (s *Storage) DeleteAlert(ctx context.Context, id string) error {
	_, err := s.DB.HDel(ctx, "alerts", id).Result()
	if err != nil {
		return err
	}

	return nil
}

// Close closes the connection to the Redis database.
func (s *Storage) Close() error {
	return s.DB.Close()
}
