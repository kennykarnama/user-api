package redisconn

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kelseyhightower/envconfig"
	"strings"
	"time"
)

var config = make(map[string]Config)

// Connect to connect to redis and return connection pooling of redis
func (cfg Config) Connect() *redis.Pool {
	// construct URL from host and port
	URL := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	// create redis pooling
	Redis := &redis.Pool{
		MaxIdle:         cfg.MaxIdle,
		MaxActive:       cfg.MaxActive,
		Wait:            cfg.Wait,
		MaxConnLifetime: time.Duration(cfg.MaxConnLifetime) * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", URL)
			if err != nil {
				return nil, err
			}
			if len(cfg.Password) > 0 {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}
	// test initialized connection pooling using PING command
	conn := Redis.Get()
	_, err := conn.Do("PING")
	if err != nil {
		panic(fmt.Sprintf("[ERR] Redis connection failed, %s", err))
	}
	defer conn.Close()
	return Redis
}

// GetConfig to get config incase config needed by service
func GetConfig(name string) Config {
	return config[name]
}

// Init to init config using default environment variable
// use <NAME>_REDIS_<ITEMS> convention to construct envar key
func Init(name string) *redis.Pool {
	if name == "" {
		panic("redis name could not be empty")
	}
	name = strings.ToUpper(name)
	var cfg Config
	envconfig.MustProcess(name, &cfg)
	return cfg.Connect()
}

// Config is configuration of redis db connection
// use <DB NAME>_REDIS_<ITEMS> convention to construct envar key
type Config struct {
	Name string `envconfig:"REDIS_NAME" default:""`
	// Host is host of redis
	Host string `envconfig:"REDIS_HOST" default:""`
	// Post of redis service
	Port int `envconfig:"REDIS_PORT" default:"6379"`
	// configured password of redis server
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	// MaxIdle is max idle connection of redis pooling
	MaxIdle int `envconfig:"REDIS_MAX_IDLE" default:"10"`
	// MaxActive is max active connection of redis pooling
	MaxActive int `envconfig:"REDIS_MAX_ACTIVE" default:"320"`
	// MaxIdelTimeout is max idle timeout of worker in pooling
	MaxIdleTimeOut int `envconfig:"REDIS_MAX_IDLE_TIMEOUT" default:"5"`
	// MaxConnectionLifetime is max life time connection in pool
	MaxConnLifetime int `envconfig:"REDIS_MAX_CONN_LIFETIME" default:"10"`
	// Wait to disable/enable redis using only connection from pooling
	Wait bool `envconfig:"REDIS_WAIT" default:"true"`
}
