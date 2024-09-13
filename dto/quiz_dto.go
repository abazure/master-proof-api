package dto

type Option struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type QuestionWithCorrectAnswer struct {
	Id            string   `json:"id"`
	Question      string   `json:"question"`
	CorrectAnswer int      `json:"actual_answer_value"`
	AnswerOptions []Option `json:"answer_options"`
}

type QuestionWithoutCorrectAnswer struct {
	Id            string   `json:"id"`
	Question      string   `json:"question"`
	AnswerOptions []Option `json:"answer_options"`
}

type RequestBodyResult struct {
	Result string `json:"result"`
}

type DiagnosticReportRequest struct {
	UserId             string `json:"user_id"`
	QuizId             string `json:"quiz_id"`
	DiagnosticReportId string `json:"diagnostic_report_id"`
}
