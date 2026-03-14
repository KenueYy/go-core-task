package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

var (
	numDecimal     int       = 42       // Десятичная система
	numOctal       int       = 052      // Восьмеричная система
	numHexadecimal int       = 0x2A     // Шестнадцатиричная система
	pi             float64   = 3.14     // Тип float64
	name           string    = "Golang" // Тип string
	isActive       bool      = true     // Тип bool
	complexNum     complex64 = 1 + 2i   // Тип complex64

	types = []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}
)

func main() {
	for _, t := range types {
		PrintType(t)
	}

	str := AllToStringAndConcatinate(types)
	fmt.Println(str)

	r := StringToRune(str)
	fmt.Println(r)

	fmt.Println(RuneToHash(r, "go-2024"))
}

func PrintType(value interface{}) string {
	s := reflect.TypeOf(value).String()
	fmt.Println(s)
	return s
}

func AllToStringAndConcatinate(values []any) string {
	var result string = ""
	for _, value := range values {
		result += fmt.Sprint(value)
	}
	return result
}

func StringToRune(s string) []rune {
	return []rune(s)
}

func RuneToHash(runes []rune, salt string) [32]byte {
	mid := len(runes) / 2
	result := make([]rune, 0, len(runes)+len([]rune(salt)))
	result = append(result, runes[:mid]...)
	result = append(result, []rune(salt)...)
	result = append(result, runes[mid:]...)

	bytes := []byte(string(result))
	return sha256.Sum256(bytes)
}
