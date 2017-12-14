package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	meses, _ := strconv.Atoi(os.Args[1])
	situacao := true
	cidade := "Sampa"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Sampa" {
		fmt.Println("O cliente vive em Sampa")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}
	//A instrucao abaixo nao ira compilar devido ao escopo da variavel estar apenas para o comando acima IF
	//fmt.Println("Me passa o status? ", descricao)
	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo"
		return
	}
	descricao = "O cliente está com o carnê em dia"
	return
}
