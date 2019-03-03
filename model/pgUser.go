package model

import "fmt"

type User struct {
	name string
}

func GetAllUsers() {
	connectPG()
	defer disconnectPG()

	rows, _ := PGDB.Query(`SELECT name FROM users`)

	for rows.Next() {
		var row User
		if err := rows.Scan(&row.name); err != nil {
			// do something with error
		} else {
			fmt.Println(row.name)
		}
	}
}
