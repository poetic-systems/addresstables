// Package puertorico holds the Puerto Rico Spanish address vocabulary as data.
//
// Project US@ keeps the Spanish forms rather than translating them to English,
// so these are not a localization of the tables in the sibling packages — they
// are their own vocabulary, and a Puerto Rico address is written in it.
//
// Five tables live here. StreetTypes are the leading type that opens a Puerto
// Rico street line, where the mainland puts its suffix at the end. Secondaries
// are the secondary address identifiers. Urbanizations are the designators
// that open the urbanization line, which the standard puts on a line of its
// own above the secondary address identifier. StandaloneUrbanizations are the
// urbanization names that are never preceded by URB. RouteWords are the words
// a Puerto Rico rural route or highway contract route address is written
// with, and the RR, HC, or BOX the standard requires in their place.
//
// Urbanization is kept separate from Secondary deliberately. URB is not a
// secondary designator — the standard gives it its own line and its own
// meaning — and listing it in both tables would make it two things at once.
//
// Rows are uppercase, as in the sibling packages, because Project US@ requires
// uppercase output.
package puertorico

import (
	"iter"
	"slices"
)

// StreetType is a Puerto Rico leading street type and its abbreviation.
// English is the translation, carried for documentation rather than for
// output: the standard does not ask for these to be translated.
type StreetType struct {
	Full    string
	Short   string
	English string
}

// streetTypes is the union of two lists the standard keeps separately, not a
// transcription of either alone. Appendix E, p. 62 ("Standard Abbreviations
// for Spanish-Language Addresses"), gives twelve rows with abbreviations. The
// body of the standard, p. 26, under "Street Names and Prefixes", gives a
// second, independent list of seventeen — the standard's own words, the line
// wrap inside CALLEJON preserved as the document has it:
//
//	CALLE AVENIDA, PASEO, PLAZA, PASAJE, CARR, PARQUE, VEREDA, VISTA, VIA, CALLE
//	JON, PATIO, BLVD, CAMINO, CAMINITO, CALETA, MARGINAL
//
// Seven words are in both lists: CALLE, AVENIDA, PASEO, VEREDA, VISTA,
// CAMINO, CAMINITO. Five are Appendix E only: CERRADA, CIRCULO, ENTRADA,
// PLACITA, RANCHO. Ten are p. 26 only: PLAZA, PASAJE, CARRETERA (CARR is the
// p. 26 spelling), PARQUE, VIA, CALLEJON, PATIO, BOULEVARD (BLVD is the p. 26
// spelling), CALETA, MARGINAL. Two independent lists that only half overlap
// is what says this is a union to take, not an errata to reconcile one list
// against the other. See #17.
//
// Eight of the ten p. 26-only rows have no abbreviation published anywhere in
// the standard, so Short is the word itself — the same convention
// standaloneUrbanizations already uses for BOSQUE. The other three —
// CARRETERA/CARR, PARQUE/PARQ, and BOULEVARD/BLVD — are each the same word in
// a second role: CARRETERA is already a Secondary, PARQUE is already a
// StandaloneUrbanization, and BOULEVARD/BLVD is Pub 28's English suffix pair.
// The StandaloneUrbanization comment already establishes that a word carrying
// a different abbreviation in a different role is normal here and not a
// collision to fix.
//
// English is documentation only, never output: the standard gives no basis
// for translating a Spanish street type, and elsewhere requires the opposite
// — Developers MUST NOT translate CALLE to the suffix ST (p. 26, the same page).
var streetTypes = []StreetType{
	{Full: "AVENIDA", Short: "AVE", English: "AVENUE"},
	{Full: "BOULEVARD", Short: "BLVD", English: "BOULEVARD"},
	{Full: "CALETA", Short: "CALETA", English: "COVE"},
	{Full: "CALLE", Short: "CLL", English: "STREET"},
	{Full: "CALLEJON", Short: "CALLEJON", English: "ALLEY"},
	{Full: "CAMINITO", Short: "CMT", English: "LITTLE ROAD"},
	{Full: "CAMINO", Short: "CAM", English: "ROAD"},
	{Full: "CARRETERA", Short: "CARR", English: "HIGHWAY"},
	{Full: "CERRADA", Short: "CER", English: "CLOSED"},
	{Full: "CIRCULO", Short: "CIR", English: "CIRCLE"},
	{Full: "ENTRADA", Short: "ENT", English: "ENTRANCE"},
	{Full: "MARGINAL", Short: "MARGINAL", English: "FRONTAGE ROAD"},
	{Full: "PARQUE", Short: "PARQ", English: "PARK"},
	{Full: "PASAJE", Short: "PASAJE", English: "PASSAGE"},
	{Full: "PASEO", Short: "PSO", English: "PATH"},
	{Full: "PATIO", Short: "PATIO", English: "COURTYARD"},
	{Full: "PLACITA", Short: "PLA", English: "LITTLE PLAZA"},
	{Full: "PLAZA", Short: "PLAZA", English: "PLAZA"},
	{Full: "RANCHO", Short: "RCH", English: "RANCH"},
	{Full: "VEREDA", Short: "VER", English: "SMALL PATH"},
	{Full: "VIA", Short: "VIA", English: "WAY"},
	{Full: "VISTA", Short: "VIS", English: "VIEW"},
}

// Secondary is a Puerto Rico secondary address identifier and its
// abbreviation. Some are Spanish and some are English words that appear in
// Puerto Rico addresses; the table does not distinguish them because the
// standard does not.
type Secondary struct {
	Full  string
	Short string
}

var secondaries = []Secondary{
	{Full: "APARTAMENTO", Short: "APT"},
	{Full: "BARRIADA", Short: "BDA"},
	{Full: "BUILDING", Short: "BLDG"},
	{Full: "BLOQUE", Short: "BL"},
	{Full: "BARRIO", Short: "BO"},
	{Full: "CARRETERA", Short: "CARR"},
	{Full: "CASERIO", Short: "CAS"},
	{Full: "CONDOMINIO", Short: "COND"},
	{Full: "COOPERATIVA", Short: "COOP"},
	{Full: "CORPORACION", Short: "CORP"},
	{Full: "DEPARTAMENTO", Short: "DEPT"},
	{Full: "EDIFICIO", Short: "EDIF"},
	{Full: "ENTREGA GENERAL", Short: "GEN DEL"},
	{Full: "EXTENCION", Short: "EXT"},
	{Full: "HOSPITAL", Short: "HOSP"},
	{Full: "INDUSTRIAL", Short: "IND"},
	{Full: "JARDINES", Short: "JARD"},
	{Full: "MANSIONES", Short: "MANS"},
	{Full: "PARCELAS", Short: "PARC"},
	{Full: "QUEBRADA", Short: "QBDA"},
	{Full: "REPARTO", Short: "REPTO"},
	{Full: "RESIDENCIAL", Short: "RES"},
	{Full: "SECTOR", Short: "SEC"},
	{Full: "TERRAZA", Short: "TERR"},
	{Full: "VILLA", Short: "VIL"},
}

// StreetTypes yields every Puerto Rico leading street type.
func StreetTypes() iter.Seq[StreetType] {
	return slices.Values(streetTypes)
}

// Secondaries yields every Puerto Rico secondary address identifier.
func Secondaries() iter.Seq[Secondary] {
	return slices.Values(secondaries)
}

// Urbanization is a spelling that opens an urbanization line, with the
// abbreviation the standard requires on output.
//
// The Spanish spelling carries an accent — URBANIZACIÓN — and this table holds
// only the unaccented form. Folding an accented input is the consumer's job,
// so a consumer that adds another accented designator needs one row here and
// no second thought about how it is typed.
type Urbanization struct {
	Full  string
	Short string
}

var urbanizations = []Urbanization{
	{Full: "URB", Short: "URB"},
	{Full: "URBANIZACION", Short: "URB"},
	{Full: "URBANIZATION", Short: "URB"},
}

// Urbanizations yields every urbanization designator spelling.
func Urbanizations() iter.Seq[Urbanization] {
	return slices.Values(urbanizations)
}

// StandaloneUrbanization is a Puerto Rico urbanization name that stands alone
// on the urbanization line: the standard requires it to be written without
// the URB designator, not with it.
//
// This is the Exceptions table from Project US@ Technical Specification
// pp. 28-29. The standard lists these under "Urbanizations" and says they
// "stand alone and MUST NOT require the use of the abbreviation URB" — so a
// consumer that sees one of these Fulls opening the urbanization line MUST
// NOT prepend URB, unlike the ordinary case covered by Urbanizations above.
// The standard's own examples: "URB EXT VISTA BELLA" is wrong, "EXT VISTA
// BELLA" is correct; "URB ALTS DE CANA" is wrong, "ALTS DE CANA" is correct.
//
// The standard writes several rows with a parenthesized S, e.g. "Altura(s)"
// abbreviated "ALT(S)". Per the standard, "Abbreviations containing the
// letter S in parentheses at the end of the abbreviation allows for the
// plural representation of the word in an abbreviated form" — the
// parenthesized notation is not itself a spelling, and its own worked
// example uses the plural (ALTS, not ALT(S)). So every such row is stored
// here as two ordinary rows, singular and plural, on both Full and Short:
// ALTURA/ALT and ALTURAS/ALTS, BRISA/BRISA and BRISAS/BRISAS, COLINA/COLINA
// and COLINAS/COLINAS, LOMA/LOMA and LOMAS/LOMAS, PARCELA/PARCELA and
// PARCELAS/PARCELAS, VILLA/VILLA and VILLAS/VILLAS, VISTA/VISTA and
// VISTAS/VISTAS.
//
// Three of these words already appear elsewhere in this package with a
// different abbreviation, and that is deliberate, not a bug to reconcile:
// the standard gives one word a different abbreviation depending on its
// role. SECTOR abbreviates to SEC in Secondaries but SECT here. PARCELA(S)
// abbreviates to PARC in Secondaries (as PARCELAS) but is unabbreviated
// here (PARCELA/PARCELAS). VILLA(S) abbreviates to VIL in Secondaries but is
// unabbreviated here (VILLA/VILLAS). Likewise PASEO and VISTA appear in
// StreetTypes abbreviated PSO and VIS; here, in the standalone-urbanization
// role, both are unabbreviated. Do not "fix" either table to match the
// other — they are the same words in two separate roles the standard treats
// differently, and each table already matches its own table's role.
//
// Rows are uppercase and diacritics are folded, as in the sibling tables:
// the standard's "Extensión" is stored as EXTENSION. Folding an accented
// input is the consumer's job.
type StandaloneUrbanization struct {
	Full  string
	Short string
}

var standaloneUrbanizations = []StandaloneUrbanization{
	{Full: "ALTURA", Short: "ALT"},
	{Full: "ALTURAS", Short: "ALTS"},
	{Full: "BARRIADA", Short: "BDA"},
	{Full: "BARRIO", Short: "BO"},
	{Full: "BOSQUE", Short: "BOSQUE"},
	{Full: "BRISA", Short: "BRISA"},
	{Full: "BRISAS", Short: "BRISAS"},
	{Full: "CHALETS", Short: "CHALETS"},
	{Full: "CIUDAD", Short: "CIUDAD"},
	{Full: "COLINA", Short: "COLINA"},
	{Full: "COLINAS", Short: "COLINAS"},
	{Full: "COMUNIDAD", Short: "COMUNIDAD"},
	{Full: "ESTANCIAS", Short: "EST"},
	{Full: "EXTENSION", Short: "EXT"},
	{Full: "HACIENDA", Short: "HACIENDA"},
	{Full: "INDUSTRIAL", Short: "IND"},
	{Full: "JARDINES", Short: "JARD"},
	{Full: "LOMA", Short: "LOMA"},
	{Full: "LOMAS", Short: "LOMAS"},
	{Full: "MANSIONES", Short: "MANS"},
	{Full: "PARCELA", Short: "PARCELA"},
	{Full: "PARCELAS", Short: "PARCELAS"},
	{Full: "PARQUE", Short: "PARQ"},
	{Full: "PASEO", Short: "PASEO"},
	{Full: "PORTAL", Short: "PORTAL"},
	{Full: "PORTALES", Short: "PORTALES"},
	{Full: "PRADERA", Short: "PRADERA"},
	{Full: "QUINTAS", Short: "QUINTAS"},
	{Full: "REPARTO", Short: "REPTO"},
	{Full: "RESIDENCIAL", Short: "RES"},
	{Full: "RIBERAS", Short: "RIBERAS"},
	{Full: "SECTOR", Short: "SECT"},
	{Full: "TERRAZA", Short: "TERR"},
	{Full: "VALLE", Short: "VALLE"},
	{Full: "VILLA", Short: "VILLA"},
	{Full: "VILLAS", Short: "VILLAS"},
	{Full: "VISTA", Short: "VISTA"},
	{Full: "VISTAS", Short: "VISTAS"},
}

// StandaloneUrbanizations yields every urbanization name that must not be
// preceded by URB.
func StandaloneUrbanizations() iter.Seq[StandaloneUrbanization] {
	return slices.Values(standaloneUrbanizations)
}

// RouteWord is a word that may be written in a Puerto Rico rural route or
// highway contract route address, and the form the standard requires in its
// place.
type RouteWord struct {
	Spelling string
	Standard string // RR, HC or BOX
}

// routeWords is the vocabulary behind RR___ BOX___ and HC____BOX____, per
// Project US@ v1.0 p. 30, Rural Routes:
//
//	A rural route address in the patient record MUST be standardized as
//	follows: RR___ BOX___
//
//	Developers MUST NOT use the words RURAL, RUTA RURAL, BUZON, or BZN. The
//	designations RFD, RD, and RT (meaning rural route) MUST be changed to RR
//	and developers MUST have a space between RR and the route number and BOX
//	and the box number.
//
//	Developers MUST NOT add a leading zero before the rural route number.
//
// The standard's own rural route examples:
//
//	RR03 BOX 9800            -> RR 3 BOX 9800
//	RFD ROUTE 4 BZN 1725     -> RR 4 BOX 1725
//	RUTA RURAL 3 BUZON 12000 -> RR 3 BOX 12000
//	RFD 1 Bzn 17-A           -> RR 1 BOX 17A
//
// pp. 30-31, Highway Contract Routes, calls itself "basically the same format
// utilized for rural routes", with its own designation in place of RR:
//
//	Highway contract route addresses MUST be standardized as HC____BOX____.
//	... Health IT developers MUST NOT include leading zeros before the route
//	number.
//
//	Ruta Estrella 1 Buzón 18 -> HC 1 BOX 18
//	HC 03 Bzn 1050           -> HC 1 BOX 1050
//
// The second highway contract example is a defect in the printed standard,
// recorded here so a reader does not have to rediscover it: HC 03 cannot
// standardize to HC 1, and the leading-zero rule on the same page gives HC 3
// BOX 1050 instead. This table holds words, not the defective worked number,
// so the wrong answer is not reproduced as a row.
//
// RFD, RD, and RT are English, and are in this Spanish-vocabulary table
// anyway: p. 30 states a different rule about the same words than p. 22 does
// for the mainland — MUST be changed to RR here, only SHOULD change there —
// and an address written RFD ROUTE 4 BZN 1725 has to find both halves, RFD
// ROUTE and BZN, in one vocabulary to be read at all. Secondaries already
// carries English words for a comparable reason. RD is also Pub 28's
// abbreviation for ROAD, so a consumer MUST NOT apply this table outside a
// recognized route pattern — the same caution StreetTypes' English column
// carries for a word that means something else out of context.
//
// Appendix F, p. 63, glosses three of these for information, not for
// substitution: RUTA RURAL = Rural Route, RUTA ESTRELLA = Highway Contract,
// BUZON = Box.
//
// BUZÓN carries an accent in the standard's own text, and this table holds
// only the unaccented BUZON — the same reason Urbanization gives for storing
// URBANIZACION unaccented: folding an accented input is the consumer's job,
// not this table's.
//
// NOTE: rows are ordered by the length of Spelling descending, then
// alphabetically — not alphabetically ascending as the sibling tables in this
// package are. A consumer scanning an address for one of these words, or
// replacing one with its Standard form, wants the first match it finds at a
// given position to be the longest one available, not merely the
// alphabetically first one: RR listed before RURAL would let a scanner stop
// at RR and leave RURAL half-read, and BOX before BUZON would do the same to
// BUZON. Sorting the longest Spelling first, and breaking ties alphabetically
// only when two words are the same length, means a caller that takes the
// first match in row order always reads a whole word, never a fragment of a
// longer one.
//
// This table has one reader, go-projectusat#119, not two: a route designator
// never appears in a TIGER street name, so zipcity has no use for it. It
// belongs here anyway — every vocabulary go-projectusat's puertorico package
// reads already comes from here, and putting this one inline would make it
// the exception inside its own package.
var routeWords = []RouteWord{
	{Spelling: "RUTA ESTRELLA", Standard: "HC"},
	{Spelling: "RUTA RURAL", Standard: "RR"},
	{Spelling: "RFD ROUTE", Standard: "RR"},
	{Spelling: "BUZON", Standard: "BOX"},
	{Spelling: "RURAL", Standard: "RR"},
	{Spelling: "BOX", Standard: "BOX"},
	{Spelling: "BZN", Standard: "BOX"},
	{Spelling: "RFD", Standard: "RR"},
	{Spelling: "HC", Standard: "HC"},
	{Spelling: "RD", Standard: "RR"},
	{Spelling: "RR", Standard: "RR"},
	{Spelling: "RT", Standard: "RR"},
}

// RouteWords yields every Puerto Rico rural route and highway contract route
// word, in longest-first order — see the NOTE on routeWords.
func RouteWords() iter.Seq[RouteWord] {
	return slices.Values(routeWords)
}
