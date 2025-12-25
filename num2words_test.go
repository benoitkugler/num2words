package num2words

import (
	"testing"
)

func TestEurosToWords(t *testing.T) {
	for _, numWord := range []struct {
		num  int
		word string
	}{
		{num: 265_45, word: "deux cent soixante-cinq euros et quarante-cinq centimes"},
		{num: 81_25, word: "quatre-vingt-un euros et vingt-cinq centimes"},
		{num: 7_72, word: "sept euros et soixante-douze centimes"},
		{num: 1_00, word: "un euro et zéro centimes"},
		{num: 2_01, word: "deux euros et un centime"},
		{num: 8_10, word: "huit euros et dix centimes"},
		{num: 12_26, word: "douze euros et vingt-six centimes"},
		{num: 21_29, word: "vingt et un euros et vingt-neuf centimes"},
		{num: 81_25, word: "quatre-vingt-un euros et vingt-cinq centimes"},
		{num: 80_00, word: "quatre-vingts euros et zéro centimes"},
		{num: 71_00, word: "soixante et onze euros et zéro centimes"},
		{num: 90_00, word: "quatre-vingt-dix euros et zéro centimes"},
		{num: 200_00, word: "deux cents euros et zéro centimes"},
		{num: 1_000_00, word: "mille euros et zéro centimes"},
		{num: 2_000_00, word: "deux mille euros et zéro centimes"},
		{num: 2_000_000_00, word: "deux millions euros et zéro centimes"},
		{num: 77_20, word: "soixante-dix-sept euros et vingt centimes"},
		{num: 1000_00, word: "mille euros et zéro centimes"},
		{num: 100_00, word: "cent euros et zéro centimes"},
		{num: -100_00, word: "moins cent euros et zéro centimes"},
	} {
		if got := EurosToWords(numWord.num); got != numWord.word {
			t.Fatalf("expected %s, got %s", numWord.word, got)
		}
	}
}
