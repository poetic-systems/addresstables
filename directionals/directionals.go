// Package directionals holds the Publication 28 directional words as data.
//
// Both the English set and the Spanish set are here. Project US@ does not
// mention the Spanish forms; Publication 28 lists them at
// https://pe.usps.com/text/pub28/28c2_044.htm, with the note that
//
//	Directionals are not commonly used in Puerto Rico addresses because other
//	descriptions, such as urbanization, identify geographic areas. ... The only
//	discrepancies between English and Spanish abbreviations occur in West
//	directionals. In the ZIP+4 file, the English equivalents are used.
//
// That note is the reason English is the row a Spanish row points at rather
// than the other way round. Six of the eight Spanish words abbreviate to the
// same letters as their English counterparts; only the three containing West
// differ — OESTE is O, NOROESTE is NO, SUDOESTE is SO. Since USPS's own ZIP+4
// file carries the English equivalents, a consumer producing Project US@
// output should accept a Spanish spelling on input and emit the English
// abbreviation, which is what English makes possible.
//
// Emitting O is a hazard rather than a nicety. There are streets in Puerto
// Rico whose name is the single letter O, so a leading O read as a directional
// turns a real address into a directional with no street name after it. The
// same shape as reading E ST as EAST ST.
//
// Rows are uppercase because Project US@ requires uppercase output. A consumer
// that wants another case should fold it; this package does not choose one for
// anybody.
package directionals

import (
	"iter"
	"slices"
)

// Directional is one directional word and its Publication 28 abbreviation.
//
// English is set only on the Spanish rows, where it names the Full of the
// English row that carries the abbreviation USPS files actually use. On an
// English row it is empty — the row is already the English one.
type Directional struct {
	Full    string
	Short   string
	Spanish bool
	English string
}

var directionals = []Directional{
	{Full: "NORTH", Short: "N"},
	{Full: "SOUTH", Short: "S"},
	{Full: "EAST", Short: "E"},
	{Full: "WEST", Short: "W"},
	{Full: "NORTHEAST", Short: "NE"},
	{Full: "SOUTHEAST", Short: "SE"},
	{Full: "NORTHWEST", Short: "NW"},
	{Full: "SOUTHWEST", Short: "SW"},

	{Full: "NORTE", Short: "N", Spanish: true, English: "NORTH"},
	{Full: "SUR", Short: "S", Spanish: true, English: "SOUTH"},
	{Full: "ESTE", Short: "E", Spanish: true, English: "EAST"},
	{Full: "OESTE", Short: "O", Spanish: true, English: "WEST"},
	{Full: "NORESTE", Short: "NE", Spanish: true, English: "NORTHEAST"},
	{Full: "SUDESTE", Short: "SE", Spanish: true, English: "SOUTHEAST"},
	{Full: "NOROESTE", Short: "NO", Spanish: true, English: "NORTHWEST"},
	{Full: "SUDOESTE", Short: "SO", Spanish: true, English: "SOUTHWEST"},
}

// All yields every directional, English rows first.
func All() iter.Seq[Directional] {
	return slices.Values(directionals)
}

// English yields only the eight English directionals.
func English() iter.Seq[Directional] {
	return func(yield func(Directional) bool) {
		for _, d := range directionals {
			if d.Spanish {
				continue
			}
			if !yield(d) {
				return
			}
		}
	}
}

// Spanish yields only the eight Spanish directionals.
func Spanish() iter.Seq[Directional] {
	return func(yield func(Directional) bool) {
		for _, d := range directionals {
			if !d.Spanish {
				continue
			}
			if !yield(d) {
				return
			}
		}
	}
}

// Len reports how many rows All will yield.
func Len() int {
	return len(directionals)
}
