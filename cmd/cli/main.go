package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Joaquim-Jambo/DoneGo/models"
	"github.com/Joaquim-Jambo/DoneGo/services"
	"github.com/Joaquim-Jambo/DoneGo/utils"
)

func main() {
	todo := services.NewTodo()
	reader := bufio.NewReader(os.Stdin)
	var menu int

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
		switch menu {
		case 1:
			for _, todo_ := range todo.GetTodo() {
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
			categoria, _ := reader.ReadString('\n')
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
			}
		case 3:
			fmt.Println("Digite o titulo da tarfa")
			titulo, _ := reader.ReadString('\n')
			fmt.Println("Digite a descrição da tarefa")
			descricao, _ := reader.ReadString('\n')
			fmt.Println("Digite a categoria da tarefa")
			categoria, _ := reader.ReadString('\n')
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
		}
	}

}
