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

func (m *URLObjectModel) Delete(id int64) (int64, error) {
	stmt := `DELETE FROM urls WHERE id = ?`
	result, err := m.DB.Exec(stmt, &id)

	if err != nil {
		return 0, err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	if aff == 0 {
		return 0, models.ErrNoRecord
	} else if aff > 1 {
		return 0, models.ErrAnother
	}

	return result.LastInsertId()
}

func (m *URLObjectModel) Get() ([]*models.URLObject, error) {

	stmt := `SELECT * FROM urls ORDER BY id`
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

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *URLObjectModel) IfExists(id int64) (bool, error) {
	var amount int64
	err := m.DB.QueryRow(`SELECT count(id) AS amount from urls where id = ?`, id).Scan(&amount)

	if err != nil {
		return false, err
	}

	switch amount {
	case 0:
		return false, models.ErrNoRecord
	case 1:
		return true, nil
	default:
		return false, models.ErrAnother
	}
}
