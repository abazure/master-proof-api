package dto

import "time"

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

type RequestGetDiagnosticResult struct {
	UserId   string `json:"user_id"`
	QuizName string `json:"quiz_name"`
}

type RequestGetCompetenceResult struct {
	UserId   string `json:"user_id"`
	QuizName string `json:"quiz_name"`
}

type DiagnosticReportRequest struct {
	UserId             string `json:"user_id"`
	QuizId             string `json:"quiz_id"`
	DiagnosticReportId string `json:"diagnostic_report_id"`
}

type RequestBodyScore struct {
	Score int `json:"score"`
}

type CompetenceReportRequest struct {
	UserId string `json:"user_id"`
	QuizId string `json:"quiz_id"`
	Score  int    `json:"score"`
}

type ResponseDiagnosticReport struct {
	StudentId string    `json:"student_id"`
	Type      string    `json:"type"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseCompetenceReport struct {
	StudentId string    `json:"student_id"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}
