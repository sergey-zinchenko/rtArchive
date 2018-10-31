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

func (h *ProcedureHandler) SaveInDB(ctx context.Context, in *protoMsg.RoundTripData) (*protoMsg.RoundtripID, error) {
	id, err := h.dbs.SaveRoundTrip(in)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (h *ProcedureHandler) AddResponseToDBEntry(ctx context.Context, in *protoMsg.IDAndResponse) (*protoMsg.Void, error) {
	err := h.dbs.AddResponse(in.Id, in.Response)
	if err != nil {
		return nil, err
	}
	return &protoMsg.Void{}, nil
}

func (h *ProcedureHandler) GetRTFromDB(ctx context.Context, in *protoMsg.RoundtripID) (*protoMsg.RoundTrip, error) {
	rt, err := h.dbs.GetRoundTrip(in.Id)
	if err != nil {
		return nil, err
	}
	return rt, nil
}
