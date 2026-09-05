// Package countries holds the country names a Project US@ address may carry.
//
// This is the one table here that is not a transcription. Project US@ gives no
// guidance for validating or normalizing an international country name beyond
// requiring that it be spelled out and capitalized, so there is no appendix to
// check these rows against. What there is instead is the standard's scope: it
// addresses US addresses and the Canadian ones the Postal Service accepts. So
// the vocabulary is the three countries that scope names, and a consumer that
// meets any other country passes the name through rather than looking it up.
//
// Domestic marks the rows Project US@ says to omit. A US address does not
// print a country line, so "UNITED STATES" is recognized in order to be
// dropped, not in order to be emitted. That distinction is a fact about the
// row; turning it into an empty output string is the consumer's business, and
// go-projectusat's NormalizeCountry is where it happens.
//
// A two-letter abbreviation is exact only within this vocabulary. CA is
// unambiguously Canada here and unambiguously California in regions, and both
// tables are right. Resolving that is a parser's job, which is why neither
// table tries.
package countries

import (
	"iter"
	"slices"
)

// Country is one row: the spelled-out name Project US@ wants printed, and
// every spelling that names it.
//
// Alt carries the primary name as well, so a consumer can build one lookup
// from Alt alone. TestAltCoversPrimary enforces that.
type Country struct {
	Primary  string
	Alt      []string
	Domestic bool
}

var countries = []Country{
	{
		Primary:  "UNITED STATES",
		Alt:      []string{"UNITED STATES", "USA", "US"},
		Domestic: true,
	},
	{
		Primary: "CANADA",
		Alt:     []string{"CANADA", "CA"},
	},
	{
		Primary: "MEXICO",
		Alt:     []string{"MEXICO", "MX"},
	},
}

// All yields every country.
func All() iter.Seq[Country] {
	return slices.Values(countries)
}
