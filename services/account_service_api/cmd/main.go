package main

import (
	"fmt"

	"github.com/joaocarvalhodb1/arch-poc/services/account_service_api/app"
	"github.com/joaocarvalhodb1/arch-poc/shared/contracts/protog"
	"github.com/joaocarvalhodb1/arch-poc/shared/grpc"
	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
	"github.com/joaocarvalhodb1/arch-poc/shared/ws"
)

const (
	webPort     = "8080"
	serviceName = "account api"
	gRPCAddress = "5055"
)

func main() {
	log := helpers.NewLoggers(serviceName)
	accountServiceConn, err := grpc.NewgRPCConnection(fmt.Sprintf("0.0.0.0:%s", gRPCAddress))
	if err != nil {
		log.Panic("account service dial error", err)
	}
	log.Debug("Connected to account service gRPC")
	defer accountServiceConn.Close()

	accountServiceClient := protog.NewAccountServiceClient(accountServiceConn)
	accountAPI := app.NewAccountAppAPI(accountServiceClient, log)

	server := ws.NewHttpServer(accountAPI.Routes(), webPort, log)
	err = server.Listen()
	if err != nil {
		log.Panic(err.Error())
	}

}
