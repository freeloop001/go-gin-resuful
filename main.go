package main

import (
	db "test/rest/database"
	."test/rest/router"
)

func main() {
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8000")
}
