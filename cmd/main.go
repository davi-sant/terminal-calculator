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
	getNum1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Qual a operacao (+ - * /)")
	operation, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Digite numero")
	getNum2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	getNum1 = strings.TrimSpace(getNum1)
	operation = strings.TrimSpace(operation)
	getNum2 = strings.TrimSpace(getNum2)

	num1, _ := strconv.ParseFloat(getNum1, 64)
	num2, _ := strconv.ParseFloat(getNum2, 64)

	var sum float64;

	switch operation {
	case "+":
		sum = num1 + num2
	case "-":
		sum = num1 - num2
	case "*":
		sum = num1 * num2
	case "/":
		if num2 != 0 {
			sum = num1 / num2
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
