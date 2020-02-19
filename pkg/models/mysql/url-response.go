package mysql

import (
	"database/sql"
	"time"
)

type URLResponseModel struct {
	DB *sql.DB
}


func (m *URLResponseModel) Insert(idUrlObject int64, response *string, duration float64, createdAt time.Time) error {
	stmt := "INSERT INTO responses (id_url, response, duration, created_at) VALUES(?, ?, ?, ?)"

	_, err := m.DB.Exec(stmt, &idUrlObject, &response, &duration, &createdAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *URLResponseModel) Get(id int64) (interface{}, error) {
	stmt := `SELECT response, duration, created_at FROM responses WHERE id_url = ? ORDER BY created_at DESC`

	type ResponseView struct {
		Response   *string
		Duration   float64
		Created_At time.Time
	}
	
	rows, err := m.DB.Query(stmt, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	data := []*ResponseView{}

	for rows.Next() {
		record := &ResponseView{}

		err := rows.Scan(&record.Response, &record.Duration, &record.Created_At)

		if err != nil {
			return nil, err
		}

		data = append(data, record)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return data, nil
}
