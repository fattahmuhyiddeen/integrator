package model

import "fmt"

type User struct {
	name string
}

func GetAllUsers() {
	// connectDB1()
	// defer disconnectDB1()
	// rows, _ := DB1.Query(`SELECT name FROM users`)

	connectDB2()
	defer disconnectDB2()
	rows, _ := DB2.Query(`SELECT name FROM users`)

	for rows.Next() {
		var row User
		if err := rows.Scan(&row.name); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(row.name)
		}
	}
}
