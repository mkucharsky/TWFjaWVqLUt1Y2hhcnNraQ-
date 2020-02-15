package mysql

import (
	"database/sql"
	"mkucharsky/wpapi/pkg/models"
	"time"
)

type URLResponseModel struct {
	DB *sql.DB
}

func (m *URLResponseModel) Insert(idUrlObject int64, response string, duration float64, createdAt time.Time) error { 
	stmt := "INSERT INTO responses (id-urlobject, response, duration, created_at) VALUES(?, ?, ?, ?)"

	_, err := m.DB.Exec(stmt, &idUrlObject, &response, &duration, &createdAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *URLResponseModel) Get(id int64) ([]*models.URLResponse, error) {
	stmt := `SELECT * FROM responses WHERE id-urlobject = ? ORDER BY created_at DESC`

	rows, err := m.DB.Query(stmt, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	data := []*models.URLResponse{}

	for rows.Next() {
		record := &models.URLResponse{}

		err := rows.Scan(&record.ID, &record.IDUrlObject, &record.Response, &record.Duration, &record.Created)

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
