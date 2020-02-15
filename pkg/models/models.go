package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	ErrAnother  = errors.New("models: operation for object faild")
)

type URLObject struct {
	ID       int64  `json:"id"`
	URL      string `json:"url"`
	Interval int64  `json:"interval"`
}

type URLResponse struct {
	ID          int64     `json:"id"`
	IDUrlObject int64     `json:"id-urlobject"`
	Response    string    `json:"response"`
	Duration    float64   `json:"duration"`
	Created     time.Time `json:"created-at"`
}
