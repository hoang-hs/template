package main

import (
	"base/src/common/configs"
	"base/src/common/log"
	"base/src/core/constant"
	"flag"
	"fmt"
)

func init() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "configs/config.yaml", "path to config file")
	flag.Parse()
	err := configs.LoadConfig(pathConfig)
	if err != nil {
		panic(err)
	}
	if !constant.IsProdEnv() {
		fmt.Println(configs.Get())
	}
	log.NewLogger()
}

func main() {
	fmt.Println("has")
}
