package model

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
