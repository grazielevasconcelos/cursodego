package main

import "fmt"

func main() {
	capitais := []string{"Lisboa"}
	// fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasília")
	//fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tokio"
	cidades[4] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))

	capitais[1] = "Brasilia"
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}

	//Primeiro item comeca com indice 0
	//Segundo item começa com o indice 1
	// capitaisAsia := cidades[3:5]
	// fmt.Println(capitaisAsia, len(capitaisAsia), cap(capitaisAsia))

	// temp1 := cidades[:2]
	// fmt.Println(temp1, len(temp1), cap(temp1))

	// temp2 := cidades[len(cidades)-2:]
	// fmt.Println(temp2, len(temp2), cap(temp2))

	indiceDoItemARetirar := 2

	temp := cidades[:indiceDoItemARetirar]
	fmt.Println("Slice original de cidades: ", cidades, len(cidades), cap(cidades))
	fmt.Println("Slice temp com as cidades iniciais que desejamos ficar: ", temp, len(temp), cap(temp))
	fmt.Println("Cidades com indice de retirada +1 ", cidades[indiceDoItemARetirar+1:])
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	fmt.Println("Cidades após o append ", cidades[indiceDoItemARetirar+1:])
	fmt.Println("Slice temp com cidades iniciais acrescentando as cidades após o item removido que desejamos ficar: ", temp, len(temp), cap(temp))
	fmt.Println("NOVAMENTE Slice original de cidades: ", cidades, len(cidades), cap(cidades))
	//copy(cidades, temp)
	fmt.Println(cidades, len(cidades), cap(cidades))
}
