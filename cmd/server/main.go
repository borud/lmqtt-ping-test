package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
	_ "github.com/lab5e/lmqtt/pkg/persistence"     // required for side-effects
	_ "github.com/lab5e/lmqtt/pkg/topicalias/fifo" // required for side-effects
)

var listenAddr = flag.String("a", ":1883", "listen address for MQTT broker")

func main() {
	flag.Parse()

	// Default config is OK for the most part.
	mqttConfig := config.DefaultConfig()
	mqttConfig.MQTT.AllowZeroLenClientID = true

	listenSocket, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	broker := lmqtt.New(
		lmqtt.WithTCPListener(listenSocket),
		lmqtt.WithConfig(mqttConfig),
		lmqtt.WithHook(lmqtt.Hooks{
			OnSubscribe:         onSubscribe,
			OnMsgArrived:        onMsgArrived,
			OnBasicAuth:         onBasicAuth,
			OnSessionCreated:    onSessionCreated,
			OnSessionTerminated: onSessionTerminated,
		}))

	lg.Infow("started server", "listenAddr", listenSocket.Addr().String())
	err = broker.Run()
	log.Printf("%v\n", err)
}

func onSessionCreated(ctx context.Context, client lmqtt.Client) {
	lg.Infow("onSessionCreated",
		"clientID", client.ClientOptions().ClientID,
		"remoteAddr", client.Connection().RemoteAddr().String(),
		"version", client.Version())
}
func onSessionTerminated(ctx context.Context, clientID string, reason lmqtt.SessionTerminatedReason) {
	lg.Infow("onSessionTerminated",
		"clientID", clientID,
		"reason", reason)
}
func onSubscribe(ctx context.Context, client lmqtt.Client, req *lmqtt.SubscribeRequest) error {
	lg.Infow("onSubscribe",
		"clientID", client.ClientOptions().ClientID,
		"remoteAddr", client.Connection().RemoteAddr().String(),
		"topics", req.Subscriptions,
	)
	return nil
}
func onMsgArrived(ctx context.Context, client lmqtt.Client, req *lmqtt.MsgArrivedRequest) error {
	lg.Infow("onMsgArrived",
		"clientID", client.ClientOptions().ClientID,
		"remoteAddr", client.Connection().RemoteAddr().String(),
		"topic", string(req.Publish.TopicName),
		"message", string(req.Message.Payload))
	return nil
}
func onBasicAuth(ctx context.Context, client lmqtt.Client, req *lmqtt.ConnectRequest) error {
	lg.Infow("onBasicAuth",
		"clientID", client.ClientOptions().ClientID,
		"remoteAddr", client.Connection().RemoteAddr().String(),
		"password", string(req.Connect.Password))
	return nil
}
