package entities

import "time"

type Auction struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	ExternalUrl string
	StartTime   time.Time
	StartDate   time.Time
}
