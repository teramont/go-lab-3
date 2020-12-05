package main

import (
	"fmt"
	"log"
)

type Machine struct {
	id   int32
	name string
}

func main() {
	sql, err := NewDbConnection()
	if err != nil {
		log.Fatalf("Cannot initialize db: %s", err)
	}

	rows, err := sql.Query("SELECT id, name FROM machines LIMIT 200")

	defer rows.Close()

	if err != nil {
		log.Fatalf("Cannot initialize db: %s", err)
	}

	for rows.Next() {
		m := Machine{}
		rows.Scan(&m.id, &m.name)
		fmt.Printf("%d %s\n", m.id, m.name)
	}
}
