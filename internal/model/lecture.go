package model

type CreateLectureBase struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Lecture struct {
	ID           int32  `json:"id"`
	Type         string `json:"type"` // TODO: костыль, пока не понятно, что это
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
	VideoURL     string `json:"video_url"`
}
