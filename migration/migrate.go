package main

import (
	"log"
	"os"

	"../src/driver"
	"../src/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
	qm := createQueryManager()
	up(qm)
}

func createQueryManager() *storage.QueryManager {
	sqlDriver := driver.NewPostgresDriver(os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBHOST"), os.Getenv("DBNAME"))
	connection := storage.NewConnection(&sqlDriver)
	queryManager := storage.NewQueryManager(&connection)

	return &queryManager
}

func up(qm *storage.QueryManager) {
	qm.Exec("CREATE TABLE exchanges(exchange_id SERIAL PRIMARY KEY, name VARCHAR(255));")
	qm.Exec("CREATE TABLE pairs(pair_id SERIAL PRIMARY KEY, symbol VARCHAR(255), exchange_id INTEGER NOT NULL);")
}
