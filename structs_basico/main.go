package main

import (
	"fmt"
)

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é : %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu Apartamento", 9000000}
	fmt.Printf("O apartamento é : %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Cocorico",
		X:     22,
		valor: 100000000,
	}

	fmt.Printf("A chacara é: %v+\r\n", chacara)

	casa.Nome = "Casa do Silvio Santos"
	casa.valor = 15000000
	casa.X = 30
	casa.Y = 43
	fmt.Printf("A casa é : %+v\r\n", casa)
}
