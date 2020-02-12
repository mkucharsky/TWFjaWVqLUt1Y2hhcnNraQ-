package mysql

import (
	"database/sql"
	"mkucharsky/wpapi/pkg/models"
)

type URLResponseModel struct {
	DB *sql.DB
}

func (m *URLResponseModel) Insert(obj *models.URLResponse) {

}

func (m *URLResponseModel) Get() {

}
