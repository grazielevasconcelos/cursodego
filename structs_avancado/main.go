package main

import (
	"encoding/json"
	"fmt"

	"github.com/grazielevasconcelos/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Grazi"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(80000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("Valor da casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}