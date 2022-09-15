package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

type EsList struct {
	Es map[string]EsConfig `mapstructure:"es"`
}

type EsConfig struct {
	Host string `mapstructure:"host"`
}

func main() {
	viper.SetConfigFile("./a.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	var conf = new(EsList)
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
}
