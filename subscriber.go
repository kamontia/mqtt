package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type ReceiveData struct {
	Data []struct {
		Room   string `json:"room"`
		Device string `json:"device"`
		Action string `json:"action"`
	} `json:"data"`
	Ispublic bool  `json:"ispublic"`
	Ts       int64 `json:"ts"`
}

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger:", log.Lshortfile|log.Ldate|log.Ltime)
	logger.SetFlags(5)
	ConnectionInfo := GetInfoForConnection()
	opts := MQTT.NewClientOptions().AddBroker(ConnectionInfo.MqttBroker).SetProtocolVersion(3).SetOnConnectHandler(OnConncectedHandler)

	opts.SetClientID("kamontia")
	opts.SetUsername(ConnectionInfo.UserName)
	// opts.SetPassword(ConnectionInfo.PassWord)
	logger.Print("Username: ", ConnectionInfo.UserName)

	client := MQTT.NewClient(opts)
	fmt.Print("Connecting...")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error:", token.Error())
	}
	client.Subscribe("ifttt/irSignal", 0, f).Wait()
	for {
		// client.Subscribe("kamontia/irSignal", 0, f)

		time.Sleep(1 * time.Nanosecond)
	}
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	// fmt.Printf("MSG: %T %s\n", msg.Payload(), msg.Payload())

	var tmp ReceiveData
	if err := json.Unmarshal([]byte(msg.Payload()), &tmp); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Room->", tmp.Data[0].Room)
		fmt.Println("Device->", tmp.Data[0].Device)
		fmt.Println("Action->", tmp.Data[0].Action)
	}
}

var OnConncectedHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("OK")
}
