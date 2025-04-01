package utils

import (
	"fmt"
)

func Menu() {

	fmt.Println("=============================================")
	fmt.Println("              📋 MENU DE TAREFAS             ")
	fmt.Println("=============================================")
	fmt.Println(" 1️)  ✅  Listar todas as tarefas")
	fmt.Println(" 2️)  📂  Listar tarefas por categoria")
	fmt.Println(" 3️)  📝  Adicionar nova tarefa")
	fmt.Println(" 4️)  ✏️  Editar tarefa")
	fmt.Println(" 5️)  ❌  Deletar tarefa")
	fmt.Println(" 6️)  ✔️  Marcar tarefa como concluída")
	fmt.Println(" 0️)  🚪  Sair do programa")
	fmt.Println("=============================================")
	fmt.Println(" 🔷 Escolha uma opção: ")
}
