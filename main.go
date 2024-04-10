package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numA string
var numB string
var op string
var usertext string

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Введите выражение (a + b), добавьте пробел и нажмите Enter")
		usertext, _ = reader.ReadString('\n')
		numA = strings.Split(usertext, " ")[0]
		numB = strings.Split(usertext, " ")[2]
		op = strings.Split(usertext, " ")[1]
		intA, _ := strconv.Atoi(numA)
		intB, _ := strconv.Atoi(numB)

		if op == "+" {
			fmt.Println("Решение: ", numA, " + ", numB, " = ", intA+intB)
		}

		if op == "-" {
			fmt.Println("Решение: ", numA, " - ", numB, " = ", intA-intB)
		}

		if op == "*" {
			fmt.Println("Решение: ", numA, " * ", numB, " = ", intA*intB)
		}

		if op == "/" {
			fmt.Println("Решение: ", numA, " / ", numB, " = ", intA/intB)
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Следующее вычисление:")
	}

}
