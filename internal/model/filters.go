package model

import "time"

type Filters struct {
	Rating         float64       `json:"rating"`
	Level          string        `json:"level"`
	LecturesLength time.Duration `json:"lectures_length"`
}
