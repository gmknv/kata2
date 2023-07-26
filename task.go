package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	fmt.Println(calc(str))
}

func romToArab(romNum string) int {
	romNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var arabNum int
	for i := range romNums {
		if romNum == romNums[i] {
			arabNum = i + 1
		}
	}
	return arabNum
}

func arabToRom(arabNum int) string {
	romDigs := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	romTens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
	return romTens[arabNum/10] + romDigs[arabNum%10]
}

func isOK(val string) bool {
	is := true
	valSplit := strings.Split(val, " ")

	if len(valSplit) == 3 {

		if !(valSplit[1] == "+" || valSplit[1] == "-" || valSplit[1] == "*" || valSplit[1] == "/") {
			is = false
		}

		if !(isRom(valSplit[0]) && isRom(valSplit[2][0:len(valSplit[2])-2]) || isArab(valSplit[0]) && isArab(valSplit[2][0:len(valSplit[2])-2])) {
			is = false
		}
	} else {
		is = false
	}
	return is
}

func isRom(val string) bool {
	romDigs := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for r := range romDigs {
		if romDigs[r] == val {
			return true
		}
	}
	return false
}

func isArab(val string) bool {
	arabDigs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for r := range arabDigs {
		if arabDigs[r] == val {
			return true
		}
	}
	return false
}

func calc(input string) string {
	inputSplit := strings.Split(input, " ")
	var (
		res, num1, num2 int
		ret             string
	)
	if !isOK(input) {
		err := errors.New("строка ввода не удовлетворяет требованиям")
		fmt.Println(err)
		return ""

	} else {

		if isArab(inputSplit[0]) {
			num1, _ = strconv.Atoi(inputSplit[0])
			num2, _ = strconv.Atoi(inputSplit[2][0 : len(inputSplit[2])-2])

		}
		if isRom(inputSplit[0]) {
			num1 = romToArab(inputSplit[0])
			num2 = romToArab(inputSplit[2][0 : len(inputSplit[2])-2])
		}

		switch inputSplit[1] {
		case "+":
			res = num1 + num2
		case "-":
			res = num1 - num2
		case "*":
			res = num1 * num2
		case "/":
			res = num1 / num2
		}

		if isRom(inputSplit[0]) {
			if res > 0 {
				ret = arabToRom(res)
			} else {
				err := errors.New("результат вычитания римских чисел неположителен")
				fmt.Println(err)
				return ""
			}
		} else {
			ret = strconv.Itoa(res)
		}
	}
	return ret

}
