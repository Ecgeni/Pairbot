package driver

import (
	sql "database/sql"
)

type DriverInterface interface {
	Query(string) *sql.Rows
	Exec(string)
}
