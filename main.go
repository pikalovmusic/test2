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

var numIntToRimsk = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var a, b *int
var opMap = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}

func intToRimsk(inputResult int) {
	var romanNum string
	for _, mapItem := range numIntToRimsk {
		for i := mapItem; i <= inputResult; {
			for in, value := range numRimsk {
				if value == mapItem {
					romanNum += in
					inputResult -= mapItem
				}
			}
		}
	}
	fmt.Println(romanNum)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		op = ""
		fmt.Println("Введите выражение и нажмите Enter")
		usertext, _ = reader.ReadString('\n')
		usertext = strings.ReplaceAll(usertext, " ", "")
		// op = strings.Split(usertext, "")[1]

		for idx := range opMap {
			for _, val := range usertext {
				if idx == string(val) {
					op += idx
					op = strings.TrimSpace(op)
				}
			}
		}
		if len(op) > 1 {
			panic("более одного операнта")
		}
		numA = strings.Split(usertext, op)[0]
		numB = strings.Split(usertext, op)[1]
		intA, _ := strconv.Atoi(strings.TrimSpace(numA))
		intB, _ := strconv.Atoi(strings.TrimSpace(numB))

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
			intA = numRimsk[strings.TrimSpace(numA)]
			intB = numRimsk[strings.TrimSpace(numB)]
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
					//fmt.Println(numIntToRimsk[res])
					intToRimsk(res)
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
