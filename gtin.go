package gtingo

import (
	"errors"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var formats = map[int]string{
	8:  "GTIN-8",
	12: "GTIN-12",
	13: "GTIN-13",
	14: "GTIN-14",
}

type Gtin string

type Country struct {
	code string
	name string
}

func Calculate(number uint64) Gtin {
	s := strconv.FormatUint(number, 10)
	_, err := getFormat(len(s) + 1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	checkDigit := getCheckDigit(number)
	g := Gtin(s + strconv.Itoa(checkDigit))

	return g
}

func Generate(format int) Gtin {
	indicator := ""
	if format == 14 {
		indicator = strconv.Itoa(rand.Intn(8-1) + 1)
	}

	randomCountry := getRandomCountry()
	countryCode := randomCountry.getCode()
	randomNumber := getRandomNumber(format - len(countryCode) - len(indicator) - 1)

	i, err := strconv.ParseUint(countryCode+randomNumber, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	checkDigit := strconv.Itoa(getCheckDigit(i))
	g := Gtin(indicator + countryCode + randomNumber + checkDigit)

	return g
}

func (g Gtin) GetFormat() (string, error) {
	return getFormat(len(string(g)))
}

func (g Gtin) GetCountry() (Country, error) {
	format, err := g.GetFormat()
	if err != nil {
		log.Fatalln(err.Error())
	}

	countryCode := string(g)[:3]
	if format == formats[14] {
		countryCode = string(g)[1:4]
	}

	for _, c := range countries {
		match, _ := regexp.MatchString("\\d{3}-\\d{3}", c.Code())
		if match {
			rangeCode := strings.Split(c.Code(), "-")
			max, _ := strconv.Atoi(rangeCode[1])
			min, _ := strconv.Atoi(rangeCode[0])
			countryCodeInt, _ := strconv.Atoi(countryCode)
			if countryCodeInt <= max && countryCodeInt >= min {
				return c, nil
			}
		}
		if c.Code() == countryCode {
			return c, nil
		}
	}

	return Country{}, errors.New("the country is not defined. Code: " + countryCode)
}

func (c Country) Code() string {
	return c.code
}

func (c Country) Name() string {
	return c.name
}

func getFormat(length int) (string, error) {
	format, ok := formats[length]
	if !ok {
		return "", errors.New("the GTIN format is not defined")
	}

	return format, nil
}

func getCheckDigit(number uint64) int {
	numberStr := strconv.FormatUint(uint64(number), 10)
	var sum int
	for i := len(numberStr) - 1; i >= 0; i-- {
		intChar, _ := strconv.Atoi(string(numberStr[i]))
		if (len(numberStr)-i+1)%2 == 0 {
			sum += intChar * 3
		} else {
			sum += intChar
		}
	}
	sumStr := strconv.Itoa(sum)
	lastChar, _ := strconv.Atoi(sumStr[len(sumStr)-1:])
	if lastChar == 0 {
		return 0
	}
	return 10 - lastChar
}

func getRandomCountry() Country {
	rand.Seed(time.Now().UnixNano())
	country := countries[rand.Intn(len(countries)-1)]
	return country
}

func getRandomNumber(len int) string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < len; i++ {
		result += strconv.Itoa(rand.Intn(9))
	}
	return result
}

func (c Country) getCode() string {
	match, _ := regexp.MatchString("\\d{3}-\\d{3}", c.Code())
	if match {
		rangeCode := strings.Split(c.Code(), "-")
		max, _ := strconv.Atoi(rangeCode[1])
		min, _ := strconv.Atoi(rangeCode[0])
		code := strconv.Itoa(rand.Intn(max-min) + min)
		for len(code) < 3 {
			code = "0" + code
		}
		return code
	}
	return c.Code()
}
