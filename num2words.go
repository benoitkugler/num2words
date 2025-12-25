package num2words

import (
	"fmt"
	"strings"
)

func pluralizeCents(n int) string {
	if n == 1 {
		return "centime"
	}
	return "centimes"
}

func pluralizeEuros(n int) string {
	if n == 1 {
		return "euro"
	}
	return "euros"
}

func parseCurrencyParts(value int) (int, int, bool) {
	negative := value < 0
	if negative {
		value = -value
	}
	return value / 100, value % 100, negative
}

func integerToTriplets(number int) []int {
	triplets := []int{}

	for number > 0 {
		triplets = append(triplets, number%1000)
		number = number / 1000
	}

	return triplets
}

// integerToFrFr converts a positive integer to French words
func integerToFrFr(input int) string {
	frenchMegas := [...]string{"", "mille", "million", "milliard", "billion", "billiard", "trillion", "trilliard", "quadrillion", "quadrilliard", "quintillion", "quintilliard"}
	frenchUnits := [...]string{"", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf"}
	frenchTens := [...]string{"", "dix", "vingt", "trente", "quarante", "cinquante", "soixante", "soixante", "quatre-vingt", "quatre-vingt"}
	frenchTeens := [...]string{"dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"}

	words := []string{}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "zéro"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// special cases
		if triplet == 1 && idx == 1 {
			words = append(words, "mille")
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		if hundreds > 0 {
			if hundreds == 1 {
				words = append(words, "cent")
			} else {
				if tens == 0 && units == 0 {
					words = append(words, frenchUnits[hundreds], "cents")
					goto tripletEnd
				} else {
					words = append(words, frenchUnits[hundreds], "cent")
				}
			}
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, frenchUnits[units])
		case 1:
			words = append(words, frenchTeens[units])
			break
		case 7:
			switch units {
			case 1:
				words = append(words, frenchTens[tens], "et", frenchTeens[units])
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchTeens[units])
				words = append(words, word)
				break
			}
			break
		case 8:
			switch units {
			case 0:
				words = append(words, frenchTens[tens]+"s")
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchUnits[units])
				words = append(words, word)
				break
			}
			break
		case 9:
			word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchTeens[units])
			words = append(words, word)
			break
		default:
			switch units {
			case 0:
				words = append(words, frenchTens[tens])
				break
			case 1:
				words = append(words, frenchTens[tens], "et", frenchUnits[units])
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchUnits[units])
				words = append(words, word)
				break
			}
			break
		}

	tripletEnd:
		// mega
		mega := frenchMegas[idx]
		if mega != "" {
			if mega != "mille" && triplet > 1 {
				mega += "s"
			}
			words = append(words, mega)
		}
	}

	return strings.Join(words, " ")
}

// EurosToWords renvoie la somme (indiquée en centimes), écrite avec des mots
func EurosToWords(val int) string {
	const separator = " et" // Cent separator
	const negword = "moins "

	left, right, isNegative := parseCurrencyParts(val)
	minusStr := ""
	if isNegative {
		minusStr = negword
	}
	cents := integerToFrFr(right)
	euros := integerToFrFr(left)
	return fmt.Sprintf("%s%s %s%s %s %s", minusStr, euros,
		pluralizeEuros(left), separator, cents, pluralizeCents(right))
}
