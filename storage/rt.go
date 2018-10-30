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
		rt                                          proto_msg.RoundTripData
	)
	err := dbs.pgSQL.QueryRow(queries.GetRtQuery, id).Scan(&source, &chatID, &username, &request, &response)
	if err != nil {
		log.Error("get db err:", err)
		return nil, err
	}
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
			rt.Source = proto_msg.RoundTripData_telegram
		case "fb":
			rt.Source = proto_msg.RoundTripData_fb
		case "viber":
			rt.Source = proto_msg.RoundTripData_viber
		case "whatsapp":
			rt.Source = proto_msg.RoundTripData_whatsapp
		case "vk":
			rt.Source = proto_msg.RoundTripData_vk
		}
	}
	return &proto_msg.RoundTrip{Id: id, Data: &rt}, nil
}

//SaveRoundTrip - func to save roundtrip in db
func (dbs *DBS) SaveRoundTrip(in *proto_msg.RoundTripData) (*proto_msg.RoundtripID, error) {
	if in == nil {
		return nil, errors.New("empty roundtrip")
	}
	var lastInsertID int64
	err := dbs.pgSQL.QueryRow(queries.SaveRtQuery, in.Source.String(), in.ChatID, in.UserName, in.Request, in.Response).Scan(&lastInsertID)
	if err != nil {
		log.Error("save db err:", err)
		return nil, err
	}
	if lastInsertID == 0 {
		return nil, errors.New("roundtrip save error. ID can't be 0")
	}
	return &proto_msg.RoundtripID{Id: lastInsertID}, nil
}

//AddResponse - add response to roundtrip in db
func (dbs *DBS) AddResponse(id int64, response string) error {
	if response == "" {
		return errors.New("empty response")
	}
	_, err := dbs.pgSQL.Exec(queries.UpdateResponseQuery, response, id)
	if err != nil {
		log.Error("update db err:", err)
		return err
	}
	return nil
}
