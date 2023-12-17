package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 1abc2        - 12
// pqr3stu8vwx  - 38
// a1b2c3d4e5f  - 15
// treb7uchet   - 77

// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	fmt.Print(string(data))

	lines := strings.Split(string(data), "\n")

	var sum = 0

	for i, row := range lines {
		if row == "" {
			continue
		}
		firstNumber := -1
		lastNumber := -1

		for i, char := range row {

			if firstNumber == -1 {
				if unicode.IsNumber(char) {
					firstNumber, _ = strconv.Atoi(string(char))
				}
			}

			if lastNumber == -1 {
				var last = row[(len(row)-1)-i]
				if unicode.IsNumber(rune(last)) {
					var last_char = string(rune(last))
					lastNumber, _ = strconv.Atoi(last_char)
				}
			}

		}

		fmt.Println(i, ":", firstNumber, lastNumber)
		var string_digit = strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber)
		var number_digit, _ = strconv.Atoi(string_digit)
		sum += number_digit

	}
	fmt.Println("sum:", sum)

}

func isInt(s string) bool {
	_, err := strconv.Atoi((s))
	return (err == nil)
}
