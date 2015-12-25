package main

import (
	"strings"
	"unicode/utf8"
)

/*
   Chapter 1
   Task 1.1
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
   Chapter 1
   Task 1.1
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

/*
   Chapter 1
   Task 1.2
   Unfortunately, there are no C-style strings in Go (with null characters),
   so I implemented just a string reverse with Unicode support.
   This implementation uses "unicode/utf8" to get runes count in the string
   and it uses additional an array to build a result.
   Space complexity: O(N).
   Time complexity: O(N).
*/
func StringReverseAdditionalArray(str string) string {
	if len(str) == 0 {
		return ""
	}
	runeCount := utf8.RuneCountInString(str)
	newStrArr := make([]rune, runeCount)
	// Additional counter because range over string returns start position of code point, not the index of rune
	i := 0
	for _, r := range str {
		newStrArr[runeCount-i-1] = r
		i++
	}
	return string(newStrArr)
}

/*
   Chapter 1
   Task 1.2
   More elegant solution with array of runes and less memory consumption.
   Space complexity: O(N).
   Time complexity: O(N).
*/
func StringReverseRunes(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
