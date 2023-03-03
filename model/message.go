package model

import "time"

type Message struct {
	Text       string
	Time       time.Time
	LineUserId string
}

type Push struct {
	Text       string
	LineUserId string
}
