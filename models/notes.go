package models

import "time"

type Notes struct {
	Id          int
	Title       string
	Content     string
	CreatedTime time.Time
	UpdatedTime time.Time
}