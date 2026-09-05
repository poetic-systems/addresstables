package secondaryunit_test

import (
	"strings"
	"testing"

	"github.com/poetic-systems/addresstables/secondaryunit"
)

// TestRowCount pins Appendix C2 at twenty-four designators. A row appearing or
// disappearing is a change to the standard, not a refactor, and should have to
// be stated.
func TestRowCount(t *testing.T) {
	if got := secondaryunit.Len(); got != 24 {
		t.Errorf("Len = %d, want 24", got)
	}
}

// TestUnnumberedRowsAreTheFootnotedNine is the invariant that makes Numbered
// worth carrying.
//
// Appendix C2 footnotes exactly nine designators as not followed by a number.
// A consumer deciding whether to expect a number after BSMT is trusting this
// list, so a row silently changing sides has to fail here.
func TestUnnumberedRowsAreTheFootnotedNine(t *testing.T) {
	footnoted := map[string]bool{
		"BSMT": true, "FRNT": true, "LBBY": true, "LOWR": true, "OFC": true,
		"PH": true, "REAR": true, "SIDE": true, "UPPR": true,
	}

	seen := 0
	for u := range secondaryunit.All() {
		if footnoted[u.Short] {
			seen++
			if u.Numbered {
				t.Errorf("%q (%s) is footnoted as taking no number but is Numbered", u.Full, u.Short)
			}
			continue
		}
		if !u.Numbered {
			t.Errorf("%q (%s) is not footnoted but is not Numbered", u.Full, u.Short)
		}
	}

	if seen != len(footnoted) {
		t.Errorf("found %d of the %d footnoted designators", seen, len(footnoted))
	}
}

// TestRowsAreUniqueAndUppercase guards the two properties a consumer builds
// its maps on: it can key by either field without collisions, and it does not
// have to fold case first.
func TestRowsAreUniqueAndUppercase(t *testing.T) {
	fulls := map[string]bool{}
	shorts := map[string]bool{}

	for u := range secondaryunit.All() {
		if u.Full == "" || u.Short == "" {
			t.Errorf("row %+v has an empty field", u)
		}
		if u.Full != strings.ToUpper(u.Full) || u.Short != strings.ToUpper(u.Short) {
			t.Errorf("row %+v is not uppercase", u)
		}
		if fulls[u.Full] {
			t.Errorf("%q appears twice as a Full", u.Full)
		}
		if shorts[u.Short] {
			t.Errorf("%q appears twice as a Short", u.Short)
		}
		fulls[u.Full] = true
		shorts[u.Short] = true
	}
}

// TestHashIsNotARow keeps the Project US@ matching designator out of the
// table. "#" is what a consumer writes when it has collapsed a numbered unit;
// the Postal Service does not list it, and a caller iterating Appendix C2 must
// not be handed it as though it did.
func TestHashIsNotARow(t *testing.T) {
	for u := range secondaryunit.All() {
		if u.Full == "#" || u.Short == "#" {
			t.Errorf("# appears as a row (%+v); it is an exchange form, not an Appendix C2 designator", u)
		}
	}
}
