package queries

const (
	//GetRtQuery - sql query to get rt from db
	GetRtQuery = `SELECT source, chat_id, username, request,response FROM rtarchive.roundtrips rt WHERE rt.id=$1`

	//UpdateResponseQuery - query that update rt response in db
	UpdateResponseQuery = `UPDATE rtarchive.roundtrips SET response = $1 WHERE id = $2`

	//SaveRtQuery - query to save roundtrip in db
	SaveRtQuery = `INSERT INTO rtarchive.roundtrips(source, chat_id, username, request, response) VALUES($1, $2, $3, $4, $5) RETURNING id;`
)
