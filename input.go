package input

import (
	"bufio"
	"os"
	"strconv"
)

/* -- Int type -- */

//Input a string and convert it to an integer
func Int() (int, error) {
	input := getString()
	inputInt, inputError := strconv.Atoi(input)
	return inputInt, inputError
}

/* -- Float type -- */

//Input a string and convert it to an float64
func Float32() (float64, error) {
	input := getString()
	inputFloat, inputError := strconv.ParseFloat(input, 32)
	return inputFloat, inputError
}

//Input a string and convert it to an float64
func Float64() (float64, error) {
	input := getString()
	inputFloat, inputError := strconv.ParseFloat(input, 64)
	return inputFloat, inputError
}

/* -- String type -- */

//Ask for a string as input
func String() string {
	input := getString()
	return input
}

/* -- Boolean type -- */

//Input a string and convert it to Boolean
func Bool() bool {
	input := getString()
	inputBool, inputError := strconv.ParseBool()
	return inputBool, inputError
}

/* -- KiritoNya's "get" function -- */

//Input a string
func getString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	string := scanner.Text()
	return string
}
