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

// TestStreetTypesIncludesTheP26OnlyRows pins the ten rows #17 adds from the
// p. 26 prefix list that Appendix E does not carry, so a future edit cannot
// silently drop one back out.
func TestStreetTypesIncludesTheP26OnlyRows(t *testing.T) {
	want := map[string]puertorico.StreetType{
		"BOULEVARD": {Full: "BOULEVARD", Short: "BLVD", English: "BOULEVARD"},
		"CALETA":    {Full: "CALETA", Short: "CALETA", English: "COVE"},
		"CALLEJON":  {Full: "CALLEJON", Short: "CALLEJON", English: "ALLEY"},
		"CARRETERA": {Full: "CARRETERA", Short: "CARR", English: "HIGHWAY"},
		"MARGINAL":  {Full: "MARGINAL", Short: "MARGINAL", English: "FRONTAGE ROAD"},
		"PARQUE":    {Full: "PARQUE", Short: "PARQ", English: "PARK"},
		"PASAJE":    {Full: "PASAJE", Short: "PASAJE", English: "PASSAGE"},
		"PATIO":     {Full: "PATIO", Short: "PATIO", English: "COURTYARD"},
		"PLAZA":     {Full: "PLAZA", Short: "PLAZA", English: "PLAZA"},
		"VIA":       {Full: "VIA", Short: "VIA", English: "WAY"},
	}
	got := map[string]puertorico.StreetType{}
	for s := range puertorico.StreetTypes() {
		got[s.Full] = s
	}
	for full, w := range want {
		if got[full] != w {
			t.Errorf("StreetTypes()[%q] = %+v, want %+v", full, got[full], w)
		}
	}
}

// TestPARQUEAndCARRETERAResolveInBothTables holds the two-role precedent the
// package doc establishes: CARRETERA/CARR is also a Secondary and PARQUE/PARQ
// is also a StandaloneUrbanization, and both abbreviations must agree with
// their StreetTypes row rather than drift from it.
func TestPARQUEAndCARRETERAResolveInBothTables(t *testing.T) {
	streetTypeShort := map[string]string{}
	for s := range puertorico.StreetTypes() {
		streetTypeShort[s.Full] = s.Short
	}
	secondaryShort := map[string]string{}
	for s := range puertorico.Secondaries() {
		secondaryShort[s.Full] = s.Short
	}
	standaloneShort := map[string]string{}
	for u := range puertorico.StandaloneUrbanizations() {
		standaloneShort[u.Full] = u.Short
	}

	if got, want := streetTypeShort["CARRETERA"], "CARR"; got != want {
		t.Errorf("StreetTypes()[CARRETERA] = %q, want %q", got, want)
	}
	if got, want := secondaryShort["CARRETERA"], "CARR"; got != want {
		t.Errorf("Secondaries()[CARRETERA] = %q, want %q", got, want)
	}

	if got, want := streetTypeShort["PARQUE"], "PARQ"; got != want {
		t.Errorf("StreetTypes()[PARQUE] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["PARQUE"], "PARQ"; got != want {
		t.Errorf("StandaloneUrbanizations()[PARQUE] = %q, want %q", got, want)
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

func TestStandaloneUrbanizationsComplete(t *testing.T) {
	seen := map[string]bool{}
	for u := range puertorico.StandaloneUrbanizations() {
		if u.Full == "" || u.Short == "" {
			t.Errorf("incomplete standalone urbanization: %+v", u)
		}
		if seen[u.Full] {
			t.Errorf("%q appears more than once in StandaloneUrbanizations", u.Full)
		}
		seen[u.Full] = true
	}
	if len(seen) == 0 {
		t.Error("StandaloneUrbanizations yielded nothing")
	}
}

// TestStandaloneUrbanizationsNoAbbreviationCollision holds the same
// no-collision property the other tables assert: two different standalone
// urbanization names should not claim the same abbreviation, because a
// consumer building a Short->Full lookup would lose one silently.
func TestStandaloneUrbanizationsNoAbbreviationCollision(t *testing.T) {
	owner := map[string]string{}
	for u := range puertorico.StandaloneUrbanizations() {
		if prior, ok := owner[u.Short]; ok {
			t.Errorf("abbreviation %q is claimed by both %q and %q", u.Short, prior, u.Full)
		}
		owner[u.Short] = u.Full
	}
}

// TestStandaloneUrbanizationsPluralExpansion pins the X(S) expansion
// decision: the standard's "Altura(s)" -> "ALT(S)" notation is stored as two
// ordinary rows, singular and plural, rather than as literal parentheses.
func TestStandaloneUrbanizationsPluralExpansion(t *testing.T) {
	want := map[string]string{
		"ALTURA":   "ALT",
		"ALTURAS":  "ALTS",
		"BRISA":    "BRISA",
		"BRISAS":   "BRISAS",
		"COLINA":   "COLINA",
		"COLINAS":  "COLINAS",
		"LOMA":     "LOMA",
		"LOMAS":    "LOMAS",
		"PARCELA":  "PARCELA",
		"PARCELAS": "PARCELAS",
		"VILLA":    "VILLA",
		"VILLAS":   "VILLAS",
		"VISTA":    "VISTA",
		"VISTAS":   "VISTAS",
	}
	got := map[string]string{}
	for u := range puertorico.StandaloneUrbanizations() {
		got[u.Full] = u.Short
	}
	for full, short := range want {
		if got[full] != short {
			t.Errorf("StandaloneUrbanizations()[%q] = %q, want %q", full, got[full], short)
		}
	}
	for _, full := range []string{"ALT(S)", "BRISA(S)", "COLINA(S)", "LOMA(S)", "PARCELA(S)", "VILLA(S)", "VISTA(S)"} {
		if _, ok := got[full]; ok {
			t.Errorf("StandaloneUrbanizations() stores the literal parenthesized notation %q; it must be expanded", full)
		}
	}
}

// TestStandaloneUrbanizationAbbreviationsAreRoleDependent pins the three
// places where this table's abbreviation for a word legitimately differs
// from the same word's abbreviation elsewhere in the package. The standard
// abbreviates these words differently depending on whether they are acting
// as a secondary designator or as a standalone urbanization name; a future
// edit must not "fix" one table to match the other.
func TestStandaloneUrbanizationAbbreviationsAreRoleDependent(t *testing.T) {
	secondaryShort := map[string]string{}
	for s := range puertorico.Secondaries() {
		secondaryShort[s.Full] = s.Short
	}
	streetTypeShort := map[string]string{}
	for s := range puertorico.StreetTypes() {
		streetTypeShort[s.Full] = s.Short
	}
	standaloneShort := map[string]string{}
	for u := range puertorico.StandaloneUrbanizations() {
		standaloneShort[u.Full] = u.Short
	}

	if got, want := secondaryShort["SECTOR"], "SEC"; got != want {
		t.Errorf("Secondaries()[SECTOR] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["SECTOR"], "SECT"; got != want {
		t.Errorf("StandaloneUrbanizations()[SECTOR] = %q, want %q", got, want)
	}

	if got, want := secondaryShort["PARCELAS"], "PARC"; got != want {
		t.Errorf("Secondaries()[PARCELAS] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["PARCELAS"], "PARCELAS"; got != want {
		t.Errorf("StandaloneUrbanizations()[PARCELAS] = %q, want %q", got, want)
	}

	if got, want := secondaryShort["VILLA"], "VIL"; got != want {
		t.Errorf("Secondaries()[VILLA] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["VILLA"], "VILLA"; got != want {
		t.Errorf("StandaloneUrbanizations()[VILLA] = %q, want %q", got, want)
	}

	if got, want := streetTypeShort["PASEO"], "PSO"; got != want {
		t.Errorf("StreetTypes()[PASEO] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["PASEO"], "PASEO"; got != want {
		t.Errorf("StandaloneUrbanizations()[PASEO] = %q, want %q", got, want)
	}

	if got, want := streetTypeShort["VISTA"], "VIS"; got != want {
		t.Errorf("StreetTypes()[VISTA] = %q, want %q", got, want)
	}
	if got, want := standaloneShort["VISTA"], "VISTA"; got != want {
		t.Errorf("StandaloneUrbanizations()[VISTA] = %q, want %q", got, want)
	}
}
