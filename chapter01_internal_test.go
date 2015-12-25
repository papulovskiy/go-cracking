package main

import "testing"

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

func TestStringReverse(t *testing.T) {
	if StringReverse("") != "" {
		t.Error("String is empty")
	}
	if StringReverse("abc") != "cba" {
		t.Error("String was not reversed")
	}
	if StringReverse("⌘aabbcc⌘") != "⌘ccbbaa⌘" {
		t.Error("String was not reversed")
	}
}
