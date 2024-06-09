package model

type Lecture struct {
	LectureID          int32  `json:"lecture_id"`
	LectureTitle       string `json:"lecture_title"`
	LectureDescription string `json:"lecture_description"`
	LectureVideoURL    string `json:"lecture_video_url"`
}
