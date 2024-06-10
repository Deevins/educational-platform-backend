package model

type Lecture struct {
	ID           int32  `json:"id"`
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
	VideoURL     string `json:"video_url"`
}

type CreateLectureBase struct {
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
}
