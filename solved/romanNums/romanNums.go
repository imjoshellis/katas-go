package main

import (
	"fmt"
	"strings"

	match "github.com/alexpantyukhin/go-pattern-match"
)

type (
	Roman string
	Int   int
)

func main() {
	// r := Roman("MMXXIV")
	// fmt.Printf("%s to Int: %d\n", r, r.toInt())

	i := Int(2970)
	m := i.toRomanMath()
	s := i.toRomanReplace()
	q := i.toRomanBiQ()
	fmt.Printf("%d to RomanReplace: %s\n", i, s)
	fmt.Printf("%d to RomanBiQ: %s\n", i, q)
	fmt.Printf("%d to RomanMath: %s\n", i, m)
	fmt.Printf("%s to Int: %d\n", s, s.toInt())
}

func (r Roman) toInt() Int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for len(r) > 0 {
		if len(r) > 1 && m[r[0]] < m[r[1]] {
			sum -= m[r[0]]
		} else {
			sum += m[r[0]]
		}
		r = r[1:]
	}
	return Int(sum)
}

func (i Int) toRomanMath() Roman {
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	r := ""
	n := int(i)
	for _, k := range keys {
		for n >= k {
			r += m[k]
			n -= k
		}
	}
	return Roman(r)
}

func (i Int) toRomanReplace() Roman {
	r := ""
	for i > 0 {
		i -= 1
		r += "I"
	}

	r = strings.ReplaceAll(r, "IIIII", "V")
	r = strings.ReplaceAll(r, "VV", "X")
	r = strings.ReplaceAll(r, "XXXXX", "L")
	r = strings.ReplaceAll(r, "LL", "C")
	r = strings.ReplaceAll(r, "CCCCC", "D")
	r = strings.ReplaceAll(r, "DD", "M")
	r = strings.ReplaceAll(r, "IIII", "IV")
	r = strings.ReplaceAll(r, "VIV", "IX")
	r = strings.ReplaceAll(r, "XXXX", "XL")
	r = strings.ReplaceAll(r, "LXL", "XC")
	r = strings.ReplaceAll(r, "CCCC", "CD")
	r = strings.ReplaceAll(r, "DCD", "CM")

	return Roman(r)
}

type Triple [3]string

func (i Int) biQuinaryDigits(place int, t Triple) Roman {
	d := int(i) % (10 * place) / place
	_, res := match.Match(d).
		When(0, "").
		When(1, t[0]).
		When(2, t[0]+t[0]).
		When(3, t[0]+t[0]+t[0]).
		When(4, t[0]+t[1]).
		When(5, t[1]).
		When(6, t[1]+t[0]).
		When(7, t[1]+t[0]+t[0]).
		When(8, t[1]+t[0]+t[0]+t[0]).
		When(9, t[0]+t[2]).
		Result()

	return Roman(res.(string))
}

func (i Int) toRomanBiQ() Roman {
	u := i.biQuinaryDigits(1, Triple{"I", "V", "X"})
	t := i.biQuinaryDigits(10, Triple{"X", "L", "C"})
	h := i.biQuinaryDigits(100, Triple{"C", "D", "M"})
	m := i.biQuinaryDigits(1000, Triple{"M", "?", "?"})
	return m + h + t + u
}
