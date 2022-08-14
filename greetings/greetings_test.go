package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.hello with a name, checking
// for a valid return.
func TestHelloName(t *testing.T) {
	name := "Kiana"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Kiana")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Kiana") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("") 
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
