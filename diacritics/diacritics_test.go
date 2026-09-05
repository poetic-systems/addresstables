package diacritics_test

import (
	"testing"
	"unicode"

	"github.com/poetic-systems/addresstables/diacritics"
)

// TestNoEmptyFields catches a row that lost a column in a move or an edit.
func TestNoEmptyFields(t *testing.T) {
	for d := range diacritics.All() {
		if d.Rune == 0 || d.Substitute == "" || d.Description == "" {
			t.Errorf("incomplete row: %+v", d)
		}
	}
}

// TestSubstituteIsOneASCIILetter holds what Appendix A actually prints. The
// appendix substitutes a single letter for every character, including the ones
// a transliteration would spell with two — AE becomes a, thorn becomes p. A
// row here growing to "ae" would be a library's preference smuggled into the
// standard's table.
func TestSubstituteIsOneASCIILetter(t *testing.T) {
	for d := range diacritics.All() {
		if len(d.Substitute) != 1 || d.Substitute[0] < 'a' || d.Substitute[0] > 'z' {
			t.Errorf("%q (%U): Substitute %q is not a single lowercase ASCII letter", d.Rune, d.Rune, d.Substitute)
		}
	}
}

// TestRunesAreUnique reports a character listed twice. A consumer keying a map
// by Rune gets last-write-wins, so a duplicate would silently pick a winner.
func TestRunesAreUnique(t *testing.T) {
	seen := map[rune]string{}
	for d := range diacritics.All() {
		if prev, ok := seen[d.Rune]; ok {
			t.Errorf("%q (%U) is listed twice: %q and %q", d.Rune, d.Rune, prev, d.Description)
		}
		seen[d.Rune] = d.Description
	}
}

// TestRunesAreNonASCII states the table's reason to exist: every row is a
// character outside ASCII that an address may carry and Project US@ output may
// not. An ASCII row would be a no-op the standard does not print.
func TestRunesAreNonASCII(t *testing.T) {
	for d := range diacritics.All() {
		if d.Rune < unicode.MaxASCII {
			t.Errorf("%q (%U) is already ASCII; Appendix A lists no such row", d.Rune, d.Rune)
		}
	}
}

// TestRowsAreOrderedByCodePoint keeps the table in the order the appendix
// prints it, which is the order a reader checking it against the PDF reads in.
func TestRowsAreOrderedByCodePoint(t *testing.T) {
	var prev rune
	for d := range diacritics.All() {
		if d.Rune <= prev {
			t.Errorf("%q (%U) follows %U; Appendix A is ordered by code point", d.Rune, d.Rune, prev)
		}
		prev = d.Rune
	}
}

// TestLenMatchesAll keeps Len honest, since consumers may size a map with it.
func TestLenMatchesAll(t *testing.T) {
	count := 0
	for range diacritics.All() {
		count++
	}
	if count != diacritics.Len() {
		t.Errorf("All yielded %d rows, Len reports %d", count, diacritics.Len())
	}
}
