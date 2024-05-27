package main

import (
	"context"
	"fmt"
	"time"

	"github.com/TropicalDog17/alert-checker/internal/handler"
	"github.com/TropicalDog17/alert-checker/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := storage.NewRedisClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	for {
		alerts := handler.AlertLoader(ctx, db)
		fmt.Println("Alerts:", alerts)

		symbols := handler.GetAllSymbols(alerts)
		if len(symbols) == 0 {
			symbols = []string{"btc", "eth"}
		}
		fmt.Println("Symbols:", symbols)
		prices, err := handler.FetchPrices(ctx, db, symbols)
		if err != nil {
			fmt.Println("Error fetching prices:", err)
		}
		for _, alert := range alerts {
			handler.AlertTriggering(db, alert, prices)
		}
		time.Sleep(30 * time.Second)
	}
}
