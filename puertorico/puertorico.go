// Package puertorico holds the Puerto Rico Spanish address vocabulary as data.
//
// Project US@ keeps the Spanish forms rather than translating them to English,
// so these are not a localization of the tables in the sibling packages — they
// are their own vocabulary, and a Puerto Rico address is written in it.
//
// Three tables live here. StreetTypes are the leading type that opens a Puerto
// Rico street line, where the mainland puts its suffix at the end. Secondaries
// are the secondary address identifiers. Urbanizations are the designators
// that open the urbanization line, which the standard puts on a line of its
// own above the secondary address identifier.
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
