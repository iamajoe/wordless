package main

import (
	"log"
	"os"

	"github.com/iamajoe/wordless/config"
	"github.com/iamajoe/wordless/infrastructure/redis"
	"github.com/iamajoe/wordless/interfaces/server"
)

func main() {
	config, err := config.Get(os.Getenv)
	if err != nil {
		log.Fatal(err)
		return
	}

	repos, err := redis.InitRepos(config.RedisHost, config.RedisSecret)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer repos.Close()

	r := server.GetRouter(repos)
	server.InitServer(":"+config.Port, r)
}
