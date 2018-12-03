package main

import "testing"

func TestFindNearDups(t *testing.T) {
	expectedCommon := "fgij"
	expectedIndex := 2

	lines := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	common, index := findNearDups(lines)

	if !(index == expectedIndex && common == expectedCommon) {
		t.Error(
			"Expected:", expectedCommon, "at", expectedIndex,
			"got:", common, "at", index,
		)
	}
}
