package conf

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type Config struct {
	Service struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"service"`

	Grpc struct {
		Address string `json:"address"`
	} `json:"grpc"`
	Etcd struct {
		Address string `json:"address"`
	} `json:"etcd"`
	Jaeger struct {
		Address string `json:"address"`
	} `json:"jaeger"`
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Root     string `json:"root"`
		Ip       string `json:"ip"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Ip   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"redis"`
	Log struct {
		FileName   string `json:"file_name"`
		MaxSize    int    `json:"max_size"`
		MaxBackups int    `json:"max_backups"`
		MaxAge     int    `json:"max_age"`
	} `json:"log"`
	LogErr struct {
		FileName   string `json:"file_name"`
		MaxSize    int    `json:"max_size"`
		MaxBackups int    `json:"max_backups"`
		MaxAge     int    `json:"max_age"`
	} `json:"log_err"`
}

func LoadConfig() *Config {

	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	c := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/config.yaml", directory)),
		),
	)
	if err = c.Load(); err != nil {
		panic(err)
	}

	var cf Config

	if err = c.Scan(&cf); err != nil {
		panic(err)
	}
	return &cf

}
