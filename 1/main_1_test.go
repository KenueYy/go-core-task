package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PrintType(t *testing.T) {

	var (
		numDecimal     int       = 42       // Десятичная система
		numOctal       int       = 052      // Восьмеричная система
		numHexadecimal int       = 0x2A     // Шестнадцатиричная система
		pi             float64   = 3.14     // Тип float64
		name           string    = "Golang" // Тип string
		isActive       bool      = true     // Тип bool
		complexNum     complex64 = 1 + 2i   // Тип complex64
	)

	require.Equal(t, PrintType(numDecimal), "int")
	require.Equal(t, PrintType(numOctal), "int")
	require.Equal(t, PrintType(numHexadecimal), "int")
	require.Equal(t, PrintType(pi), "float64")
	require.Equal(t, PrintType(name), "string")
	require.Equal(t, PrintType(isActive), "bool")
	require.Equal(t, PrintType(complexNum), "complex64")
}

func Test_AllToStringAndConcatinate(t *testing.T) {

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
	require.Equal(t, AllToStringAndConcatinate(types), "4242423.14Golangtrue(1+2i)")
}

func Test_StringToRune(t *testing.T) {
	str := "4242423.14Golangtrue(1+2i)"
	r := []rune{
		52, 50, 52, 50, 52, 50, 51, 46, 49, 52,
		71, 111, 108, 97, 110, 103, 116, 114, 117, 101,
		40, 49, 43, 50, 105, 41,
	}

	require.Equal(t, StringToRune(str), r)
}

func Test_RuneToHash(t *testing.T) {
	r := []rune{
		52, 50, 52, 50, 52, 50, 51, 46, 49, 52,
		71, 111, 108, 97, 110, 103, 116, 114, 117, 101,
		40, 49, 43, 50, 105, 41,
	}

	b := [32]byte{
		83, 242, 246, 10, 198, 196, 19, 137,
		211, 237, 61, 132, 216, 141, 140, 40,
		96, 191, 137, 129, 198, 119, 190, 24,
		36, 58, 111, 53, 166, 182, 161, 179,
	}

	require.Equal(t, RuneToHash(r, "go-2024"), b)
}
