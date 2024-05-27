package handler

import (
	"context"
	"fmt"

	"github.com/TropicalDog17/alert-checker/internal/model"
	"github.com/TropicalDog17/alert-checker/internal/storage"
)

type PriceHandler interface {
	FetchPrice(symbol string) (*model.PriceInfo, error)
	FetchPrices(symbols []string) (map[string]*model.PriceInfo, error)
	CheckPriceHandler(alert *model.Alert, priceInfo *model.PriceInfo) (string, bool)
}

func FetchPrice(ctx context.Context, db storage.StorageInterface, symbol string) (*model.PriceInfo, error) {
	price, err := db.GetDB().HGet(ctx, "prices", symbol).Float32()
	if err != nil {
		return nil, err
	}

	change24h, err := db.GetDB().HGet(ctx, "change24h", symbol).Float32()
	if err != nil {
		return nil, err
	}

	change1h, err := db.GetDB().HGet(ctx, "change1h", symbol).Float32()
	if err != nil {
		return nil, err
	}

	change7d, err := db.GetDB().HGet(ctx, "change7d", symbol).Float32()
	if err != nil {
		return nil, err
	}

	high24h, err := db.GetDB().HGet(ctx, "high24h", symbol).Float32()
	if err != nil {
		return nil, err
	}
	low24h, err := db.GetDB().HGet(ctx, "low24h", symbol).Float32()
	if err != nil {
		return nil, err
	}

	return &model.PriceInfo{
		Symbol:    symbol,
		Price:     price,
		Change24h: change24h,
		Change1h:  change1h,
		Change7d:  change7d,
		High24h:   high24h,
		Low24h:    low24h,
	}, nil

}

func FetchPrices(ctx context.Context, db storage.StorageInterface, symbols []string) (map[string]*model.PriceInfo, error) {
	// Fetch the current price of the symbols from the database
	prices := make(map[string]*model.PriceInfo)
	for _, symbol := range symbols {
		price, err := FetchPrice(ctx, db, symbol)
		if err != nil {
			return nil, err
		}
		prices[symbol] = price
	}
	return prices, nil
}

func CheckPriceAbove(alert *model.Alert, priceInfo *model.PriceInfo) bool {
	return priceInfo.Price >= alert.Value
}

func CheckPriceBelow(alert *model.Alert, priceInfo *model.PriceInfo) bool {
	fmt.Println("Price:", priceInfo.Price, "Alert value:", alert.Value)
	return priceInfo.Price <= alert.Value
}

func CheckPricePercentageAbove(alert *model.Alert, priceInfo *model.PriceInfo) bool {
	return alert.Value >= priceInfo.Change24h && priceInfo.Change24h >= 0
}

func CheckPricePercentageBelow(alert *model.Alert, priceInfo *model.PriceInfo) bool {
	return alert.Value <= priceInfo.Change24h && priceInfo.Change24h <= 0
}

// const (
// 	Condition_CONDITION_UNSPECIFIED      Condition = 0
// 	Condition_PRICE_ABOVE                Condition = 1
// 	Condition_PRICE_BELOW                Condition = 2
// 	Condition_PRICE_EQUAL                Condition = 3
// 	Condition_PRICE_PERCENT_CHANGE_ABOVE Condition = 4
// 	Condition_PRICE_PERCENT_CHANGE_BELOW Condition = 5
// )

// // Enum value maps for Condition.
// var (
//
//	Condition_name = map[int32]string{
//		0: "CONDITION_UNSPECIFIED",
//		1: "PRICE_ABOVE",
//		2: "PRICE_BELOW",
//		3: "PRICE_EQUAL",
//		4: "PRICE_PERCENT_CHANGE_ABOVE",
//		5: "PRICE_PERCENT_CHANGE_BELOW",
//	}
func CheckPriceHandler(alert *model.Alert, priceInfo *model.PriceInfo) (bool, string) {
	switch alert.Condition {
	case model.Condition_PRICE_ABOVE:
		return CheckPriceAbove(alert, priceInfo), "price is above the alert value"
	case model.Condition_PRICE_BELOW:
		return CheckPriceBelow(alert, priceInfo), "price is below the alert value"
	case model.Condition_PRICE_PERCENT_CHANGE_ABOVE:
		return CheckPricePercentageAbove(alert, priceInfo), "price percentage change is above the alert value"
	case model.Condition_PRICE_PERCENT_CHANGE_BELOW:
		return CheckPricePercentageBelow(alert, priceInfo), "price percentage change is below the alert value"
	default:
		return false, "condition not supported"
	}
}
