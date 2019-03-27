package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	Connection := GetInfoForConnection()
	opts := MQTT.NewClientOptions().AddBroker(Connection.MqttBroker).SetProtocolVersion(3).SetOnConnectHandler(OnConncectedHandler)

	opts.SetClientID("kamontia")
	opts.SetUsername(Connection.UserName)
	opts.SetPassword(Connection.PassWord)
	client := MQTT.NewClient(opts)

	fmt.Print("Connecting...")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error:", token.Error())
	}

	fmt.Println("Publish...")

	token := client.Publish("ifttt/irSignal", 0, true, "{\"data\":\"abcdefg\",\"ispublic\":true,\"ts\":8552749246946 }")
	token.Wait()
	fmt.Println("Published")

	client.Disconnect(250)

}

var OnConncectedHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("OK")
}
