package main

import (
	"encoding/json"
	"fmt"

	"github.com/grazielevasconcelos/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Grazi"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(4000000000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelInvalido {
			fmt.Println("Corretor, por favor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("Valor da casa é: %d\r\n", casa.GetValor())

	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] Houve um erro na geracao do objeto json", err.Error())
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))

}
