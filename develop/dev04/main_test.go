package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnagramDict(t *testing.T) {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}
	result := AnagramDict(input)
	require.Equal(t, result, map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}})
	fmt.Println(result)
}

func TestUniqLower(t *testing.T) {
	input := []string{"тест", "тест", "листок", "пятак", "пятак"}
	result := UniqLower(input)
	require.Equal(t, result, []string{"тест", "листок", "пятак"})
}
