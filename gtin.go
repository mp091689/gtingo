package gtingo

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/MykytaPopov/gtingo/internal/country"
)

const Gtin8 = 8
const Gtin12 = 12
const Gtin13 = 13
const Gtin14 = 14

var formats = map[int]string{
	Gtin8:  "GTIN-8",
	Gtin12: "GTIN-12",
	Gtin13: "GTIN-13",
	Gtin14: "GTIN-14",
}

type Gtin struct{}

func NewGtin() Gtin {
	rand.Seed(time.Now().UnixNano())

	return Gtin{}
}

func (g *Gtin) Generate(format int) (string, error) {
	_, err := getFormat(format)
	if err != nil {
		return "", err
	}

	gtin := make(Number, format)

	loadIndicator(gtin)
	loadCountryCode(gtin)
	loadBody(gtin)
	loadCheckSum(gtin)

	return gtin.Stringify(), nil
}

func (g Gtin) Calculate(input string) (string, error) {
	input += "0"

	format := len(input)

	_, err := getFormat(format)
	if err != nil {
		return "", err
	}

	gtin := NewNumber(input)

	loadCheckSum(gtin)

	return gtin.Stringify(), nil
}

func (g Gtin) Validate(number string) bool {
	_, err := getFormat(len(number))
	if err != nil {
		return false
	}

	n := NewNumber(number)

	loadCheckSum(n)

	return n.Stringify() == number
}

// Package type: range 0 - 8
func loadIndicator(g Number) {
	if len(g) == 14 {
		g[0] = rand.Intn(8)
	}
}

func loadCountryCode(g Number) {
	key := 0
	if len(g) == 14 {
		key = 1
	}

	randomCountry := country.NewCountry()
	countryCode := randomCountry.GetCode()

	for i, rune := range countryCode {
		g[key+i], _ = strconv.Atoi(string(rune))
	}
}

func loadBody(g Number) {
	startIdx := 3
	if len(g) == 14 {
		startIdx = 4
	}

	for i := startIdx; i < len(g)-1; i++ {
		g[i] = rand.Intn(9)
	}
}

func loadCheckSum(g Number) {
	sum := 0
	multiply := true

	for i := len(g) - 1; i > 0; i-- {
		if multiply {
			sum += g[i-1] * 3
		} else {
			sum += g[i-1]
		}

		multiply = !multiply
	}

	sumIncrement := sum
	for {
		if sumIncrement%10 == 0 {
			g[len(g)-1] = sumIncrement - sum
			break
		}
		sumIncrement++
	}
}

func getFormat(length int) (string, error) {
	format, ok := formats[length]
	if !ok {
		return "", errors.New("the GTIN format is not defined")
	}

	return format, nil
}
