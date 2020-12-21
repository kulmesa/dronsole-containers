package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/VolantMQ/volantmq/configuration"
	"github.com/VolantMQ/volantmq/server"
	"github.com/VolantMQ/volantmq/transport"
	persistenceMem "gitlab.com/VolantMQ/vlplugin/persistence/mem"
)

func transportStatus(id string, status string) {
	log.Println("Listener status:", id, status)
}
func onDuplicate(s string, b bool) {
	log.Printf("Session duplicate: clientId: %v allowed: %v", s, b)
}

func main() {
	RegisterAuthManagers()

	persist, _ := persistenceMem.Load(nil, nil)
	healthChecks := NewHealthChecks()
	mqttConfig, acceptorConfig := GetConfigs()

	serverConfig := server.Config{
		Health:          healthChecks,
		MQTT:            mqttConfig,
		Acceptor:        acceptorConfig,
		TransportStatus: transportStatus,
		Persistence:     persist,
		OnDuplicate:     onDuplicate,
	}

	srv, err := server.NewServer(serverConfig)
	if err != nil {
		log.Fatalf("Could not create mqtt server: %v", err)
	}

	internalAuthManager := NewInternalAuthManager()
	internalTransportConfig := transport.Config{
		Port:        "8883",
		AuthManager: internalAuthManager,
	}
	err = srv.ListenAndServe(transport.NewConfigTCP(&internalTransportConfig))
	if err != nil {
		log.Fatalf("Could not listen tcp: %v", err)
	}
	/*
		gcpAuthManager:= NewGCPAuthManager()
		gcpTransportConfig := transport.Config{
			Port:        "8884",
			AuthManager: gcpAuthManager,
		}
		err = srv.ListenAndServe(transport.NewConfigTCP(&gcpTransportConfig))
		if err != nil {
			log.Fatalf("Could not listen tcp: %v", err)
		}
	*/

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	log.Printf("Received quit signal: %v", sig.String())
}

func GetConfigs() (configuration.MqttConfig, configuration.AcceptorConfig) {
	mqttConfig := configuration.MqttConfig{}
	mqttConfig.Version = []string{"v3.1.1"}
	mqttConfig.KeepAlive.Force = true
	mqttConfig.KeepAlive.Period = 60
	mqttConfig.Systree.Enabled = true
	mqttConfig.Systree.UpdateInterval = 10
	mqttConfig.Options.ConnectTimeout = 2
	mqttConfig.Options.OfflineQoS0 = true
	mqttConfig.Options.SessionPreempt = false
	mqttConfig.Options.RetainAvailable = true
	mqttConfig.Options.SubsOverlap = false
	mqttConfig.Options.SubsID = false
	mqttConfig.Options.ReceiveMax = 65535
	mqttConfig.Options.MaxPacketSize = 268435455
	mqttConfig.Options.MaxTopicAlias = 65535
	mqttConfig.Options.MaxQoS = 2

	acceptorConfig := configuration.AcceptorConfig{
		MaxIncoming: 1000,
		PreSpawn:    10,
	}

	return mqttConfig, acceptorConfig
}
