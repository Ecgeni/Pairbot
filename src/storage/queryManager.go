package storage

import "fmt"

type QueryManager struct {
	connection DriverConnection
}

func NewQueryManager(Connection DriverConnection) QueryManager {
	queryManager := QueryManager{}
	queryManager.connection = Connection

	return queryManager
}

func (qm *QueryManager) query(sql string) []string {
	rows := qm.connection.Query(sql)

	return rows
}

func (qm *QueryManager) exec(sql string) {
	qm.connection.Exec(sql)
}

func (qm *QueryManager) FindSymbolsByExchangeName(exchangeName string) []string {
	rows := qm.query(fmt.Sprintf("SELECT symbol FROM pairs LEFT JOIN exchanges ON exchanges.exchange_id = pairs.exchange_id WHERE exchanges.name = '%s';", exchangeName))

	return rows
}
