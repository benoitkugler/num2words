package num2words

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	negword     = "moins "
	GIGA_SUFFIX = "illiard"
	MEGA_SUFFIX = "illion"
)

var (
	cards     = map[int64]string{}
	keysCards []int64 // triée par ordre décroissant
	MAXVAL    int64
)

func init() {
	high_numwords := []string{"b", "m"}

	mid_numwords := []numWord{
		{num: 1000, word: "mille"},
		{num: 100, word: "cent"},
		{num: 80, word: "quatre-vingts"},
		{num: 60, word: "soixante"},
		{num: 50, word: "cinquante"},
		{num: 40, word: "quarante"},
		{num: 30, word: "trente"},
	}

	low_numwords := []string{"vingt", "dix-neuf", "dix-huit", "dix-sept",
		"seize", "quinze", "quatorze", "treize", "douze",
		"onze", "dix", "neuf", "huit", "sept", "six",
		"cinq", "quatre", "trois", "deux", "un", "zéro"}

	for _, numWord := range mid_numwords {
		cards[int64(numWord.num)] = numWord.word
	}

	for index, word := range low_numwords {
		cards[int64(len(low_numwords)-1-index)] = word
	}

	cap := 3 + 6*len(high_numwords)
	for index, word := range high_numwords {
		n := cap - index*6
		cards[int64(math.Pow10(n))] = word + GIGA_SUFFIX
		cards[int64(math.Pow10(n-3))] = word + MEGA_SUFFIX
	}
	for key := range cards {
		keysCards = append(keysCards, key)
	}
	sort.Slice(keysCards, func(i, j int) bool { return keysCards[i] > keysCards[j] })
	MAXVAL = keysCards[0]
}

type numWord struct {
	num  int
	word string
}

type textNum struct {
	text string
	num  int64
}

// si List vaut nil, représente un textNum.
type tNorList struct {
	tN   textNum
	list []tNorList
}

func merge(curr, next textNum) textNum {
	if curr.num == 1 {
		if next.num < 1000000 {
			return next
		}
	} else {
		if (((curr.num-80)%100 == 0) || (curr.num%100 == 0 && curr.num < 1000)) && next.num < 1000000 && curr.text[:len(curr.text)-1] == "s" {
			curr.text = curr.text[:len(curr.text)-1]
		}
		if curr.num < 1000 && next.num != 1000 && next.text[len(next.text)-1:] != "s" && next.num%100 == 0 {
			next.text += "s"
		}
	}
	if next.num < curr.num && curr.num < 100 {
		if next.num%10 == 1 && curr.num != 80 {
			return textNum{text: fmt.Sprintf("%s et %s", curr.text, next.text), num: curr.num + next.num}
		}
		return textNum{text: fmt.Sprintf("%s-%s", curr.text, next.text), num: curr.num + next.num}
	}
	if next.num > curr.num {
		return textNum{text: fmt.Sprintf("%s %s", curr.text, next.text), num: curr.num * next.num}
	}
	return textNum{text: fmt.Sprintf("%s %s", curr.text, next.text), num: curr.num + next.num}
}

func pluralizeCents(n int64) string {
	if n == 1 {
		return "centime"
	}
	return "centimes"
}

func pluralizeEuros(n int64) string {
	if n == 1 {
		return "euro"
	}
	return "euros"
}

func parse_currency_parts(value float64) (int64, int64, bool) {
	negative := value < 0
	inCents := int64(value * 100)
	cents := inCents % 100
	integer := (inCents - cents) / 100
	return integer, cents, negative
}

func splitnum(value int64) []tNorList {
	for _, elem := range keysCards {
		if elem > value {
			continue
		}
		var div, mod int64
		out := []tNorList{}
		if value == 0 {
			div, mod = 1, 0
		} else {
			div, mod = value/elem, value%elem
		}

		if div == 1 {
			out = append(out, tNorList{tN: textNum{text: cards[1], num: 1}})
		} else {
			if div == value { // The system tallies, eg Roman Numerals
				return []tNorList{{tN: textNum{text: strings.Repeat(cards[elem], int(div)), num: div * elem}}}
			}
			out = append(out, tNorList{list: splitnum(div)})
		}
		out = append(out, tNorList{tN: textNum{text: cards[elem], num: elem}})

		if mod != 0 {
			out = append(out, tNorList{list: splitnum(mod)})
		}
		return out
	}
	return nil
}

func clean(val []tNorList) tNorList {
	out := val
	for len(val) != 1 {
		out = []tNorList{}
		left, right := val[0], val[1]
		// if isinstance(left, tuple) and isinstance(right, tuple):
		if left.list == nil && right.list == nil {
			out = append(out, tNorList{tN: merge(left.tN, right.tN)})
			if len(val) > 2 {
				out = append(out, tNorList{list: val[2:]})
			}
		} else {
			for _, elem := range val {
				if elem.list != nil {
					if len(elem.list) == 1 {
						out = append(out, elem.list[0])
					} else {
						out = append(out, clean(elem.list))
					}
				} else {
					out = append(out, elem)
				}
			}
		}
		val = out
	}
	return out[0]
}

func to_cardinal(value int64) string {
	out := ""
	if value < 0 {
		value = int64(math.Abs(float64(value)))
		out = negword
	}

	if value >= MAXVAL { // fail silencieux
		return ""
	}

	val := splitnum(value)
	cleaned := clean(val)
	return out + cleaned.tN.text
}

// EurosToWords renvoi la somme indiqué, écrite avec des mots
func EurosToWords(val float64) string {
	separator := " et" //Cent separator

	left, right, is_negative := parse_currency_parts(val)

	minus_str := ""
	if is_negative {
		minus_str = fmt.Sprintf("%s ", negword)
	}
	cents_str := to_cardinal(right)

	return fmt.Sprintf("%s%s %s%s %s %s", minus_str, to_cardinal(left),
		pluralizeEuros(left), separator, cents_str, pluralizeCents(right))
}
