package models

import(
	"time"
)

type URLObject struct {
	URL string `json:"url"`
	Interval int32 `json:"interval"`
}

type URLResponse struct {
	Response string `json:"response"`
	Duration float32 `json:"duration"`
	Created  time.Time `json:"created_at"`
}

