package fizzbuzz_test

import (
	"testing"

	"github.com/kapomman191/fizzbuzz"
)
func TestFizzBuzzShouldSayOne(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if result =/= expected {
		t.Error("It should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if result =/= expected {
		t.Error("It should say %q but get %q", expected, result)
	}
}

