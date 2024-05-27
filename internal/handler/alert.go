package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/TropicalDog17/alert-checker/internal/model"
	"github.com/TropicalDog17/alert-checker/internal/storage"
	tele "gopkg.in/telebot.v3"
)

func AlertLoader(ctx context.Context, db storage.StorageInterface) []*model.Alert {

	alerts, err := db.GetAlerts(ctx)
	if err != nil {
		fmt.Println("Error getting alerts:", err)
	}

	return alerts
}

func AlertTriggering(db storage.StorageInterface, alert *model.Alert, priceInfos map[string]*model.PriceInfo) {
	symbol := alert.Symbol
	priceInfo, ok := priceInfos[symbol]

	if !ok {
		fmt.Println("Price not found for", symbol)
		return
	}

	true, message := CheckPriceHandler(alert, priceInfo)
	if true {
		go NotifyUser(alert, priceInfos)
		err := db.DeleteAlert(context.Background(), alert.Id)
		if err != nil {
			fmt.Println("Error deleting alert:", err)
		}
		fmt.Println(message)
		fmt.Println("Trigger for user id", alert.UserId)
	}
}

// Get unique symbols from the alerts
func GetAllSymbols(alerts []*model.Alert) []string {
	symbols := make(map[string]bool)
	for _, alert := range alerts {
		symbols[alert.Symbol] = true
	}

	var result []string
	for symbol := range symbols {
		result = append(result, symbol)
	}

	return result
}

func convertSymbol(symbol string) string {
	switch symbol {
	case "btc":
		return "bitcoin"
	case "eth":
		return "ethereum"
	case "inj":
		return "injective-protocol"
	case "atom":
		return "cosmos"
	case "wif":
		return "dogwifcoin"
	case "doge":
		return "dogecoin"
	default:
		return symbol
	}
}

type User struct {
	UserId string
}

func (u *User) Recipient() string {
	return u.UserId
}

func NotifyUser(alert *model.Alert, priceInfos map[string]*model.PriceInfo) {
	// Send a notification to the user, given user id and telegram bot token
	token := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	user := &User{UserId: alert.UserId}
	msg := FormatNotification(alert, priceInfos)
	_, err = bot.Send(user, msg, tele.ModeMarkdown)
	if err != nil {
		log.Fatal(err)
	}

}

// 🚨🚨 **TRUMPE is up 149% within the last 6 hours** 🚨🚨

// Current price: **$0.02149**

// Price change: **1H:** -3.09%  **24H:** 7.83%  **7D:** 149%
// Trump Pepe - TRUMPE/WETH

// Ethereum / Uniswap

func FormatNotification(alert *model.Alert, priceInfos map[string]*model.PriceInfo) string {
	headingText := ""
	switch alert.Condition {
	case model.Condition_PRICE_ABOVE:
		headingText = fmt.Sprintf("**🚨🚨%s is above %.2f 🚨🚨**", strings.ToUpper(alert.Symbol), alert.Value)
	case model.Condition_PRICE_BELOW:
		headingText = fmt.Sprintf("🚨🚨 **%s is below %.2f** 🚨🚨", strings.ToUpper(alert.Symbol), alert.Value)
	case model.Condition_PRICE_PERCENT_CHANGE_ABOVE:
		headingText = fmt.Sprintf("🚨🚨 **%s is up %.2f%% within the last 6 hours** 🚨🚨", strings.ToUpper(alert.Symbol), alert.Value)
	case model.Condition_PRICE_PERCENT_CHANGE_BELOW:
		headingText = fmt.Sprintf("🚨🚨 **%s is down %.2f%% within the last 6 hours** 🚨🚨", strings.ToUpper(alert.Symbol), alert.Value)
	}

	fmt.Println("Price info:", alert.Symbol)
	priceText := fmt.Sprintf("Current price: *$%.5f*\n", priceInfos[alert.Symbol].Price)

	priceChange := fmt.Sprintf("Price change: *1H:* %.2f%%  *24H:* %.2f%%  *7D:* %.2f%%  \n", priceInfos[alert.Symbol].Change1h, priceInfos[alert.Symbol].Change24h, priceInfos[alert.Symbol].Change7d)

	symbolText := fmt.Sprintf("*%s - %s*\n", strings.ToTitle(alert.Symbol), strings.ToUpper(alert.Symbol))

	return fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", headingText, priceText, priceChange, symbolText)

}
