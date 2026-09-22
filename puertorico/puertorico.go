// Package puertorico holds the Puerto Rico Spanish address vocabulary as data.
//
// Project US@ keeps the Spanish forms rather than translating them to English,
// so these are not a localization of the tables in the sibling packages — they
// are their own vocabulary, and a Puerto Rico address is written in it.
//
// Four tables live here. StreetTypes are the leading type that opens a Puerto
// Rico street line, where the mainland puts its suffix at the end. Secondaries
// are the secondary address identifiers. Urbanizations are the designators
// that open the urbanization line, which the standard puts on a line of its
// own above the secondary address identifier. StandaloneUrbanizations are the
// urbanization names that are never preceded by URB.
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

var streetTypes = []StreetType{
	{Full: "AVENIDA", Short: "AVE", English: "AVENUE"},
	{Full: "CALLE", Short: "CLL", English: "STREET"},
	{Full: "CAMINITO", Short: "CMT", English: "LITTLE ROAD"},
	{Full: "CAMINO", Short: "CAM", English: "ROAD"},
	{Full: "CERRADA", Short: "CER", English: "CLOSED"},
	{Full: "CIRCULO", Short: "CIR", English: "CIRCLE"},
	{Full: "ENTRADA", Short: "ENT", English: "ENTRANCE"},
	{Full: "PASEO", Short: "PSO", English: "PATH"},
	{Full: "PLACITA", Short: "PLA", English: "LITTLE PLAZA"},
	{Full: "RANCHO", Short: "RCH", English: "RANCH"},
	{Full: "VEREDA", Short: "VER", English: "SMALL PATH"},
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
