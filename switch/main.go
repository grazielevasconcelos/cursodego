package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da família do unix ?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim !")
	default:
		fmt.Println("Deixa esssa pergunta pra lá ... ")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK, descansa aí ;)")
	default:
		fmt.Println("Hoje é dia de trabalho")
	}

	numero = 100
	fmt.Println("Este numero cabe em um dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Pode ser com dois dígitos...")
	case numero >= 100:
		fmt.Println("Não vai dar pois o número é muito grande :(")
	}
}
