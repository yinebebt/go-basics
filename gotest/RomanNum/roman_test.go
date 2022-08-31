package RomanNum

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"3 gets converted to II", 3, "III"},
		{"5 gets converted to V", 5, "V"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

//Refactor and try to come up with a better implementation
//to consider for all case; see @

//strings.builder
//quick.check -- from testing
