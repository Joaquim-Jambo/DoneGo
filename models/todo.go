package models

import "time"

type Todo struct {
	ID         string    `json:"id"`
	Titulo     string    `json:"titulo"`
	Descricao  string    `json:"descricao"`
	Categoria  string    `json:"categoria"`
	Estado     bool      `json:"estado"`
	DateCreat  time.Time `json:"date_creat"`
	DateUpdate time.Time `json:"date_update"`
}

type Todos struct {
	ListTodo map[string]Todo
}
