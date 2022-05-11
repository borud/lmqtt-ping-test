package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/borud/hwid"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client defaults
const (
	packetChanLen               = 128
	defaultKeepAlive            = 5 * time.Second
	defaultPingTimeout          = 5 * time.Second
	defaultAutoReconnect        = true
	defaultConnectRetryInterval = 5 * time.Second
	defaultWriteTimeout         = 5 * time.Second
	defaultConnectionTimeout    = 5 * time.Second
)

var (
	brokerURL = flag.String("addr", "tcp://:1883", "listen address for MQTT broker")
	clientID  = flag.String("client-id", "", "client ID")
	username  = flag.String("user", "", "username")
	password  = flag.String("pass", "", "password")
)

func main() {
	if *clientID == "" {
		var err error
		*clientID, err = hwid.ID()
		if err != nil {
			lg.Fatal(err)

		}
	}

	clientOpts := mqtt.NewClientOptions().
		AddBroker(*brokerURL).
		SetClientID(*clientID).
		SetUsername(*username).
		SetPassword(*password).
		SetKeepAlive(defaultKeepAlive).
		SetPingTimeout(defaultPingTimeout).
		SetConnectRetryInterval(defaultConnectRetryInterval).
		SetWriteTimeout(defaultWriteTimeout).
		SetAutoReconnect(defaultAutoReconnect).
		SetConnectTimeout(defaultConnectionTimeout)

	// Log when we initiate reconnection
	clientOpts.SetReconnectingHandler(func(client mqtt.Client, options *mqtt.ClientOptions) {
		lg.Infow("ReconnectHandler",
			"brokerURL", brokerURL,
			"client", client,
			"options", options)
	})

	// Log when we have lost the connection
	clientOpts.SetConnectionLostHandler(func(c mqtt.Client, e error) {
		lg.Info()
		lg.Infow("ConnectionLostHandler",
			"client", c,
			"err", e)
	})

	// Log when connection succeeds
	clientOpts.SetOnConnectHandler(func(c mqtt.Client) {
		lg.Infow("OnConnectHandler",
			"client", c)
	})

	client := mqtt.NewClient(clientOpts)
	token := client.Connect()
	res := token.Wait()
	lg.Infow("token.Wait", "res", res)

	err := token.Error()
	if err != nil {
		lg.Infow("connect failed", "err", err)
	}

	for i := 0; ; i++ {
		payload := fmt.Sprintf("payload %d", i)

		token := client.Publish("sometopic/sub", byte(1), true, payload)
		ret := token.Wait()
		//lg.Infow("publish token.Wait", "ret", ret)

		err := token.Error()
		if err != nil {
			lg.Infow("publish token.Error", "ret", ret, "err", err)
		}

		time.Sleep(time.Second)
	}

}
