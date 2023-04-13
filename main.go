package main

import (
	"app-mahasiswa-api/database"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDB()
}