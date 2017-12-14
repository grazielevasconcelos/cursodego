package main

import (
	"encoding/json"
	"fmt"

	"github.com/grazielevasconcelos/gobuild/model"
)

/*
	Comandos para compilação
		$go build 			> gera o executável gobuild
		$go build -o meuapp > gera o executável conforme o nome escolhido

	Cross compiler
		Gerar executável para windows
			$GOOS=windows GOARCH=386 go build -o meuapp.exe

		Gerar executável para linux(arquitetura ARM exemplo: raspberry) parâmetro -v exibe os fontes que estão sendo compilados
		parametro path (github.com/grazielevasconcelos/gobuild) serve para quando não se está no pacote direcionar o pacote
		que deseja a compilação:
			$GOOS=linux   GOARCH=arm go build -o meuappraspberry -v
			$GOOS=linux   GOARCH=arm go build -o meuappraspberry -v github.com/grazielevasconcelos/gobuild
*/
func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Grazi"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(400000); err != nil {
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
