package model

type Category struct {
	ID            int32       `json:"id"`
	Name          string      `json:"name"`
	Subcategories []*Category `json:"subcategories"`
}

type Language struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type MetasCount struct {
	RegistrationsCount int32 `json:"registrations_count"`
	CoursesCount       int32 `json:"courses_count"`
	StudentsCount      int32 `json:"students_count"`
}
