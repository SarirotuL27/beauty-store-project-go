package main

import (
	"github.com/SarirotuL27/beauty-store-project-go/config"
	"github.com/SarirotuL27/beauty-store-project-go/routes"
)

func main() {
	db := config.DB

	r := routes.Router(db)
	r.Run()
}
