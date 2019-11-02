package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/spf13/viper"
)

func main() {
	c, err := ReadConfig("./config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", c)
}

type Config struct {
	Server Server `mapstructure:"server"`
	Other  Other  `mapstructure:"other"`
}

type Server struct {
	Name string `env:"SERVER_NAME" mapstructure:"name"`
	Port string `env:"SERVER_NAME" mapstructure:"port"`
}

type Other struct {
	Text string `env:"OTHER_TEXT" mapstructure:"port"`
}

func ReadConfig(path string) (Config, error) {

	config := Config{}

	// ファイルから設定の読み取り
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("read config error[path: %v]: %w", path, err)
	}

	// ファイルから読み取ったデータを構造体にunmarshal
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("unmarchal error[path: %v]: %w", path, err)
	}

	// 環境変数で上書き、環境変数があった場合のみ上書きされる
	if err := env.Parse(&config); err != nil {
		return Config{}, fmt.Errorf("parse env value error: %w", err)
	}

	return config, nil
}
