package main

import (
	db "abcpirntf/go-crud-api/db"
	router "abcpirntf/go-crud-api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
