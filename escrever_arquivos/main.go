package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/grazielevasconcelos/escrever_arquivos/model"
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

	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o JSON do item ", item, " .Erro: ", err.Error())
			}
			escritor.WriteString("  " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}

	escritor.WriteString("\r\n]")
	escritor.Flush()
	arquivoJSON.Close()
	arquivo.Close()

}
