package main

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"

	"github.com/spf13/viper"
)

type Config struct {
	Name     string `env:"NAME"`
	Database Database
}

type Database struct {
	Pass     string `env:"DATABASE_PASS"`
	PassWord string `env:"DATABASE_PASSWORD"`
}

func main() {
	var c Config
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("pre")
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := env.Parse(&c); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", c)
}

func SetupViper(v *viper.Viper, envPrefix string) {
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
}
