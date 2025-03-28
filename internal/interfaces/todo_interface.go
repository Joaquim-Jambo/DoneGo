package interfaces

import (
	"github.com/Joaquim-Jambo/DoneGo/models"
)
type gerenTodo interface {
	AddTodo(titulo string, descricao string, categoria string) models.Todo
	UpdateTodo(id string, todo models.Todo) (models.Todo, error)
	DeleteTodo(id string) (models.Todo, error)
	GetTodo() []models.Todo
	CompletedTodo(id string) (models.Todo, error)
	GetByCategory(categoria string) models.Todo
	GetById(id string) models.Todo
}
