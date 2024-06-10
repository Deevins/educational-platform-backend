package model

type Question struct {
	IsCorrect    bool   `json:"is_correct"`
	QuestionBody string `json:"question_body"`
}

type Test struct {
	TestID    int32      `json:"test_id"`
	TestName  string     `json:"test_name"`
	Questions []Question `json:"questions"`
}

type InsertQuestion struct {
	QuestionBody string `json:"question_body"`
	IsCorrect    bool   `json:"is_correct"`
}

type Response struct {
	TestID       int32  `json:"test_id"`
	ResponseText string `json:"response_text"`
	IsCorrect    bool   `json:"is_correct"`
}

type CreateTestBase struct {
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
}
