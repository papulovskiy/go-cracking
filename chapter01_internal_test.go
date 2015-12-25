package main

import "testing"

/*
	Chapter 1
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
	Chapter 1
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
	Chapter 1
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
