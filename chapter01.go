package main

import "fmt"

/*
This implementation uses map to count all characters occurrences
Space complexity (assuming Unicode usage) in a worst case would be O(N)
but the maximum size of the map is 1111998 elements due to Unicode limitations.
Time complexity is O(N).
*/
func StringUniqueCharacters(str string) bool {
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
