package recursion

import "fmt"

var (
	mapping = map[rune][]rune{
		'0': {'0'},
		'1': {'1'},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

// PhoneNumberMnemonics ...
//
// T -> O(4 ** (n * n))where n is the length of the phone number.
// S -> O(4 ** (n * n)) where n is the length of the phone number.
func PhoneNumberMnemonics(phoneNumber string) []string {
	var mnemonics []string
	if len(phoneNumber) == 0 {
		return mnemonics
	}

	head, tail := phoneNumber[0], phoneNumber[1:]

	// Seed mnemonics.
	opts := mapping[rune(head)]
	for _, o := range opts {
		mnemonics = append(mnemonics, string(o))
	}

	// Permutate the rest.
	for _, c := range tail {
		mnemonics = permutate(mnemonics, mapping[c])
	}

	return mnemonics
}

func permutate(mnemonics []string, opts []rune) []string {
	if len(opts) == 0 {
		return []string{}
	}

	var out []string
	for _, m := range mnemonics {
		out = append(out, fmt.Sprintf("%s%s", m, string(opts[0])))
	}

	return append(out, permutate(mnemonics, opts[1:])...)
}
