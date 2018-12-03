package main

import "testing"

type testPair struct {
	id      string
	results []bool
}

var tests = []testPair{
	{"abcdef", []bool{false, false}},
	{"bababc", []bool{true, true}},
	{"abbcde", []bool{true, false}},
	{"abcccd", []bool{false, true}},
}

func TestCheckID(t *testing.T) {
	for _, test := range tests {
		hasDouble, hasTriple := checkID(test.id)

		if !(hasDouble == test.results[0] && hasTriple == test.results[1]) {
			t.Error(
				"For", test.id,
				"expected", test.results,
				"got", []bool{hasDouble, hasTriple},
			)
		}
	}
}
