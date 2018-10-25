package rpc

import (
	"context"
	protoMsg "rtArchive/proto_msg"
)

type (
	ProcedureHandler struct {
	}
)

func (h *ProcedureHandler) Save(ctx context.Context, in *protoMsg.RoundTrip) (*protoMsg.RoundTrip, error) {
	return nil, nil
}

func (h *ProcedureHandler) AddResponse(ctx context.Context, in *protoMsg.IDAndResponse) (*protoMsg.RoundTrip, error) {

}

func (h *ProcedureHandler) Get(ctx context.Context, in *protoMsg.IDMessage) (*protoMsg.RoundTrip, error) {
	return nil, nil
}
