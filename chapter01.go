package main

import (
	"math"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
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
   Task 1.1
   This implementation does not use additional data structure.
   Requires strings module to operate runes for Unicode support.
   If there was no Unicode requirement I'd implement inner for loop.
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

/*
   Task 1.3
   Strings are immutable in Go, so I converted it to an array of runes,
   despite the fact of memory limit in the task.
   Space complexity: O(1) (if don't count array of runes).
   Time complexity: O(N²).
*/
func StringDuplicateCharactersCleanup(str string) string {
	s := []rune(str)
	l := len(s)
	arrayLen := l
	i := 0
	for {
		currentArrayLen := arrayLen - 1
		for j := i + 1; j <= currentArrayLen; j++ {
			if s[i] == s[j] {
				// Duplicate found
				// Now we need to move remaining subarray one position left
				for k := j; k < arrayLen-1; k++ {
					s[k] = s[k+1]
				}
				arrayLen--
			}
		}
		// Exit condition: we reached the end of array
		if i <= arrayLen {
			i++
		} else {
			break
		}
	}
	return string(s[0:arrayLen])
}

/*
   Task 1.4
*/
// Set of functions for runes slice sort
type RunesSlice []rune

func (rs RunesSlice) Len() int {
	return len(rs)
}
func (rs RunesSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}
func (rs RunesSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func SortRunesSlice(r []rune) []rune {
	a := RunesSlice(r)
	sort.Sort(a)
	return a
}

// Slices comparison
func AreRunesSlicesEqual(a, b []rune) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Characters filter
func RunesFilter(str []rune) []rune {
	result := make([]rune, 0)
	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, unicode.ToLower(r))
		}
	}
	return result
}

// And anagram function itself
func AreAnagrams(s1, s2 string) bool {
	if s1 == "" && s2 == "" {
		// An assumption that empty strings cannot be anagrams
		return false
	}
	a1 := SortRunesSlice(RunesFilter([]rune(s1)))
	a2 := SortRunesSlice(RunesFilter([]rune(s2)))
	return AreRunesSlicesEqual(a1, a2)
}

/*
   Task 1.6
*/
func AreMatricesEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i := 0; i < len(m1); i++ {
		if len(m1[i]) != len(m2[i]) {
			return false
		}
		for j := 0; j < len(m1[i]); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func MatrixRotateCopy(m1 [][]int) [][]int {
	if len(m1) == 0 {
		return [][]int{}
	}
	s := len(m1)
	m2 := make([][]int, s)
	for i := 0; i < s; i++ {
		m2[i] = make([]int, s)
	}
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			m2[j][s-1-i] = m1[i][j]
		}
	}
	return m2
}

func TranslatePosition(x, y, size int) (x1, y1 int) {
	x1 = y
	y1 = size - 1 - x
	return
}

func MatrixRotateInPlace(m [][]int) {
	size := len(m)
	half := int(math.Ceil(float64(size) / 2.0))

	for x0 := 0; x0 < half; x0++ {
		for y0 := x0; y0 < half; y0++ {
			x1, y1 := TranslatePosition(x0, y0, size)
			x2, y2 := TranslatePosition(x1, y1, size)
			x3, y3 := TranslatePosition(x2, y2, size)
			value := m[x0][y0]
			m[x0][y0] = m[x3][y3]
			m[x3][y3] = m[x2][y2]
			m[x2][y2] = m[x1][y1]
			m[x1][y1] = value
		}
	}
}

/*
   Task 1.7
*/
func NullLines(m [][]int) {
	size := len(m)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m[i][j] == 0 {
				for k := 0; k < size; k++ {
					m[i][k] = 0
					m[k][j] = 0
				}
			}
		}
	}
}

/*
   Task 1.8
*/
func IsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

func IsRotation(s1, s2 string) bool {
	s3 := s1 + s1
	return IsSubstring(s3, s2)
}
