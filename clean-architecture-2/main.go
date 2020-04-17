package main

import (
	"github.com/tanimutomo/go-samples/clean-architecture-2/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouter(db)
	r.Run(":8080")
}
