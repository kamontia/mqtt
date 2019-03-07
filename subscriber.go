package main

import (
	"fmt"
	"time"

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
	for {
		client.Subscribe("light", 0, f)
		// client.Subscribe("kamontia/irSignal", 0, f)

		time.Sleep(1 * time.Second)
	}
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
