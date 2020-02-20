package main

import (
	"bitcoin/bitflyer"
	"bitcoin/config"
	"bitcoin/utils"
	"fmt"
	"log"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
