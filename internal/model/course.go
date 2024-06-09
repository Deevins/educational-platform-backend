package model

import "time"

// Course represents a course
type Course struct {
	ID               int32            `json:"id"`
	Title            string           `json:"title"`
	Subtitle         string           `json:"subtitle"`
	Description      string           `json:"description"`
	Language         string           `json:"language"`
	AvatarURL        string           `json:"avatar_url"`
	Requirements     []string         `json:"requirements"`
	Level            string           `json:"level"`
	LecturesLength   time.Duration    `json:"lectures_length"`
	LecturesCount    int              `json:"lectures_count"`
	StudentsCount    int              `json:"students_count"`
	ReviewsCount     int              `json:"reviews_count"`
	Rating           float64          `json:"rating"`
	PreviewVideoURL  string           `json:"preview_video_URL"`
	TargetAudience   []string         `json:"target_audience"`
	WhatYouWillLearn []string         `json:"what_you_will_learn"`
	CourseGoals      []string         `json:"course_goals"`
	Instructor       CourseInstructor `json:"instructor"`
}

// CourseInstructor represents an instructor of a course
type CourseInstructor struct {
	ID            int32   `json:"id"`
	FullName      string  `json:"full_name"`
	Avatar        []byte  `json:"avatar"`
	Description   string  `json:"description"`
	StudentsCount int32   `json:"students_count"`
	CoursesCount  int32   `json:"courses_count"`
	Rating        float64 `json:"rating"` // count all ratings of instructor and divide by count of courses
}

// CoursePageReview represents a review of a course on the course page
type CoursePageReview struct {
	ID         int32   `json:"id"`
	FullName   string  `json:"full_name"`
	Avatar     []byte  `json:"avatar"`
	Rating     float64 `json:"rating"`
	ReviewText string  `json:"review_text"`
	CreatedAt  string  `json:"created_at"`
}

type ShortCourse struct {
	ID             int32         `json:"id,omitempty" `
	Title          string        `json:"title,omitempty" `
	AvatarURL      string        `json:"avatar_url,omitempty" `
	Subtitle       string        `json:"subtitle,omitempty" `
	Rating         float64       `json:"rating,omitempty" ` // count all ratings of instructor and divide by count of courses
	StudentsCount  int32         `json:"students_count,omitempty" `
	LecturesLength time.Duration `json:"lectures_length,omitempty" `
	Description    string        `json:"description,omitempty" `
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
