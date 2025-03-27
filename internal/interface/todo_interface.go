package interface

import (
	"fmt"
	"github.com/Joaquim-Jambo/DoneGo/models"
)
type gerenTodo interface {
	addTodo(titulo string, descricao string, categoria string) models.Todo
	updateTodo(id string, todo models.Todo) (Todo, error)
	deleteTodo(id string) (models.Todo, error)
	getTodo() []models.Todo
	completedTodo(id string) (models.Todo, error)
	getByCategory(categoria string) models.Todo
	getById(id string) models.Todo
}
