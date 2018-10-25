package app

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		rpc *rpc.ProcedureHandler
	}
)

func NewApp() (a *App) {
	a = new(App)
	a.dbs = &storage.DBS{}
	return a
}

func (a *App) ConnectDB() error {
	err := a.dbs.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ConnectGRPC() {
	lis, err := net.Listen(config.GRPCNetwork, config.GRPCPort)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	serv := rpc.ProcedureHandler{}
	proto_service.RegisterArchiveServiceServer(s, &serv)
	reflection.Register(s)
	a.setProcedureHandler(&serv)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

func (a *App) setProcedureHandler(h *rpc.ProcedureHandler) {
	a.rpc = h
}
