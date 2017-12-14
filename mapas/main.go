package main

import (
	"fmt"

	"github.com/grazielevasconcelos/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa da Grazi"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(80000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]

	if achou {
		fmt.Printf("Sim, e o que ache foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto da Grazi"
	apto.X = 68
	apto.Y = 75
	apto.SetValor(180001)

	cache[apto.Nome] = apto

	rancho := model.Imovel{}
	rancho.Nome = "Rancho da Grazi"
	rancho.X = 105
	rancho.Y = 36
	rancho.SetValor(300000)

	cache[rancho.Nome] = rancho

	imovel, achou = cache[rancho.Nome]

	fmt.Println("Quanto itens ha no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave [%s] = %+v\n\r", chave, imovel)
	}

	if achou {
		delete(cache, rancho.Nome)
		fmt.Println("Quanto itens ha no cache? ", len(cache))
	}

}
