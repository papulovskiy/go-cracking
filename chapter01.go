package main

import (
	"strings"
)

/*
This implementation uses map to count all characters occurrences.
Space complexity (assuming Unicode usage) in a worst case would be O(N)
but the maximum size of the map is 1111998 elements due to Unicode limitations.
Time complexity is O(N).
*/
func StringUniqueCharactersMap(str string) bool {
	if len(str) == 0 {
		// String is empty, so it does not contain any characters
		return false
	}
	charMap := make(map[rune]int)
	for _, c := range str {
		charMap[c] += 1
	}
	for _, count := range charMap {
		if count > 1 {
			return false
		}
	}
	return true
}

/*
This implementation does not use additional data structure.
Requires strings module to operate runes for Unicode support.
Space complexity O(1).
Time complexity O(N²).
*/
func StringUniqueCharacters(str string) bool {
	if len(str) == 0 {
		// String is empty, so it does not contain any characters
		return false
	}
	for index, c := range str {
		if index != strings.LastIndexFunc(str, func(r rune) bool { return r == c }) {
			return false
		}
	}
	return true
}
