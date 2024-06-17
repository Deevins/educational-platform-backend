package dto

type TestSubmission struct {
	UserID               int32 `json:"user_id"`
	TotalQuestionsCount  int32 `json:"total_questions_count"`
	CorrectAnsweredCount int32 `json:"correct_answered_count"`
}
