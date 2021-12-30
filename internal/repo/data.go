package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"go-kratos/conf"
	mLog "go-kratos/log"
	mGorm "go-kratos/pkg/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Data struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewData(db *gorm.DB, rds *redis.Client) *Data {
	return &Data{
		db:  db,
		rds: rds,
	}
}

func NewMysql(c *conf.Config, log *mLog.Logger) *gorm.DB {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Ip+":"+strconv.Itoa(c.Mysql.Port),
		c.Mysql.Database)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	_ = db.Use(&mGorm.OpentracingPlugin{})

	log.Debug("mysql start successful")

	return db
}

func NewRedis(c *conf.Config, log *mLog.Logger) *redis.Client {

	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Ip, c.Redis.Port),
		Password: "",
	})

	_, err := rds.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	rds.AddHook(redisotel.TracingHook{})

	log.Debug("redis start successful")

	return rds
}
