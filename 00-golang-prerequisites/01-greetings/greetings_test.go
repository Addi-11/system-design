package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Glady"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Glady")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Glady") = %q. %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
