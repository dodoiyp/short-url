package models

import "time"

type Url struct {
	ShortUrl  string    `gorm:"size:20;uniqueIndex"`
	Url       string    `gorm:"size:1024" binding:"required"`
	ExpireAt  time.Time `binding:"required"`
	CreatedAt time.Time `binding:"required"`
}
