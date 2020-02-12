package mysql

import (
	"database/sql"
	"mkucharsky/wpapi/pkg/models"
)

type URLObjectModel struct {
	DB *sql.DB
}

func (m *URLObjectModel) Insert(url string, interval int64) (int64, error) {
	stmt := `INSERT INTO urls (url, interval) VALUES(?,?)`
	result, err := m.DB.Exec(stmt, &url, &interval )

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil


}

func (m *URLObjectModel) Delete(id int64) error {
	stmt := `DELETE FROM urls WHERE id = ?`
	_, err := m.DB.Exec(stmt, &id)

	if err != nil {
		return err
	}

	return nil
	
}

func (m *URLObjectModel) Get(id int64) ([]*models.URLObject, error) {

	stmt := `SELECT * FROM urls WHERE id = ? `
	rows, err := m.DB.Query(stmt, &id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	data := []*models.URLObject{}

	for rows.Next() {
		record := &models.URLObject{}

		err := rows.Scan(&record.ID, &record.URL, &record.Interval)

		if err != nil {
			return nil, err
		}

		data = append(data, record)
	}

	return data, nil
}
