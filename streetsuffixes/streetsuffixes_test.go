package streetsuffixes_test

import (
	"slices"
	"testing"

	"github.com/poetic-systems/addresstables/streetsuffixes"
)

// TestAltCoversPrimaryAndShort enforces the invariant the package doc states:
// a row is reachable by its own Primary and its own Short through Alt alone.
//
// Consumers build their lookup maps from Alt, so a row that omits either
// spelling would be unfindable by the very forms the standard prints for it.
func TestAltCoversPrimaryAndShort(t *testing.T) {
	for s := range streetsuffixes.All() {
		if !slices.Contains(s.Alt, s.Primary) {
			t.Errorf("%q: Alt does not contain its own Primary %q; a lookup built from Alt cannot find it", s.Primary, s.Primary)
		}
		if !slices.Contains(s.Alt, s.Short) {
			t.Errorf("%q: Alt does not contain its own Short %q; a lookup built from Alt cannot find it", s.Primary, s.Short)
		}
	}
}

// TestNoEmptyFields catches a row that lost a column in a move or an edit.
func TestNoEmptyFields(t *testing.T) {
	for s := range streetsuffixes.All() {
		if s.Primary == "" || s.Short == "" || len(s.Alt) == 0 {
			t.Errorf("incomplete row: %+v", s)
		}
	}
}

// TestLenMatchesAll keeps Len honest, since consumers may size a map with it.
func TestLenMatchesAll(t *testing.T) {
	count := 0
	for range streetsuffixes.All() {
		count++
	}
	if count != streetsuffixes.Len() {
		t.Errorf("All yielded %d rows, Len reports %d", count, streetsuffixes.Len())
	}
}

// TestAltIsUnambiguous reports any spelling claimed by more than one row.
//
// A consumer keying a map by Alt gets last-write-wins on a collision, silently
// resolving a suffix to whichever row happened to come later in the table.
func TestAltIsUnambiguous(t *testing.T) {
	owner := map[string]string{}
	for s := range streetsuffixes.All() {
		for _, a := range s.Alt {
			if prev, ok := owner[a]; ok && prev != s.Primary {
				t.Errorf("%q is claimed by both %q and %q", a, prev, s.Primary)
			}
			owner[a] = s.Primary
		}
	}
}
