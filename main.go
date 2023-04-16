package main

import (
	"github.com/ReygaFitra/app-mahasiswa-api/database"
	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDB()
}