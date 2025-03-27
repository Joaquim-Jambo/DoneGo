package services

import (
	"errors"
	"time"
	"github.com/Joaquim-Jambo/DoneGo/models"
	"github.com/google/uuid"
)

func NewTodo() *models.Todos {
	return &models.Todos{
		listTodo: make(map[string]Todo),
	}
}

func (t *models.Todos) addTodo(titulo string, descricao string, categoria string) models.Todo {
	id := uuid.New().String()
	t.listTodo[id] = models.Todo{
		id:         id,
		titulo:     titulo,
		descricao:  descricao,
		estado:     false,
		categoria:  categoria,
		DateCreat:  time.Now(),
		DateUpdate: time.Now(),
	}
	return t.listTodo[id]
}
func (t *models.Todos) getByCategory(categoria string) ([]models.Todo, error) {
	var todos []models.Todo
	for _, todo := range t.listTodo {
		if todo.categoria == categoria {
			todos = append(todos, todo)
		}
	}
	if len(todos) == 0 {
		return nil, errors.New("❌ Nenhuma tarefa encontrada para essa categoria")
	}
	return todos, nil
}
func (t *models.Todos) getById(id string) (models.Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	return Todo_, nil
}
func (t *models.Todos) completedTodo(id string) (models.Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	Todo_.estado = true
	Todo_.DateUpdate = time.Now()
	t.listTodo[id] = Todo_
	return Todo_, nil

}
func (t *models.Todos) deleteTodo(id string) (models.Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	delete(t.listTodo, id)
	return Todo_, nil
}
func (t *Todos) updateTodo(id string, todo models.Todo) (models.Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado !")
	}
	Todo_.titulo = todo.titulo
	Todo_.descricao = todo.descricao
	Todo_.categoria = todo.categoria
	Todo_.DateUpdate = time.Now()
	t.listTodo[id] = Todo_
	return Todo_, nil
}
func (t *models.Todos) getTodo() []models.Todo {
	var todos []models.Todo

	for _, todo := range t.listTodo {
		todos = append(todos, todo)
	}
	return todos
}
