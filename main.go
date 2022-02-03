package main

import (
	"github.com/mathewsmacedo/go_api/config"
	"github.com/mathewsmacedo/go_api/db"
)

func main() {
	db.Init()
	config.Routes()
}
