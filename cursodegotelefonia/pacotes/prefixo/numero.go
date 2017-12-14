package prefixo

import (
	"strconv"
)

//Capital representa o numero do prefixo de telefone da capital de um estado / provincia
var Capital = 11

var teste = "teste"

//TesteComPrefixo isto e so um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)

//NomeOperadoraPrefixoInterior nome da operadora e seu prefixo no anterior
//var NomeOperadoraPrefixoInterior = operadora.NomeOperadora + " - 13"
