package directionals_test

import (
	"testing"

	"github.com/poetic-systems/addresstables/directionals"
)

// TestSpanishRowsNameARealEnglishRow is the guard that makes English usable.
//
// A consumer emitting Project US@ output resolves a Spanish spelling to the
// English abbreviation by following English to its row. A Spanish row naming a
// word that is not in the table would send it nowhere.
func TestSpanishRowsNameARealEnglishRow(t *testing.T) {
	english := map[string]string{}
	for d := range directionals.English() {
		english[d.Full] = d.Short
	}

	for d := range directionals.Spanish() {
		if d.English == "" {
			t.Errorf("%q is Spanish but names no English equivalent", d.Full)
			continue
		}
		if _, ok := english[d.English]; !ok {
			t.Errorf("%q names English equivalent %q, which is not an English row", d.Full, d.English)
		}
	}
}

// TestEnglishRowsNameNoEquivalent keeps the direction of the pointer one way.
func TestEnglishRowsNameNoEquivalent(t *testing.T) {
	for d := range directionals.English() {
		if d.English != "" {
			t.Errorf("%q is English and should not name an equivalent, got %q", d.Full, d.English)
		}
	}
}

// TestOnlyWestFormsDisagree encodes Publication 28's own statement that "the
// only discrepancies between English and Spanish abbreviations occur in West
// directionals".
//
// It is here because that sentence is the whole reason a consumer can treat
// the two vocabularies as interchangeable on output. If a future edit breaks
// it, every consumer that assumed it needs to know.
func TestOnlyWestFormsDisagree(t *testing.T) {
	english := map[string]string{}
	for d := range directionals.English() {
		english[d.Full] = d.Short
	}

	disagree := map[string]bool{"OESTE": true, "NOROESTE": true, "SUDOESTE": true}

	for d := range directionals.Spanish() {
		same := d.Short == english[d.English]
		if disagree[d.Full] && same {
			t.Errorf("%q was expected to abbreviate differently from %q, both are %q", d.Full, d.English, d.Short)
		}
		if !disagree[d.Full] && !same {
			t.Errorf("%q abbreviates to %q but %q abbreviates to %q; only West forms should differ",
				d.Full, d.Short, d.English, english[d.English])
		}
	}
}

// TestEnglishAndSpanishPartitionAll makes sure neither filtered iterator drops
// a row or yields one twice.
func TestEnglishAndSpanishPartitionAll(t *testing.T) {
	var e, s int
	for range directionals.English() {
		e++
	}
	for range directionals.Spanish() {
		s++
	}
	if e+s != directionals.Len() {
		t.Errorf("English (%d) + Spanish (%d) = %d, want Len %d", e, s, e+s, directionals.Len())
	}
	if e != 8 || s != 8 {
		t.Errorf("want 8 English and 8 Spanish directionals, got %d and %d", e, s)
	}
}

// TestOIsSpanishWestOnly pins the specific collision hazard documented in the
// package: O is a directional only in Spanish, and there are Puerto Rico
// streets named O, so a consumer must never emit it.
func TestOIsSpanishWestOnly(t *testing.T) {
	for d := range directionals.English() {
		if d.Short == "O" {
			t.Errorf("%q abbreviates to O in the English table; O must be Spanish-only", d.Full)
		}
	}
}
