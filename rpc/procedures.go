package rpc

import (
	"context"
	protoMsg "rtArchive/proto_msg"
	"rtArchive/storage"
)

type (
	ProcedureHandler struct {
		dbs *storage.DBS
	}
)

func (h *ProcedureHandler) Dbs(dbs *storage.DBS) {
	h.dbs = dbs
}

func (h *ProcedureHandler) Save(ctx context.Context, in *protoMsg.RoundTripWithoutID) (*protoMsg.RoundTrip, error) {
	rt, err := h.dbs.SaveRoundTrip(in)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (h *ProcedureHandler) AddResponse(ctx context.Context, in *protoMsg.IDAndResponse) (*protoMsg.RoundTrip, error) {
	rt, err := h.dbs.AddResponse(in.Id, in.Response)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (h *ProcedureHandler) Get(ctx context.Context, in *protoMsg.RoundtripID) (*protoMsg.RoundTrip, error) {
	rt, err := h.dbs.GetRoundTrip(in.Id)
	if err != nil {
		return nil, err
	}
	return rt, nil
}
