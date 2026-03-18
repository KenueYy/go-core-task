package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterSlice_BasicExample(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	got := FilterSlice(slice1, slice2)

	want := []string{"apple", "cherry", "43", "lead", "gno1"}

	require.ElementsMatch(t, want, got)
}

func TestFilterSlice_SecondEmpty(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{}

	got := FilterSlice(slice1, slice2)
	want := []string{"apple", "banana", "cherry"}

	require.Equal(t, want, got)
}

func TestFilterSlice_FirstEmpty(t *testing.T) {
	var slice1 []string
	slice2 := []string{"date", "43"}

	got := FilterSlice(slice1, slice2)
	require.Nil(t, got)
}

func TestFilterSlice_AllInSecond(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"apple", "banana", "cherry"}

	got := FilterSlice(slice1, slice2)
	require.Empty(t, got)
}

func TestFilterSlice_DuplicatesInFirst(t *testing.T) {
	slice1 := []string{"apple", "apple", "cherry"}
	slice2 := []string{"banana"}

	got := FilterSlice(slice1, slice2)
	want := []string{"apple", "apple", "cherry"}

	require.Equal(t, want, got)
}
