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

//Input a string and convert it to an float32
func Float32() (float32, error) {
	input := getString()
	inputFloat64, inputError := strconv.ParseFloat(input, 32)
	return float32(inputFloat64), inputError
}

//Input a string and convert it to an float64
func Float64() (float64, error) {
	input := getString()
	inputFloat64, inputError := strconv.ParseFloat(input, 64)
	return inputFloat64, inputError
}

/* -- String type -- */

//Input a string and return it
func String() string {
	input := getString()
	return input
}

/* -- Boolean type -- */

//Input a string and convert it to Boolean
func Bool() (bool, error) {
	input := getString()
	inputBool, inputError := strconv.ParseBool(input)
	return inputBool, inputError
}

/* -- Complex type -- */

//Input a string and convert it to an complex64
func Complex64() (complex64, error) {
	input := getString()
	inputComplex128, coversionError := strconv.ParseComplex(input, 64)
	return complex64(inputComplex128), coversionError
}

//Input a string and convert it to an complex128
func Complex128() (complex128, error) {
	input := getString()
	inputComplex128, coversionError := strconv.ParseComplex(input, 64)
	return inputComplex128, coversionError
}

/* -- Byte type -- */

//Input a string and convert it to an byte (uint8)
func Byte() (byte, error) {
	input := getString()
	inputUint8, coversionError := strconv.ParseUint(input, 10, 8)
	return byte(inputUint8), coversionError
}

/* -- Rune type -- */

//Input a string and convert it to an rune (int32)
func Rune() (rune, error) {
	input := getString()
	inputInt32, coversionError := strconv.ParseInt(input, 10, 32)
	return rune(inputInt32), coversionError
}

/* -- KiritoNya's "get" function -- */

//Input a string
func getString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	string := scanner.Text()
	return string
}
