package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"server/component/config"
	"time"
)

var pool redis.Pool

func init() {
	address := fmt.Sprintf("%s:%s", config.Config.Cache.Ip, config.Config.Cache.Port)
	var err error
	pool = redis.Pool{
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", address)
			if err != nil {
				panic(err)
			}
			return
		},
		TestOnBorrow:    nil,
		MaxIdle:         3,
		MaxActive:       20,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	if err != nil {
		panic(err)
	}
}

type connAdapter struct {
	conn redis.Conn
}

func GetConn() (adapter *connAdapter) {
	connection := pool.Get()
	connection.Do("SELECT", config.Config.Cache.Db)
	adapter = &connAdapter{conn: connection}
	return
}

func GetConnForDb(dbNum int) (adapter *connAdapter) {
	connection := pool.Get()
	connection.Do("SELECT", dbNum)
	adapter = &connAdapter{conn: connection}
	return
}

func (adapter *connAdapter) Expire(key string, duration time.Duration) (err error) {
	defer adapter.conn.Close()
	_, err = adapter.conn.Do("EXPIRE", key, duration.Seconds())
	return
}

func (adapter *connAdapter) Exist(key string) (exist bool, err error) {
	defer adapter.conn.Close()
	exist, err = redis.Bool(adapter.conn.Do("EXISTS", key))
	return
}

func (adapter *connAdapter) Set(key string, value interface{}) (err error) {
	defer adapter.conn.Close()
	_, err = adapter.conn.Do("SET", key, value)
	return
}

func (adapter *connAdapter) SetEx(key string, duration time.Duration, value interface{}) (err error) {
	defer adapter.conn.Close()
	_, err = adapter.conn.Do("SETEX", key, duration.Seconds(), value)
	return
}

func (adapter *connAdapter) Get(key string) (value []byte, err error) {
	defer adapter.conn.Close()
	value, err = redis.Bytes(adapter.conn.Do("GET", key))
	return
}

func (adapter *connAdapter) Del(key string) (err error) {
	defer adapter.conn.Close()
	_, err = adapter.conn.Do("DEL", key)
	return
}
