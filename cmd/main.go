package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite um numero")
	num1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num1 = strings.TrimSpace(num1)
	fmt.Println("-----------------------------")
	fmt.Println("Qual a operacao (+ - * /)")
	operacao, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	operacao = strings.TrimSpace(operacao)
	fmt.Println("-----------------------------")
	fmt.Println("Digite numero")
	num2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num2 = strings.TrimSpace(num2)
	convertNum1, _ := strconv.ParseFloat(num1, 64)
	convertNum2, _ := strconv.ParseFloat(num2, 64)

	var sum float64;

	switch operacao {
	case "+":
		sum = convertNum1 + convertNum2
	case "-":
		sum = convertNum1 - convertNum2
	case "*":
		sum = convertNum1 * convertNum2
	case "/":
		if convertNum2 != 0 {
			sum = convertNum1 / convertNum2
		} else {
			fmt.Println("Error: divisão por zero.")
		}
	default:
		fmt.Println("Operacao invalida.")
		return
	}

	println("-----------------------------")
	fmt.Println("Resultado: ", sum)

}
