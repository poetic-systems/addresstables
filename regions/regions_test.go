package regions_test

import (
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/poetic-systems/addresstables/regions"
)

// TestAltCoversPrimaryAndShort enforces the invariant the package doc states:
// a row is reachable by its own Primary and its own Short through Alt alone.
//
// Consumers build their lookup maps from Alt, so a row that omits either
// spelling would be unfindable by the very forms the standard prints for it.
func TestAltCoversPrimaryAndShort(t *testing.T) {
	for r := range regions.All() {
		if !slices.Contains(r.Alt, r.Primary) {
			t.Errorf("%q: Alt does not contain its own Primary %q; a lookup built from Alt cannot find it", r.Primary, r.Primary)
		}
		if !slices.Contains(r.Alt, r.Short) {
			t.Errorf("%q: Alt does not contain its own Short %q; a lookup built from Alt cannot find it", r.Primary, r.Short)
		}
	}
}

// TestNoEmptyFields catches a row that lost a column in a move or an edit.
func TestNoEmptyFields(t *testing.T) {
	for r := range regions.All() {
		if r.Primary == "" || r.Short == "" || len(r.Alt) == 0 {
			t.Errorf("incomplete row: %+v", r)
		}
	}
}

// TestShortIsTwoLetters holds the one shape rule the standard states outright:
// every abbreviation in these three tables is two letters.
func TestShortIsTwoLetters(t *testing.T) {
	for r := range regions.All() {
		if len(r.Short) != 2 {
			t.Errorf("%q: Short %q is not two letters", r.Primary, r.Short)
		}
	}
}

// TestAltIsUnambiguous reports any spelling claimed by more than one row.
//
// A consumer keying a map by Alt gets last-write-wins on a collision, silently
// resolving a region to whichever row happened to come later in the table.
// The three tables are one namespace for this purpose because a consumer
// reading a last line does not know in advance which of them it is looking at.
func TestAltIsUnambiguous(t *testing.T) {
	owner := map[string]string{}
	for r := range regions.All() {
		for _, a := range r.Alt {
			if prev, ok := owner[a]; ok && prev != r.Primary {
				t.Errorf("%q is claimed by both %q and %q", a, prev, r.Primary)
			}
			owner[a] = r.Primary
		}
	}
}

// TestAllCoversEveryTable keeps All honest against the three group accessors.
func TestAllCoversEveryTable(t *testing.T) {
	grouped := 0
	for _, seq := range []func() iter.Seq[regions.Region]{
		regions.UnitedStates, regions.Canada, regions.Military,
	} {
		for range seq() {
			grouped++
		}
	}

	all := 0
	for range regions.All() {
		all++
	}

	if all != grouped {
		t.Errorf("All yielded %d rows, the three tables together yield %d", all, grouped)
	}
}

// TestRowsAreUppercase holds the module rule: Project US@ requires uppercase
// output, so the rows are stored uppercase and no consumer has to fold them.
func TestRowsAreUppercase(t *testing.T) {
	for r := range regions.All() {
		for _, s := range append([]string{r.Primary, r.Short}, r.Alt...) {
			if s != strings.ToUpper(s) {
				t.Errorf("%q: %q is not uppercase", r.Primary, s)
			}
		}
	}
}

// TestRowsAreTrimmed catches transcription whitespace, which ToUpper cannot
// see and which turns an exact-match lookup into a silent miss.
func TestRowsAreTrimmed(t *testing.T) {
	for r := range regions.All() {
		for _, s := range append([]string{r.Primary, r.Short}, r.Alt...) {
			if s != strings.TrimSpace(s) {
				t.Errorf("%q: %q has leading or trailing whitespace", r.Primary, s)
			}
		}
	}
}
