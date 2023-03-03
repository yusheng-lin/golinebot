package model

import "time"

type Receive struct {
	Text       string
	Time       time.Time
	LineUserId string
}

type Push struct {
	Text       string
	LineUserId string
}
