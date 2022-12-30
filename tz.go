package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var romanInput = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var arabInput = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var romanOutput = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
var arabOutput = []int{1, 4, 5, 9, 20, 40, 50, 90, 100}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		s := strings.Split(txt, " ")
		if len(s) != 3 {
			fmt.Println("Not correct input!")
			break
		}

		exp1, type1 := Identify(s[0])
		exp2, type2 := Identify(s[2])

		if type1 != type2 {
			fmt.Println("not arab and roman together!")
			break
		}

		exp := 0
		switch s[1] {
		case "+":
			exp = exp1 + exp2
		case "-":
			exp = exp1 - exp2
		case "*":
			exp = exp1 * exp2
		case "/":
			exp = exp1 / exp2
		default:
			fmt.Println("Not valid operation!")
			break
		}

		if type1 == "arab" {
			fmt.Println(exp)
		} else {
			if exp < 0 {
				fmt.Println("There are no negatives in roman!")
				break
			}
			fmt.Println(intToRoman(exp))
		}

	}
}

func Identify(x string) (int, string) {
	for i, n := range romanInput {
		if x == n {
			return i + 1, "roman"
		}
	}
	for i, n := range arabInput {
		if x == n {
			return i + 1, "arab"
		}
	}
	return 0, "Nan"
}

// подсмотрел в интернете
func intToRoman(num int) string {
	roman := ""
	index := len(romanOutput) - 1
	for num > 0 {
		for arabOutput[index] <= num {
			roman += romanOutput[index]
			num -= arabOutput[index]
		}
		index -= 1
	}
	return roman
}
