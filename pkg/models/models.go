package models

import (
	"time"
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	ErrAnother  = errors.New("models: operation faild")
)

type URLObject struct {
	ID       int64  `json:"id"`
	URL      string `json:"url"`
	Interval int64  `json:"interval"`
}

type URLResponse struct {
	Response string    `json:"response"`
	Duration float64   `json:"duration"`
	Created  time.Time `json:"created_at"`
}
