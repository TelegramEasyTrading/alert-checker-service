package tests

import (
	"context"
	"fmt"
	"testing"

	rand "math/rand"

	"github.com/TropicalDog17/alert-checker/internal/handler"
	"github.com/TropicalDog17/alert-checker/internal/model"
	"github.com/TropicalDog17/alert-checker/internal/storage"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMain_E2E(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()
	storage := &storage.Storage{
		DB: rdb,
	}
	require.NoError(t, err)

	// Create a new alert
	alert := NewMockAlert("AAPL", 100)
	alertBytes, err := proto.Marshal(alert)
	require.NoError(t, err)
	_, err = rdb.HSet(context.Background(), "alerts", alert.Id, alertBytes).Result()
	require.NoError(t, err)

	// Get the alert
	alertString, err := rdb.HGet(context.Background(), "alerts", alert.Id).Result()
	require.NoError(t, err)
	alert2 := &model.Alert{}
	err = proto.Unmarshal([]byte(alertString), alert2)
	require.NoError(t, err)
	require.Equal(t, alert.Value, alert2.Value)
	alerts := []*model.Alert{alert2}
	priceInfos := map[string]*model.PriceInfo{
		"AAPL": {
			Symbol: "AAPL",
			Price:  120,
		},
		"GOOGL": {
			Symbol: "GOOGL",
			Price:  2000,
		},
	}

	symbols := handler.GetAllSymbols(alerts)
	require.Equal(t, []string{"AAPL"}, symbols)

	handler.AlertTriggering(storage, alert2, priceInfos)

	// Check if the alert was deleted
	alertString, err = rdb.HGet(context.Background(), "alerts", alert.Id).Result()
	require.Error(t, err)
	require.Equal(t, "", alertString)

}

// random id,
func NewMockAlert(symbol string, price float32) *model.Alert {
	return &model.Alert{
		Id:        fmt.Sprintf("%d", rand.Intn(1000)),
		Symbol:    symbol,
		Value:     price,
		Condition: model.Condition_PRICE_ABOVE,
	}
}
