package redacted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {
	value := "test"
	redacted := New(value)
	require.Equal(t, value, redacted.Value())
}

func TestStringer(t *testing.T) {
	redacted := New("test")
	redactedString := "****"
	args := []struct {
		name  string
		value any
	}{
		{"Value", redacted},
		{"Pointer", &redacted},
	}

	for _, a := range args {
		t.Run(a.name, func(t *testing.T) {
			formatted := fmt.Sprint(a.value)
			require.Equal(t, redactedString, formatted)
		})
	}
}

func TestFmt(t *testing.T) {
	redacted := New("test")
	formats := []string{"%v", "%+v", "%#v"}
	args := []struct {
		name  string
		value any
	}{
		{"Value", redacted},
		{"Pointer", &redacted},
	}

	for _, a := range args {
		t.Run(a.name, func(t *testing.T) {
			for _, f := range formats {
				t.Run(f, func(t *testing.T) {
					formatted := fmt.Sprintf(f, a.value)
					require.Equal(t, "****", formatted)
				})
			}
		})
	}
}
