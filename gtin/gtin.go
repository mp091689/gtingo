// Package gtin implement functionality to calculate checksum for GTIN13(EAN) code or generate it randomly.
package gtin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type country struct {
	Code string
	Name string
}

// CheckSum calculates the checksum for passed part of gtin code.
func CheckSum(s string) (int, error) {
	if len(s) != 12 {
		return 0, fmt.Errorf("expected the len of argument 12 chars, got %d", len(s))
	}
	even := 0
	odd := 0
	for i := len(s) - 1; i >= 0; i-- {
		intChar, _ := strconv.Atoi(string(s[i]))
		if i%2 == 0 {
			even += intChar
		} else {
			odd += intChar
		}
	}
	sum := strconv.Itoa((odd * 3) + even)
	lastChar, _ := strconv.Atoi(string(sum[len(sum)-1]))
	if lastChar == 0 {
		return 0, nil
	}
	return 10 - lastChar, nil
}

// Generate generates GTIN13(EAN) code randomly with calculated checksum.
func Generate() string {
	countryCode, _ := randomCountry()
	randomNums := randomNumbers(9)
	sum, _ := CheckSum(countryCode + randomNums)
	checkSum := strconv.Itoa(sum)
	return countryCode + randomNums + checkSum
}

// randomCountry returns randomly selected struct of country within countryCode and countryName.
func randomCountry() (string, string) {
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/gtin/country_codes.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteCountries, _ := ioutil.ReadAll(jsonFile)

	var countries []country
	_ = json.Unmarshal(byteCountries, &countries)

	rand.Seed(time.Now().UnixNano())
	country := countries[rand.Intn(len(countries)-1)]
	match, _ := regexp.MatchString("\\d{3}-\\d{3}", country.Code)
	if match {
		rangeCode := strings.Split(country.Code, "-")
		max, _ := strconv.Atoi(rangeCode[1])
		min, _ := strconv.Atoi(rangeCode[0])
		code := strconv.Itoa(rand.Intn(max-min) + min)
		for len(code) < 3 {
			code = "0" + code
		}
		return code, country.Name
	}
	return country.Code, country.Name
}

// randomNumber generates string of specified length within randomly generated numbers.
func randomNumbers(len int) string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < len; i++ {
		result += strconv.Itoa(rand.Intn(9))
	}
	return result
}
