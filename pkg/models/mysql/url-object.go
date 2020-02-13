package mysql

import (
	"database/sql"
	"errors"
	"mkucharsky/wpapi/pkg/models"
)

type URLObjectModel struct {
	DB *sql.DB
}

func (m *URLObjectModel) Insert(url string, interval int64) (int64, error) {
	stmt := `INSERT INTO urls (url, interval) VALUES(?,?)`
	result, err := m.DB.Exec(stmt, &url, &interval)

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
	row := m.DB.QueryRow(stmt, &id)

	u := models.URLObject{}

	err := row.Scan(&u.ID, &u.URL, &u.Interval)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoRecord
		} else {
			return models.ErrAnother
		}

	}
	return nil
}

func (m *URLObjectModel) Get() ([]*models.URLObject, error) {

	stmt := `SELECT * FROM urls`
	rows, err := m.DB.Query(stmt)
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
