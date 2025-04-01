package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Joaquim-Jambo/DoneGo/models"
	"github.com/Joaquim-Jambo/DoneGo/services"
	"github.com/Joaquim-Jambo/DoneGo/utils"
)

func main() {
	todo := services.NewTodo()
	reader := bufio.NewReader(os.Stdin)
	var menu int
	var categoria string
	var id string

	for {
		utils.Menu()
		fmt.Println("Digite a opção desejada: ")
		fmt.Scan(&menu)
		reader.ReadString('\n')

		switch menu {
		case 1:
			utils.ClearScreen()
			fmt.Println("📜 Lista de Todas as Tarefas")
			fmt.Println("---------------------------------")
			todos, err := services.GetTodo()
			if err != nil {
				fmt.Println(err)
			}
			for _, todo_ := range todos {
				emojiCategoria := utils.ObterEmojiCategoria(todo_.Categoria)
				fmt.Printf("%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n\n",
					emojiCategoria, todo_.Categoria, todo_.Titulo, todo_.Descricao,
					todo_.DateCreat.Format("02/01/2006 15:04"), utils.StatusEmoji(todo_.Estado))
			}

		case 2:
			utils.ClearScreen()
			fmt.Println("🔍 Filtrar Tarefas por Categoria")
			fmt.Print("Digite a categoria desejada: ")
			fmt.Scan(&categoria)
			reader.ReadString('\n')
			todos, err := services.GetByCategory(categoria)
			if err != nil {
				fmt.Println(err)
			}
			for _, todoN := range todos {
				emojiCategoria := utils.ObterEmojiCategoria(todoN.Categoria)
				fmt.Printf("%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n\n",
					emojiCategoria, todoN.Categoria, todoN.Titulo, todoN.Descricao,
					todoN.DateCreat.Format("02/01/2006 15:04"), utils.StatusEmoji(todoN.Estado))
			}

		case 3:
			utils.ClearScreen()
			fmt.Println("📝 Criar Nova Tarefa")
			fmt.Print("Digite o título da tarefa: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			fmt.Print("Digite a descrição da tarefa: ")
			descricao, _ := reader.ReadString('\n')
			descricao = strings.TrimSpace(descricao)
			fmt.Print("Digite a categoria da tarefa: ")
			categoria, _ := reader.ReadString('\n')
			categoria = strings.TrimSpace(categoria)
			todo.AddTodo(titulo, descricao, categoria)
			fmt.Println("✅ Tarefa adicionada com sucesso!")

		case 4:
			utils.ClearScreen()
			fmt.Println("✏️ Editar Tarefa")
			fmt.Print("Digite o ID da tarefa: ")
			fmt.Scan(&id)
			reader.ReadString('\n')
			fmt.Print("Digite o novo título [OPCIONAL]: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			fmt.Print("Digite a nova descrição [OPCIONAL]: ")
			descricao, _ := reader.ReadString('\n')
			descricao = strings.TrimSpace(descricao)
			data, err := services.UpdateTodo(id, models.Todo{Titulo: titulo, Descricao: descricao})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("✅ Tarefa atualizada com sucesso!", data)
			}

		case 5:
			utils.ClearScreen()
			fmt.Println("🗑️ Apagar Tarefa")
			fmt.Print("Digite o ID da tarefa: ")
			fmt.Scan(&id)
			data, err := services.DeleteTodo(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("✅ Tarefa apagada com sucesso!", data)
			}

		case 6:
			utils.ClearScreen()
			fmt.Println("✅ Concluir Tarefa")
			fmt.Print("Digite o ID da tarefa: ")
			fmt.Scan(&id)
			data, err := services.CompletedTodo(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("🎉 Tarefa concluída com sucesso!", data)
			}

		case 0:
			utils.ClearScreen()
			fmt.Println("👋 Saindo... Até logo!")
			os.Exit(0)
		}
	}
}
