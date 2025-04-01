package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Joaquim-Jambo/DoneGo/models"
	"github.com/Joaquim-Jambo/DoneGo/utils"
	"github.com/google/uuid"
)

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

func (t *TodoManager) AddTodo(titulo string, descricao string, categoria string) (models.Todo, error) {
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
	err := utils.CreateDb(t.ListTodo[id])
	if err != nil {
		return models.Todo{}, err
	}
	return t.ListTodo[id], nil
}

func GetByCategory(categoria string) ([]models.Todo, error) {
	var todos []models.Todo
	var listTodo []models.Todo
	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return nil, errors.New("Erro ao carregar o arquivo db.json")
	}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, errors.New("Erro ao deserealizar a lista")
	}

	for _, todo := range todos {
		if todo.Categoria == categoria {
			listTodo = append(listTodo, todo)
		}
	}
	if len(listTodo) == 0 {
		return nil, errors.New("❌ Nenhuma tarefa encontrada para essa categoria")
	}
	return listTodo, nil
}

func CompletedTodo(id string) (string, error) {
	var todos []models.Todo
	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return "", errors.New("Erro ao ler o arquivo db.json")
	}
	err = json.Unmarshal(data, &todos)
	found := false
	var titulo string
	// pegar o id
	for i, t := range todos {
		if t.ID == id {
			if !todos[i].Estado {
				todos[i].Estado = true
				titulo = todos[i].Titulo
				todos[i].DateUpdate = time.Now()
				found = true
				break
			} else {
				return fmt.Sprintf("⚠️ A tarefa '%s' já está concluída!", t.Titulo), nil
			}
		}
	}
	if !found {
		return "", errors.New("Nenhuma tarefa encontrada com este ID")
	}
	ndata, err := json.Marshal(todos)
	if err != nil {
		return "", errors.New("Erro ao serializar os dados")
	}
	err = os.WriteFile("data/db.json", ndata, 0644)
	if err != nil {
		return "", errors.New("Erro ao gravar no arquivo db.json")
	}

	return fmt.Sprintf("✅ A tarefa '%s' foi concluída com sucesso!", titulo), nil
}

func DeleteTodo(id string) (string, error) {
	var todos []models.Todo
	found := false
	var UpdateTodo []models.Todo
	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return "", errors.New("Erro ao ler o arquivo db.json")
	}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return "", errors.New("Erro ao serializar os dados")
	}
	for _, t := range todos {
		if t.ID == id {
			found = true
			continue
		}
		UpdateTodo = append(UpdateTodo, t)
	}
	if !found {
		return "", errors.New("Nenhuma tarefa encontrada com este ID")
	}
	ndata, err := json.Marshal(UpdateTodo)
	if err != nil {
		return "", errors.New("Erro ao codificar os dados")
	}
	err = os.WriteFile("data/db.json", ndata, 0644)
	if err != nil {
		return "", errors.New("Erro ao salvar os dados")
	}
	return "Tarefa apaga com sucesso!", nil
}

func UpdateTodo(id string, todo models.Todo) (string, error) {
	var todos []models.Todo

	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return "", errors.New("Erro ao ler o arquivo db.json")
	}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return "", errors.New("Erro ao deserealizar a lista")
	}
	found := false
	for i, t := range todos {
		if t.ID == id {
			if todo.Categoria != "" && todo.Categoria != todos[i].Categoria {
				todos[i].Categoria = todo.Categoria
			}
			if todo.Descricao != "" && todo.Descricao != todos[i].Descricao {
				todos[i].Descricao = todo.Descricao
			}
			if todo.Titulo != "" && todo.Titulo != todos[i].Titulo {
				todos[i].Titulo = todo.Titulo
			}
			if todo.Categoria != "" || todo.Descricao != "" || todo.Titulo != "" {
				todos[i].DateUpdate = time.Now()
			}
			found = true
			break
		}
	}
	if !found {
		return "", errors.New("Nenhuma tarefa encontrada com este ID")
	}
	ndata, err := json.Marshal(todos)
	if err != nil {
		return "", errors.New("Erro ao codificar")
	}
	err = os.WriteFile("data/db.json", ndata, 0644)
	if err != nil {
		return "", errors.New("Erro ao acatualizar")
	}
	return "Tarefa actualizada com sucesso", nil
}

func GetTodo() ([]models.Todo, error) {
	var todos []models.Todo

	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return nil, errors.New("Erro ao ler o arquivo db.json")
	}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, errors.New("Erro ao deserealizar")
	}
	if len(todos) == 0 {
		return nil, errors.New("❌ Nenhum tarefa cadastrada")
	}
	return todos, nil
}
