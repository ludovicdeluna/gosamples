package docme

import (
	"testing"
)

type showFail struct {
	t *testing.T
}

func (s showFail) Println(expected, result interface{}) {
	s.t.Errorf("\ngot:\n%s\nwant:\n%s\n", result, expected)
}

// This exemple show how use Hello() to print a greeting, and it also used
// as unit test by go test -v
// Be aware that the latest comment is used by test as expected string in return
func ExampleHello() {
	fmt.Println(Hello("Ludovic"))
	// Output: Hello, my name is Ludovic
}

// This test is used as unit test
func TestHello(t *testing.T) {
	failed := showFail{t}
	name := "Ludovic"
	expected := "Hello, my name is " + name
	result := Hello(name)
	if expected != result {
		failed.Println(expected, result)
	}
}
