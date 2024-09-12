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
