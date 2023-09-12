package redacted_test

import (
	"fmt"

	"github.com/mech-pig/go-redacted"
)

func ExampleRedacted_Value() {
	value := struct {
		example string
	}{
		example: "test",
	}

	fmt.Printf("%+v", redacted.New(value).Value())
	// Output: {example:test}
}

func ExampleRedacted_String_formatterDefaultImplicit() {
	fmt.Print(redacted.New("test"))
	// Output: ****
}

func ExampleRedacted_String_formatterDefaultExplicit() {
	fmt.Printf("%v", redacted.New("test"))
	// Output: ****
}

func ExampleRedacted_String_formatterString() {
	fmt.Printf("%s", redacted.New("test"))
	// Output: ****
}

func ExampleRedacted_String_formatterStructFields() {
	fmt.Printf("%+v", redacted.New("test"))
	// Output: ****
}

func ExampleRedacted_String_formatterGoRepr() {
	fmt.Printf("%#v", redacted.New("test"))
	// Output: ****
}
