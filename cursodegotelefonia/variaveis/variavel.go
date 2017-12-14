package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string
	//Telefone é um valor importante para nossa aula
	Telefone            string
	quantidade, estoque int     //default 0
	comprou             bool    // default false
	valor               float64 // default 0.00
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("telefone: %v\r\n", Telefone)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
