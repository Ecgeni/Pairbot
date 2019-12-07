package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type postgresDriver struct {
	connectionParams string
}

func NewPostgresDriver(Dbuser string, Dbpassword string, Dbhost string, Dbname string) postgresDriver {
	driver := postgresDriver{}
	driver.connectionParams = fmt.Sprintf("postgres://%s:%s@%s/%s", Dbuser, Dbpassword, Dbhost, Dbname)

	return driver
}

func (d *postgresDriver) connect() *sql.DB {
	db, err := sql.Open("postgres", d.connectionParams)
	if err != nil {
		panic(err)
	}

	return db
}

func (d *postgresDriver) Query(sql string) *sql.Rows {
	db := d.connect()
	defer db.Close()

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	return rows
}

func (d *postgresDriver) Exec(sql string) {
	db := d.connect()
	defer db.Close()

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
