package storage

import (
	"log"

	driver "../driver"
)

type DriverConnection interface {
	Query(string) []string
	Exec(string)
}

type connection struct {
	driver driver.DriverInterface
}

func NewConnection(Driver driver.DriverInterface) connection {
	conn := connection{}
	conn.driver = Driver

	return conn
}

func (c *connection) Query(sql string) []string {
	rows := c.driver.Query(sql)
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return names
}

func (c *connection) Exec(sql string) {
	c.driver.Exec(sql)
}
