package services

import (
	"errors"
	"time"

	"github.com/Joaquim-Jambo/DoneGo/models"
	"github.com/google/uuid"
)

// Interface para os métodos do serviço de tarefas
type TodoService interface {
	AddTodo(titulo string, descricao string, categoria string) models.Todo
	GetByCategory(categoria string) ([]models.Todo, error)
	GetById(id string) (models.Todo, error)
	CompletedTodo(id string) (models.Todo, error)
	DeleteTodo(id string) (models.Todo, error)
	UpdateTodo(id string, todo models.Todo) (models.Todo, error)
	GetTodo() []models.Todo
}

// Struct para gerenciar a lista de tarefas
type TodoManager struct {
	ListTodo map[string]models.Todo
}

// Construtor para inicializar o TodoManager
func NewTodo() *TodoManager {
	return &TodoManager{
		ListTodo: make(map[string]models.Todo),
	}
}

func (t *TodoManager) AddTodo(titulo string, descricao string, categoria string) models.Todo {
	id := uuid.New().String()
	t.ListTodo[id] = models.Todo{
		ID:         id,
		Titulo:     titulo,
		Descricao:  descricao,
		Categoria:  categoria,
		Estado:     false,
		DateCreat:  time.Now(),
		DateUpdate: time.Now(),
	}
	return t.ListTodo[id]
}

func (t *TodoManager) GetByCategory(categoria string) ([]models.Todo, error) {
	var todos []models.Todo
	for _, todo := range t.ListTodo {
		if todo.Categoria == categoria {
			todos = append(todos, todo)
		}
	}
	if len(todos) == 0 {
		return nil, errors.New("❌ Nenhuma tarefa encontrada para essa categoria")
	}
	return todos, nil
}

func (t *TodoManager) GetById(id string) (models.Todo, error) {
	todo, exist := t.ListTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	return todo, nil
}

func (t *TodoManager) CompletedTodo(id string) (models.Todo, error) {
	todo, exist := t.ListTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	todo.Estado = true
	todo.DateUpdate = time.Now()
	t.ListTodo[id] = todo
	return todo, nil
}

func (t *TodoManager) DeleteTodo(id string) (models.Todo, error) {
	todo, exist := t.ListTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	delete(t.ListTodo, id)
	return todo, nil
}

func (t *TodoManager) UpdateTodo(id string, todo models.Todo) (models.Todo, error) {
	todoAtual, exist := t.ListTodo[id]
	if !exist {
		return models.Todo{}, errors.New("Todo não encontrado")
	}
	todoAtual.Titulo = todo.Titulo
	todoAtual.Descricao = todo.Descricao
	todoAtual.Categoria = todo.Categoria
	todoAtual.DateUpdate = time.Now()
	t.ListTodo[id] = todoAtual
	return todoAtual, nil
}

func (t *TodoManager) GetTodo() []models.Todo {
	var todos []models.Todo
	for _, todo := range t.ListTodo {
		todos = append(todos, todo)
	}
	return todos
}
