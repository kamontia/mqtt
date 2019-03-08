package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	Connection := GetInfoForConnection()
	opts := MQTT.NewClientOptions().AddBroker(Connection.MqttBroker)

	opts.SetClientID("publisher")
	opts.SetUsername(Connection.UserName)
	opts.SetPassword(Connection.PassWord)
	client := MQTT.NewClient(opts)

	fmt.Println("Connecting...")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error:", token.Error())
	}

	fmt.Println("Publish...")
	token := client.Publish("light", 0, true, "{\"message\": HelloWorld}")
	token.Wait()

	client.Disconnect(250)

}
