package models
import "time"
type Todo struct{
	id string `json:"id"`
	titulo string `json:titulo`
	descricao string `json:descricao`
	categoria string `json:categoria`
	estado bool `json:estado`
	DateCreat time.Time `json:date_creat`	
	DateUpdate time.Time `json:date_update`
}

type Todos struct{
	listTodo map[string]Todo
}