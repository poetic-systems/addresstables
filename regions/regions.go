// Package regions holds the Publication 28 state, possession, Canadian
// province and military "state" abbreviations as data.
//
// Three tables, because they are three lists in the standard and a consumer
// may legitimately want one without the others: the states and possessions of
// Appendix B, the Canadian provinces and territories Publication 28 carries
// for the addresses it accepts across the border, and the three military
// "states" that stand in for an APO/FPO/DPO destination.
//
// Alt is the lookup surface, as in streetsuffixes: every spelling a consumer
// might see, including the abbreviation and the primary name themselves, so a
// caller can build one map from Alt alone and reach every row.
// TestAltCoversPrimaryAndShort enforces that.
//
// There is deliberately no matcher here. go-projectusat resolves a written
// region to a canonical form and wants an edit-distance fallback when the
// input is misspelled; a filter generator wants an exact key. Those are
// different semantics over the same rows, so the rows are what is shared.
package regions

import (
	"iter"
	"slices"
)

// Region is one row: the spelled-out name Publication 28 prints, the
// two-letter abbreviation it standardizes on, and every spelling that names
// the same place.
//
// There is no field saying which of the three tables a row came from. That is
// what the three accessors are for; a consumer that needs the distinction asks
// for one list rather than filtering All.
type Region struct {
	Primary string
	Short   string
	Alt     []string
}

var usStatesAndPossessions = []Region{
	{
		Primary: "ALABAMA",
		Short:   "AL",
		Alt: []string{
			"ALABAMA", "AL",
		},
	},
	{
		Primary: "ALASKA",
		Short:   "AK",
		Alt: []string{
			"ALASKA", "AK",
		},
	},
	{
		Primary: "AMERICAN SAMOA",
		Short:   "AS",
		Alt: []string{
			"AMERICAN SAMOA", "AS",
		},
	},
	{
		Primary: "ARIZONA",
		Short:   "AZ",
		Alt: []string{
			"ARIZONA", "AZ",
		},
	},
	{
		Primary: "ARKANSAS",
		Short:   "AR",
		Alt: []string{
			"ARKANSAS", "AR",
		},
	},
	{
		Primary: "CALIFORNIA",
		Short:   "CA",
		Alt: []string{
			"CALIFORNIA", "CA",
		},
	},
	{
		Primary: "COLORADO",
		Short:   "CO",
		Alt: []string{
			"COLORADO", "CO",
		},
	},
	{
		Primary: "CONNECTICUT",
		Short:   "CT",
		Alt: []string{
			"CONNECTICUT", "CT", "CONN",
		},
	},
	{
		Primary: "DELAWARE",
		Short:   "DE",
		Alt: []string{
			"DELAWARE", "DELEWARE", "DE",
		},
	},
	{
		Primary: "DISTRICT OF COLUMBIA",
		Short:   "DC",
		Alt: []string{
			"DISTRICT OF COLUMBIA", "DC",
		},
	},
	{
		Primary: "FEDERATED STATES OF MICRONESIA",
		Short:   "FM",
		Alt: []string{
			"FEDERATED STATES OF MICRONESIA", "MICRONESIA", "FM",
		},
	},
	{
		Primary: "FLORIDA",
		Short:   "FL",
		Alt: []string{
			"FLORIDA", "FL",
		},
	},
	{
		Primary: "GEORGIA",
		Short:   "GA",
		Alt: []string{
			"GEORGIA", "GA",
		},
	},
	{
		Primary: "GUAM",
		Short:   "GU",
		Alt: []string{
			"GUAM", "GU",
		},
	},
	{
		Primary: "HAWAII",
		Short:   "HI",
		Alt: []string{
			"HAWAII", "HI",
		},
	},
	{
		Primary: "IDAHO",
		Short:   "ID",
		Alt: []string{
			"IDAHO", "ID",
		},
	},
	{
		Primary: "ILLINOIS",
		Short:   "IL",
		Alt: []string{
			"ILLINOIS", "IL",
		},
	},
	{
		Primary: "INDIANA",
		Short:   "IN",
		Alt: []string{
			"INDIANA", "IN",
		},
	},
	{
		Primary: "IOWA",
		Short:   "IA",
		Alt: []string{
			"IOWA", "IA",
		},
	},
	{
		Primary: "KANSAS",
		Short:   "KS",
		Alt: []string{
			"KANSAS", "KS",
		},
	},
	{
		Primary: "KENTUCKY",
		Short:   "KY",
		Alt: []string{
			"KENTUCKY", "KY",
		},
	},
	{
		Primary: "LOUISIANA",
		Short:   "LA",
		Alt: []string{
			"LOUISIANA", "LA",
		},
	},
	{
		Primary: "MAINE",
		Short:   "ME",
		Alt: []string{
			"MAINE", "ME",
		},
	},
	{
		Primary: "MARSHALL ISLANDS",
		Short:   "MH",
		Alt: []string{
			"MARSHALL ISLANDS", "MARSHALL IS", "MARSHALL ISL", "MARSHALL ISLS",
			"MARSHALL ISS", "MARSHALL ISLD", "MH",
		},
	},
	{
		Primary: "MARYLAND",
		Short:   "MD",
		Alt: []string{
			"MARYLAND", "MD",
		},
	},
	{
		Primary: "MASSACHUSETTS",
		Short:   "MA",
		Alt: []string{
			"MASSACHUSETTS", "MA", "MASS",
		},
	},
	{
		Primary: "MICHIGAN",
		Short:   "MI",
		Alt: []string{
			"MICHIGAN", "MI",
		},
	},
	{
		Primary: "MINNESOTA",
		Short:   "MN",
		Alt: []string{
			"MINNESOTA", "MN", "MINN",
		},
	},
	{
		Primary: "MISSISSIPPI",
		Short:   "MS",
		Alt: []string{
			"MISSISSIPPI", "MS",
		},
	},
	{
		Primary: "MISSOURI",
		Short:   "MO",
		Alt: []string{
			"MISSOURI", "MO",
		},
	},
	{
		Primary: "MONTANA",
		Short:   "MT",
		Alt: []string{
			"MONTANA", "MT",
		},
	},
	{
		Primary: "NEBRASKA",
		Short:   "NE",
		Alt: []string{
			"NEBRASKA", "NE",
		},
	},
	{
		Primary: "NEVADA",
		Short:   "NV",
		Alt: []string{
			"NEVADA", "NV",
		},
	},
	{
		Primary: "NEW HAMPSHIRE",
		Short:   "NH",
		Alt: []string{
			"NEW HAMPSHIRE", "NH",
		},
	},
	{
		Primary: "NEW JERSEY",
		Short:   "NJ",
		Alt: []string{
			"NEW JERSEY", "NJ",
		},
	},
	{
		Primary: "NEW MEXICO",
		Short:   "NM",
		Alt: []string{
			"NEW MEXICO", "NM",
		},
	},
	{
		Primary: "NEW YORK",
		Short:   "NY",
		Alt: []string{
			"NEW YORK", "NY",
		},
	},
	{
		Primary: "NORTH CAROLINA",
		Short:   "NC",
		Alt: []string{
			"NORTH CAROLINA", "N CAROLINA", "NC",
		},
	},
	{
		Primary: "NORTH DAKOTA",
		Short:   "ND",
		Alt: []string{
			"NORTH DAKOTA", "N DAKOTA", "ND",
		},
	},
	{
		Primary: "NORTHERN MARIANA ISLANDS",
		Short:   "MP",
		Alt: []string{
			"NORTHERN MARIANA ISLANDS", "NORTHERN MARIANA IS",
			"NORTHERN MARIANA ISL", "NORTHERN MARIANA ISLS",
			"NORTHERN MARIANA ISS", "NORTHERN MARIANA ISLD",
			"N MARIANA ISLANDS", "N MARIANA IS", "N MARIANA ISL",
			"N MARIANA ISLS", "N MARIANA ISS", "N MARIANA ISLD", "MP",
		},
	},
	{
		Primary: "OHIO",
		Short:   "OH",
		Alt: []string{
			"OHIO", "OH",
		},
	},
	{
		Primary: "OKLAHOMA",
		Short:   "OK",
		Alt: []string{
			"OKLAHOMA", "OK",
		},
	},
	{
		Primary: "OREGON",
		Short:   "OR",
		Alt: []string{
			"OREGON", "OR",
		},
	},
	{
		Primary: "PALAU",
		Short:   "PW",
		Alt: []string{
			"PALAU", "PW",
		},
	},
	{
		Primary: "PENNSYLVANIA",
		Short:   "PA",
		Alt: []string{
			"PENNSYLVANIA", "PENN", "PA",
		},
	},
	{
		Primary: "PUERTO RICO",
		Short:   "PR",
		Alt: []string{
			"PUERTO RICO", "PR",
		},
	},
	{
		Primary: "RHODE ISLAND",
		Short:   "RI",
		Alt: []string{
			"RHODE ISLAND", "RHODE IS", "RHODE ISL", "RHODE ISLD", "RI",
		},
	},
	{
		Primary: "SOUTH CAROLINA",
		Short:   "SC",
		Alt: []string{
			"SOUTH CAROLINA", "S CAROLINA", "SC",
		},
	},
	{
		Primary: "SOUTH DAKOTA",
		Short:   "SD",
		Alt: []string{
			"SOUTH DAKOTA", "S DAKOTA", "SD",
		},
	},
	{
		Primary: "TENNESSEE",
		Short:   "TN",
		Alt: []string{
			"TENNESSEE", "TENN", "TN",
		},
	},
	{
		Primary: "TEXAS",
		Short:   "TX",
		Alt: []string{
			"TEXAS", "TX",
		},
	},
	{
		Primary: "UTAH",
		Short:   "UT",
		Alt: []string{
			"UTAH", "UT",
		},
	},
	{
		Primary: "VERMONT",
		Short:   "VT",
		Alt: []string{
			"VERMONT", "VT",
		},
	},
	{
		Primary: "VIRGIN ISLANDS",
		Short:   "VI",
		Alt: []string{
			"VIRGIN ISLANDS", "VIRGIN IS", "VIRGIN ISL", "VIRGIN ISLS",
			"VIRGIN ISS", "VIRGIN ISLD", "US VIRGIN ISLANDS", "US VIRGIN IS",
			"US VIRGIN ISL", "US VIRGIN ISLS", "US VIRGIN ISS",
			"US VIRGIN ISLD", "USVI", "VIS", "USA VI", "VI USA", "VI",
		},
	},
	{
		Primary: "VIRGINIA",
		Short:   "VA",
		Alt: []string{
			"VIRGINIA", "VA",
		},
	},
	{
		Primary: "WASHINGTON",
		Short:   "WA",
		Alt: []string{
			"WASHINGTON", "WA",
		},
	},
	{
		Primary: "WEST VIRGINIA",
		Short:   "WV",
		Alt: []string{
			"WEST VIRGINIA", "W VIRGINIA", "WV",
		},
	},
	{
		Primary: "WISCONSIN",
		Short:   "WI",
		Alt: []string{
			"WISCONSIN", "WI",
		},
	},
	{
		Primary: "WYOMING",
		Short:   "WY",
		Alt: []string{
			"WYOMING", "WY",
		},
	},
}

// canadianProvincesAndTerritories comes from the Project US@ technical
// specification, pp. 31-32, not from Publication 28. Appendix B of Publication
// 28 is states and possessions, the eight directionals and the three military
// "states"; it carries no Canadian list at all.
//
// The specification prints "Nunavat Territory". That spelling is retained in
// Alt so a record written to the standard still looks up, but Primary is the
// spelling Canada Post uses. Nothing conforming is emitted from Primary
// anyway - the standard requires the two-character abbreviation - so the
// choice only decides what a caller reading the row back sees.
var canadianProvincesAndTerritories = []Region{
	{
		Primary: "ALBERTA",
		Short:   "AB",
		Alt: []string{
			"ALBERTA", "AB",
		},
	},
	{
		Primary: "BRITISH COLUMBIA",
		Short:   "BC",
		Alt: []string{
			"BRITISH COLUMBIA", "BC",
		},
	},
	{
		Primary: "MANITOBA",
		Short:   "MB",
		Alt: []string{
			"MANITOBA", "MB",
		},
	},
	{
		Primary: "NEW BRUNSWICK",
		Short:   "NB",
		Alt: []string{
			"NEW BRUNSWICK", "NB",
		},
	},
	{
		Primary: "NEWFOUNDLAND AND LABRADOR",
		Short:   "NL",
		Alt: []string{
			"NEWFOUNDLAND AND LABRADOR", "NEWFOUNDLAND", "LABRADOR", "NL",
		},
	},
	{
		Primary: "NORTHWEST TERRITORIES",
		Short:   "NT",
		Alt: []string{
			"NORTHWEST TERRITORIES", "NORTHWEST TERR", "NW TERRITORIES",
			"NW TERR", "NT",
		},
	},
	{
		Primary: "NOVA SCOTIA",
		Short:   "NS",
		Alt: []string{
			"NOVA SCOTIA", "NS",
		},
	},
	{
		Primary: "NUNAVUT TERRITORY",
		Short:   "NU",
		Alt: []string{
			"NUNAVUT TERRITORY", "NUNAVUT TERR", "NUNAVUT",
			"NUNAVAT TERRITORY", "NUNAVAT TERR", "NUNAVAT", "NU",
		},
	},
	{
		Primary: "ONTARIO",
		Short:   "ON",
		Alt: []string{
			"ONTARIO", "ON",
		},
	},
	{
		Primary: "PRINCE EDWARD ISLAND",
		Short:   "PE",
		Alt: []string{
			"PRINCE EDWARD ISLAND", "PRINCE EDWARD IS", "PRINCE EDWARD ISL",
			"PRINCE EDWARD ISLD", "PE",
		},
	},
	{
		Primary: "QUEBEC",
		Short:   "QC",
		Alt: []string{
			"QUEBEC", "QC",
		},
	},
	{
		Primary: "SASKATCHEWAN",
		Short:   "SK",
		Alt: []string{
			"SASKATCHEWAN", "SK",
		},
	},
	{
		Primary: "YUKON TERRITORY",
		Short:   "YT",
		Alt: []string{
			"YUKON TERRITORY", "YUKON TERR", "YUKON", "YT",
		},
	},
}

var militaryRegions = []Region{
	{
		Primary: "ARMED FORCES EUROPE THE MIDDLE EAST AND CANADA",
		Short:   "AE",
		Alt: []string{
			"ARMED FORCES EUROPE THE MIDDLE EAST AND CANADA",
			"ARMED FORCES EUROPE", "AE",
		},
	},
	{
		Primary: "ARMED FORCES PACIFIC",
		Short:   "AP",
		Alt: []string{
			"ARMED FORCES PACIFIC", "AP",
		},
	},
	{
		Primary: "ARMED FORCES AMERICAS",
		Short:   "AA",
		Alt: []string{
			"ARMED FORCES AMERICAS", "ARMED FORCES AMERICA", "AA",
		},
	},
}

// UnitedStates yields the states and possessions, alphabetically.
func UnitedStates() iter.Seq[Region] {
	return slices.Values(usStatesAndPossessions)
}

// Canada yields the Canadian provinces and territories, alphabetically.
func Canada() iter.Seq[Region] {
	return slices.Values(canadianProvincesAndTerritories)
}

// Military yields the three military "state" abbreviations.
func Military() iter.Seq[Region] {
	return slices.Values(militaryRegions)
}

// All yields every row: states and possessions, then Canada, then military.
func All() iter.Seq[Region] {
	return func(yield func(Region) bool) {
		for _, table := range [][]Region{
			usStatesAndPossessions,
			canadianProvincesAndTerritories,
			militaryRegions,
		} {
			for _, r := range table {
				if !yield(r) {
					return
				}
			}
		}
	}
}
