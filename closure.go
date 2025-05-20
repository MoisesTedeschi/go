// O que é uma Closure?
// Uma closure é uma função que captura o ambiente em que foi criada.
// Isso significa que ela pode acessar variáveis que estão fora de seu escopo,
// mesmo após a função externa ter terminado sua execução.

// Como Funciona
// Em Go, uma closure é criada quando você define uma função dentro de outra função.
// A função interna pode acessar as variáveis da função externa.

// Explicação
// 1 - Função Externa: A função main define uma função anônima que retorna outra função.
// Variável count: A variável count é criada no escopo da função externa, mas é capturada
// pela closure.
// 2 - Função Interna: A função interna incrementa count e a retorna.
// 3 - Chamada da Closure: Cada chamada para counter() mantém o estado de count, permitindo
// que ele seja incrementado a cada vez.

// Benefícios das Closures
// - Encapsulamento: Permite ocultar estados e variáveis, mantendo a interface limpa.
// - Funções de Retorno: Útil para criar funções que mantêm estado entre chamadas.

// Exemplos Práticos
// As closures são frequentemente usadas em:

// - Manipulação de eventos.
// - Funções de retorno de chamadas (callbacks).
// - Implementação de funções que precisam de estado interno.
//
// Conclusão
// Closures são uma poderosa ferramenta em Go que permite criar funções que mantêm e manipulam estados.
// Se precisar de mais exemplos ou quiser explorar outro tópico, é só avisar!

package main

import "fmt"

func main() {
	//Função que retorna uma Closure
	counter := func() func() int {
		count := 0 // Variável capturada pela Closure
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}