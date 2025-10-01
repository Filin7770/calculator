package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input, operator string
	var num1, num2, result float64
	var operatorIndex, firstMinus int

	fmt.Println("Введите выражение")

	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	if input[0] == '-' {
		firstMinus = 1
	}

	for i := firstMinus; i < len(input); i++ {
		char := input[i]
		if char == '+' || char == '-' || char == '*' || char == '/' {
			operator = string(char)
			operatorIndex = i
			break
		}
	}

	sliceNum1 := input[0:operatorIndex]
	sliceNum2 := input[operatorIndex+1:]

	num1, err1 := strconv.ParseFloat(sliceNum1, 64)
	num2, err2 := strconv.ParseFloat(sliceNum2, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка: некорректный ввод чисел")
		return
	}

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Ошибка: знак не найден")
		return
	}

	fmt.Println(result)

}
