package request

import "time"

type ShortURLRequest struct {
	Url      string     `json:"url"`
	ExpireAt *time.Time `json:"expireAt"`
}
