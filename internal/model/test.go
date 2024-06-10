package model

type Test struct {
	TestID    int32      `json:"test_id"`
	TestName  string     `json:"test_name"`
	Questions []Question `json:"questions"`
}

type Question struct {
	QuestionBody string     `json:"question_body"`
	Answers      []Response `json:"answers"`
}

type Response struct {
	ResponseText string `json:"response_text"`
	Description  string `json:"description"`
	IsCorrect    bool   `json:"is_correct"`
}

type CreateTestBase struct {
	Title        string `json:"title"`
	SerialNumber int32  `json:"serial_number"`
	Description  string `json:"description"`
}
