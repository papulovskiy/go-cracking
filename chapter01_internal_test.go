package main

import "testing"

func TestStringUniqueCharacters(t *testing.T) {
    if StringUniqueCharacters("abcdef") != true {
        t.Error("String contains unique characters")
    }
    if StringUniqueCharacters("abcdefa") != false {
        t.Error("String does not contain unique characters")
    }
    if StringUniqueCharacters("") != false {
        t.Error("String is empty")
    }
    if StringUniqueCharacters("⌘") != true {
        t.Error("String contains unique characters")
    }
    if StringUniqueCharacters("⌘⌘") != false {
        t.Error("String does not contain unique characters")
    }
}
