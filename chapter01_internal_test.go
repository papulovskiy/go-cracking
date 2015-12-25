package main

import "testing"

/*
	Task 1.1
*/
func testStringUniqueCharactersFunction(f func(string) bool, t *testing.T) {
	if f("abcdef") != true {
		t.Error("String contains unique characters")
	}
	if f("abcdefa") != false {
		t.Error("String does not contain unique characters")
	}
	if f("") != false {
		t.Error("String is empty")
	}
	if f("⌘") != true {
		t.Error("String contains unique characters")
	}
	if f("⌘⌘") != false {
		t.Error("String does not contain unique characters")
	}
}

func TestStringUniqueCharactersMap(t *testing.T) {
	testStringUniqueCharactersFunction(func(str string) bool { return StringUniqueCharactersMap(str) }, t)
}

func TestStringUniqueCharactersWithoutMap(t *testing.T) {
	testStringUniqueCharactersFunction(func(str string) bool { return StringUniqueCharacters(str) }, t)
}

/*
	Task 1.2
*/
func testStringReverse(f func(string) string, t *testing.T) {
	if f("") != "" {
		t.Error("String is empty")
	}
	if f("abc") != "cba" {
		t.Error("String was not reversed")
	}
	if f("⌘aabbcc⌘") != "⌘ccbbaa⌘" {
		t.Error("String was not reversed")
	}
}

func TestStringReverseAdditionalArray(t *testing.T) {
	testStringReverse(func(str string) string { return StringReverseAdditionalArray(str) }, t)
}
func TestStringReverseRunes(t *testing.T) {
	testStringReverse(func(str string) string { return StringReverseRunes(str) }, t)
}

/*
	Task 1.3
*/
func TestStringDuplicateCharactersCleanup(t *testing.T) {
	if StringDuplicateCharactersCleanup("abcdef") != "abcdef" {
		t.Error("String does not contain duplicate characters")
	}
	if StringDuplicateCharactersCleanup("abcdefa") != "abcdef" {
		t.Error("String contains duplicate characters")
	}
	if StringDuplicateCharactersCleanup("ab⌘cde⌘fa") != "ab⌘cdef" {
		t.Error("String contains duplicate characters")
	}
	if StringDuplicateCharactersCleanup("") != "" {
		t.Error("String is empty")
	}
	if StringDuplicateCharactersCleanup("aa") != "a" {
		t.Error("String contains duplicate characters")
	}
	if StringDuplicateCharactersCleanup("aaaaaa") != "a" {
		t.Error("String contains duplicate characters")
	}
}

/*
	Task 1.4
*/
func TestAreRunesSlicesEqual(t *testing.T) {
	if AreRunesSlicesEqual([]rune(""), []rune("")) != true {
		t.Error("Empty runes slices must be equal")
	}
	if AreRunesSlicesEqual([]rune("aoe"), []rune("")) != false {
		t.Error("Runes slices must not be equal")
	}
	if AreRunesSlicesEqual([]rune("aoe"), []rune("aoe")) != true {
		t.Error("Runes slices must be equal")
	}
}

func TestRunesFilter(t *testing.T) {
	str := "aB, aB ! cdEf"
	strFiltered := RunesFilter([]rune(str))
	if AreRunesSlicesEqual(strFiltered, []rune("ababcdef")) != true {
		t.Error("Runes were not filtered")
	}
	if AreRunesSlicesEqual(RunesFilter([]rune("abcdef")), []rune("abcdef")) != true {
		t.Error("Runes must be equal")
	}
	if AreRunesSlicesEqual(RunesFilter([]rune("")), []rune("")) != true {
		t.Error("Runes must be equal")
	}
	if AreRunesSlicesEqual(RunesFilter([]rune("abcdef!")), []rune("abcde")) != false {
		t.Error("Runes must not be equal")
	}
}

func TestRunesSliceSort(t *testing.T) {
	if AreRunesSlicesEqual(SortRunesSlice([]rune("cba")), []rune("abc")) != true {
		t.Error("Runes slice is not sorted")
	}
	if AreRunesSlicesEqual(SortRunesSlice([]rune("abc")), []rune("abc")) != true {
		t.Error("Runes slice is not sorted")
	}
}

func TestAreAnagrams(t *testing.T) {
	if AreAnagrams("abcdef", "bcdefa") != true {
		t.Error("Strings are anagrams")
	}
	if AreAnagrams("", "bcdefa") != false {
		t.Error("Strings are not anagrams")
	}
	if AreAnagrams("", "") != false {
		t.Error("Strings are not anagrams")
	}
}
