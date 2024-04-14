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
var rimsk bool
var inttest int
var res int

var numRimsk = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var numIntToRimsk = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	8:   "VIII",
	7:   "VII",
	6:   "VI",
	5:   "V",
	4:   "IV",
	3:   "III",
	2:   "II",
	1:   "I",
}

var opMap = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

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

		if opMap[op] == false {
			panic("Ошибка. Неверная арифметическая операция")
		}

		if intA == 0 && intB == 0 {
			rimsk = true
		}
		if intA > 0 && intB > 0 {
			rimsk = false
		}

		if intA >= 1 && intB == 0 {
			panic("Ошибка. Можно использовать только римские или арабские цифры от 1 до 10 (I - X)")
		}
		if intA == 0 && intB >= 1 {
			panic("Ошибка. Можно использовать только римские или арабские цифры от 1 до 10 (I - X)")

		}

		if rimsk == true {
			intA = numRimsk[numA]
			intB = numRimsk[numB]
		}

		if intA >= 1 && intA <= 10 && intB >= 1 && intB <= 10 {

			if op == "+" {
				res = intA + intB
			}

			if op == "-" {
				res = intA - intB
			}

			if op == "*" {
				res = intA * intB
			}

			if op == "/" {
				res = intA / intB
			}

			if rimsk == false {
				fmt.Println(res)
			} else {
				if res < 1 {
					panic("Результат меньше единицы")
				} else {
					fmt.Println(numIntToRimsk[res])
				}
			}

		} else {
			panic("Ошибка. Вводимые числа должны быть от 1 до 10 (I - X)")
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Следующее вычисление:")

	}

}
