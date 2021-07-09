package models

import "time"

type Url struct {
	ShortUrl string
	Url      string
	ExpireAt *time.Time
}

type ShortURLRequest struct {
	Url      string     `json:"url"`
	ExpireAt *time.Time `json:"expireAt"`
}

type ShortUrlResponse struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
