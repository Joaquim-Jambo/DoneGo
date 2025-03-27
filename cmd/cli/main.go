package main
import (
	"github.com/google/uuid"
	"time"
	"errors"
	"fmt"
	"bufio"
	"os"
)

type gerenTodo interface {
	addTodo(titulo string, descricao string, categoria string) Todo
	updateTodo(id string, todo Todo) (Todo, error)
	deleteTodo(id string) (Todo, error)
	getTodo() []Todo
	completedTodo(id string) (Todo, error)
	getByCategory(categoria string) Todo
	getById(id string) Todo
}


func main(){
	todo := NewTodo()
	reader := bufio.NewReader(os.Stdin)
	var menu int
	
	for count:=0 ; ;count++{
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
		switch menu{
			case 1:
				for _, todo_:= range todo.getTodo(){
					emojiCategoria := obterEmojiCategoria(todo_.categoria)
					fmt.Println("🗂️ Lista de Todos:")
					fmt.Println("-------------------------------")
					fmt.Printf(
						"%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n",
						emojiCategoria,
						todo_.categoria,
						todo_.titulo,
						todo_.descricao,
						todo_.DateCreat.Format("02/01/2006 15:04"),
						statusEmoji(todo_.estado),
					)
					
				}
			case 2:
				fmt.Println("Digite a categoria que deseja a listar: ")
				categoria, _:= reader.ReadString('\n')
				todos, err := todo.getByCategory(categoria)
				if err != nil{
					fmt.Println(err)
				}
				for _, todoN := range todos{
					emojiCategoria := obterEmojiCategoria(todoN.categoria)
					fmt.Println("🗂️ Lista de Todos:")
					fmt.Println("-------------------------------")
					fmt.Printf(
						"%s [%s] - %s\n📋 Descrição: %s\n🕒 Criado em: %s\n⚙️ Status: %s\n	",
						emojiCategoria,
						todoN.categoria,
						todoN.titulo,
						todoN.descricao,
						todoN.DateCreat.Format("02/01/2006 15:04"),
						statusEmoji(todoN.estado),
					)
				} 
			case 3:
				fmt.Println("Digite o titulo da tarfa")
				titulo, _:= reader.ReadString('\n')
				fmt.Println("Digite a descrição da tarefa")
				descricao, _:= reader.ReadString('\n')
				fmt.Println("Digite a categoria da tarefa")
				categoria, _:= reader.ReadString('\n')
				todo.addTodo(titulo, descricao, categoria)
			case 4:{
				fmt.Println("Digite o id do Todo que deseja editar: ")
				id,_:= reader.ReadString('\n')
				fmt.Println("Digite o titulo da tarfa")
				titulo, _:= reader.ReadString('\n')
				fmt.Println("Digite a descrição da tarefa")
				descricao, _:= reader.ReadString('\n')
				todo.updateTodo(id, Todo{titulo: titulo, descricao: descricao})
			}
			}
	}
	
}