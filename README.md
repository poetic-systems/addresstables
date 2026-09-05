# addresstables

USPS Publication 28 and Census Bureau TIGER/Line address vocabularies, as data.

## Copyright and License

Copyright 2026 Poetic Systems

Unless otherwise specified all code and related artifacts in this repository are
made available under the Apache 2 License. See the [license](./LICENSE) for
details.

### Data

The tables here are transcribed from
[USPS Publication 28](https://pe.usps.com/text/pub28/welcome.htm) and from the
appendices of the Census Bureau's
[TIGER/Line technical documentation](https://www2.census.gov/geo/pdfs/maps-data/data/tiger/tgrshp2025/TGRSHP2025_TechDoc.pdf),
both U.S. government publications in the public domain.

## Purpose

Two libraries need the same tables and disagree about them.

[go-projectusat](https://github.com/PortobelloAuth/go-projectusat) parses an
address into its Project US@ components. [zipcity](https://github.com/poetic-systems/zipcity)
compresses street, city, region, and postal code data into bloom filters so a
parser can ask whether a street name is plausible for a place without calling a
service. A bloom filter answers "maybe present" or "definitely absent", so the
key the caller builds has to be byte-for-byte the key the filter was built with.
A mismatch in the key form is not a soft failure — it is a confident, silent
false negative.

That makes the abbreviation tables an interop contract rather than a
convenience. When zipcity's Census-derived expansion and go-projectusat's
Publication 28 abbreviation disagree about a single suffix, one library builds
`FOX PARK DR` and the other looks up `FOX PARK DRIVE`, and the address is
rejected by a filter that contains it. Seven such disagreements were measured
across the shipped suffix table. The tables live here so there is one row for
each word and both sides read it.

## What this package is not

It does not normalize, match, or parse. There are no `Normalize` or
`Abbreviate` functions and there will not be.

That is deliberate. The two consumers legitimately need different matching
semantics — go-projectusat resolves an input spelling to a canonical output
form and has to rate its confidence; zipcity builds an index key and has to be
exact. Putting a matcher here would force one of those to be wrong, and would
put the interesting behaviour in the place neither team reviews. What both
sides genuinely share is the rows.

So each package exposes plain structs and Go iterators over them:

```go
suffixes := map[string]string{}
for s := range streetsuffixes.All() {
    for _, alt := range s.Alt {
        suffixes[alt] = s.Short
    }
}
```

A consumer collects the shape it wants — a map keyed by abbreviation, a map
keyed by every spelling, a slice sorted by length for longest-match — and owns
the semantics of the lookup it built.

## Packages

| Package | Contents |
| --- | --- |
| `streetsuffixes` | The 206 Publication 28 street suffixes: primary name, standard abbreviation, and every commonly used spelling. |
| `directionals` | The eight English directionals and the eight Spanish ones, each Spanish row naming its English equivalent. |
| `countries` | The three countries Project US@ scopes itself to, each with the spellings and abbreviations that name it, and a flag for the domestic rows the standard says to omit. |
| `puertorico` | Puerto Rico Spanish street types, secondary designators, and urbanization keywords. |
| `ustigerfile/featuretypes` | The 503 Census Bureau feature types (TechDoc Appendix D): the codes TIGER writes into PRETYP and SUFTYP, and the abbreviation it displays for each. |
| `ustigerfile/qualifiers` | The 17 Census Bureau feature name qualifiers (TechDoc Appendix C): BUSINESS, BYPASS, SPUR, and the rest, as PREQUAL and SUFQUAL codes. |

The `ustigerfile` packages are the Census Bureau's tables, kept apart from the
Postal Service's because they are a different vocabulary that happens to
describe the same streets. They overlap on the ordinary suffixes and disagree
on several abbreviations, and the disagreement is the reason both are here:
reconciling them is a consumer's job, and it cannot be done from one table.

TIGER's own directional table (TechDoc Appendix B) is not among them. Its
sixteen words and abbreviations are the Publication 28 sixteen, row for row,
so `directionals` serves both. It adds only a numeric code, which no TIGER
field a consumer reads is expressed in — PREDIRABRV and SUFDIRABRV carry the
abbreviation itself.

## Uppercase

Every row is uppercase, because Project US@ requires uppercase output. A
consumer that wants another case folds it. This package does not choose one for
anybody.

## Addresses are patient data

These tables describe addresses, and in the systems that consume them an
address identifies a patient. Nothing here holds an address, and nothing here
should: no fixture, test, or example may contain a real one. Invented streets
and the examples printed in the standard itself are fine.
