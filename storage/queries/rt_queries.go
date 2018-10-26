package queries

const (
	//GetRtQuery - sql query to get rt from db
	GetRtQuery = `SELECT source, chat_id, username, request,response FROM rtarchive.roundtrips rt WHERE rt.id=?`

	//UpdateResponseQuery - query that update rt response in db
	UpdateResponseQuery = `UPDATE rtarchive.roundtrips SET response = ? WHERE id = ?`

	//SaveRtQuery - query to save roundtrip in db
	SaveRtQuery = `INSERT INTO rtarchive.roundtrips (source, chat_id, username, request, response) VALUES (?, ?, ?, ?, ?);`
)
