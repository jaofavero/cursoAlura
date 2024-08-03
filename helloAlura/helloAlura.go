package main

import "fmt"

func main() {
	// Declaração de variáveis
	// declaração inteira, exemplo usando const
	const versao float32 = 2.0
	// declaração curta sem especificar o tipo
	var idade = 24
	// usando := para declaração curta
	nome := "Alura"

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	//Scan reduzido
	fmt.Scan(&comando)
	/*
		// while loop
		for comando < 0 || comando > 2 {
			fmt.Println("Comando inválido")
			fmt.Scan(&comando)
		}
	*/
	fmt.Println("O comando escolhido foi", comando)
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")

	}
}
