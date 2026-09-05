// Package qualifiers holds Appendix C of the Census Bureau's TIGER/Line
// technical documentation as data.
//
// A qualifier is the word a road carries alongside its name rather than as
// part of it — BUSINESS, BYPASS, SPUR, OLD. TIGER writes it as a code in
// PREQUAL and SUFQUAL, apart from the base name, which is the reason this
// table exists at all: a parser reading the assembled name has to decide
// whether BUSINESS in US HIGHWAY 41 BUSINESS is a street name or a qualifier,
// and the Census Bureau has already answered that in the record.
//
// Publication 28 has no equivalent table. It abbreviates BYPASS as BYP in the
// suffix list and says nothing about the other sixteen, so a consumer wanting
// to know that CONNECTOR and OVERPASS are qualifiers has only this.
//
// Rows are uppercase, following the rest of this module. The published table
// prints them in title case.
//
// Transcribed from
// https://www2.census.gov/geo/pdfs/maps-data/data/tiger/tgrshp2025/TGRSHP2025_TechDoc_C.pdf
// (TIGER/Line Shapefiles, 2025, Appendix C: Feature Name Qualifiers).
package qualifiers

import (
	"iter"
	"slices"
)

// Qualifier is one row of Appendix C.
//
// Code is the value TIGER writes into PREQUAL and SUFQUAL. Prefix and Suffix
// report which of the two the Census Bureau permits the qualifier in, and no
// row permits neither.
type Qualifier struct {
	Code   string
	Full   string
	Short  string
	Prefix bool
	Suffix bool
}

var qualifiers = []Qualifier{
	{Code: "11", Full: "ACCESS", Short: "ACC", Suffix: true},
	{Code: "12", Full: "ALTERNATE", Short: "ALT", Prefix: true, Suffix: true},
	{Code: "13", Full: "BUSINESS", Short: "BUS", Prefix: true, Suffix: true},
	{Code: "14", Full: "BYPASS", Short: "BYP", Prefix: true, Suffix: true},
	{Code: "15", Full: "CONNECTOR", Short: "CON", Suffix: true},
	{Code: "16", Full: "EXTENDED", Short: "EXD", Prefix: true, Suffix: true},
	{Code: "17", Full: "EXTENSION", Short: "EXN", Suffix: true},
	{Code: "18", Full: "HISTORIC", Short: "HST", Prefix: true},
	{Code: "19", Full: "LOOP", Short: "LP", Prefix: true, Suffix: true},
	{Code: "20", Full: "OLD", Short: "OLD", Prefix: true},
	{Code: "21", Full: "PRIVATE", Short: "PVT", Prefix: true, Suffix: true},
	{Code: "22", Full: "PUBLIC", Short: "PUB", Prefix: true, Suffix: true},
	{Code: "23", Full: "SCENIC", Short: "SCN", Suffix: true},
	{Code: "24", Full: "SPUR", Short: "SPR", Prefix: true, Suffix: true},
	{Code: "25", Full: "RAMP", Short: "RMP", Suffix: true},
	{Code: "26", Full: "UNDERPASS", Short: "UNP", Suffix: true},
	{Code: "27", Full: "OVERPASS", Short: "OVP", Suffix: true},
}

// All yields every qualifier, in the order the appendix prints them.
func All() iter.Seq[Qualifier] {
	return slices.Values(qualifiers)
}

// Len reports how many rows All will yield.
func Len() int {
	return len(qualifiers)
}
