package models

type Sequence struct {
	ID int `gorm:"size:20;uniqueIndex"`
}
