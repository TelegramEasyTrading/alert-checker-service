package storage

import "context"

func (s *Storage) GetPrice(ctx context.Context, symbol string) (float64, error) {
	price, err := s.DB.HGet(ctx, "prices", symbol).Float64()
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (s *Storage) GetBatchPrices(ctx context.Context, symbols []string) (map[string]float64, error) {
	prices, err := s.DB.HMGet(ctx, "prices", symbols...).Result()
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)
	for i, price := range prices {
		if price == nil {
			result[symbols[i]] = 0
			continue
		}
		result[symbols[i]] = price.(float64)
	}

	return result, nil
}
