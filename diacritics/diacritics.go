// Package diacritics holds Project US@ Appendix A, the diacritic mapping
// guidance, as data.
//
// The standard prints one row per character: the character, its decimal code,
// the ASCII letter to substitute for it, its Unicode code point, and a
// description. Rune carries the character, and the decimal code and code point
// are the same fact stated three ways, so only the rune is here.
//
// This is the table, not a transliteration. Substitute is what Appendix A
// prints, which is a single ASCII letter for every row — the appendix maps ae
// to a and thorn to p. A library may reasonably prefer ae and th, and
// go-projectusat has a Transliterate that does, but that is a choice made on
// top of the standard rather than a reading of it, so it does not live here.
//
// # Case
//
// Every other table in this module is uppercase, because Project US@ requires
// uppercase output. This one is not, and cannot be: case is data here. The
// appendix gives capital A with grave and small a with grave separate rows,
// and folding the substitutions to uppercase would state something the
// appendix does not. Rows are exactly as printed; a consumer producing
// Project US@ output uppercases its result, as it does everywhere else.
package diacritics

import (
	"iter"
	"slices"
)

// Diacritic is one row of Appendix A: the character, the ASCII letter the
// standard substitutes for it, and the description the appendix prints.
type Diacritic struct {
	Rune        rune
	Substitute  string
	Description string
}

var diacritics = []Diacritic{
	{Rune: 'À', Substitute: "a", Description: "Capital letter A with grave accent"},
	{Rune: 'Á', Substitute: "a", Description: "Capital letter A with acute accent"},
	{Rune: 'Â', Substitute: "a", Description: "Capital letter A with circumflex accent"},
	{Rune: 'Ã', Substitute: "a", Description: "Capital letter A with tilde"},
	{Rune: 'Ä', Substitute: "a", Description: "Capital letter A with dieresis or umlaut mark"},
	{Rune: 'Å', Substitute: "a", Description: "Capital letter A with ring above"},
	{Rune: 'Æ', Substitute: "a", Description: "Capital letter AE diphthong"},
	{Rune: 'Ç', Substitute: "c", Description: "Capital letter C with cedilla"},
	{Rune: 'È', Substitute: "e", Description: "Capital letter E with grave accent"},
	{Rune: 'É', Substitute: "e", Description: "Capital letter E with acute accent"},
	{Rune: 'Ê', Substitute: "e", Description: "Capital letter E with circumflex accent"},
	{Rune: 'Ë', Substitute: "e", Description: "Capital letter E with dieresis or umlaut mark"},
	{Rune: 'Ì', Substitute: "i", Description: "Capital letter I with grave accent"},
	{Rune: 'Í', Substitute: "i", Description: "Capital letter I with acute accent"},
	{Rune: 'Î', Substitute: "i", Description: "Capital letter I with circumflex"},
	{Rune: 'Ï', Substitute: "i", Description: "Capital letter I with dieresis or umlaut mark"},
	{Rune: 'Ð', Substitute: "e", Description: "Capital letter ETH (Icelandic)"},
	{Rune: 'Ñ', Substitute: "n", Description: "Capital letter N with tilde"},
	{Rune: 'Ò', Substitute: "o", Description: "Capital letter O with grave accent"},
	{Rune: 'Ó', Substitute: "o", Description: "Capital letter O with acute accent"},
	{Rune: 'Ô', Substitute: "o", Description: "Capital letter O with circumflex"},
	{Rune: 'Õ', Substitute: "o", Description: "Capital letter O with tilde"},
	{Rune: 'Ö', Substitute: "o", Description: "Capital letter O with dieresis or umlaut mark"},
	{Rune: 'Ø', Substitute: "o", Description: "Capital letter O with slash"},
	{Rune: 'Ù', Substitute: "u", Description: "Capital letter U with grave accent"},
	{Rune: 'Ú', Substitute: "u", Description: "Capital letter U with acute accent"},
	{Rune: 'Û', Substitute: "u", Description: "Capital letter U with circumflex"},
	{Rune: 'Ü', Substitute: "u", Description: "Capital letter U with dieresis or umlaut mark"},
	{Rune: 'Ý', Substitute: "y", Description: "Capital letter Y with acute accent"},
	{Rune: 'Þ', Substitute: "p", Description: "Capital letter THORN"},
	{Rune: 'ß', Substitute: "s", Description: "Small letter sharp s - ess-zed"},
	{Rune: 'à', Substitute: "a", Description: "Small letter a with grave accent"},
	{Rune: 'á', Substitute: "a", Description: "Small letter a with acute accent"},
	{Rune: 'â', Substitute: "a", Description: "Small letter a with circumflex"},
	{Rune: 'ã', Substitute: "a", Description: "Small letter a with tilde"},
	{Rune: 'ä', Substitute: "a", Description: "Small letter a with dieresis or umlaut mark"},
	{Rune: 'å', Substitute: "a", Description: "Small letter a with ring above"},
	{Rune: 'æ', Substitute: "a", Description: "Small letter ae"},
	{Rune: 'ç', Substitute: "c", Description: "Small letter c with cedilla"},
	{Rune: 'è', Substitute: "e", Description: "Small letter e with grave accent"},
	{Rune: 'é', Substitute: "e", Description: "Small letter e with acute accent"},
	{Rune: 'ê', Substitute: "e", Description: "Small letter e with circumflex"},
	{Rune: 'ë', Substitute: "e", Description: "Small letter e with dieresis"},
	{Rune: 'ì', Substitute: "i", Description: "Small letter i with grave accent"},
	{Rune: 'í', Substitute: "i", Description: "Small letter i with acute accent"},
	{Rune: 'î', Substitute: "i", Description: "Small letter i with circumflex"},
	{Rune: 'ï', Substitute: "i", Description: "Small letter i with diaresis"},
	{Rune: 'ð', Substitute: "e", Description: "Small letter eth"},
	{Rune: 'ñ', Substitute: "n", Description: "Small letter n with tilde"},
	{Rune: 'ò', Substitute: "o", Description: "Small letter o with grave accent"},
	{Rune: 'ó', Substitute: "o", Description: "Small letter o with acute accent"},
	{Rune: 'ô', Substitute: "o", Description: "Small letter o with circumflex"},
	{Rune: 'õ', Substitute: "o", Description: "Small letter o with tilde"},
	{Rune: 'ö', Substitute: "o", Description: "Small letter o with dieresis"},
	{Rune: 'ø', Substitute: "o", Description: "Small letter o with slash"},
	{Rune: 'ù', Substitute: "u", Description: "Small letter u with grave accent"},
	{Rune: 'ú', Substitute: "u", Description: "Small letter u with acute accent"},
	{Rune: 'û', Substitute: "u", Description: "Small letter u with circumflex"},
	{Rune: 'ü', Substitute: "u", Description: "Small letter u with dieresis"},
	{Rune: 'ý', Substitute: "y", Description: "Small letter y with acute accent"},
	{Rune: 'þ', Substitute: "p", Description: "Small letter thorn"},
	{Rune: 'ÿ', Substitute: "y", Description: "Small letter y with dieresis"},
	{Rune: 'Œ', Substitute: "o", Description: "Capital letter OE"},
	{Rune: 'œ', Substitute: "o", Description: "Small letter oe"},
	{Rune: 'Š', Substitute: "s", Description: "Capital letter S with caron"},
	{Rune: 'š', Substitute: "s", Description: "Small letter s with caron"},
	{Rune: 'Ÿ', Substitute: "y", Description: "Capital letter Y with dieresis"},
}

// All yields every row, in the order Appendix A prints them: by code point.
func All() iter.Seq[Diacritic] {
	return slices.Values(diacritics)
}

// Len reports how many rows Appendix A has.
func Len() int {
	return len(diacritics)
}
