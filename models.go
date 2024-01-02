package main

import "time"

type URLInfo struct {
	LongURL   string    `json:"longURL"`
	UserAgent string    `json:"userAgent"`
	CreatedAt time.Time `json:"createdAt"`
}
