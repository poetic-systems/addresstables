package featuretypes_test

import (
	"strings"
	"testing"

	"github.com/poetic-systems/addresstables/ustigerfile/featuretypes"
)

// TestAppendixDIsWhole counts Appendix D's 503 rows.
//
// A missing row is invisible downstream rather than loud: a reader of PRETYP
// or SUFTYP that finds no match leaves the type out of the assembled name, so
// the street becomes MARTIN LUTHER KING instead of MARTIN LUTHER KING BLVD and
// every index built from it misses.
func TestAppendixDIsWhole(t *testing.T) {
	if featuretypes.Len() != 503 {
		t.Errorf("Appendix D has 503 rows, featuretypes.Len() is %d", featuretypes.Len())
	}
}

// TestCodesAreUnique guards the map every consumer will build from All.
func TestCodesAreUnique(t *testing.T) {
	seen := map[string]string{}
	for f := range featuretypes.All() {
		if first, ok := seen[f.Code]; ok {
			t.Errorf("code %q is on both %q and %q", f.Code, first, f.Full)
			continue
		}
		seen[f.Code] = f.Full
	}
}

// TestEveryRowIsPermittedSomewhere: a type allowed in neither position could
// never appear in a record, so a row with both false is a transcription error.
func TestEveryRowIsPermittedSomewhere(t *testing.T) {
	for f := range featuretypes.All() {
		if !f.Prefix && !f.Suffix {
			t.Errorf("%q is permitted as neither a prefix nor a suffix", f.Full)
		}
	}
}

// TestOnlySpanishRowsAreTranslated keeps Translation meaning one thing. The
// appendix prints N/A on every English row, and a consumer branching on
// Spanish must be able to trust that the field is empty there.
func TestOnlySpanishRowsAreTranslated(t *testing.T) {
	for f := range featuretypes.All() {
		if f.Spanish && f.Translation == "" {
			t.Errorf("%q is Spanish and carries no translation", f.Full)
		}
		if !f.Spanish && f.Translation != "" {
			t.Errorf("%q is not Spanish but carries translation %q", f.Full, f.Translation)
		}
	}
}

// TestRowsAreUppercase holds the module-wide rule. These rows are keys in an
// index a consumer builds, and a title-case row would miss every lookup.
func TestRowsAreUppercase(t *testing.T) {
	for f := range featuretypes.All() {
		for _, word := range []string{f.Full, f.Short, f.Translation} {
			if word != strings.ToUpper(word) {
				t.Errorf("%q is not uppercase", word)
			}
		}
	}
}
