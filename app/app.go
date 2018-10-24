package app

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"rtArchive/config"
	"rtArchive/proto_service"
	"rtArchive/rpc"
	"rtArchive/storage"
)

type (
	//App structure - connects databases with the middleware and handlers of router
	App struct {
		dbs *storage.DBS
		//admin *admingrpc.AdminServiceServer
	}
)

func NewApp() *App {
	return &App{}
}

func (a *App) ConnectDB() error {
	err := a.dbs.Connect()
	if err != nil {
		return err
	}
}

func (a *App) ConnectGRPC() {
	lis, err := net.Listen(config.GRPCNetwork, config.GRPCPort)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	serv := rpc.ProcedureHandler{}
	proto_service.RegisterArchiveServiceServer(s, &serv)
}
