package models

type Student struct {
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Profession string `json:"profession"`
}

type Grade struct {
	Subject string `json:"subject"`
	Grade   int64  `json:"grade"`
}
