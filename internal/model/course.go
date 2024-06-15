package model

import "time"

// Course represents a course
type Course struct {
	ID              int32             `json:"id"`
	Title           string            `json:"title"`
	Subtitle        string            `json:"subtitle"`
	Description     string            `json:"description"`
	Language        string            `json:"language"`
	AvatarURL       string            `json:"avatar_url"`
	Requirements    []string          `json:"requirements"`
	Level           string            `json:"level"`
	LecturesLength  time.Duration     `json:"lectures_length"`
	LecturesCount   int               `json:"lectures_count"`
	StudentsCount   int               `json:"students_count"`
	ReviewsCount    int               `json:"reviews_count"`
	Rating          float64           `json:"rating"`
	PreviewVideoURL string            `json:"preview_video_URL"`
	TargetAudience  []string          `json:"target_audience"`
	CourseGoals     []string          `json:"course_goals"`
	Instructor      *CourseInstructor `json:"instructor"`
	CreatedAt       time.Time         `json:"created_at"`
	Status          string            `json:"status"`
	Reviews         []*CourseReview   `json:"reviews"`
	Category        string            `json:"category"`
}

// CourseInstructor represents an instructor of a course
type CourseInstructor struct {
	ID            int32               `json:"id"`
	FullName      string              `json:"full_name"`
	AvatarURL     string              `json:"avatar_url"`
	Description   string              `json:"description"`
	StudentsCount int32               `json:"students_count"`
	CoursesCount  int32               `json:"courses_count"`
	RatingsCount  int32               `json:"ratings_count"`
	Courses       []*InstructorCourse `json:"courses"`
	Rating        float64             `json:"rating"` // count all ratings of instructor and divide by count of courses
}

type ShortCourse struct {
	ID              int32         `json:"id" `
	Title           string        `json:"title" `
	CourseAvatarURL string        `json:"course_avatar_url" `
	Status          string        `json:"status" `
	Subtitle        string        `json:"subtitle" `
	Level           string        `json:"level" `
	AuthorFullName  string        `json:"author_full_name" `
	Rating          float64       `json:"rating" ` // count all ratings of instructor and divide by count of courses
	StudentsCount   int32         `json:"students_count" `
	ReviewsCount    int32         `json:"reviews_count" `
	LecturesCount   int32         `json:"lectures_count" `
	LecturesLength  time.Duration `json:"lectures_length" `

	Description string `json:"description,omitempty" `
}

// CourseReview represents a review of a course on the course page
type CourseReview struct {
	FullName   string    `json:"full_name"`
	AvatarURL  string    `json:"avatar"`
	Rating     int32     `json:"rating"`
	ReviewText string    `json:"review_text"`
	CreatedAt  time.Time `json:"created_at"`
}

type InstructorCourse struct {
	ID           int32   `json:"id"`
	Title        string  `json:"title"`
	AvatarURL    string  `json:"avatar_url"`
	Rating       float64 `json:"rating"`
	ReviewsCount int32   `json:"reviews_count"`
}

type GetAllCoursesResponse struct {
	Courses        []*ShortCourse `json:"courses"`
	AuthorFullName string         `json:"author_full_name"`
	Tags           []string       `json:"tags"`
}

type CourseBase struct {
	AuthorID    int32  `json:"author_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	CategoryID  int32  `json:"category_id"`
	TimePlanned string `json:"time_planned"`
}

type CourseIDWithResourceLink struct {
	CourseID int32  `json:"course_id"`
	Link     string `json:"link"`
}

type UpdateCourseGoals struct {
	Goals          []string `json:"goals"`
	Requirements   []string `json:"requirements"`
	TargetAudience []string `json:"target_audience"`
}

type CourseSection struct {
	SectionID          int32      `json:"section_id"`
	SectionTitle       string     `json:"section_title"`
	SectionDescription string     `json:"section_description"`
	Lectures           []*Lecture `json:"lectures"`
	Tests              []*Test    `json:"tests"`
}

type SectionCreation struct {
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
}
