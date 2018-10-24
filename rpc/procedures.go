package rpc

import (
	"context"
	"google.golang.org/grpc"
	protoMsg "rtArchive/proto_msg"
)

type (
	ProcedureHandler struct {
	}
)

func (h *ProcedureHandler) Get(ctx context.Context, in *protoMsg.IDMessage, opts ...grpc.CallOption) (*protoMsg.RoundTrip, error) {
	return nil, nil
}

func (h *ProcedureHandler) Save(ctx context.Context, in *protoMsg.RoundTrip, opts ...grpc.CallOption) (*protoMsg.RoundTrip, error) {
	return nil, nil
}
