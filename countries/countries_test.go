package countries_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/poetic-systems/addresstables/countries"
)

// TestAltCoversPrimary enforces the invariant the package doc states: a row is
// reachable by its own Primary through Alt alone, since consumers build their
// lookups from Alt.
func TestAltCoversPrimary(t *testing.T) {
	for c := range countries.All() {
		if !slices.Contains(c.Alt, c.Primary) {
			t.Errorf("%q: Alt does not contain its own Primary; a lookup built from Alt cannot find it", c.Primary)
		}
	}
}

// TestNoEmptyFields catches a row that lost a column in a move or an edit.
func TestNoEmptyFields(t *testing.T) {
	for c := range countries.All() {
		if c.Primary == "" || len(c.Alt) == 0 {
			t.Errorf("incomplete row: %+v", c)
		}
	}
}

// TestAltIsUnambiguous reports any spelling claimed by more than one row. A
// consumer keying a map by Alt gets last-write-wins on a collision.
func TestAltIsUnambiguous(t *testing.T) {
	owner := map[string]string{}
	for c := range countries.All() {
		for _, a := range c.Alt {
			if prev, ok := owner[a]; ok && prev != c.Primary {
				t.Errorf("%q is claimed by both %q and %q", a, prev, c.Primary)
			}
			owner[a] = c.Primary
		}
	}
}

// TestRowsAreUppercase holds the module rule: Project US@ requires uppercase
// output, so the rows are stored uppercase and no consumer has to fold them.
func TestRowsAreUppercase(t *testing.T) {
	for c := range countries.All() {
		for _, s := range append([]string{c.Primary}, c.Alt...) {
			if s != strings.ToUpper(s) {
				t.Errorf("%q: %q is not uppercase", c.Primary, s)
			}
		}
	}
}

// TestRowsAreTrimmed catches transcription whitespace, which ToUpper cannot
// see and which turns an exact-match lookup into a silent miss.
func TestRowsAreTrimmed(t *testing.T) {
	for c := range countries.All() {
		for _, s := range append([]string{c.Primary}, c.Alt...) {
			if s != strings.TrimSpace(s) {
				t.Errorf("%q: %q has leading or trailing whitespace", c.Primary, s)
			}
		}
	}
}
