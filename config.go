package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConnectionInfo struct {
	UserName   string
	PassWord   string
	MqttBroker string
}

func GetInfoForConnection() ConnectionInfo {
	info := ConnectionInfo{}

	err := godotenv.Load("config.ini")
	if err != nil {
		log.Fatal("Error: can't open config.ini")
	}

	info.UserName = os.Getenv("MQTT_USERNAME")
	info.PassWord = os.Getenv("MQTT_PASSWORD")
	info.MqttBroker = os.Getenv("MQTT_BROKER")

	return info
}
