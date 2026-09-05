package qualifiers_test

import (
	"strings"
	"testing"

	"github.com/poetic-systems/addresstables/ustigerfile/qualifiers"
)

// TestAppendixCIsWhole counts Appendix C's seventeen rows.
//
// A code the table does not have is not an error a consumer can see: a reader
// of PREQUAL or SUFQUAL that finds no row treats the field as empty and drops
// the qualifier out of the name silently.
func TestAppendixCIsWhole(t *testing.T) {
	if qualifiers.Len() != 17 {
		t.Errorf("Appendix C has 17 rows, qualifiers.Len() is %d", qualifiers.Len())
	}
}

// TestCodesAreUnique guards the map every consumer will build from All.
func TestCodesAreUnique(t *testing.T) {
	seen := map[string]string{}
	for q := range qualifiers.All() {
		if first, ok := seen[q.Code]; ok {
			t.Errorf("code %q is on both %q and %q", q.Code, first, q.Full)
			continue
		}
		seen[q.Code] = q.Full
	}
}

// TestEveryRowIsPermittedSomewhere encodes the appendix's own shape: a
// qualifier the Census Bureau allows in neither position could never appear in
// a record, so a row with both false is a transcription error.
func TestEveryRowIsPermittedSomewhere(t *testing.T) {
	for q := range qualifiers.All() {
		if !q.Prefix && !q.Suffix {
			t.Errorf("%q is permitted as neither a prefix nor a suffix", q.Full)
		}
	}
}

// TestRowsAreUppercase holds the module-wide rule. These rows are keys in an
// index a consumer builds, and a title-case row would miss every lookup.
func TestRowsAreUppercase(t *testing.T) {
	for q := range qualifiers.All() {
		for _, word := range []string{q.Full, q.Short} {
			if word != strings.ToUpper(word) {
				t.Errorf("%q is not uppercase", word)
			}
		}
	}
}
