package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("cidades.csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	// scanner := bufio.NewScanner(arquivo)

	// for scanner.Scan() {
	// 	linha := scanner.Text()
	// 	fmt.Println("O conteúdo da linha é: ", linha)
	// }

	leitorCSV := csv.NewReader(arquivo)

	//conteudo 	primeiro slice: linhas
	//	 		segundo slice: elemento com split
	conteudo, err := leitorCSV.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo com leitorCSV. Erro: ", err.Error())
		return
	}

	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indiceItem, item)
		}
	}

	arquivo.Close()

}
