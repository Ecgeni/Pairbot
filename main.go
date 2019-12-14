package main

import (
	"log"
	"os"

	"./src/api"
	"./src/driver"
	"./src/storage"
	bot "./src/telegram-bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	tokenBot := os.Getenv("BOT-TOKEN")

	qm := createQueryManager()
	alfabankSymbols := qm.FindSymbolsByExchangeName("alfabank")
	binanceSymbols := qm.FindSymbolsByExchangeName("binance")

	var a api.GetPairData = api.NewAlfaBank(alfabankSymbols)
	var b api.GetPairData = api.NewBinance(binanceSymbols)
	var pairCharges = []api.GetPairData{a, b}
	exchangeHandler := api.NewExchange(pairCharges)
	telegbot := bot.New(&exchangeHandler)

	telegbot.Process(tokenBot)
}

func createQueryManager() *storage.QueryManager {
	sqlDriver := driver.NewPostgresDriver(os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBHOST"), os.Getenv("DBNAME"))
	connection := storage.NewConnection(&sqlDriver)
	queryManager := storage.NewQueryManager(&connection)

	return &queryManager
}
