package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/Hashmath-19/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.Fail()
	}
	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v should  have returned true", 3)
		t.Fail()
	}
	if result != "fizz" {
		t.Logf("the result shou've been Fizz")
		t.Fail()
	}

}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	//Output: fizz
}
