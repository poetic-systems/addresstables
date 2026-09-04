package puertorico_test

import (
	"testing"

	"github.com/poetic-systems/addresstables/puertorico"
)

func TestStreetTypesComplete(t *testing.T) {
	seen := map[string]bool{}
	for s := range puertorico.StreetTypes() {
		if s.Full == "" || s.Short == "" {
			t.Errorf("incomplete street type: %+v", s)
		}
		if seen[s.Short] {
			t.Errorf("abbreviation %q is claimed by more than one street type", s.Short)
		}
		seen[s.Short] = true
	}
}

func TestSecondariesComplete(t *testing.T) {
	seen := map[string]bool{}
	for s := range puertorico.Secondaries() {
		if s.Full == "" || s.Short == "" {
			t.Errorf("incomplete secondary: %+v", s)
		}
		if seen[s.Short] {
			t.Errorf("abbreviation %q is claimed by more than one secondary", s.Short)
		}
		seen[s.Short] = true
	}
}

// TestUrbanizationIsNotASecondary holds the separation the package doc
// describes: URB has its own line in the standard and its own table here, and
// listing it in both would make it two things at once.
func TestUrbanizationIsNotASecondary(t *testing.T) {
	for s := range puertorico.Secondaries() {
		if s.Short == "URB" {
			t.Errorf("%q appears in Secondaries with short URB; urbanization belongs only in Urbanizations", s.Full)
		}
	}
}

func TestUrbanizationsAllAbbreviateToURB(t *testing.T) {
	count := 0
	for u := range puertorico.Urbanizations() {
		count++
		if u.Short != "URB" {
			t.Errorf("%q abbreviates to %q, want URB", u.Full, u.Short)
		}
	}
	if count == 0 {
		t.Error("Urbanizations yielded nothing")
	}
}
