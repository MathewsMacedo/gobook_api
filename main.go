package main

import (
	"fmt"

	"github.com/mathewsmacedo/go_api/config"
	"github.com/mathewsmacedo/go_api/config/environment"
	"github.com/mathewsmacedo/go_api/config/initializers"
	"github.com/mathewsmacedo/go_api/db"
)

func main() {
	environment.SetEnv("GO_ENV", environment.Current())
	fmt.Println("GO_API | Environment: ", environment.Current())
	initializers.RedisInit()
	db.Init()
	config.Routes()
}
