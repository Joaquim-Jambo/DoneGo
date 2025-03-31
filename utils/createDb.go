package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Joaquim-Jambo/DoneGo/models"
)

func CreateDb(todo models.Todo) error {
	//verificar se o arquivo exista e caso nao exista criar ele
	if _, err := os.Stat("data/db.json"); os.IsNotExist(err) {
		err = os.WriteFile("data/db.json", []byte("[]"), 0644)
		if err != nil {
			return errors.New("Erro ao criar o arquivo")
		}
	}
	data, err := os.ReadFile("data/db.json")
	if err != nil {
		return errors.New("Erro ao ler o arquivo db.json")
	}
	var todos []models.Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		return errors.New("Erro ao deserealizar a lista")
	}
	todos = append(todos, todo)
	ndata, err := json.Marshal(todos)
	if err != nil {
		return errors.New("Erro ao serealizar a lista")
	}
	err = os.WriteFile("data/db.json", ndata, 0644)
	if err != nil {
		return errors.New("Erro ao gravar o arquivo db")
	}
	return nil
}
