package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo"
)

type Config struct {
	Server Server `mapstructure:"server"`
	Other  Other  `mapstructure:"other"`
}

type Server struct {
	Name string `env:"SERVER_NAME" mapstructure:"name"`
	Port string `env:"SERVER_PORT" mapstructure:"port"`
}

type Other struct {
	Text string `env:"TEXT"`
}

func ReadConfig(path string) (*Config, error) {

	config := Config{}

	// ファイルから設定の読み取り
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config error[path: %v]: %w", path, err)
	}

	// ファイルから読み取ったデータを構造体にunmarshal
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarchal error[path: %v]: %w", path, err)
	}

	// 環境変数で上書き (環境変数があった場合のみ上書きされ、環境変数が設定されていない場合はfileから読み取った設定が使われる)
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("parse env value error: %w", err)
	}

	return &config, nil
}

func main() {
	c, err := ReadConfig("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", c)

	//server := echo.New()
	//server.GET("/", c.hello)
	//if err := server.Start(":" + c.Server.Port); err != nil {
	//	log.Fatal(err)
	//}
}

func (cfg *Config) hello(c echo.Context) error {
	return c.String(http.StatusOK, cfg.Other.Text)
}
