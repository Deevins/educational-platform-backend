package model

import "time"

// Course represents a course
type Course struct {
	ID              int32            `json:"id"`
	Title           string           `json:"title"`
	Subtitle        string           `json:"subtitle"`
	Description     string           `json:"description"`
	Language        string           `json:"language"`
	AvatarURL       string           `json:"avatar_url"`
	Requirements    []string         `json:"requirements"`
	Level           string           `json:"level"`
	LecturesLength  time.Duration    `json:"lectures_length"`
	LecturesCount   int              `json:"lectures_count"`
	StudentsCount   int              `json:"students_count"`
	ReviewsCount    int              `json:"reviews_count"`
	Rating          float64          `json:"rating"`
	PreviewVideoURL string           `json:"preview_video_URL"`
	TargetAudience  []string         `json:"target_audience"`
	CourseGoals     []string         `json:"course_goals"`
	Instructor      CourseInstructor `json:"instructor"`
	CreatedAt       time.Time        `json:"created_at"`
	Status          string           `json:"status"`
	Category        string           `json:"category"`
}

// CourseInstructor represents an instructor of a course
type CourseInstructor struct {
	ID            int32   `json:"id"`
	FullName      string  `json:"full_name"`
	AvatarURL     string  `json:"avatar_url"`
	Description   string  `json:"description"`
	StudentsCount int32   `json:"students_count"`
	CoursesCount  int32   `json:"courses_count"`
	Rating        float64 `json:"rating"` // count all ratings of instructor and divide by count of courses
}

type ShortCourse struct {
	ID             int32         `json:"id,omitempty" `
	Title          string        `json:"title,omitempty" `
	AvatarURL      string        `json:"avatar_url,omitempty" `
	Status         string        `json:"status,omitempty" `
	Subtitle       string        `json:"subtitle,omitempty" `
	Rating         float64       `json:"rating,omitempty" ` // count all ratings of instructor and divide by count of courses
	StudentsCount  int32         `json:"students_count,omitempty" `
	LecturesLength time.Duration `json:"lectures_length,omitempty" `
	Description    string        `json:"description,omitempty" `
}

// CourseReview represents a review of a course on the course page
type CourseReview struct {
	ID         int32   `json:"id"`
	FullName   string  `json:"full_name"`
	AvatarURL  string  `json:"avatar"`
	Rating     float64 `json:"rating"`
	ReviewText string  `json:"review_text"`
	CreatedAt  string  `json:"created_at"`
}

type InstructorCourse struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	AvatarURL string `json:"avatar_url"`
	Status    string `json:"status"`
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

type Lecture struct {
	LectureID          int32  `json:"lecture_id"`
	LectureTitle       string `json:"lecture_title"`
	LectureDescription string `json:"lecture_description"`
	LectureVideoURL    string `json:"lecture_video_url"`
}

type Question struct {
	QuestionID   int32  `json:"question_id"`
	IsCorrect    bool   `json:"is_correct"`
	QuestionBody string `json:"question_body"`
}

type Test struct {
	TestID    int32      `json:"test_id"`
	TestName  string     `json:"test_name"`
	Questions []Question `json:"questions"`
}

type CourseSection struct {
	SectionID          int32      `json:"section_id"`
	SectionTitle       string     `json:"section_title"`
	SectionDescription string     `json:"section_description"`
	Lectures           []*Lecture `json:"lectures"`
	Tests              []*Test    `json:"tests"`
}
