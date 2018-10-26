package storage

import (
	"database/sql"
	"github.com/kataras/iris/core/errors"
	log "github.com/sirupsen/logrus"
	"rtArchive/proto_msg"
	"rtArchive/storage/queries"
)

//GetRoundTrip - func to get roundtrip from db
func (dbs *DBS) GetRoundTrip(id int64) (*proto_msg.RoundTrip, error) {
	var (
		chatID, username, request, response, source sql.NullString
		rt                                          proto_msg.RoundTrip
	)
	err := dbs.pgSQL.QueryRow(queries.GetRtQuery, id).Scan(&source, &chatID, &username, &request, &response)
	if err != nil {
		return nil, err
	}
	rt.Id = id
	if chatID.Valid {
		rt.ChatID = chatID.String
	}
	if username.Valid {
		rt.UserName = username.String
	}
	if request.Valid {
		rt.Request = request.String
	}
	if response.Valid {
		rt.Response = response.String
	}
	if source.Valid {
		switch source.String {
		case "telegram":
			rt.Source = proto_msg.RoundTrip_telegram
		case "fb":
			rt.Source = proto_msg.RoundTrip_fb
		case "viber":
			rt.Source = proto_msg.RoundTrip_viber
		case "whatsapp":
			rt.Source = proto_msg.RoundTrip_whatsapp
		case "vk":
			rt.Source = proto_msg.RoundTrip_vk
		}
	}
	return &rt, nil
}

//SaveRoundTrip - func to save roundtrip in db
func (dbs *DBS) SaveRoundTrip(in *proto_msg.RoundTripWithoutID) (*proto_msg.RoundTrip, error) {
	if in == nil {
		return nil, errors.New("empty roundtrip")
	}
	var transaction = func(tx *sql.Tx, in *proto_msg.RoundTripWithoutID) (*proto_msg.RoundTrip, error) {
		var lastInsertID int64
		err := tx.QueryRow(queries.SaveRtQuery, in.Source.String(), in.ChatID, in.UserName, in.Request, in.Response).Scan(&lastInsertID)
		if err != nil {
			log.Warn("save db err:", err)
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		rt, err := txGet(tx, lastInsertID)
		if err != nil {
			return nil, err
		}
		return rt, nil
	}
	tx, err := dbs.pgSQL.Begin()
	if err != nil {
		return nil, err
	}
	rt, err := transaction(tx, in)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return nil, e
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return rt, nil
}

//AddResponse - add response to roundtrip in db
func (dbs *DBS) AddResponse(id int64, response string) (*proto_msg.RoundTrip, error) {
	if response == "" {
		return nil, errors.New("empty response")
	}

	var transaction = func(tx *sql.Tx, id int64, response string) (*proto_msg.RoundTrip, error) {
		_, err := tx.Exec(queries.UpdateResponseQuery, response, id)
		if err != nil {
			return nil, err
		}
		rt, err := txGet(tx, id)
		if err != nil {
			return nil, err
		}
		return rt, nil
	}

	tx, err := dbs.pgSQL.Begin()
	if err != nil {
		return nil, err
	}
	rt, err := transaction(tx, id, response)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return nil, e
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return rt, nil
}

//txGet - func to get roundtrip in transaction
func txGet(tx *sql.Tx, id int64) (*proto_msg.RoundTrip, error) {
	var (
		chatID, username, request, response, source sql.NullString
		rt                                          proto_msg.RoundTrip
	)
	err := tx.QueryRow(queries.GetRtQuery, id).Scan(&source, &chatID, &username, &request, &response)
	if err != nil {
		return nil, err
	}
	rt.Id = id
	if chatID.Valid {
		rt.ChatID = chatID.String
	}
	if username.Valid {
		rt.UserName = username.String
	}
	if request.Valid {
		rt.Request = request.String
	}
	if response.Valid {
		rt.Response = response.String
	}
	if source.Valid {
		switch source.String {
		case "telegram":
			rt.Source = proto_msg.RoundTrip_telegram
		case "fb":
			rt.Source = proto_msg.RoundTrip_fb
		case "viber":
			rt.Source = proto_msg.RoundTrip_viber
		case "whatsapp":
			rt.Source = proto_msg.RoundTrip_whatsapp
		case "vk":
			rt.Source = proto_msg.RoundTrip_vk
		}
	}
	return &rt, nil
}
