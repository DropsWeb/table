package main

import (
	"github.com/DropsWeb/table/internal/database"
	"github.com/DropsWeb/table/internal/server"
)

func main() {
	database.Db()
	server.Server()
}
