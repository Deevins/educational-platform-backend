package dto

type CourseBasicInfo struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Level       string `json:"level"`
	Category    string `json:"category"`
}

type CourseGoals struct {
	Goals          []string `json:"goals"`
	Requirements   []string `json:"requirements"`
	TargetAudience []string `json:"target_audience"`
}
