// Package cityabbreviations holds the words that city names are commonly
// written abbreviated with, and the spelling Publication 28 requires.
//
// Publication 28 §223 says
//
//	Spell city names in their entirety. When abbreviations must be used due
//	to labeling constraints, use only the approved 13–character abbreviations
//	provided in the City State file.
//
// and Project US@ (p. 20) says the same: city names SHALL be spelled out in
// their entirety. Neither prints a table of the words that get abbreviated in
// the first place. In the data both consumers read — GeoNames postal names,
// TIGER place names, addresses as people type them — the words that actually
// turn up abbreviated are four: SAINT, SAINTE, MOUNT and FORT. ST ALBANS,
// STE GENEVIEVE, MT VERNON, FT WORTH. This is that table, so both libraries
// spell them out the same way and a city key built on one side is the key the
// other side looks up.
//
// The rows are whole words. ST. with a period is the same word after
// punctuation is dropped (Publication 28 §222, Project US@ p. 15), which a
// consumer does first; it is not a separate row here. Where a word stands in
// a name is the consumer's rule too: ST is SAINT at the head of ST LOUIS and
// nothing at the end of it.
//
// Rows are uppercase because Project US@ requires uppercase output. A consumer
// that wants another case should fold it; this package does not choose one for
// anybody.
package cityabbreviations

import (
	"iter"
	"slices"
)

// Abbreviation is one word a city name is commonly abbreviated with, and the
// spelling Publication 28 §223 requires in its place.
type Abbreviation struct {
	Full  string
	Short string
}

var abbreviations = []Abbreviation{
	{Full: "SAINT", Short: "ST"},
	{Full: "SAINTE", Short: "STE"},
	{Full: "MOUNT", Short: "MT"},
	{Full: "FORT", Short: "FT"},
}

// All yields every abbreviation.
func All() iter.Seq[Abbreviation] {
	return slices.Values(abbreviations)
}

// Len reports how many rows All will yield.
func Len() int {
	return len(abbreviations)
}
