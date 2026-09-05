// Package secondaryunit holds the Publication 28 Appendix C2 secondary unit
// designators as data.
//
// A secondary unit designator names the part of an address below the street
// line — the apartment, the floor, the suite. Appendix C2 lists twenty-four of
// them, and marks with a footnote the nine that are not followed by a number,
// which is what Numbered records.
//
// That distinction is the reason this is a table rather than a list of words.
// Publication 28 says a designator that takes a number should carry one, and a
// designator that does not should stand alone; a consumer reading FL 3 and
// BSMT has to know which shape it is looking at, and the appendix is where
// that is written down.
//
// Two things a caller might expect are deliberately absent.
//
// The Project US@ exchange form writes an unknown or collapsed numbered unit
// as "#". That is a matching designator rather than an Appendix C2 row — the
// Postal Service does not list it — so it belongs to the consumer that decided
// to collapse the unit, not to this table.
//
// There is no lookup, matching or normalization here, for the reason the
// sibling packages give: the exact matching a caller wants is a decision about
// how an address is being read, and two callers reading for different purposes
// will make it differently.
//
// The Puerto Rico counterpart to this table is puertorico.Secondaries. Project
// US@ keeps the Spanish designators rather than translating them, so the two
// are separate vocabularies and not two spellings of one.
//
// Rows are uppercase because Project US@ requires uppercase output.
package secondaryunit

import (
	"iter"
	"slices"
)

// SecondaryUnit is one Appendix C2 designator.
//
// Numbered reports whether the designator is followed by a number. It is false
// for the nine the appendix footnotes — the ones that describe a position in a
// building rather than one of a series.
type SecondaryUnit struct {
	Full     string
	Short    string
	Numbered bool
}

var secondaryUnits = []SecondaryUnit{
	{Full: "APARTMENT", Short: "APT", Numbered: true},
	{Full: "BASEMENT", Short: "BSMT"},
	{Full: "BUILDING", Short: "BLDG", Numbered: true},
	{Full: "DEPARTMENT", Short: "DEPT", Numbered: true},
	{Full: "FLOOR", Short: "FL", Numbered: true},
	{Full: "FRONT", Short: "FRNT"},
	{Full: "HANGER", Short: "HNGR", Numbered: true},
	{Full: "KEY", Short: "KEY", Numbered: true},
	{Full: "LOBBY", Short: "LBBY"},
	{Full: "LOT", Short: "LOT", Numbered: true},
	{Full: "LOWER", Short: "LOWR"},
	{Full: "OFFICE", Short: "OFC"},
	{Full: "PENTHOUSE", Short: "PH"},
	{Full: "PIER", Short: "PIER", Numbered: true},
	{Full: "REAR", Short: "REAR"},
	{Full: "ROOM", Short: "RM", Numbered: true},
	{Full: "SIDE", Short: "SIDE"},
	{Full: "SLIP", Short: "SLIP", Numbered: true},
	{Full: "SPACE", Short: "SPC", Numbered: true},
	{Full: "STOP", Short: "STOP", Numbered: true},
	{Full: "SUITE", Short: "STE", Numbered: true},
	{Full: "TRAILER", Short: "TRLR", Numbered: true},
	{Full: "UNIT", Short: "UNIT", Numbered: true},
	{Full: "UPPER", Short: "UPPR"},
}

// All yields every designator, in the order Appendix C2 lists them.
func All() iter.Seq[SecondaryUnit] {
	return slices.Values(secondaryUnits)
}

// Len reports how many rows All will yield.
func Len() int {
	return len(secondaryUnits)
}
