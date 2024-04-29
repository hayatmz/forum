package model

// Prepare the query which returns a stmt.
//
// With the stmt, execute the request with the args which return a sql result and with this one,
// return the id of this sql Result.
func execQuery(query string, args ...any) (int64, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	
	idRes, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return idRes, nil
}