package cityabbreviations_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/poetic-systems/addresstables/cityabbreviations"
)

// TestRowsAreWholeUppercaseWords pins the contract a consumer keys on: a row
// is one word, uppercase, with no punctuation for the consumer to have to
// strip. ST. is ST after the punctuation rule, not a row of its own.
func TestRowsAreWholeUppercaseWords(t *testing.T) {
	for a := range cityabbreviations.All() {
		for _, w := range []string{a.Full, a.Short} {
			if w == "" {
				t.Errorf("row %+v has an empty word", a)
			}
			if strings.ToUpper(w) != w {
				t.Errorf("%q is not uppercase", w)
			}
			if strings.ContainsFunc(w, func(r rune) bool { return !unicode.IsLetter(r) }) {
				t.Errorf("%q is not a single word of letters", w)
			}
		}
		if len(a.Short) >= len(a.Full) {
			t.Errorf("%q is no shorter than %q", a.Short, a.Full)
		}
	}
}

// TestShortFormsAreDistinct makes sure a consumer's map keyed by Short has one
// spelling to expand to. STE has to stay its own row rather than a spelling of
// ST, or SAULT STE MARIE becomes SAULT SAINT MARIE.
func TestShortFormsAreDistinct(t *testing.T) {
	seen := map[string]string{}
	for a := range cityabbreviations.All() {
		if full, dup := seen[a.Short]; dup {
			t.Errorf("%q abbreviates both %q and %q", a.Short, full, a.Full)
		}
		seen[a.Short] = a.Full
	}
	if len(seen) != cityabbreviations.Len() {
		t.Errorf("All yielded %d distinct rows, Len says %d", len(seen), cityabbreviations.Len())
	}
}
