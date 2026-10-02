package businesswords_test

import (
	"testing"

	"github.com/poetic-systems/addresstables/businesswords"
)

// TestNoEmptyFields catches a row that lost a column in a move or an edit.
func TestNoEmptyFields(t *testing.T) {
	for w := range businesswords.All() {
		if w.Primary == "" || w.Short == "" || len(w.Alt) == 0 {
			t.Errorf("incomplete row: %+v", w)
		}
	}
}

// TestLenMatchesAll keeps Len honest, since consumers may size a map with it.
func TestLenMatchesAll(t *testing.T) {
	count := 0
	for range businesswords.All() {
		count++
	}
	if count != businesswords.Len() {
		t.Errorf("All yielded %d rows, Len reports %d", count, businesswords.Len())
	}
}

// TestNoDuplicatePrimary enforces the one uniqueness guarantee this table
// makes. Unlike streetsuffixes, Appendix G does not promise Alt is
// unambiguous across rows (see the package doc), so this checks Primary only.
func TestNoDuplicatePrimary(t *testing.T) {
	seen := map[string]bool{}
	for w := range businesswords.All() {
		if seen[w.Primary] {
			t.Errorf("%q appears as Primary in more than one row", w.Primary)
		}
		seen[w.Primary] = true
	}
}

// TestAltContainsPrimary enforces the one promise Alt does make: every row's
// Alt includes its own Primary, so a consumer can always find a row by it.
func TestAltContainsPrimary(t *testing.T) {
	for w := range businesswords.All() {
		found := false
		for _, a := range w.Alt {
			if a == w.Primary {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%q: Alt does not contain its own Primary", w.Primary)
		}
	}
}
