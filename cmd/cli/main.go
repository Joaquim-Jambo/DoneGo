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

	for count := 0; ; count++ {
		fmt.Println("===================================")
		fmt.Println()
		fmt.Println("1)✅ Listar todos os Todos")
		fmt.Println("2)📂 Listar por Categoria")
		fmt.Println("3)📝 Adicionar Novo Todo")
		fmt.Println("4)✏️ Editar Todo")
		fmt.Println("5)❌ Deletar Todo")
		fmt.Println()
		fmt.Println("===================================")
		fmt.Print("Digite a opção desejada: ")
		fmt.Scan(&menu)
		reader.ReadString('\n')
		switch menu {
		case 1:
			todos, err := todo.GetTodo()
			if err != nil {
				fmt.Println(err)
			}
			for _, todo_ := range todos {
				emojiCategoria := utils.ObterEmojiCategoria(todo_.Categoria)
				fmt.Println("🗂️ Lista de Todos:")
				fmt.Println("-------------------------------")
				fmt.Printf(
					"%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n",
					emojiCategoria,
					todo_.Categoria,
					todo_.Titulo,
					todo_.Descricao,
					todo_.DateCreat.Format("02/01/2006 15:04"),
					utils.StatusEmoji(todo_.Estado),
				)
			}
		case 2:
			fmt.Println("Digite a categoria que deseja a listar: ")
			fmt.Scan(&categoria)
			todos, err := todo.GetByCategory(categoria)
			if err != nil {
				fmt.Println(err)
			}
			for _, todoN := range todos {
				emojiCategoria := utils.ObterEmojiCategoria(todoN.Categoria)
				fmt.Println("🗂️ Lista de Todos:")
				fmt.Println("-------------------------------")
				fmt.Printf(
					"%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n	",
					emojiCategoria,
					todoN.Categoria,
					todoN.Titulo,
					todoN.Descricao,
					todoN.DateCreat.Format("02/01/2006 15:04"),
					utils.StatusEmoji(todoN.Estado),
				)
				fmt.Println()
			}
		case 3:
			fmt.Println("Digite o titulo da tarfa")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			fmt.Println("Digite a descrição da tarefa")
			descricao, _ := reader.ReadString('\n')
			descricao = strings.TrimSpace(descricao)
			fmt.Println("Digite a categoria da tarefa")
			categoria, _ := reader.ReadString('\n')
			categoria = strings.TrimSpace(categoria)
			todo.AddTodo(titulo, descricao, categoria)
		case 4:
			{
				fmt.Println("Digite o id do Todo que deseja editar: ")
				id, _ := reader.ReadString('\n')
				fmt.Println("Digite o titulo da tarfa")
				titulo, _ := reader.ReadString('\n')
				fmt.Println("Digite a descrição da tarefa")
				descricao, _ := reader.ReadString('\n')
				todo.UpdateTodo(id, models.Todo{Titulo: titulo, Descricao: descricao})
			}
		case 0:
			{
				os.Exit(0)
			}
		}

	}

}
