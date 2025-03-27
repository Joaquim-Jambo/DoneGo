package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func NewTodo() *Todos {
	return &Todos{
		listTodo: make(map[string]Todo),
	}
}

func (t *Todos) addTodo(titulo string, descricao string, categoria string) Todo {
	id := uuid.New().String()
	t.listTodo[id] = Todo{
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
func (t *Todos) getByCategory(categoria string) ([]Todo, error) {
	var todos []Todo
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
func (t *Todos) getById(id string) (Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return Todo{}, errors.New("Todo não encontrado")
	}
	return Todo_, nil
}
func (t *Todos) completedTodo(id string) (Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return Todo{}, errors.New("Todo não encontrado")
	}
	Todo_.estado = true
	Todo_.DateUpdate = time.Now()
	t.listTodo[id] = Todo_
	return Todo_, nil

}
func (t *Todos) deleteTodo(id string) (Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return Todo{}, errors.New("Todo não encontrado")
	}
	delete(t.listTodo, id)
	return Todo_, nil
}
func (t *Todos) updateTodo(id string, todo Todo) (Todo, error) {
	Todo_, exist := t.listTodo[id]
	if !exist {
		return Todo{}, errors.New("Todo não encontrado !")
	}
	Todo_.titulo = todo.titulo
	Todo_.descricao = todo.descricao
	Todo_.categoria = todo.categoria
	Todo_.DateUpdate = time.Now()
	t.listTodo[id] = Todo_
	return Todo_, nil
}
func (t *Todos) getTodo() []Todo {
	var todos []Todo

	for _, todo := range t.listTodo {
		todos = append(todos, todo)
	}
	return todos
}
