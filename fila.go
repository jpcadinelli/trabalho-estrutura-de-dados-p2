package fila

import "fmt"

// Estrutura para representar a fila
type Queue struct {
    elementos []interface{} // Armazena os elementos da fila
}

// Função para enfileirar um elemento
func (fila *Queue) Enqueue(elemento interface{}) {
    // Adiciona o elemento no final da fila
    fila.elementos = append(fila.elementos, elemento)
}

// Função para desenfileirar um elemento
func (fila *Queue) Dequeue() interface{} {
    // Verifica se a fila está vazia
    if len(fila.elementos) == 0 {
        return nil
    }

    // Remove o elemento do início da fila e o retorna
    elementoFrente := fila.elementos[0]
    fila.elementos = fila.elementos[1:]
    return elementoFrente
}

// Função para verificar o elemento na frente da fila
func (fila *Queue) Front() interface{} {
    // Verifica se a fila está vazia
    if len(fila.elementos) == 0 {
        return nil
    }

    // Retorna o elemento na frente da fila
    return fila.elementos[0]
}

// Função principal para demonstrar o uso da fila
func fila() {
    // Cria uma nova fila
    fila := Queue{}

    // Enfileira alguns elementos
    fila.Enqueue("Elemento 1")
    fila.Enqueue("Elemento 2")
    fila.Enqueue("Elemento 3")

    // Desenfileira e imprime os elementos
    fmt.Println("Elemento desenfileirado:", fila.Dequeue())
    fmt.Println("Elemento na frente:", fila.Front())
    fmt
