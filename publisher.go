package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	ConnectionInfo := GetInfoForConnection()
	opts := MQTT.NewClientOptions().AddBroker(ConnectionInfo.MqttBroker)

	opts.SetClientID("publisher")
	opts.SetUsername(ConnectionInfo.UserName)
	opts.SetPassword(ConnectionInfo.PassWord)
	client := MQTT.NewClient(opts)

	fmt.Println("Connecting...")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error:", token.Error())
	}

	fmt.Println("Publish...")
	token := client.Publish("light", 0, true, "{\"message\": 123}")
	token.Wait()

	client.Disconnect(250)

}
