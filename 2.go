package main

import "time"

type Question struct {
	ID         int64
	Content    string
	CreatedAt  time.Time
	DeleteAt   time.Time
	UpdateAt   time.Time
	UserSearch string
	Draft      string
	Comments   int64
	Approvals  int64
	Like       int64
	Collects   int64
	Share      int64
}
