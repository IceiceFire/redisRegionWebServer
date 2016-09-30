package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var Regionredis *redisClient

type redisClient struct {
	pool *redis.Pool
}

//获取所有 key
func (c *redisClient) Getallkey() {
	beego.Info("Getallkey ............................Close")
	conn := c.pool.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"))

	if err != nil {
		panic(err)
	}
	fmt.Printf("keys : %d \n", len(keys))
	beego.Info(len(keys))
	for no, key := range keys {
		fmt.Println("No:", no, "--keys:", key)
	}
}

//设定key [ string ]
func (r *redisClient) Setkey(key string, val string) {
	// 从池里获取连接
	conn := r.pool.Get()
	// 用完后将连接放回连接池
	defer conn.Close()
	client, err := redis.String(conn.Do("SET", key, val))
	if err != nil {
		beego.Error(err)
		panic(err)
	}

	beego.Info(val)
	fmt.Println(client)
}

//获取key的值 [ string ]
func (r *redisClient) Getkey(key string) string {
	conn := r.pool.Get()
	defer conn.Close()
	client, err := redis.String(conn.Do("GET", key))
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	beego.Info(client)
	return client
}

//关闭redis连接池
func (c *redisClient) Close() {
	beego.Info("RedisClient ............................Close")
	if c.pool != nil {
		c.pool.Close()
	}
}

//获取redis链接
func newClient() *redisClient {
	return &redisClient{
		pool: newPool(beego.AppConfig.String("redisserver"), beego.AppConfig.String("redispassword")),
	}
}

//创建redis connection pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		// 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxIdle: 3,
		// 最大的激活连接数，表示同时最多有N个连接
		MaxActive: 100,
		// 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 240 * time.Second,
		// 建立连接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//初期化 init
func init() {
	beego.Info("Redis Connection ..................................................")
	Regionredis = newClient()
}
