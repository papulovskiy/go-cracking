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

/*
	Task 1.6
*/
func TestAreMatricesEqual(t *testing.T) {
	m1 := [][]int{[]int{1, 2}, []int{3, 4}}
	m2 := [][]int{[]int{1, 2}, []int{3, 4}}
	if AreMatricesEqual(m1, m2) != true {
		t.Error("Matrices must be equal")
	}
	m2 = [][]int{[]int{1, 3}, []int{2, 4}}
	if AreMatricesEqual(m1, m2) != false {
		t.Error("Matrices must not be equal")
	}
	m2 = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	if AreMatricesEqual(m1, m2) != false {
		t.Error("Matrices must not be equal")
	}
}

func TestMatrixRotate(t *testing.T) {
	m1 := [][]int{[]int{1, 2}, []int{3, 4}}
	m2 := MatrixRotateCopy(m1)
	m3 := [][]int{[]int{3, 1}, []int{4, 2}}
	if AreMatricesEqual(m1, m2) != false {
		t.Error("Matrices must not be equal")
	}
	if AreMatricesEqual(m2, m3) != true {
		t.Error("Matrices must be equal")
	}

	m5 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	m6 := [][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}
	if AreMatricesEqual(MatrixRotateCopy(m5), m6) != true {
		t.Error("Matrices must be equal")
	}

	size := 4
	value := 0
	m7 := make([][]int, size)
	for i := 0; i < size; i++ {
		m7[i] = make([]int, size)
		for j := 0; j < size; j++ {
			m7[i][j] = value
			value++
		}
	}
	if AreMatricesEqual(MatrixRotateCopy(MatrixRotateCopy(MatrixRotateCopy(MatrixRotateCopy(m7)))), m7) != true {
		t.Error("Matrices must be equal")
	}
}

func TestTranslatePosition(t *testing.T) {
	x1, y1 := TranslatePosition(0, 0, 2)
	if x1 != 0 || y1 != 1 {
		t.Error("Incorrect position translation")
	}
	x1, y1 = TranslatePosition(0, 1, 2)
	if x1 != 1 || y1 != 1 {
		t.Error("Incorrect position translation")
	}
	x1, y1 = TranslatePosition(1, 0, 2)
	if x1 != 0 || y1 != 0 {
		t.Error("Incorrect position translation")
	}
	x1, y1 = TranslatePosition(1, 1, 2)
	if x1 != 1 || y1 != 0 {
		t.Error("Incorrect position translation")
	}

	x1, y1 = TranslatePosition(0, 0, 4)
	if x1 != 0 || y1 != 3 {
		t.Error("Incorrect position translation")
	}
	x1, y1 = TranslatePosition(2, 2, 4)
	if x1 != 2 || y1 != 1 {
		t.Error("Incorrect position translation")
	}
}

func TestMatrixRotateInPlace(t *testing.T) {
	m1 := [][]int{[]int{1, 2}, []int{3, 4}}
	m2 := MatrixRotateCopy(m1)
	MatrixRotateInPlace(m1)
	if AreMatricesEqual(m1, m2) != true {
		t.Error("Matrix was incorrectly rotated")
	}

	m1 = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	m2 = [][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}
	MatrixRotateInPlace(m1)
	if AreMatricesEqual(m1, m2) != true {
		t.Error("Matrix was incorrectly rotated")
	}
}

/*
	Task 1.7
*/
func TestNullLines(t *testing.T) {
	m1 := [][]int{[]int{0, 2}, []int{3, 4}}
	NullLines(m1)
	if m1[0][0] != 0 || m1[0][1] != 0 || m1[1][0] != 0 {
		t.Error("Not all values set to 0")
	}

	m1 = [][]int{[]int{7, 4, 1}, []int{8, 0, 2}, []int{9, 6, 3}}
	NullLines(m1)
	if m1[0][1] != 0 || m1[1][0] != 0 || m1[1][2] != 0 || m1[2][2] != 0 {
		t.Error("Not all values set to 0")
	}

}
