package num2words

import (
	"fmt"
	"testing"
)

var (
	TEST_CASES_TO_CURRENCY_EUR = []struct {
		num  float64
		word string
	}{
		{num: 1.00, word: "un euro et zéro centimes"},
		{num: 2.01, word: "deux euros et un centime"},
		{num: 8.10, word: "huit euros et dix centimes"},
		{num: 12.26, word: "douze euros et vingt-six centimes"},
		{num: 21.29, word: "vingt et un euros et vingt-neuf centimes"},
		{num: 81.25, word: "quatre-vingt-un euros et vingt-cinq centimes"},
		{num: 100.00, word: "cent euros et zéro centimes"},
	}
)

func TestBasic(t *testing.T) {
	// fmt.Println(EurosToWords(265.45))
	fmt.Println(EurosToWords(81.25))
}

func TestCurrency(t *testing.T) {
	for _, numWord := range TEST_CASES_TO_CURRENCY_EUR {
		if EurosToWords(numWord.num) != numWord.word {
			t.Fatalf(EurosToWords(numWord.num))
		}
	}
}
