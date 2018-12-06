package gtin

import (
	"testing"
)

func TestCheckSum(t *testing.T) {
	code := "100123456789"
	check := 4
	result, err := CheckSum(code)
	if err != nil {
		t.Error(err)
	} else if result != check {
		t.Errorf("\nFor: CheckSum(%s)\nExpected: %d\nGot: %d\n", code, check, result)
	}
}
