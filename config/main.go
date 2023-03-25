package config

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type config struct {
	Env         string
	AuthSecret  string
	Port        string
	RedisHost   string
	RedisSecret string
}

// string mapping
const (
	EnvDevelopment string = "development"
	EnvProduction         = "production"
	EnvTesting            = "testing"
)

func Get(getenv func(string) string) (config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}

	env := getenv("ENV")
	if strings.Contains(os.Args[0], "/_test/") || strings.HasSuffix(os.Args[0], ".test") || env == EnvTesting {
		env = EnvTesting
	} else if env == EnvProduction {
		env = EnvProduction
	} else {
		env = EnvDevelopment
	}

	authSecret := getenv("AUTH_SECRET")
	if authSecret == "" {
		rand.Seed(time.Now().Unix())
		authSecret = fmt.Sprintf("%d", rand.Intn(100000000))
	}

	port := getenv("PORT")
	if len(port) == 0 {
		port = "4040"
	}

	redisHost := getenv("REDIS_HOST")
	if env == EnvTesting {
		redisHost = "localhost:6380"
	} else if len(redisHost) == 0 {
		redisHost = "localhost:6379"
	}

	redisSecret := getenv("REDIS_SECRET")
	if env == EnvTesting {
		redisSecret = ""
	} else if len(redisSecret) == 0 {
		redisSecret = "redis"
	}

	return config{
		Env:         env,
		AuthSecret:  authSecret,
		Port:        port,
		RedisHost:   redisHost,
		RedisSecret: redisSecret,
	}, nil
}
