# ustigerfile

The Census Bureau's address vocabularies, transcribed from the appendices of the
[TIGER/Line technical documentation](https://www2.census.gov/geo/pdfs/maps-data/data/tiger/tgrshp2025/TGRSHP2025_TechDoc.pdf).

They are kept apart from the Publication 28 packages in the root of this module
because they are a different vocabulary that happens to describe the same
streets. The two overlap on the ordinary suffixes and disagree on several
abbreviations, and that disagreement is the reason both are here: reconciling
them is a consumer's job, and it cannot be done from one table.

| Package | Appendix | Rows | Contents |
| --- | --- | --- | --- |
| `featuretypes` | D | 503 | The feature types TIGER writes into `PRETYP` and `SUFTYP`, and the abbreviation it displays for each. |
| `qualifiers` | C | 17 | The feature name qualifiers TIGER writes into `PREQUAL` and `SUFQUAL` — BUSINESS, BYPASS, SPUR, and the rest. |

## Appendix B is deliberately not here

TIGER publishes its own directional table as Appendix B, and there is no
`ustigerfile/directionals` for it.

Its sixteen words and abbreviations are the Publication 28 sixteen, row for
row, so the module's [`directionals`](../directionals) package already serves
both readers. Appendix B adds only a numeric code, and no TIGER field a
consumer reads is expressed in it — `PREDIRABRV` and `SUFDIRABRV` carry the
abbreviation itself.

Transcribing it would mean two tables that must agree forever, and a consumer
having to pick one. One table, no choice.

## Reading a street name out of TIGER

A `FEATNAMES` record carries a rendered name in `FULLNAME`, but the parts are
safer: a base `NAME` with six codes around it — `PREQUAL`, `PREDIRABRV`,
`PRETYP` before it, and `SUFTYP`, `SUFDIRABRV`, `SUFQUAL` after it. `FULLNAME`
is one abbreviated rendering of those parts and is not always complete, so a
reader that wants the whole name expands the codes itself. These two tables
expand the qualifiers and the types; the two directional fields carry the
abbreviation itself, and [`directionals`](../directionals) is what expands
those.

A consumer producing Project US@ output does not stop there. The Census
abbreviation is not the Publication 28 abbreviation, so the word is read here
and then resolved through [`streetsuffixes`](../streetsuffixes); reading only
one of the two is what produces a key that misses.
