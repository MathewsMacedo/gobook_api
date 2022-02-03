package initializers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/mathewsmacedo/go_api/config/environment"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func RedisInit() {
	viper.SetConfigName("redis")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	env := environment.Current()

	db, _ := strconv.Atoi(viper.GetString(env + ".database"))

	addr := viper.GetString(env+".host") + ":" + viper.GetString(env+".port")

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString(env + ".password"),
		DB:       db,
	})

	if _, err := rdb.Ping().Result(); err != nil {
		fmt.Println("Error Initialize Redis: ", err)
	} else {
		fmt.Println("Redis Running: ", addr)
	}

}

func RedisConnection() *redis.Client {
	return rdb
}
