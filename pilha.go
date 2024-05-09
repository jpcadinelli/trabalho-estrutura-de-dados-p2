package pilha

import "fmt"

// Estrutura para representar a pilha
type Stack struct {
    elementos []interface{} // Armazena os elementos da pilha
}

// Função para empilhar um elemento
func (pilha *Stack) Push(elemento interface{}) {
    // Adiciona o elemento no topo da pilha
    pilha.elementos = append(pilha.elementos, elemento)
}

// Função para desempilhar um elemento
func (pilha *Stack) Pop() interface{} {
    // Verifica se a pilha está vazia
    if len(pilha.elementos) == 0 {
        return nil
    }

    // Remove o elemento do topo da pilha e o retorna
    elementoTopo := pilha.elementos[len(pilha.elementos)-1]
    pilha.elementos = pilha.elementos[:len(pilha.elementos)-1]
    return elementoTopo
}

// Função para verificar o elemento no topo da pilha
func (pilha *Stack) Top() interface{} {
    // Verifica se a pilha está vazia
    if len(pilha.elementos) == 0 {
        return nil
    }

    // Retorna o elemento no topo da pilha
    return pilha.elementos[len(pilha.elementos)-1]
}

// Função principal para demonstrar o uso da pilha
func pilha() {
    // Cria uma nova pilha
    pilha := Stack{}

    // Empilha alguns elementos
    pilha.Push("Elemento 1")
    pilha.Push("Elemento 2")
    pilha.Push("Elemento 3")

    // Desempilha e imprime os elementos
    fmt.Println("Elemento desempilhado:", pilha.Pop())
    fmt.Println("Elemento no topo:", pilha.Top())
    fmt.Println("Elemento desempilhado:", pilha.Pop())
    fmt.Println("Elemento desempilhado:", pilha.Pop())

    // Verifica se a pilha está vazia
    if pilha.Top() == nil {
        fmt.Println("Pilha vazia!")
    }
}
