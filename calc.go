package main

import (
	"fmt"
	"strings"
)

func operations(number1 string, number2 string, expr int) string {
	s1 := strings.TrimSpace(number1)
	s2 := strings.TrimSpace(number2)
	n1 := 0
	n2 := 0
	roman1 := false
	roman2 := false

	if s1 == "1" {
		n1 = 1
	}
	if s1 == "2" {
		n1 = 2
	}
	if s1 == "3" {
		n1 = 3
	}
	if s1 == "4" {
		n1 = 4
	}
	if s1 == "5" {
		n1 = 5
	}
	if s1 == "6" {
		n1 = 6
	}
	if s1 == "7" {
		n1 = 7
	}
	if s1 == "8" {
		n1 = 8
	}
	if s1 == "9" {
		n1 = 9
	}
	if s1 == "10" {
		n1 = 10
	}

	if s2 == "1" {
		n2 = 1
	}
	if s2 == "2" {
		n2 = 2
	}
	if s2 == "3" {
		n2 = 3
	}
	if s2 == "4" {
		n2 = 4
	}
	if s2 == "5" {
		n2 = 5
	}
	if s2 == "6" {
		n2 = 6
	}
	if s2 == "7" {
		n2 = 7
	}
	if s2 == "8" {
		n2 = 8
	}
	if s2 == "9" {
		n2 = 9
	}
	if s2 == "10" {
		n2 = 10
	}

	if s1 == "I" {
		roman1 = true
		n1 = 1
	}
	if s1 == "II" {
		roman1 = true
		n1 = 2
	}
	if s1 == "III" {
		roman1 = true
		n1 = 3
	}
	if s1 == "IV" {
		roman1 = true
		n1 = 4
	}
	if s1 == "V" {
		roman1 = true
		n1 = 5
	}
	if s1 == "VI" {
		roman1 = true
		n1 = 6
	}
	if s1 == "VII" {
		roman1 = true
		n1 = 7
	}
	if s1 == "VIII" {
		roman1 = true
		n1 = 8
	}
	if s1 == "IX" {
		roman1 = true
		n1 = 9
	}
	if s1 == "X" {
		roman1 = true
		n1 = 10
	}

	if s2 == "I" {
		roman2 = true
		n2 = 1
	}
	if s2 == "II" {
		roman2 = true
		n2 = 2
	}
	if s2 == "III" {
		roman2 = true
		n2 = 3
	}
	if s2 == "IV" {
		roman2 = true
		n2 = 4
	}
	if s2 == "V" {
		roman2 = true
		n2 = 5
	}
	if s2 == "VI" {
		roman2 = true
		n2 = 6
	}
	if s2 == "VII" {
		roman2 = true
		n2 = 7
	}
	if s2 == "VIII" {
		roman2 = true
		n2 = 8
	}
	if s2 == "IX" {
		roman2 = true
		n2 = 9
	}
	if s2 == "X" {
		roman2 = true
		n2 = 10
	}

	if n1 == 0 {
		fmt.Println("Первое значение меньше 1 или больше 10 или не число. Начните сначала")
		return "not ok"
	}
	if n2 == 0 {
		fmt.Println("Второе значение меньше 1 или больше 10 или не число. Начните сначала")
		return "not ok"
	}
	if (roman1 == false && roman2 == true) || (roman1 == true && roman2 == false) {
		fmt.Println("Одно из чисел арабское, другое - римское. Начните сначала")
		return "not ok"
	}

	var result int
	if expr == 0 {
		result = n1 + n2
	}
	if expr == 1 {
		result = n1 - n2
	}
	if expr == 2 {
		result = n1 * n2
	}
	if expr == 3 {
		result = n1 / n2
	}

	if roman1 == false && roman2 == false {
		fmt.Println(result)
		return "ok"
	}

	if roman1 == true && roman2 == true {
		fmt.Println(toRoman(result))
		return "ok"
	}

	return ""
}

func toRoman(n int) string {
	if n < 1 {
		return "Латинское число всегда больше нуля!"
	}
	dozens := [10]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
	units := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var roman string
	if n > 9 {
		roman = dozens[n/10-1]
	}
	if n%10 > 0 {
		roman += units[n%10-1]
	}
	return roman
}

func main() {
	for {
		var exp string
		fmt.Println("Введите 2 числа и знак операции (только +, -, *, /) между ними (или введите 0 для выхода):")
		fmt.Scanf("%s\n", &exp)
		if (exp=="0") {
			fmt.Println("Остановлено\n")
			break
		}
		exp = strings.ToUpper(exp)
		splittedPlus := strings.Split(exp, "+")
		splittedMinus := strings.Split(exp, "-")
		splittedMultiply := strings.Split(exp, "*")
		splittedDivide := strings.Split(exp, "/")

		if len(splittedPlus) == 1 && len(splittedMinus) == 1 && len(splittedMultiply) == 1 && len(splittedDivide) == 1 {
			fmt.Println("Нет знака операции в строке, начните сначала\n")
		}

		if len(splittedPlus) > 2 || len(splittedMinus) > 2 || len(splittedMultiply) > 2 || len(splittedDivide) > 2 {
			fmt.Println("Больше одного знака операции в строке, начните сначала\n")
		}

		if len(splittedPlus) == 2 {
			operations(splittedPlus[0], splittedPlus[1], 0)
		}
		if len(splittedMinus) == 2 {
			operations(splittedMinus[0], splittedMinus[1], 1)
		}
		if len(splittedMultiply) == 2 {
			operations(splittedMultiply[0], splittedMultiply[1], 2)
		}
		if len(splittedDivide) == 2 {
			operations(splittedDivide[0], splittedDivide[1], 3)
		}
	}
}
