package dal

import (
	"fmt"
	"gin_mani_center/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"github.com/gomodule/redigo/redis"
)

var dbProxy *gorm.DB

var redisProxy  *redis.Pool

func InitDB() (err error) {
	cf := conf.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cf.Database.User, cf.Database.Pass, cf.Server.Ip, cf.Database.Port, cf.Database.Name)
	dbProxy, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func GetDBProxy() (*gorm.DB, error) {
	var err error
	if dbProxy == nil {
		err = InitDB()
	}
	return dbProxy, err
}

func GetRedisProxy()*redis.Pool{
	if redisProxy == nil{
		redisProxy = InitRedisPool()
	}
	return redisProxy
}

func InitRedisPool() *redis.Pool{
	server :=  "127.0.0.1:6379"
	pool := &redis.Pool{
		MaxIdle:     3,                    //最大空闲连接数
		IdleTimeout: 240 * time.Second,    //最大空闲连接时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
}
