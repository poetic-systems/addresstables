// Package businesswords holds Publication 28 Appendix G as data.
//
// This package deliberately does no normalization, for the same reason
// described in streetsuffixes: the exact matching a caller wants is a
// decision about how a business name is being read, and callers reading for
// different purposes will make it differently. So the contract is All, and
// each caller builds the maps it needs from the rows it gets.
//
// Unlike streetsuffixes, this table does not guarantee that Alt is an
// unambiguous or complete lookup surface, and both are properties of the
// source rather than gaps in how it was transcribed:
//
//   - Appendix G reuses short common presentations across unrelated words
//     ("ACCT" appears under both ACCOUNT and ACCOUNTANT), so a map keyed by
//     Alt can collide. TestNoDuplicatePrimary enforces uniqueness of Primary
//     only; it does not and cannot promise the same for Alt.
//   - A handful of rows give an official Short that the standard's own list
//     of common presentations does not otherwise spell out, so Alt is not
//     guaranteed to contain Short the way it is in streetsuffixes.
//
// A caller that needs Short to be reachable should look it up directly
// rather than assume it is always one of the Alt spellings.
package businesswords

import (
	"iter"
	"slices"
)

// BusinessWord is one row of Publication 28 Appendix G.
//
// Alt lists the common presentations the standard prints for Primary,
// always including Primary itself. See the package doc for what Alt does
// not promise.
type BusinessWord struct {
	Primary string
	Short   string
	Alt     []string
}

var businessWords = []BusinessWord{
	{
		Primary: "ABACUS",
		Short:   "ABCS",
		Alt: []string{
			"ABACUS", "ABCS",
		},
	},
	{
		Primary: "ABOVE",
		Short:   "ABV",
		Alt: []string{
			"ABOVE", "ABV",
		},
	},
	{
		Primary: "ABRASIVE",
		Short:   "ABR",
		Alt: []string{
			"ABR", "ABRASIVE", "ABRSV",
		},
	},
	{
		Primary: "ABROAD",
		Short:   "ABRD",
		Alt: []string{
			"ABRD", "ABROAD",
		},
	},
	{
		Primary: "ABSOLUTE",
		Short:   "ABSLT",
		Alt: []string{
			"ABSLT", "ABSOLUTE",
		},
	},
	{
		Primary: "ABSTRACT",
		Short:   "ABSTRCT",
		Alt: []string{
			"ABSTRACT", "ABSTRCT",
		},
	},
	{
		Primary: "ACADEMIC",
		Short:   "ACDMC",
		Alt: []string{
			"ACADEMIC", "ACDMC",
		},
	},
	{
		Primary: "ACADEMY",
		Short:   "ACDMY",
		Alt: []string{
			"ACAD", "ACADEM", "ACADEMY", "ACDMY",
		},
	},
	{
		Primary: "ACCESS",
		Short:   "ACCSS",
		Alt: []string{
			"ACCESS", "ACCSS",
		},
	},
	{
		Primary: "ACCESSORY",
		Short:   "ACC",
		Alt: []string{
			"ACC", "ACCESSORY",
		},
	},
	{
		Primary: "ACCIDENT",
		Short:   "ACDNT",
		Alt: []string{
			"ACC", "ACCIDENT", "ACDNT",
		},
	},
	{
		Primary: "ACCOMPLISHMENT",
		Short:   "ACCMPLSMNT",
		Alt: []string{
			"ACCMPLSSMNT", "ACCOMPLISHMENT",
		},
	},
	{
		Primary: "ACCOUNT",
		Short:   "ACCT",
		Alt: []string{
			"AC", "ACC", "ACCNT", "ACCONT", "ACCOUNT",
			"ACCT", "ACCUNT", "ACNT",
		},
	},
	{
		Primary: "ACCOUNTANCY",
		Short:   "ACCTNCY",
		Alt: []string{
			"ACC", "ACCOUNTANC", "ACCOUNTANCY", "ACCOUNTY", "ACCTNCY",
		},
	},
	{
		Primary: "ACCOUNTANT",
		Short:   "ACCNT",
		Alt: []string{
			"AC", "ACC", "ACCNT", "ACCOUNTANT", "ACCT",
			"ACCTANT", "ACCTNT", "ACT",
		},
	},
	{
		Primary: "ACCOUNTING",
		Short:   "ACCTG",
		Alt: []string{
			"ACCOUNTING", "ACCTG", "ACCTNG", "ACTG",
		},
	},
	{
		Primary: "ACCREDITED",
		Short:   "ACCRDTD",
		Alt: []string{
			"ACCRDTD", "ACCREDITED",
		},
	},
	{
		Primary: "ACCREDITATION",
		Short:   "ACCRDTN",
		Alt: []string{
			"ACCRDTN", "ACCREDITATION",
		},
	},
	{
		Primary: "ACCURACY",
		Short:   "ACCRCY",
		Alt: []string{
			"ACCRCY", "ACCURACY",
		},
	},
	{
		Primary: "ACCURATE",
		Short:   "ACCRT",
		Alt: []string{
			"ACCRT", "ACCURATE",
		},
	},
	{
		Primary: "ACHIEVEMENT",
		Short:   "ACHVMNT",
		Alt: []string{
			"ACHIEVEMENT", "ACHVMNT",
		},
	},
	{
		Primary: "ACOUSTIC",
		Short:   "ACSTC",
		Alt: []string{
			"ACOUSTIC", "ACSTC",
		},
	},
	{
		Primary: "ACQUISITION",
		Short:   "ACQSTN",
		Alt: []string{
			"ACQSTN", "ACQUIS", "ACQUISITION",
		},
	},
	{
		Primary: "ACROSS",
		Short:   "ACR",
		Alt: []string{
			"ACR", "ACROSS",
		},
	},
	{
		Primary: "ACTING",
		Short:   "ACTNG",
		Alt: []string{
			"ACTING", "ACTNG",
		},
	},
	{
		Primary: "ACTION",
		Short:   "ACTN",
		Alt: []string{
			"ACTION", "ACTN",
		},
	},
	{
		Primary: "ACTIVITY",
		Short:   "ACTVTY",
		Alt: []string{
			"ACTIVITY", "ACTVT", "ACTVTY",
		},
	},
	{
		Primary: "ACTOR",
		Short:   "ACTR",
		Alt: []string{
			"ACTOR", "ACTR",
		},
	},
	{
		Primary: "ACTUARY",
		Short:   "ACTRY",
		Alt: []string{
			"ACTRY", "ACTUARY",
		},
	},
	{
		Primary: "ACTUARIAL",
		Short:   "ACTRL",
		Alt: []string{
			"ACTRL", "ACTUARIAL", "ACTURIAL",
		},
	},
	{
		Primary: "ACUPUNCTURE",
		Short:   "ACPNCTR",
		Alt: []string{
			"ACPNCTR", "ACUPUNCTURE",
		},
	},
	{
		Primary: "ADDITION",
		Short:   "ADDTN",
		Alt: []string{
			"ADDITION", "ADDTN",
		},
	},
	{
		Primary: "ADDRESS",
		Short:   "ADDR",
		Alt: []string{
			"ADDR", "ADDRESS",
		},
	},
	{
		Primary: "ADHESIVE",
		Short:   "ADHSV",
		Alt: []string{
			"ADHESIVE", "ADHSV",
		},
	},
	{
		Primary: "ADJUSTER",
		Short:   "ADJTER",
		Alt: []string{
			"ADJ", "ADJT", "ADJTER", "ADJUSTER",
		},
	},
	{
		Primary: "ADJUSTMENT",
		Short:   "ADJMT",
		Alt: []string{
			"ADJMT", "ADJUSTMENT",
		},
	},
	{
		Primary: "ADJUSTOR",
		Short:   "ADJTOR",
		Alt: []string{
			"ADJ", "ADJT", "ADJTOR", "ADJUSTOR",
		},
	},
	{
		Primary: "ADJUTANT",
		Short:   "ADJT",
		Alt: []string{
			"ADJ", "ADJT", "ADJUTANT",
		},
	},
	{
		Primary: "ADMINISTRATION",
		Short:   "ADMN",
		Alt: []string{
			"AD", "ADM", "ADMIN", "ADMINIST", "ADMINISTRATI",
			"ADMINISTRATION", "ADMINISTRATN", "ADMN", "ADMSTRN",
		},
	},
	{
		Primary: "ADMINISTRATIVE",
		Short:   "ADMNSTRV",
		Alt: []string{
			"AD", "ADMIN", "ADMINI", "ADMINISTRATIVE", "ADMINISTRATV",
			"ADMSTR",
		},
	},
	{
		Primary: "ADMINISTRATOR",
		Short:   "ADMNSTR",
		Alt: []string{
			"ADMIN", "ADMINISTER", "ADMINISTR", "ADMINISTRA", "ADMINISTRATOR",
			"ADMINSTR", "ADMR", "ADMSTR",
		},
	},
	{
		Primary: "ADMIRAL",
		Short:   "ADM",
		Alt: []string{
			"ADM", "ADMIRAL",
		},
	},
	{
		Primary: "ADOPTION",
		Short:   "ADPTN",
		Alt: []string{
			"ADOPTION", "ADPTN",
		},
	},
	{
		Primary: "ADROIT",
		Short:   "ADRT",
		Alt: []string{
			"ADROIT", "ADRT",
		},
	},
	{
		Primary: "ADULT",
		Short:   "ADLT",
		Alt: []string{
			"ADLT", "ADULT",
		},
	},
	{
		Primary: "ADVANCE",
		Short:   "ADVNC",
		Alt: []string{
			"ADVANCE", "ADVNC",
		},
	},
	{
		Primary: "ADVANCED",
		Short:   "ADVNCD",
		Alt: []string{
			"ADV", "ADVANCED", "ADVNCD",
		},
	},
	{
		Primary: "ADVANCEMENT",
		Short:   "ADVMNT",
		Alt: []string{
			"ADVANCEMENT", "ADVMNT",
		},
	},
	{
		Primary: "ADVENTURE",
		Short:   "ADVNTR",
		Alt: []string{
			"ADVENTURE", "ADVNTR",
		},
	},
	{
		Primary: "ADVERTISE",
		Short:   "ADVT",
		Alt: []string{
			"ADVERTISE", "ADVT",
		},
	},
	{
		Primary: "ADVERTISEMENT",
		Short:   "AD",
		Alt: []string{
			"AD", "ADV", "ADVERTISEMENT",
		},
	},
	{
		Primary: "ADVERTISING",
		Short:   "ADVTSNG",
		Alt: []string{
			"AD", "ADV", "ADVERT", "ADVERTISIN", "ADVERTISING",
			"ADVERTISNG", "ADVG", "ADVR", "ADVTG", "ADVTNG",
			"ADVTSNG",
		},
	},
	{
		Primary: "ADVISER",
		Short:   "ADVSR",
		Alt: []string{
			"ADV", "ADVISER", "ADVISOR", "ADVSER", "ADVSOR",
			"ADVSR",
		},
	},
	{
		Primary: "ADVISORY",
		Short:   "ADVRY",
		Alt: []string{
			"ADV", "ADVISORY",
		},
	},
	{
		Primary: "AERIAL",
		Short:   "ARL",
		Alt: []string{
			"AERIAL", "ARL",
		},
	},
	{
		Primary: "AERONAUTICAL",
		Short:   "ARNTCL",
		Alt: []string{
			"AERONAUTICAL", "ARNTCL",
		},
	},
	{
		Primary: "AEROSPACE",
		Short:   "ARSPC",
		Alt: []string{
			"AEROSPACE", "ARSPC", "AS",
		},
	},
	{
		Primary: "AEROSTAT",
		Short:   "ARSTT",
		Alt: []string{
			"AEROSTAT", "ARSTT",
		},
	},
	{
		Primary: "AESTHETIC",
		Short:   "ASTHTC",
		Alt: []string{
			"AESTHETIC", "ASTHTC",
		},
	},
	{
		Primary: "AFFAIR",
		Short:   "AFFR",
		Alt: []string{
			"AFFAIR", "AFFR",
		},
	},
	{
		Primary: "AFFILIATE",
		Short:   "AFFLT",
		Alt: []string{
			"AFFILIATE", "AFFLT",
		},
	},
	{
		Primary: "AFFILIATED",
		Short:   "AFFLTD",
		Alt: []string{
			"AFFILIATED", "AFFLTD",
		},
	},
	{
		Primary: "AFRICAN",
		Short:   "AFRCN",
		Alt: []string{
			"AFRCN", "AFRICAN",
		},
	},
	{
		Primary: "AGENCY",
		Short:   "AGCY",
		Alt: []string{
			"AGCY", "AGE", "AGENC", "AGENCY", "AGNCY",
		},
	},
	{
		Primary: "AGENT",
		Short:   "AGNT",
		Alt: []string{
			"AGEN", "AGENT", "AGNT", "AGT",
		},
	},
	{
		Primary: "AGGREGATE",
		Short:   "AGGRGT",
		Alt: []string{
			"AGGREGATE", "AGGRGT",
		},
	},
	{
		Primary: "AGING",
		Short:   "AGNG",
		Alt: []string{
			"AGING", "AGNG",
		},
	},
	{
		Primary: "AGRICULTURAL",
		Short:   "AGRCLTL",
		Alt: []string{
			"AG", "AGRCLTRL", "AGRICULTURAL",
		},
	},
	{
		Primary: "AGRICULTURE",
		Short:   "AGRCLT",
		Alt: []string{
			"AG", "AGRCLT", "AGRICULTURE",
		},
	},
	{
		Primary: "AIDED",
		Short:   "AID",
		Alt: []string{
			"AID", "AIDED",
		},
	},
	{
		Primary: "AIR CONDITIONING",
		Short:   "AC",
		Alt: []string{
			"AC", "AIR CONDITIONING", "AIRCONDITIONING", "ARCNDTNG",
		},
	},
	{
		Primary: "AIRCRAFT",
		Short:   "ARCRFT",
		Alt: []string{
			"AIRCRAFT", "AIRCRFT", "ARCRFT",
		},
	},
	{
		Primary: "AIRLINE",
		Short:   "ARLN",
		Alt: []string{
			"AIRLINE", "ARLN",
		},
	},
	{
		Primary: "AIRMAN",
		Short:   "ARMN",
		Alt: []string{
			"AIRMAN", "AMN", "ARMN",
		},
	},
	{
		Primary: "AIRPORT",
		Short:   "ARPRT",
		Alt: []string{
			"AIRP", "AIRPORT", "AIRPT", "ARPRT", "ARPT",
		},
	},
	{
		Primary: "AIRWAY",
		Short:   "ARWY",
		Alt: []string{
			"AIRWAY", "ARWY",
		},
	},
	{
		Primary: "ALARM",
		Short:   "ALRM",
		Alt: []string{
			"ALARM", "ALRM",
		},
	},
	{
		Primary: "ALCOHOLIC",
		Short:   "ALCHLC",
		Alt: []string{
			"ALCHLC", "ALCOHOLIC",
		},
	},
	{
		Primary: "ALCOHOLISM",
		Short:   "ALCHLSM",
		Alt: []string{
			"ALCHLSM", "ALCOHOLISM",
		},
	},
	{
		Primary: "ALDERMAN",
		Short:   "ALDM",
		Alt: []string{
			"ALDERMAN", "ALDM",
		},
	},
	{
		Primary: "ALIGNER",
		Short:   "ALGNR",
		Alt: []string{
			"ALGNR", "ALIGNER",
		},
	},
	{
		Primary: "ALIGNING",
		Short:   "ALGNNG",
		Alt: []string{
			"ALGNNG", "ALIGNING",
		},
	},
	{
		Primary: "ALIGNMENT",
		Short:   "ALIGN",
		Alt: []string{
			"ALGNMNT", "ALGNMT", "ALIG", "ALIGN", "ALIGNMENT",
			"ALIGNMNT", "ALIGNMT", "ALIMENT",
		},
	},
	{
		Primary: "ALLERGIST",
		Short:   "ALLRGST",
		Alt: []string{
			"ALLERGIST", "ALLRGST",
		},
	},
	{
		Primary: "ALLERGY",
		Short:   "ALLRGY",
		Alt: []string{
			"ALLERGY", "ALLRGY",
		},
	},
	{
		Primary: "ALLIANCE",
		Short:   "ALLNCE",
		Alt: []string{
			"ALLIANCE", "ALLIE", "ALLNCE",
		},
	},
	{
		Primary: "ALLIED",
		Short:   "ALLD",
		Alt: []string{
			"ALLD", "ALLIE", "ALLIED",
		},
	},
	{
		Primary: "ALLOCATE",
		Short:   "ALLCT",
		Alt: []string{
			"ALLCT", "ALLOCATE",
		},
	},
	{
		Primary: "ALLOCATION",
		Short:   "ALLCTN",
		Alt: []string{
			"ALLCTN", "ALLOCATION",
		},
	},
	{
		Primary: "ALLOY",
		Short:   "ALLY",
		Alt: []string{
			"ALLOY", "ALLY",
		},
	},
	{
		Primary: "ALPHA",
		Short:   "ALPH",
		Alt: []string{
			"ALPH", "ALPHA",
		},
	},
	{
		Primary: "ALTER",
		Short:   "ALTR",
		Alt: []string{
			"ALTER", "ALTR",
		},
	},
	{
		Primary: "ALTERATION",
		Short:   "ALTRN",
		Alt: []string{
			"ALTER", "ALTERATION",
		},
	},
	{
		Primary: "ALTERNATIVE",
		Short:   "ALTRNTV",
		Alt: []string{
			"ALTERNATIVE", "ALTRNTV",
		},
	},
	{
		Primary: "ALTERNATOR",
		Short:   "ALTRNTR",
		Alt: []string{
			"ALTERNATOR", "ALTRNTR",
		},
	},
	{
		Primary: "ALTITUDE",
		Short:   "ALTTD",
		Alt: []string{
			"ALTITUDE", "ALTTD",
		},
	},
	{
		Primary: "ALUMINUM",
		Short:   "ALUMN",
		Alt: []string{
			"AL", "ALUM", "ALUMINUM",
		},
	},
	{
		Primary: "AMATEUR",
		Short:   "AMTR",
		Alt: []string{
			"AMATEUR", "AMTR",
		},
	},
	{
		Primary: "AMBASSADOR",
		Short:   "AMB",
		Alt: []string{
			"AMB", "AMBASSADOR",
		},
	},
	{
		Primary: "AMBIANCE",
		Short:   "AMBNC",
		Alt: []string{
			"AMBIANCE", "AMBNC",
		},
	},
	{
		Primary: "AMBULANCE",
		Short:   "AMBL",
		Alt: []string{
			"AMB", "AMBL", "AMBULANCE",
		},
	},
	{
		Primary: "AMELIORATION",
		Short:   "AMLRTN",
		Alt: []string{
			"AMELIORATION", "AMLRTN",
		},
	},
	{
		Primary: "AMERICA",
		Short:   "AMER",
		Alt: []string{
			"AMER", "AMERICA",
		},
	},
	{
		Primary: "AMERICAN",
		Short:   "AMERCN",
		Alt: []string{
			"AMER", "AMERCN", "AMERICAN",
		},
	},
	{
		Primary: "AMMONIA",
		Short:   "AMMN",
		Alt: []string{
			"AMMN", "AMMONIA",
		},
	},
	{
		Primary: "AMMUNITION",
		Short:   "AMMUN",
		Alt: []string{
			"AMMUN", "AMMUNITION",
		},
	},
	{
		Primary: "AMOUNT",
		Short:   "AMNT",
		Alt: []string{
			"AMNT", "AMOUNT",
		},
	},
	{
		Primary: "AMPHIBIOUS",
		Short:   "AMPHBS",
		Alt: []string{
			"AMPHBS", "AMPHIBIOUS",
		},
	},
	{
		Primary: "AMUSEMENT",
		Short:   "AMUSE",
		Alt: []string{
			"AMUS", "AMUSE", "AMUSEMENT",
		},
	},
	{
		Primary: "ANALOG",
		Short:   "ANLG",
		Alt: []string{
			"ANALOG", "ANLG",
		},
	},
	{
		Primary: "ANALYSIS",
		Short:   "ANLYS",
		Alt: []string{
			"ANALYSIS", "ANLYS",
		},
	},
	{
		Primary: "ANALYST",
		Short:   "ANLYST",
		Alt: []string{
			"ANAL", "ANALY", "ANALYS", "ANALYST", "ANL",
			"ANLST", "ANLYS", "ANLYST",
		},
	},
	{
		Primary: "ANALYTIC",
		Short:   "ANLYTC",
		Alt: []string{
			"ANALYTIC", "ANLYTC",
		},
	},
	{
		Primary: "ANALYTICAL",
		Short:   "ANLYTCL",
		Alt: []string{
			"ANALYTICAL", "ANLYTCL",
		},
	},
	{
		Primary: "ANCHOR",
		Short:   "ANCHR",
		Alt: []string{
			"ANCHOR", "ANCHR",
		},
	},
	{
		Primary: "ANCIENT",
		Short:   "ANCNT",
		Alt: []string{
			"ANCIENT", "ANCNT",
		},
	},
	{
		Primary: "AND",
		Short:   "&",
		Alt: []string{
			"&", "&&", "AND",
		},
	},
	{
		Primary: "ANESTHESIA",
		Short:   "ANSTHS",
		Alt: []string{
			"ANESTHESIA", "ANSTHS",
		},
	},
	{
		Primary: "ANESTHESIOLOGY",
		Short:   "ANSTHSLGY",
		Alt: []string{
			"ANESTHESIOLOGY", "ANSTHSLGY",
		},
	},
	{
		Primary: "ANGLE",
		Short:   "ANGL",
		Alt: []string{
			"ANGL", "ANGLE",
		},
	},
	{
		Primary: "ANGLER",
		Short:   "ANGLR",
		Alt: []string{
			"ANGLER", "ANGLR",
		},
	},
	{
		Primary: "ANGLICAN",
		Short:   "ANGLCN",
		Alt: []string{
			"ANGLCN", "ANGLICAN",
		},
	},
	{
		Primary: "ANIMAL",
		Short:   "ANML",
		Alt: []string{
			"ANIMAL", "ANML",
		},
	},
	{
		Primary: "ANIMATED",
		Short:   "ANMTD",
		Alt: []string{
			"ANIMATED", "ANMTD",
		},
	},
	{
		Primary: "ANNEX",
		Short:   "ANX",
		Alt: []string{
			"ANNEX", "ANNX",
		},
	},
	{
		Primary: "ANONYMOUS",
		Short:   "ANON",
		Alt: []string{
			"ANNYMS", "ANONYMOUS",
		},
	},
	{
		Primary: "ANNUAL",
		Short:   "ANNL",
		Alt: []string{
			"ANNL", "ANNUAL",
		},
	},
	{
		Primary: "ANODIZING",
		Short:   "ANDZNG",
		Alt: []string{
			"ANDZNG", "ANODIZING",
		},
	},
	{
		Primary: "ANSWERING",
		Short:   "ANSWRNG",
		Alt: []string{
			"ANS", "ANSWERING", "ANSWRNG",
		},
	},
	{
		Primary: "ANTIQUE",
		Short:   "ANTQ",
		Alt: []string{
			"ANTIQUE", "ANTQ",
		},
	},
	{
		Primary: "APARTMENT",
		Short:   "APT",
		Alt: []string{
			"APART", "APARTMENT", "APT",
		},
	},
	{
		Primary: "APOSTOLATE",
		Short:   "APSTLT",
		Alt: []string{
			"APOSTOLATE", "APSTLT",
		},
	},
	{
		Primary: "APOSTOLIC",
		Short:   "APSTLC",
		Alt: []string{
			"APOSTOLIC", "APSTLC",
		},
	},
	{
		Primary: "APPARATUS",
		Short:   "APPRTS",
		Alt: []string{
			"APPARATUS", "APPRTS",
		},
	},
	{
		Primary: "APPAREL",
		Short:   "APPRL",
		Alt: []string{
			"AP", "APPAREL", "APPRL",
		},
	},
	{
		Primary: "APPLE",
		Short:   "APPLE",
		Alt:     []string{"APPLE"},
	},
	{
		Primary: "APPLIANCE",
		Short:   "APPLNC",
		Alt: []string{
			"APPL", "APPLIANC", "APPLIANCE", "APPLNC",
		},
	},
	{
		Primary: "APPLICATION",
		Short:   "APPLCTN",
		Alt: []string{
			"APPLCTN", "APPLICATION",
		},
	},
	{
		Primary: "APPLICATOR",
		Short:   "APPLCTR",
		Alt: []string{
			"APPLCTR", "APPLICATOR",
		},
	},
	{
		Primary: "APPLIED",
		Short:   "APPLD",
		Alt: []string{
			"APPLD", "APPLIED",
		},
	},
	{
		Primary: "APPLIQUE",
		Short:   "APPLQ",
		Alt: []string{
			"APPLIQUE", "APPLQ",
		},
	},
	{
		Primary: "APPOINTED",
		Short:   "APPNTD",
		Alt: []string{
			"APPNTD", "APPOINTED",
		},
	},
	{
		Primary: "APPRAISAL",
		Short:   "APPRSL",
		Alt: []string{
			"APPRAISAL", "APPRSL", "APRSL",
		},
	},
	{
		Primary: "APPRAISER",
		Short:   "APPRSER",
		Alt: []string{
			"APPRAISER", "APPRSER", "APPRSR",
		},
	},
	{
		Primary: "APPRAISOR",
		Short:   "APPRSOR",
		Alt: []string{
			"APPRAISOR", "APPRSOR", "APPRSR",
		},
	},
	{
		Primary: "APPRENTICE",
		Short:   "APPRNTC",
		Alt: []string{
			"APPRENTICE", "APPRNTC",
		},
	},
	{
		Primary: "APPROACHER",
		Short:   "APPRCHR",
		Alt: []string{
			"APPRCHR", "APPROACHER",
		},
	},
	{
		Primary: "ARABIAN",
		Short:   "ARBN",
		Alt: []string{
			"ARABIAN", "ARBN",
		},
	},
	{
		Primary: "ARCADE",
		Short:   "ARC",
		Alt: []string{
			"ARC", "ARCADE", "ARCD",
		},
	},
	{
		Primary: "ARCHBISHOP",
		Short:   "ABP",
		Alt: []string{
			"AB", "ABP", "ARCHBISHOP", "ARCHS",
		},
	},
	{
		Primary: "ARCHERY",
		Short:   "ARCHRY",
		Alt: []string{
			"ARCHERY", "ARCHRY",
		},
	},
	{
		Primary: "ARCHITECT",
		Short:   "ARCHT",
		Alt: []string{
			"ARCHITECT", "ARCHT", "ARCHTCT",
		},
	},
	{
		Primary: "ARCHITECTURAL",
		Short:   "ARCHL",
		Alt: []string{
			"ARCH", "ARCHITECTURAL", "ARCHL",
		},
	},
	{
		Primary: "ARCHITECTURE",
		Short:   "ARCH",
		Alt: []string{
			"ARCH", "ARCHITECTURE",
		},
	},
	{
		Primary: "ARCHIVE",
		Short:   "ARCHV",
		Alt: []string{
			"ARCHIVE", "ARCHV",
		},
	},
	{
		Primary: "ARENA",
		Short:   "ARN",
		Alt: []string{
			"ARENA", "ARN",
		},
	},
	{
		Primary: "ARISTOCRAT",
		Short:   "ARSTCRT",
		Alt: []string{
			"ARISTOCAT", "ARISTOCRAT", "ARSTCRT",
		},
	},
	{
		Primary: "ARMADILLO",
		Short:   "ARMDLL",
		Alt: []string{
			"ARMADILLO", "ARMDLL",
		},
	},
	{
		Primary: "ARMATURE",
		Short:   "ARMTR",
		Alt: []string{
			"ARMATURE", "ARMTR",
		},
	},
	{
		Primary: "ARMED",
		Short:   "ARMD",
		Alt: []string{
			"ARMD", "ARMED",
		},
	},
	{
		Primary: "ARMORED",
		Short:   "ARMRD",
		Alt: []string{
			"ARMORED", "ARMRD",
		},
	},
	{
		Primary: "ARMORY",
		Short:   "ARMRY",
		Alt: []string{
			"ARMORY", "ARMRY",
		},
	},
	{
		Primary: "ARROW",
		Short:   "ARW",
		Alt: []string{
			"ARROW", "ARW",
		},
	},
	{
		Primary: "ARSENAL",
		Short:   "ARSNL",
		Alt: []string{
			"ARSENAL", "ARSNL",
		},
	},
	{
		Primary: "ARTERY",
		Short:   "ARTRY",
		Alt: []string{
			"ARTERY", "ARTRY",
		},
	},
	{
		Primary: "ARTIFICIAL",
		Short:   "ARTFL",
		Alt: []string{
			"ARTFL", "ARTIFCAL", "ARTIFICIAL",
		},
	},
	{
		Primary: "ARTISAN",
		Short:   "ARTSN",
		Alt: []string{
			"ARTISAN", "ARTSN",
		},
	},
	{
		Primary: "ARTIST",
		Short:   "ART",
		Alt: []string{
			"ART", "ARTIST",
		},
	},
	{
		Primary: "ARTISTIC",
		Short:   "ARTSTC",
		Alt: []string{
			"ARTISTIC", "ARTSTC",
		},
	},
	{
		Primary: "ARTISTRY",
		Short:   "ARTSTRY",
		Alt: []string{
			"ARTISTRY", "ARTSTRY",
		},
	},
	{
		Primary: "ASBESTOS",
		Short:   "ASB",
		Alt: []string{
			"ASB", "ASBESTOS",
		},
	},
	{
		Primary: "ASPHALT",
		Short:   "ASPHLT",
		Alt: []string{
			"ASP", "ASPHALT", "ASPHLT",
		},
	},
	{
		Primary: "ASSEMBLE",
		Short:   "ASSMBL",
		Alt:     []string{"ASSEMBLE"},
	},
	{
		Primary: "ASSEMBLER",
		Short:   "ASSMBLR",
		Alt: []string{
			"ASSEMBLER", "ASSMBLR",
		},
	},
	{
		Primary: "ASSEMBLY",
		Short:   "ASMBLY",
		Alt: []string{
			"ASMBLY", "ASSEM", "ASSEMBLY",
		},
	},
	{
		Primary: "ASSET",
		Short:   "ASST",
		Alt: []string{
			"ASSET", "ASST",
		},
	},
	{
		Primary: "ASSIGNEE",
		Short:   "ASSGN",
		Alt: []string{
			"ASSGN", "ASSIGNEE",
		},
	},
	{
		Primary: "ASSISTANCE",
		Short:   "ASSTNCE",
		Alt: []string{
			"ASSISTANCE", "ASSTNCE",
		},
	},
	{
		Primary: "ASSISTANT",
		Short:   "ASSIST",
		Alt: []string{
			"ASSIST", "ASSISTANT", "ASST", "AST",
		},
	},
	{
		Primary: "ASSOCIATE",
		Short:   "ASSOC",
		Alt: []string{
			"ASO", "ASOC", "ASS", "ASSC", "ASSCE",
			"ASSO", "ASSOC", "ASSOCATE", "ASSOCI", "ASSOCIA",
			"ASSOCIAT", "ASSOCIATE", "ASST",
		},
	},
	{
		Primary: "ASSOCIATED",
		Short:   "ASSOCD",
		Alt: []string{
			"ASOC", "ASSCD", "ASSOC", "ASSOCATED", "ASSOCD",
			"ASSOCIATED", "ASSOD",
		},
	},
	{
		Primary: "ASSOCIATION",
		Short:   "ASSN",
		Alt: []string{
			"ASSCO", "ASSN", "ASSOC", "ASSOCIATION",
		},
	},
	{
		Primary: "ASSUMPTION",
		Short:   "ASSMPTN",
		Alt: []string{
			"ASSMPTN", "ASSUMPTION",
		},
	},
	{
		Primary: "ASSURANCE",
		Short:   "ASSURNC",
		Alt: []string{
			"ASRN", "ASSRNC", "ASSUR", "ASSURANCE", "ASSURNC",
		},
	},
	{
		Primary: "ASSURE",
		Short:   "ASSUR",
		Alt: []string{
			"ASSR", "ASSUR", "ASSURE",
		},
	},
	{
		Primary: "ASTRONAUTIC",
		Short:   "ASTRNTC",
		Alt: []string{
			"ASTRNTC", "ASTRONAUTIC",
		},
	},
	{
		Primary: "ATHLETIC",
		Short:   "ATHL",
		Alt: []string{
			"ATHC", "ATHL", "ATHLETIC",
		},
	},
	{
		Primary: "ATLANTIC",
		Short:   "ATL",
		Alt: []string{
			"ATL", "ATLANTIC", "ATLNTC",
		},
	},
	{
		Primary: "ATLAS",
		Short:   "ATLS",
		Alt: []string{
			"ATLAS", "ATLS",
		},
	},
	{
		Primary: "ATOMIC",
		Short:   "ATMC",
		Alt: []string{
			"ATMC", "ATOMIC",
		},
	},
	{
		Primary: "ATTACHE",
		Short:   "ATT",
		Alt: []string{
			"ATT", "ATTACHE",
		},
	},
	{
		Primary: "ATTENDANT",
		Short:   "ATTNDNT",
		Alt: []string{
			"ATTENDANT", "ATTNDNT",
		},
	},
	{
		Primary: "ATTENTION",
		Short:   "ATTN",
		Alt: []string{
			"ATN", "ATT", "ATTENTION", "ATTN", "ATTNTN",
		},
	},
	{
		Primary: "ATTIC",
		Short:   "ATTC",
		Alt: []string{
			"ATTC", "ATTIC",
		},
	},
	{
		Primary: "ATTITUDE",
		Short:   "ATTTD",
		Alt: []string{
			"ATTITUDE", "ATTTD",
		},
	},
	{
		Primary: "ATTORNEY",
		Short:   "ATTY",
		Alt: []string{
			"AT", "ATRNY", "ATT", "ATTNY", "ATTORNEY",
			"ATTY", "ATY",
		},
	},
	{
		Primary: "AUCTION",
		Short:   "AUCT",
		Alt: []string{
			"AUCT", "AUCTION", "AUCTN",
		},
	},
	{
		Primary: "AUCTIONEER",
		Short:   "AUCTNR",
		Alt: []string{
			"AUCTIONEER", "AUCTNR",
		},
	},
	{
		Primary: "AUCTIONEERING",
		Short:   "ACTNRG",
		Alt: []string{
			"ACTNRG", "AUCTIONEERING",
		},
	},
	{
		Primary: "AUDIO",
		Short:   "AUD",
		Alt: []string{
			"AUD", "AUDIO",
		},
	},
	{
		Primary: "AUDIOLOGIST",
		Short:   "AUDLGST",
		Alt: []string{
			"AUD", "AUDIOLOGIST", "AUDLGST",
		},
	},
	{
		Primary: "AUDIOLOGY",
		Short:   "AUDLGY",
		Alt: []string{
			"AUD", "AUDIOLOGY", "AUDLGY",
		},
	},
	{
		Primary: "AUDIOPROTHEISISTE",
		Short:   "AUDIOPR",
		Alt: []string{
			"AUD", "AUDIOPR", "AUDIOPROTH", "AUDIOPROTHEISISTE", "AUDPROT",
		},
	},
	{
		Primary: "AUDIT",
		Short:   "AUDT",
		Alt: []string{
			"AUD", "AUDIT", "AUDT",
		},
	},
	{
		Primary: "AUDITING",
		Short:   "ADTNG",
		Alt: []string{
			"ADTNG", "AUDITING",
		},
	},
	{
		Primary: "AUDITOR",
		Short:   "AUDTR",
		Alt: []string{
			"ADTR", "AUD", "AUDITOR", "AUDTR",
		},
	},
	{
		Primary: "AUDITORIUM",
		Short:   "ADTRM",
		Alt: []string{
			"ADTRM", "AUDITORIUM",
		},
	},
	{
		Primary: "AUTHORITY",
		Short:   "ATHRTY",
		Alt: []string{
			"ATHRTY", "AUT", "AUTH", "AUTHORI", "AUTHORITY",
			"AUTHY",
		},
	},
	{
		Primary: "AUTOMATED",
		Short:   "AUTOM",
		Alt: []string{
			"AUTOM", "AUTOMATED",
		},
	},
	{
		Primary: "AUTOMATIC",
		Short:   "AUTOMTC",
		Alt: []string{
			"AUTMTC", "AUTO", "AUTOMATIC", "AUTOMTC",
		},
	},
	{
		Primary: "AUTOMATION",
		Short:   "AUTOMTN",
		Alt: []string{
			"ATMTN", "AUTO", "AUTOMATION", "AUTOMTN",
		},
	},
	{
		Primary: "AUTOMOBILE",
		Short:   "AUTO",
		Alt: []string{
			"AUTO", "AUTOMOBILE",
		},
	},
	{
		Primary: "AUTOMOTIVE",
		Short:   "AUTOMTV",
		Alt: []string{
			"AUT", "AUTO", "AUTOMOTIVE", "AUTOMTV",
		},
	},
	{
		Primary: "AUXILIARY",
		Short:   "AUX",
		Alt: []string{
			"AUX", "AUXIL", "AUXILARY", "AUXILIARY", "AUXILRY",
		},
	},
	{
		Primary: "AVAILABILITY",
		Short:   "AVLBLTY",
		Alt: []string{
			"AVAILABILITY", "AVLBLTY",
		},
	},
	{
		Primary: "AVENUE",
		Short:   "AVE",
		Alt: []string{
			"AV", "AVE", "AVENUE",
		},
	},
	{
		Primary: "AVIATION",
		Short:   "AVN",
		Alt: []string{
			"AVI", "AVIATION", "AVN",
		},
	},
	{
		Primary: "AVIONIC",
		Short:   "AVNC",
		Alt: []string{
			"AVIONIC", "AVNC",
		},
	},
	{
		Primary: "AWARD",
		Short:   "AWRD",
		Alt: []string{
			"AWARD", "AWRD",
		},
	},
	{
		Primary: "AWNING",
		Short:   "AWN",
		Alt: []string{
			"AWN", "AWNG", "AWNING",
		},
	},
	{
		Primary: "BACHELOR",
		Short:   "BCHLR",
		Alt: []string{
			"BACHELOR", "BCHLR",
		},
	},
	{
		Primary: "BACKHOE",
		Short:   "BCKHOE",
		Alt: []string{
			"BACKHOE", "BCKHOE",
		},
	},
	{
		Primary: "BAGATELLE",
		Short:   "BGTTL",
		Alt: []string{
			"BAGATELLE", "BGTTL",
		},
	},
	{
		Primary: "BAILING",
		Short:   "BLG",
		Alt: []string{
			"BAILING", "BLG",
		},
	},
	{
		Primary: "BAKED",
		Short:   "BKD",
		Alt: []string{
			"BAKED", "BKD",
		},
	},
	{
		Primary: "BAKER",
		Short:   "BKR",
		Alt: []string{
			"BAKER", "BKR",
		},
	},
	{
		Primary: "BAKERY",
		Short:   "BKRY",
		Alt: []string{
			"BAKERY", "BKRY", "BKY",
		},
	},
	{
		Primary: "BAKING",
		Short:   "BKG",
		Alt: []string{
			"BAKING", "BKG",
		},
	},
	{
		Primary: "BALANCE",
		Short:   "BAL",
		Alt: []string{
			"BAL", "BALANCE",
		},
	},
	{
		Primary: "BALANCING",
		Short:   "BALNCNG",
		Alt: []string{
			"BALANCING", "BALNCNG",
		},
	},
	{
		Primary: "BALLER",
		Short:   "BLLR",
		Alt: []string{
			"BALLER", "BLLR",
		},
	},
	{
		Primary: "BALLOON",
		Short:   "BLN",
		Alt: []string{
			"BALLOON", "BLN",
		},
	},
	{
		Primary: "BALLROOM",
		Short:   "BLLRM",
		Alt: []string{
			"BALLROOM", "BLLRM",
		},
	},
	{
		Primary: "BANK",
		Short:   "BK",
		Alt: []string{
			"BANK", "BK",
		},
	},
	{
		Primary: "BANKER",
		Short:   "BNKR",
		Alt: []string{
			"BANKER", "BKR", "BNKR",
		},
	},
	{
		Primary: "BANKING",
		Short:   "BNKNG",
		Alt: []string{
			"BANKING", "BNKG", "BNKNG",
		},
	},
	{
		Primary: "BAPTIST",
		Short:   "BAPT",
		Alt: []string{
			"BAPT", "BAPTIST", "BPTST",
		},
	},
	{
		Primary: "BARBEQUE",
		Short:   "BBQ",
		Alt: []string{
			"BAR B Q", "BAR BQ", "BARBEQUE", "BARBQUE", "BBQ",
		},
	},
	{
		Primary: "BARBER",
		Short:   "BARB",
		Alt: []string{
			"BARB", "BARBER", "BARBR",
		},
	},
	{
		Primary: "BARGAIN",
		Short:   "BRGN",
		Alt: []string{
			"BARGAIN", "BRGN",
		},
	},
	{
		Primary: "BARREL",
		Short:   "BRL",
		Alt: []string{
			"BARREL", "BRL",
		},
	},
	{
		Primary: "BARRISTER",
		Short:   "BRRSTR",
		Alt: []string{
			"BARRISTER", "BRRSTR",
		},
	},
	{
		Primary: "BASEBALL",
		Short:   "BSBLL",
		Alt: []string{
			"BASEBALL", "BSBLL",
		},
	},
	{
		Primary: "BASEMENT",
		Short:   "BSMNT",
		Alt: []string{
			"BASEMENT", "BSMNT",
		},
	},
	{
		Primary: "BASIC",
		Short:   "BSC",
		Alt: []string{
			"BASIC", "BSC",
		},
	},
	{
		Primary: "BASKET",
		Short:   "BSK",
		Alt: []string{
			"BASKET", "BSK",
		},
	},
	{
		Primary: "BASKETBALL",
		Short:   "BSKTBLL",
		Alt: []string{
			"BASKETBALL", "BSKTBLL",
		},
	},
	{
		Primary: "BATTERY",
		Short:   "BATT",
		Alt: []string{
			"BATT", "BATTERY", "BTRY",
		},
	},
	{
		Primary: "BAZAAR",
		Short:   "BZR",
		Alt: []string{
			"BAZAAR", "BZR",
		},
	},
	{
		Primary: "BEACH",
		Short:   "BCH",
		Alt: []string{
			"BCH", "BEACH",
		},
	},
	{
		Primary: "BEARING",
		Short:   "BRNG",
		Alt: []string{
			"BEARING", "BRNG",
		},
	},
	{
		Primary: "BEAUTICIAN",
		Short:   "BTCN",
		Alt: []string{
			"BEAUTICIAN", "BTCN",
		},
	},
	{
		Primary: "BEAUTY",
		Short:   "BTY",
		Alt: []string{
			"BEAUTY", "BTY", "BUTY",
		},
	},
	{
		Primary: "BEAVER",
		Short:   "BVR",
		Alt: []string{
			"BEAVER", "BVR",
		},
	},
	{
		Primary: "BEDDING",
		Short:   "BEDG",
		Alt: []string{
			"BEDDING", "BEDG",
		},
	},
	{
		Primary: "BEGINNING",
		Short:   "BGNG",
		Alt: []string{
			"BEGINNING", "BGNG",
		},
	},
	{
		Primary: "BEHAVIORAL",
		Short:   "BHVRL",
		Alt: []string{
			"BEHAVIORAL", "BHVRL",
		},
	},
	{
		Primary: "BENEFICE",
		Short:   "BNFC",
		Alt: []string{
			"BENEFICE", "BNFC",
		},
	},
	{
		Primary: "BENEFICIAL",
		Short:   "BNFCL",
		Alt: []string{
			"BENEFICIAL", "BNFCL",
		},
	},
	{
		Primary: "BENEFIT",
		Short:   "BNFT",
		Alt: []string{
			"BENEFIT", "BNFT",
		},
	},
	{
		Primary: "BENEVOLENT",
		Short:   "BNVLNT",
		Alt: []string{
			"BENEVOLENT", "BNVLNT",
		},
	},
	{
		Primary: "BERRY",
		Short:   "BRY",
		Alt: []string{
			"BERRY", "BRY",
		},
	},
	{
		Primary: "BETTER",
		Short:   "BETR",
		Alt: []string{
			"BETR", "BETTER", "BTR",
		},
	},
	{
		Primary: "BEVERAGE",
		Short:   "BEV",
		Alt: []string{
			"BEV", "BEVERAGE",
		},
	},
	{
		Primary: "BIBLE",
		Short:   "BB",
		Alt: []string{
			"BB", "BIBLE",
		},
	},
	{
		Primary: "BICYCLE",
		Short:   "BIKE",
		Alt: []string{
			"BICYCLE", "BIKE",
		},
	},
	{
		Primary: "BIJOU",
		Short:   "BIJ",
		Alt: []string{
			"BIJ", "BIJOU",
		},
	},
	{
		Primary: "BIJOUTERIE",
		Short:   "BIJTR",
		Alt: []string{
			"BIJOUTERIE", "BIJTR",
		},
	},
	{
		Primary: "BILLETING",
		Short:   "BLLTNG",
		Alt: []string{
			"BILLETING", "BLLTNG",
		},
	},
	{
		Primary: "BILLIARD",
		Short:   "BILLD",
		Alt: []string{
			"BILLD", "BILLIARD",
		},
	},
	{
		Primary: "BILLING",
		Short:   "BLLNG",
		Alt: []string{
			"BILLING", "BLLNG",
		},
	},
	{
		Primary: "BINDER",
		Short:   "BDR",
		Alt: []string{
			"BDR", "BINDER",
		},
	},
	{
		Primary: "BINDERY",
		Short:   "BDRY",
		Alt: []string{
			"BDRY", "BINDERY",
		},
	},
	{
		Primary: "BINDING",
		Short:   "BDNG",
		Alt: []string{
			"BDNG", "BINDING",
		},
	},
	{
		Primary: "BINGO",
		Short:   "BNG",
		Alt: []string{
			"BINGO", "BNG",
		},
	},
	{
		Primary: "BIOCHEMISTRY",
		Short:   "BIOCHEM",
		Alt: []string{
			"BIOCHEM", "BIOCHEMISTRY",
		},
	},
	{
		Primary: "BIOLOGICAL",
		Short:   "BIOL",
		Alt: []string{
			"BIO", "BIOL", "BIOLGCL", "BIOLOGICAL",
		},
	},
	{
		Primary: "BIOLOGIST",
		Short:   "BIOGST",
		Alt: []string{
			"BIO", "BIOGST", "BIOL", "BIOLOGIST",
		},
	},
	{
		Primary: "BIOLOGY",
		Short:   "BIO",
		Alt: []string{
			"BIO", "BIOL", "BIOLOGY",
		},
	},
	{
		Primary: "BIOMEDICAL",
		Short:   "BIOMDCL",
		Alt: []string{
			"BIOMDCL", "BIOMEDICAL",
		},
	},
	{
		Primary: "BIONOMIC",
		Short:   "BIONMC",
		Alt: []string{
			"BIONMC", "BIONOMIC",
		},
	},
	{
		Primary: "BIOTECHNOLOGY",
		Short:   "BIOTECH",
		Alt: []string{
			"BIOTECH", "BIOTECHNOLOGY",
		},
	},
	{
		Primary: "BISCUIT",
		Short:   "BSCT",
		Alt: []string{
			"BISCUIT", "BSCT",
		},
	},
	{
		Primary: "BISHOP",
		Short:   "BP",
		Alt: []string{
			"BISHOP", "BP",
		},
	},
	{
		Primary: "BISTRO",
		Short:   "BSTR",
		Alt: []string{
			"BISTRO", "BSTR",
		},
	},
	{
		Primary: "BLACK",
		Short:   "BLCK",
		Alt: []string{
			"BLACK", "BLCK", "BLK",
		},
	},
	{
		Primary: "BLACKSMITH",
		Short:   "BSMITH",
		Alt: []string{
			"BLACKSMITH", "BSMITH",
		},
	},
	{
		Primary: "BLAZON",
		Short:   "BLZN",
		Alt: []string{
			"BLAZON", "BLZN",
		},
	},
	{
		Primary: "BLEND",
		Short:   "BLEN",
		Alt: []string{
			"BLEN", "BLEND",
		},
	},
	{
		Primary: "BLESSED",
		Short:   "BLSSD",
		Alt: []string{
			"BLESSED", "BLSSD",
		},
	},
	{
		Primary: "BLIND",
		Short:   "BLND",
		Alt: []string{
			"BLIND", "BLND",
		},
	},
	{
		Primary: "BLOCK",
		Short:   "BLK",
		Alt: []string{
			"BLK", "BLOCK",
		},
	},
	{
		Primary: "BLUEPRINT",
		Short:   "BLPRNT",
		Alt: []string{
			"BLPRNT", "BLUEPRINT",
		},
	},
	{
		Primary: "BOARD",
		Short:   "BD",
		Alt: []string{
			"BD", "BOARD", "BRD",
		},
	},
	{
		Primary: "BOARDING",
		Short:   "BRDNG",
		Alt: []string{
			"BOARDING", "BRDNG",
		},
	},
	{
		Primary: "BOMBER",
		Short:   "BMBR",
		Alt: []string{
			"BMBR", "BOMBER",
		},
	},
	{
		Primary: "BOOKBINDER",
		Short:   "BKBNDR",
		Alt: []string{
			"BKBNDR", "BOOKBINDER",
		},
	},
	{
		Primary: "BOOKBINDING",
		Short:   "BKBNDNG",
		Alt: []string{
			"BKBNDNG", "BOOKBINDING",
		},
	},
	{
		Primary: "BOOKKEEPER",
		Short:   "BKPR",
		Alt: []string{
			"BKKP", "BKKPR", "BKPR", "BOOKKEEPER",
		},
	},
	{
		Primary: "BOOKKEEPING",
		Short:   "BKPG",
		Alt: []string{
			"BKKP", "BKKPG", "BKKPNG", "BKPG", "BOOKKEEPING",
			"BOOKKPING",
		},
	},
	{
		Primary: "BOOKSELLER",
		Short:   "BKSLLR",
		Alt: []string{
			"BKSLLR", "BOOKSELLER",
		},
	},
	{
		Primary: "BOOKSHELF",
		Short:   "BKSHLF",
		Alt: []string{
			"BKSHLF", "BOOKSHELF",
		},
	},
	{
		Primary: "BOOKSTORE",
		Short:   "BKSTR",
		Alt: []string{
			"BKSTR", "BOOKSTOR", "BOOKSTORE",
		},
	},
	{
		Primary: "BOROUGH",
		Short:   "BORO",
		Alt: []string{
			"BORO", "BOROUGH",
		},
	},
	{
		Primary: "BOTTLED",
		Short:   "BOTLD",
		Alt: []string{
			"BOTLD", "BOTTLED",
		},
	},
	{
		Primary: "BOTTLER",
		Short:   "BTTLR",
		Alt: []string{
			"BOTTLER", "BTLR", "BTTLR",
		},
	},
	{
		Primary: "BOTTLING",
		Short:   "BTLG",
		Alt: []string{
			"BOTLNG", "BOTTLING", "BTG", "BTLG", "BTLNG",
		},
	},
	{
		Primary: "BOTTOM",
		Short:   "BTM",
		Alt: []string{
			"BOT", "BOTTOM", "BTM",
		},
	},
	{
		Primary: "BOULEVARD",
		Short:   "BLVD",
		Alt: []string{
			"BLVD", "BOULEVARD",
		},
	},
	{
		Primary: "BOUTIQUE",
		Short:   "BTQ",
		Alt: []string{
			"BOUTIQUE", "BTQ", "BTQUE",
		},
	},
	{
		Primary: "BOWLING",
		Short:   "BOWL",
		Alt: []string{
			"BOWL", "BOWLING",
		},
	},
	{
		Primary: "BRAIN",
		Short:   "BRN",
		Alt: []string{
			"BRAIN", "BRN",
		},
	},
	{
		Primary: "BRAKE",
		Short:   "BRK",
		Alt: []string{
			"BRAKE", "BRK",
		},
	},
	{
		Primary: "BRANCH",
		Short:   "BR",
		Alt: []string{
			"BR", "BRANCH", "BRCH", "BRNCH",
		},
	},
	{
		Primary: "BRASSERIE",
		Short:   "BRSSR",
		Alt: []string{
			"BRASSERIE", "BRSSR",
		},
	},
	{
		Primary: "BREEDER",
		Short:   "BRDR",
		Alt: []string{
			"BRDR", "BREEDER",
		},
	},
	{
		Primary: "BREWERY",
		Short:   "BRWRY",
		Alt: []string{
			"BREWERY", "BRWRY",
		},
	},
	{
		Primary: "BREWING",
		Short:   "BRWNG",
		Alt: []string{
			"BREWING", "BRWNG",
		},
	},
	{
		Primary: "BRICK",
		Short:   "BRCK",
		Alt: []string{
			"BRCK", "BRICK", "BRK",
		},
	},
	{
		Primary: "BRIDAL",
		Short:   "BRDL",
		Alt: []string{
			"BRDL", "BRIDAL",
		},
	},
	{
		Primary: "BRIDGE",
		Short:   "BRG",
		Alt: []string{
			"BDG", "BR", "BRDGE", "BRIDGE",
		},
	},
	{
		Primary: "BRIEF",
		Short:   "BRF",
		Alt: []string{
			"BRF", "BRIEF",
		},
	},
	{
		Primary: "BRIGADIER",
		Short:   "BRIG",
		Alt: []string{
			"BRIG", "BRIGADIER",
		},
	},
	{
		Primary: "BRIQUETTE",
		Short:   "BRQTT",
		Alt: []string{
			"BRIQUETTE", "BRQTT",
		},
	},
	{
		Primary: "BRITISH",
		Short:   "BRTSH",
		Alt: []string{
			"BRITISH", "BRTSH",
		},
	},
	{
		Primary: "BROADCAST",
		Short:   "BRDCST",
		Alt: []string{
			"BRDCST", "BROADCAST",
		},
	},
	{
		Primary: "BROADCASTER",
		Short:   "BRDCSTR",
		Alt: []string{
			"BRDCST", "BRDCSTR", "BROADCASTER",
		},
	},
	{
		Primary: "BROADCASTING",
		Short:   "BRDCSTG",
		Alt: []string{
			"BROADCASTING", "BROCSTG",
		},
	},
	{
		Primary: "BROADWAY",
		Short:   "BRDWY",
		Alt: []string{
			"BRDWY", "BROADWAY",
		},
	},
	{
		Primary: "BROKER",
		Short:   "BRKR",
		Alt: []string{
			"BRK", "BRKR", "BROKER",
		},
	},
	{
		Primary: "BROKERAGE",
		Short:   "BRKRGE",
		Alt: []string{
			"BRKG", "BRKRGE", "BROKERAGE",
		},
	},
	{
		Primary: "BROTHER",
		Short:   "BRO",
		Alt: []string{
			"BRO", "BROTHER",
		},
	},
	{
		Primary: "BROTHERHOOD",
		Short:   "BRTHD",
		Alt: []string{
			"BROTHERHOOD", "BRTHD",
		},
	},
	{
		Primary: "BROWN",
		Short:   "BRWN",
		Alt: []string{
			"BRN", "BROWN", "BRWN",
		},
	},
	{
		Primary: "BUCCANEER",
		Short:   "BCCNR",
		Alt: []string{
			"BCCNR", "BUCCANEER",
		},
	},
	{
		Primary: "BUCKET",
		Short:   "BCKT",
		Alt: []string{
			"BCKT", "BUCKET",
		},
	},
	{
		Primary: "BUCKEYE",
		Short:   "BCKEYE",
		Alt: []string{
			"BCKEYE", "BUCKEYE",
		},
	},
	{
		Primary: "BUDDY",
		Short:   "BDDY",
		Alt: []string{
			"BDDY", "BUDDY",
		},
	},
	{
		Primary: "BUDGET",
		Short:   "BGT",
		Alt: []string{
			"BDGT", "BGT", "BUDG", "BUDGET", "BUG",
			"BUGT",
		},
	},
	{
		Primary: "BUFFALO",
		Short:   "BUFF",
		Alt:     []string{"BUFFALO"},
	},
	{
		Primary: "BUILDER",
		Short:   "BLDR",
		Alt: []string{
			"BLDR", "BUILDER",
		},
	},
	{
		Primary: "BUILDING",
		Short:   "BLDG",
		Alt: []string{
			"BLD", "BLDG", "BUILDING",
		},
	},
	{
		Primary: "BUILT",
		Short:   "BLT",
		Alt: []string{
			"BLT", "BUILT",
		},
	},
	{
		Primary: "BULLDOZING",
		Short:   "BLLDZG",
		Alt: []string{
			"BLLDZG", "BULLDOZING",
		},
	},
	{
		Primary: "BULLET",
		Short:   "BLLT",
		Alt: []string{
			"BLLT", "BULLET",
		},
	},
	{
		Primary: "BULLETIN",
		Short:   "BLLTN",
		Alt: []string{
			"BLLTN", "BULLETIN",
		},
	},
	{
		Primary: "BUREAU",
		Short:   "BUR",
		Alt: []string{
			"BUR", "BUREAU",
		},
	},
	{
		Primary: "BURGER",
		Short:   "BGR",
		Alt: []string{
			"BGR", "BURGER",
		},
	},
	{
		Primary: "BURNING",
		Short:   "BRNNG",
		Alt: []string{
			"BRNNG", "BURNING",
		},
	},
	{
		Primary: "BURSAR",
		Short:   "BRSR",
		Alt: []string{
			"BRSR", "BURSAR",
		},
	},
	{
		Primary: "BUSINESS",
		Short:   "BUS",
		Alt: []string{
			"BSNS", "BUS", "BUSINES", "BUSINESS", "BUSN",
		},
	},
	{
		Primary: "BUTCHER",
		Short:   "BTCHR",
		Alt:     []string{"BUTCHER"},
	},
	{
		Primary: "BUTLER",
		Short:   "BTLR",
		Alt: []string{
			"BTLR", "BUTLER", "BUTLR",
		},
	},
	{
		Primary: "BUTTER",
		Short:   "BUTR",
		Alt: []string{
			"BTR", "BUTR", "BUTTER",
		},
	},
	{
		Primary: "BUTTON",
		Short:   "BUTN",
		Alt: []string{
			"BUTN", "BUTTON",
		},
	},
	{
		Primary: "BUYER",
		Short:   "BUYR",
		Alt: []string{
			"BUYER", "BYR",
		},
	},
	{
		Primary: "BYPASS",
		Short:   "BYP",
		Alt: []string{
			"BYP", "BYPASS",
		},
	},
	{
		Primary: "CABARET",
		Short:   "CBRT",
		Alt: []string{
			"CABARET", "CBRT",
		},
	},
	{
		Primary: "CABIN",
		Short:   "CBN",
		Alt: []string{
			"CABIN", "CBN",
		},
	},
	{
		Primary: "CABINET",
		Short:   "CBNT",
		Alt: []string{
			"CAB", "CABINET", "CBNT",
		},
	},
	{
		Primary: "CABINETMAKER",
		Short:   "CABMKR",
		Alt: []string{
			"CABINETMAKER", "CABMKR",
		},
	},
	{
		Primary: "CABLE",
		Short:   "CABL",
		Alt: []string{
			"CABL", "CABLE", "CBL",
		},
	},
	{
		Primary: "CADET",
		Short:   "CDT",
		Alt: []string{
			"CADET", "CDT",
		},
	},
	{
		Primary: "CADRE",
		Short:   "CDR",
		Alt: []string{
			"CADRE", "CDR",
		},
	},
	{
		Primary: "CAFETERIA",
		Short:   "CAFTRA",
		Alt: []string{
			"CAFETERIA", "CAFTRA", "CFTR",
		},
	},
	{
		Primary: "CALIPER",
		Short:   "CLPR",
		Alt: []string{
			"CALIPER", "CLPR",
		},
	},
	{
		Primary: "CALLIGRAPHER",
		Short:   "CLLGRPHR",
		Alt: []string{
			"CALLIGRAPHER", "CLLGRPHR",
		},
	},
	{
		Primary: "CALVARY",
		Short:   "CLVRY",
		Alt: []string{
			"CALV", "CALVARY", "CLVRY",
		},
	},
	{
		Primary: "CAMERA",
		Short:   "CAM",
		Alt: []string{
			"CAM", "CAMERA",
		},
	},
	{
		Primary: "CAMPAIGN",
		Short:   "CMPGN",
		Alt: []string{
			"CAMPAIGN", "CMPGN",
		},
	},
	{
		Primary: "CAMPER",
		Short:   "CMPR",
		Alt: []string{
			"CAMPER", "CMPR",
		},
	},
	{
		Primary: "CAMPGROUND",
		Short:   "CMPGRND",
		Alt: []string{
			"CAMPGROUND", "CMPGRND",
		},
	},
	{
		Primary: "CAMPING",
		Short:   "CMPNG",
		Alt: []string{
			"CAMPING", "CMPNG",
		},
	},
	{
		Primary: "CAMPSITE",
		Short:   "CMPST",
		Alt: []string{
			"CAMPSITE", "CMPST",
		},
	},
	{
		Primary: "CAMPUS",
		Short:   "CMPS",
		Alt: []string{
			"CAMPUS", "CMPS", "CMPUS",
		},
	},
	{
		Primary: "CANADIAN",
		Short:   "CNDN",
		Alt: []string{
			"CANADIAN", "CNDN",
		},
	},
	{
		Primary: "CANAL",
		Short:   "CNL",
		Alt: []string{
			"CANAL", "CNL",
		},
	},
	{
		Primary: "CANDLELIGHT",
		Short:   "CNDLLGHT",
		Alt: []string{
			"CANDLELIGHT", "CNDLLGHT",
		},
	},
	{
		Primary: "CANDY",
		Short:   "CNDY",
		Alt: []string{
			"CANDY", "CNDY",
		},
	},
	{
		Primary: "CANNERY",
		Short:   "CAN",
		Alt: []string{
			"CAN", "CANNERY",
		},
	},
	{
		Primary: "CANNING",
		Short:   "CNNNG",
		Alt: []string{
			"CANNING", "CNNNG",
		},
	},
	{
		Primary: "CANTONMENT",
		Short:   "CNTNMNT",
		Alt: []string{
			"CANTONMENT", "CNTNMNT",
		},
	},
	{
		Primary: "CANTOR",
		Short:   "CANTR",
		Alt: []string{
			"CANTOR", "CANTR", "CNTR",
		},
	},
	{
		Primary: "CANVAS",
		Short:   "CANV",
		Alt: []string{
			"CANV", "CANVAS",
		},
	},
	{
		Primary: "CANYON",
		Short:   "CYN",
		Alt: []string{
			"CANYON", "CYN",
		},
	},
	{
		Primary: "CAPITAL",
		Short:   "CPTAL",
		Alt: []string{
			"CAPITAL", "CPTAL", "CPTL",
		},
	},
	{
		Primary: "CAPITOL",
		Short:   "CPTOL",
		Alt: []string{
			"CAPITOL", "CPTL", "CPTOL",
		},
	},
	{
		Primary: "CAPTAIN",
		Short:   "CAPT",
		Alt: []string{
			"CAPT", "CAPTAIN", "CPT",
		},
	},
	{
		Primary: "CARBONATED",
		Short:   "CARB",
		Alt: []string{
			"CARB", "CARBONATED",
		},
	},
	{
		Primary: "CARBURETOR",
		Short:   "CARBTR",
		Alt: []string{
			"CARBTR", "CARBURETOR",
		},
	},
	{
		Primary: "CARDIAC",
		Short:   "CRDC",
		Alt: []string{
			"CARDIAC", "CRDC",
		},
	},
	{
		Primary: "CARDINAL",
		Short:   "CARD",
		Alt: []string{
			"CARD", "CARDINAL",
		},
	},
	{
		Primary: "CARDIOGRAPHIC",
		Short:   "CRDGRPHC",
		Alt: []string{
			"CARDIOGRAPHIC", "CRDGRPHC",
		},
	},
	{
		Primary: "CARDIOLOGY",
		Short:   "CRDLGY",
		Alt: []string{
			"CARDIOLOGY", "CRDLGY",
		},
	},
	{
		Primary: "CARDIOVASCULAR",
		Short:   "CRDVSCLR",
		Alt: []string{
			"CARDIOVASCULAR", "CRDVSCLR",
		},
	},
	{
		Primary: "CAREER",
		Short:   "CAR",
		Alt: []string{
			"CAR", "CAREER",
		},
	},
	{
		Primary: "CARGO",
		Short:   "CRG",
		Alt: []string{
			"CARGO", "CRG",
		},
	},
	{
		Primary: "CARIBBEAN",
		Short:   "CRBBN",
		Alt: []string{
			"CARIBBEAN", "CRBBN",
		},
	},
	{
		Primary: "CARLOADING",
		Short:   "CRLDNG",
		Alt: []string{
			"CARLOADING", "CRLDNG",
		},
	},
	{
		Primary: "CARPENTER",
		Short:   "CARPTR",
		Alt: []string{
			"CARPENTER", "CARPTR", "CPTR",
		},
	},
	{
		Primary: "CARPENTRY",
		Short:   "CRPNTRY",
		Alt: []string{
			"CARPENTRY", "CRPNTRY",
		},
	},
	{
		Primary: "CARPET",
		Short:   "CPT",
		Alt: []string{
			"CARPET", "CPT", "CRPT",
		},
	},
	{
		Primary: "CARRIAGE",
		Short:   "CARR",
		Alt: []string{
			"CARR", "CARRIAGE", "CGE",
		},
	},
	{
		Primary: "CASCADE",
		Short:   "CASC",
		Alt: []string{
			"CASC", "CASCADE",
		},
	},
	{
		Primary: "CASHIER",
		Short:   "CAS",
		Alt: []string{
			"CAS", "CASH", "CASHIER",
		},
	},
	{
		Primary: "CASKET",
		Short:   "CSKT",
		Alt: []string{
			"CASKET", "CSKT",
		},
	},
	{
		Primary: "CASSETTE",
		Short:   "CASSTT",
		Alt: []string{
			"CASSETTE", "CASSTT",
		},
	},
	{
		Primary: "CASTING",
		Short:   "CAST",
		Alt: []string{
			"CAST", "CASTING",
		},
	},
	{
		Primary: "CASTLE",
		Short:   "CASTL",
		Alt: []string{
			"CASTLE", "CSTL",
		},
	},
	{
		Primary: "CASUAL",
		Short:   "CSL",
		Alt: []string{
			"CASUAL", "CSL",
		},
	},
	{
		Primary: "CASUALTY",
		Short:   "CSLTY",
		Alt: []string{
			"CAS", "CASUALTY", "CSLTY",
		},
	},
	{
		Primary: "CATALOG",
		Short:   "CATLG",
		Alt: []string{
			"CATALOG", "CATLG", "CTLG",
		},
	},
	{
		Primary: "CATALOGUE",
		Short:   "CTLG",
		Alt: []string{
			"CATALOGUE", "CTLG",
		},
	},
	{
		Primary: "CATERER",
		Short:   "CATR",
		Alt: []string{
			"CATERER", "CATR",
		},
	},
	{
		Primary: "CATERING",
		Short:   "CTRG",
		Alt: []string{
			"CATERING", "CATRG", "CTRG",
		},
	},
	{
		Primary: "CATFISH",
		Short:   "CTFSH",
		Alt: []string{
			"CATFISH", "CTFSH",
		},
	},
	{
		Primary: "CATHEDRAL",
		Short:   "CATHDRL",
		Alt: []string{
			"CATH", "CATHDRL", "CATHEDRAL",
		},
	},
	{
		Primary: "CATHOLIC",
		Short:   "CATH",
		Alt: []string{
			"CATH", "CATHOLIC", "CTHLC",
		},
	},
	{
		Primary: "CATTLE",
		Short:   "CTTL",
		Alt: []string{
			"CATTLE", "CTTL",
		},
	},
	{
		Primary: "CAUSEWAY",
		Short:   "CSWY",
		Alt: []string{
			"CAUSEWAY", "CSWY",
		},
	},
	{
		Primary: "CEDAR",
		Short:   "CEDR",
		Alt: []string{
			"CDR", "CEDAR", "CEDR",
		},
	},
	{
		Primary: "CELEBRITY",
		Short:   "CLBRTY",
		Alt: []string{
			"CELEBRITY", "CLBRTY",
		},
	},
	{
		Primary: "CELLULAR",
		Short:   "CELL",
		Alt: []string{
			"CELL", "CELLULAR",
		},
	},
	{
		Primary: "CEMENT",
		Short:   "CEM",
		Alt: []string{
			"CEM", "CEMENT",
		},
	},
	{
		Primary: "CEMETERY",
		Short:   "CMTRY",
		Alt:     []string{"CEMETERY"},
	},
	{
		Primary: "CENTENNIAL",
		Short:   "CENT",
		Alt: []string{
			"CENT", "CENTENNAL", "CENTENNIAL", "CNTNNL",
		},
	},
	{
		Primary: "CENTER",
		Short:   "CTR",
		Alt: []string{
			"CEN", "CENT", "CENTER", "CENTR", "CNTR",
			"CTR",
		},
	},
	{
		Primary: "CENTRAL",
		Short:   "CTRL",
		Alt: []string{
			"CENTL", "CENTR", "CENTRAL", "CNTL", "CNTRL",
			"CTRL",
		},
	},
	{
		Primary: "CENTRE",
		Short:   "CTR",
		Alt: []string{
			"CENTRE", "CTR",
		},
	},
	{
		Primary: "CENTURY",
		Short:   "CEN",
		Alt: []string{
			"CEN", "CENTURY",
		},
	},
	{
		Primary: "CERAMIC",
		Short:   "CRMC",
		Alt: []string{
			"CERAMIC", "CRMC", "CRMIC",
		},
	},
	{
		Primary: "CEREMONY",
		Short:   "CRMNY",
		Alt: []string{
			"CEREMONY", "CRMNY",
		},
	},
	{
		Primary: "CERTIFICATION",
		Short:   "CTRFCTN",
		Alt: []string{
			"CERTIFICATION", "CTRFCTN",
		},
	},
	{
		Primary: "CERTIFIED",
		Short:   "CERT",
		Alt: []string{
			"CERTD", "CERTIF", "CERTIFIE", "CERTIFIED",
		},
	},
	{
		Primary: "CHAIN",
		Short:   "CHN",
		Alt: []string{
			"CH", "CHAIN", "CHN",
		},
	},
	{
		Primary: "CHAIR",
		Short:   "CHR",
		Alt: []string{
			"CHAIR", "CHR",
		},
	},
	{
		Primary: "CHAIRED",
		Short:   "CHRD",
		Alt: []string{
			"CHAIRED", "CHRD",
		},
	},
	{
		Primary: "CHAIRMAN",
		Short:   "CHRMN",
		Alt: []string{
			"CH", "CHAIR", "CHAIRMAN", "CHARMN", "CHM",
			"CHMN", "CHRM", "CHRMN",
		},
	},
	{
		Primary: "CHAIRPERSON",
		Short:   "CHRPRSN",
		Alt: []string{
			"CHAIRPERSON", "CHRPRSN",
		},
	},
	{
		Primary: "CHAIRWOMAN",
		Short:   "CHRWMN",
		Alt: []string{
			"CHAIRWOMAN", "CHRWMN",
		},
	},
	{
		Primary: "CHAMBER",
		Short:   "CHMBR",
		Alt: []string{
			"CHAMB", "CHAMBER", "CHMBR",
		},
	},
	{
		Primary: "CHAMPION",
		Short:   "CHAMP",
		Alt: []string{
			"CHAMP", "CHAMPION",
		},
	},
	{
		Primary: "CHANCELLOR",
		Short:   "CHANCLLR",
		Alt: []string{
			"CH", "CHAN", "CHANCELLOR", "CHANCLLR",
		},
	},
	{
		Primary: "CHANCELOR",
		Short:   "CHANCLR",
		Alt: []string{
			"CH", "CHAN", "CHANCELOR", "CHANCLR",
		},
	},
	{
		Primary: "CHANDLER",
		Short:   "CHANL",
		Alt: []string{
			"CHANDLER", "CHANL",
		},
	},
	{
		Primary: "CHANGE",
		Short:   "CHNG",
		Alt: []string{
			"CHANGE", "CHNG",
		},
	},
	{
		Primary: "CHANNEL",
		Short:   "CHNNL",
		Alt: []string{
			"CHANNEL", "CHNNL",
		},
	},
	{
		Primary: "CHAPEL",
		Short:   "CPL",
		Alt: []string{
			"CHAPEL", "CPL",
		},
	},
	{
		Primary: "CHAPLAIN",
		Short:   "CHAP",
		Alt: []string{
			"CHAP", "CHAPLAIN",
		},
	},
	{
		Primary: "CHAPTER",
		Short:   "CHPTR",
		Alt: []string{
			"CHAPTER", "CHPTR",
		},
	},
	{
		Primary: "CHARACTER",
		Short:   "CHAR",
		Alt: []string{
			"CHAR", "CHARACTER",
		},
	},
	{
		Primary: "CHARCOAL",
		Short:   "CHRCL",
		Alt: []string{
			"CHARCOAL", "CHRCL",
		},
	},
	{
		Primary: "CHARGE",
		Short:   "CHRG",
		Alt:     []string{"CHARGE"},
	},
	{
		Primary: "CHARITABLE",
		Short:   "CHRTBL",
		Alt: []string{
			"CHARITABLE", "CHRTBL",
		},
	},
	{
		Primary: "CHARTER",
		Short:   "CHRTR",
		Alt: []string{
			"CHAR", "CHARTER", "CHRTR",
		},
	},
	{
		Primary: "CHARTERED",
		Short:   "CHRTRD",
		Alt: []string{
			"CHARTERED", "CHRTRD",
		},
	},
	{
		Primary: "CHAUFFEUR",
		Short:   "CHFFR",
		Alt: []string{
			"CHAUFFEUR", "CHFFR",
		},
	},
	{
		Primary: "CHAUSSURE",
		Short:   "CHSSR",
		Alt: []string{
			"CHAUSSURE", "CHSSR",
		},
	},
	{
		Primary: "CHECK",
		Short:   "CHK",
		Alt: []string{
			"CHECK", "CHK",
		},
	},
	{
		Primary: "CHEESE",
		Short:   "CHES",
		Alt: []string{
			"CHEESE", "CHES", "CHS",
		},
	},
	{
		Primary: "CHEMICAL",
		Short:   "CHEML",
		Alt: []string{
			"CHEM", "CHEMICAL", "CHEML",
		},
	},
	{
		Primary: "CHEMIST",
		Short:   "CHEM",
		Alt: []string{
			"CHEM", "CHEMIST", "CHMST",
		},
	},
	{
		Primary: "CHERRY",
		Short:   "CHRY",
		Alt: []string{
			"CHERRY", "CHRY",
		},
	},
	{
		Primary: "CHESS",
		Short:   "CHSS",
		Alt: []string{
			"CHESS", "CHSS",
		},
	},
	{
		Primary: "CHESTNUT",
		Short:   "CHSTNT",
		Alt: []string{
			"CHESTNUT", "CHSTNT",
		},
	},
	{
		Primary: "CHEVROLET",
		Short:   "CHEVY",
		Alt: []string{
			"CHEVROLET", "CHEVY",
		},
	},
	{
		Primary: "CHICKEN",
		Short:   "CHICK",
		Alt: []string{
			"CHC", "CHCKN", "CHICK", "CHICKEN", "CHKN",
		},
	},
	{
		Primary: "CHIEF",
		Short:   "CHF",
		Alt: []string{
			"CHF", "CHIEF",
		},
	},
	{
		Primary: "CHILDREN",
		Short:   "CHLD",
		Alt: []string{
			"CHILDREN", "CHLD", "CHLDRN",
		},
	},
	{
		Primary: "CHILDRENS",
		Short:   "CHLDS",
		Alt: []string{
			"CHILD", "CHILDRENS",
		},
	},
	{
		Primary: "CHIMNEY",
		Short:   "CHMNY",
		Alt: []string{
			"CHIM", "CHIMNEY", "CHMNY",
		},
	},
	{
		Primary: "CHINESE",
		Short:   "CHIN",
		Alt: []string{
			"CHIN", "CHINESE",
		},
	},
	{
		Primary: "CHIROPRACTIC",
		Short:   "CHIROPRCTC",
		Alt: []string{
			"CHIRO", "CHIROPRAC", "CHIROPRACTIC", "CHIROPRCTC",
		},
	},
	{
		Primary: "CHIROPRACTOR",
		Short:   "CHIRO",
		Alt: []string{
			"CHIRO", "CHIROPRACTOR",
		},
	},
	{
		Primary: "CHOCOLATE",
		Short:   "CHOC",
		Alt: []string{
			"CHOC", "CHOCOLATE",
		},
	},
	{
		Primary: "CHOICE",
		Short:   "CHCE",
		Alt: []string{
			"CHCE", "CHOICE",
		},
	},
	{
		Primary: "CHOSE",
		Short:   "CHS",
		Alt: []string{
			"CHOSE", "CHS",
		},
	},
	{
		Primary: "CHRIST",
		Short:   "CHRST",
		Alt: []string{
			"CHR", "CHRIST", "CHRST",
		},
	},
	{
		Primary: "CHRISTIAN",
		Short:   "CHRSTN",
		Alt: []string{
			"CHR", "CHRIST", "CHRISTIAN", "CHRISTN", "CHRSTN",
		},
	},
	{
		Primary: "CHRONICLE",
		Short:   "CHRNCL",
		Alt: []string{
			"CHRNCL", "CHRONICLE",
		},
	},
	{
		Primary: "CHRYSLER",
		Short:   "CHRYSLR",
		Alt: []string{
			"CHRY", "CHRYSLER", "CHRYSLR",
		},
	},
	{
		Primary: "CHURCH",
		Short:   "CHURCH",
		Alt: []string{
			"CHR", "CHUR", "CHURC", "CHURCH",
		},
	},
	{
		Primary: "CIGAR",
		Short:   "CG",
		Alt: []string{
			"CG", "CIGAR",
		},
	},
	{
		Primary: "CIGARETTE",
		Short:   "CIG",
		Alt: []string{
			"CIG", "CIGARETTE",
		},
	},
	{
		Primary: "CINEMA",
		Short:   "CINE",
		Alt: []string{
			"CINE", "CINEMA",
		},
	},
	{
		Primary: "CIRCLE",
		Short:   "CIR",
		Alt: []string{
			"CIR", "CIRCLE", "CRCL",
		},
	},
	{
		Primary: "CIRCUIT",
		Short:   "CRCT",
		Alt: []string{
			"CIRCUIT", "CRCT",
		},
	},
	{
		Primary: "CIRCULAR",
		Short:   "CRCLR",
		Alt: []string{
			"CIRCULAR", "CRCLR",
		},
	},
	{
		Primary: "CIRCUS",
		Short:   "CRCS",
		Alt: []string{
			"CIRCUS", "CRCS",
		},
	},
	{
		Primary: "CIRQUE",
		Short:   "CRQ",
		Alt: []string{
			"CIRQUE", "CRQ",
		},
	},
	{
		Primary: "CITIZEN",
		Short:   "CITZN",
		Alt: []string{
			"CITIZEN", "CITZN", "CTZN",
		},
	},
	{
		Primary: "CITRUS",
		Short:   "CTRS",
		Alt: []string{
			"CITRUS", "CTRS",
		},
	},
	{
		Primary: "CIVIC",
		Short:   "CVC",
		Alt: []string{
			"CIVIC", "CVC",
		},
	},
	{
		Primary: "CIVIL",
		Short:   "CVL",
		Alt: []string{
			"CIVIL", "CVL",
		},
	},
	{
		Primary: "CLAIM",
		Short:   "CLM",
		Alt: []string{
			"CLAIM", "CLM",
		},
	},
	{
		Primary: "CLASS",
		Short:   "CLAS",
		Alt: []string{
			"CLAS", "CLASS",
		},
	},
	{
		Primary: "CLASSIC",
		Short:   "CLSC",
		Alt: []string{
			"CLASSIC", "CLSC",
		},
	},
	{
		Primary: "CLASSIFICATION",
		Short:   "CLASS",
		Alt: []string{
			"CLASS", "CLASSIFICATION", "CLSFCTN",
		},
	},
	{
		Primary: "CLEAN",
		Short:   "CLN",
		Alt: []string{
			"CLEAN", "CLN",
		},
	},
	{
		Primary: "CLEANER",
		Short:   "CLNR",
		Alt: []string{
			"CLEANER", "CLNR", "CLR",
		},
	},
	{
		Primary: "CLEANING",
		Short:   "CLNG",
		Alt: []string{
			"CLEANG", "CLEANING", "CLG", "CLNG",
		},
	},
	{
		Primary: "CLEANSER",
		Short:   "CLNSR",
		Alt: []string{
			"CLEANSER", "CLNSR",
		},
	},
	{
		Primary: "CLEARING",
		Short:   "CLRNG",
		Alt: []string{
			"CLEARING", "CLRNG",
		},
	},
	{
		Primary: "CLERGY",
		Short:   "CLER",
		Alt: []string{
			"CL", "CLER", "CLERGY",
		},
	},
	{
		Primary: "CLERK",
		Short:   "CLRK",
		Alt: []string{
			"CK", "CL", "CLERK", "CLK", "CLRK",
		},
	},
	{
		Primary: "CLIFF",
		Short:   "CLFS",
		Alt: []string{
			"CLF", "CLIFF",
		},
	},
	{
		Primary: "CLIMATE",
		Short:   "CLIMAT",
		Alt: []string{
			"CLIMAT", "CLIMATE",
		},
	},
	{
		Primary: "CLINIC",
		Short:   "CLNC",
		Alt: []string{
			"CL", "CLIN", "CLINI", "CLINIC", "CLNC",
		},
	},
	{
		Primary: "CLINICAL",
		Short:   "CLINIC",
		Alt: []string{
			"CLINIC", "CLINICA", "CLINICAL",
		},
	},
	{
		Primary: "CLIPPER",
		Short:   "CLPPR",
		Alt: []string{
			"CLIPPER", "CLPPR",
		},
	},
	{
		Primary: "CLOCK",
		Short:   "CLCK",
		Alt: []string{
			"CLCK", "CLK", "CLOCK",
		},
	},
	{
		Primary: "CLOSET",
		Short:   "CLOS",
		Alt: []string{
			"CLOS", "CLOSET",
		},
	},
	{
		Primary: "CLOTHES",
		Short:   "CLTHS",
		Alt: []string{
			"CLOS", "CLOTHES", "CLTHS",
		},
	},
	{
		Primary: "CLOTHIER",
		Short:   "CLTHR",
		Alt: []string{
			"CLOTHIER", "CLTHR",
		},
	},
	{
		Primary: "CLOTHING",
		Short:   "CLTHNG",
		Alt: []string{
			"CL", "CLOTHING", "CLTHNG",
		},
	},
	{
		Primary: "CLUBHOUSE",
		Short:   "CLBHS",
		Alt: []string{
			"CLBHS", "CLUBHOUSE",
		},
	},
	{
		Primary: "CLUTCH",
		Short:   "CLTCH",
		Alt: []string{
			"CLTCH", "CLUTCH",
		},
	},
	{
		Primary: "COACH",
		Short:   "CH",
		Alt: []string{
			"CCH", "COACH",
		},
	},
	{
		Primary: "COAST",
		Short:   "CST",
		Alt: []string{
			"COAST", "CST",
		},
	},
	{
		Primary: "COASTAL",
		Short:   "CSTL",
		Alt: []string{
			"COASTAL", "CSTL",
		},
	},
	{
		Primary: "COATED",
		Short:   "CTD",
		Alt: []string{
			"COATED", "CTD",
		},
	},
	{
		Primary: "COATING",
		Short:   "CTNG",
		Alt: []string{
			"COATING", "CTNG",
		},
	},
	{
		Primary: "COCKPIT",
		Short:   "CCKPT",
		Alt: []string{
			"CCKPT", "COCKPIT",
		},
	},
	{
		Primary: "COCOA",
		Short:   "CCO",
		Alt: []string{
			"CCO", "COCOA",
		},
	},
	{
		Primary: "COFFEE",
		Short:   "COF",
		Alt: []string{
			"COF", "COFFEE",
		},
	},
	{
		Primary: "COIFFEUR",
		Short:   "CFFR",
		Alt: []string{
			"CFFR", "COIFFEUR",
		},
	},
	{
		Primary: "COIFFEUSE",
		Short:   "CFFS",
		Alt: []string{
			"CFFS", "COIFFEUSE",
		},
	},
	{
		Primary: "COIFFURE",
		Short:   "COIFF",
		Alt: []string{
			"COIFF", "COIFFURE",
		},
	},
	{
		Primary: "COLLABORATIVE",
		Short:   "CLLBRTV",
		Alt: []string{
			"CLLBRTV", "COLL", "COLLABORATIVE",
		},
	},
	{
		Primary: "COLLATERAL",
		Short:   "CLLTRL",
		Alt: []string{
			"CLLTRL", "COLLATERAL",
		},
	},
	{
		Primary: "COLLECTABLE",
		Short:   "CLLCTABL",
		Alt: []string{
			"CLLCTABL", "CLLCTBL", "COLLECTABLE",
		},
	},
	{
		Primary: "COLLECTIBLE",
		Short:   "CLLCTIBL",
		Alt: []string{
			"CLLCTBL", "CLLCTIBL", "COLLECTIBLE",
		},
	},
	{
		Primary: "COLLECTION",
		Short:   "COLLECT",
		Alt: []string{
			"CLCTN", "COLLECT", "COLLECTION", "COLLECTN",
		},
	},
	{
		Primary: "COLLECTIVE",
		Short:   "CLLCTV",
		Alt: []string{
			"CLLCTV", "COLLECTIVE",
		},
	},
	{
		Primary: "COLLECT0R",
		Short:   "COLL",
		Alt: []string{
			"COLL", "COLLECT0R",
		},
	},
	{
		Primary: "COLLEGE",
		Short:   "COLG",
		Alt: []string{
			"CLG", "CLGE", "COL", "COLG", "COLL",
			"COLLEG", "COLLEGE",
		},
	},
	{
		Primary: "COLLEGIATE",
		Short:   "COLGT",
		Alt: []string{
			"COLGT", "COLLEGIATE",
		},
	},
	{
		Primary: "COLLISION",
		Short:   "CLLSN",
		Alt: []string{
			"CLLSN", "COLLISION",
		},
	},
	{
		Primary: "COLONEL",
		Short:   "COL",
		Alt: []string{
			"COL", "COLONEL",
		},
	},
	{
		Primary: "COLONIAL",
		Short:   "CLNL",
		Alt: []string{
			"CLNL", "COL", "COLONIAL",
		},
	},
	{
		Primary: "COLONY",
		Short:   "CLNY",
		Alt: []string{
			"CLNY", "COLONY",
		},
	},
	{
		Primary: "COLOR",
		Short:   "CLR",
		Alt: []string{
			"CLR", "COLOR",
		},
	},
	{
		Primary: "COLOSSAL",
		Short:   "CLSSL",
		Alt: []string{
			"CLSSL", "COLOSSAL",
		},
	},
	{
		Primary: "COMBINED",
		Short:   "COMB",
		Alt: []string{
			"COM", "COMB", "COMBINED",
		},
	},
	{
		Primary: "COMBUSTION",
		Short:   "COMBSTN",
		Alt: []string{
			"CMBSTN", "COMBSTN", "COMBUSTION",
		},
	},
	{
		Primary: "COMFORT",
		Short:   "CMFRT",
		Alt: []string{
			"CMFRT", "CMFT", "COMFORT",
		},
	},
	{
		Primary: "COMMAND",
		Short:   "CMND",
		Alt: []string{
			"CMND", "COM", "COMMAND",
		},
	},
	{
		Primary: "COMMANDANT",
		Short:   "COMDT",
		Alt: []string{
			"COM", "COMDT", "COMMANDANT", "COMMDT",
		},
	},
	{
		Primary: "COMMANDER",
		Short:   "CMDR",
		Alt: []string{
			"CDR", "CMDR", "COM", "COMM", "COMMANDER",
			"COMMDR",
		},
	},
	{
		Primary: "COMMANDING",
		Short:   "COMDG",
		Alt: []string{
			"COMDG", "COMMANDING",
		},
	},
	{
		Primary: "COMMENCEMENT",
		Short:   "COMMNCMNT",
		Alt: []string{
			"COMMENCEMENT", "COMMNCMNT",
		},
	},
	{
		Primary: "COMMERCE",
		Short:   "COMMRCE",
		Alt: []string{
			"CMMRC", "COMM", "COMMERC", "COMMERCE", "COMMRCE",
		},
	},
	{
		Primary: "COMMERCIAL",
		Short:   "COMRCL",
		Alt: []string{
			"CMRCL", "COMMERCIAL", "COMRCL",
		},
	},
	{
		Primary: "COMMISSARY",
		Short:   "COMSY",
		Alt: []string{
			"COMMISSARY", "COMSY",
		},
	},
	{
		Primary: "COMMISSION",
		Short:   "COMM",
		Alt: []string{
			"COMM", "COMMISSION",
		},
	},
	{
		Primary: "COMMISSIONER",
		Short:   "COMMR",
		Alt: []string{
			"COMMISSIONER", "COMMR",
		},
	},
	{
		Primary: "COMMITTEE",
		Short:   "CMMTE",
		Alt: []string{
			"CMMTE", "COM", "COMITE", "COMM", "COMMITTEE",
		},
	},
	{
		Primary: "COMMODITY",
		Short:   "COM",
		Alt: []string{
			"COM", "COMMODITY",
		},
	},
	{
		Primary: "COMMODORE",
		Short:   "COMD",
		Alt: []string{
			"COMD", "COMMODORE", "COMO",
		},
	},
	{
		Primary: "COMMON",
		Short:   "CMMN",
		Alt: []string{
			"CMMN", "COMMON",
		},
	},
	{
		Primary: "COMMONWEALTH",
		Short:   "CMNWLTH",
		Alt: []string{
			"CMNWLTH", "COMMONWEALTH", "COMMONWLTH",
		},
	},
	{
		Primary: "COMMUNE",
		Short:   "COMMN",
		Alt: []string{
			"COMMN", "COMMUNE",
		},
	},
	{
		Primary: "COMMUNICATE",
		Short:   "COMMUN",
		Alt: []string{
			"CCOMMNCTE", "COMM", "COMMUNICAT", "COMMUNICATE",
		},
	},
	{
		Primary: "COMMUNICATION",
		Short:   "COMMCTN",
		Alt: []string{
			"COMM", "COMMCTN", "COMMUN", "COMMUNICATI", "COMMUNICATION",
			"COMMUNICTN",
		},
	},
	{
		Primary: "COMMUNIQUE",
		Short:   "COMMNQ",
		Alt: []string{
			"COMMNQ", "COMMUNIQUE",
		},
	},
	{
		Primary: "COMMUNITY",
		Short:   "CMNTY",
		Alt: []string{
			"CMMNTY", "CMNTY", "CMTY", "COM", "COMM",
			"COMMUNITY", "COMNTY", "CTY",
		},
	},
	{
		Primary: "COMPANY",
		Short:   "CO",
		Alt: []string{
			"CO", "COMP", "COMPAN", "COMPANY", "COMPNY",
		},
	},
	{
		Primary: "COMPARATIVE",
		Short:   "COMPRTV",
		Alt: []string{
			"COMPARATIVE", "COMPRTV",
		},
	},
	{
		Primary: "COMPATIBLE",
		Short:   "COMPTBL",
		Alt: []string{
			"COMPATIBLE", "COMPTBL",
		},
	},
	{
		Primary: "COMPENSATION",
		Short:   "CMPNSTN",
		Alt: []string{
			"CMPNSTN", "COMPENSATION",
		},
	},
	{
		Primary: "COMPILER",
		Short:   "COMPLR",
		Alt: []string{
			"COMP", "COMPILER", "COMPLR",
		},
	},
	{
		Primary: "COMPLETE",
		Short:   "CMPLT",
		Alt: []string{
			"CMPLT", "COMPLET", "COMPLETE",
		},
	},
	{
		Primary: "COMPLEX",
		Short:   "COMPLX",
		Alt: []string{
			"COMPLEX", "COMPLX",
		},
	},
	{
		Primary: "COMPONENT",
		Short:   "COMPNNT",
		Alt: []string{
			"COMPNNT", "COMPONENT",
		},
	},
	{
		Primary: "COMPOSE",
		Short:   "COMPS",
		Alt: []string{
			"COMPOSE", "COMPS",
		},
	},
	{
		Primary: "COMPOSITE",
		Short:   "COMPST",
		Alt: []string{
			"COMPOSITE", "COMPST",
		},
	},
	{
		Primary: "COMPOSITION",
		Short:   "COMP",
		Alt: []string{
			"COMP", "COMPOSITION",
		},
	},
	{
		Primary: "COMPOUNDING",
		Short:   "COMPNDNG",
		Alt: []string{
			"COMPNDNG", "COMPOUNDING",
		},
	},
	{
		Primary: "COMPREHENSIVE",
		Short:   "CMPRHNSV",
		Alt: []string{
			"CMPRHNSV", "COMPREHENSIVE",
		},
	},
	{
		Primary: "COMPRESS",
		Short:   "COMPRSS",
		Alt: []string{
			"COMPRESS", "COMPRSS",
		},
	},
	{
		Primary: "COMPRESSOR",
		Short:   "CMPSR",
		Alt: []string{
			"CMPSR", "COMPRESSOR",
		},
	},
	{
		Primary: "COMPTABLE",
		Short:   "COMPTBLE",
		Alt: []string{
			"COMPTABLE", "COMPTBLE",
		},
	},
	{
		Primary: "COMPTROLLER",
		Short:   "COMPTLR",
		Alt: []string{
			"CMPTRLR", "COMP", "COMPT", "COMPTLR", "COMPTRLR",
			"COMPTROLL", "COMPTROLLER",
		},
	},
	{
		Primary: "COMPUTER",
		Short:   "CMPTR",
		Alt: []string{
			"CMP", "CMPTR", "COM", "COMP", "COMPTR",
			"COMPU", "COMPUTER",
		},
	},
	{
		Primary: "COMPUTERIZED",
		Short:   "COMPTRZD",
		Alt: []string{
			"COMPTRZD", "COMPUTERIZED",
		},
	},
	{
		Primary: "COMPUTING",
		Short:   "CMPTG",
		Alt: []string{
			"CMPTG", "COMPUTING",
		},
	},
	{
		Primary: "CONCENTRATE",
		Short:   "CONCNTRT",
		Alt: []string{
			"CON", "CONCENTRATE", "CONCNTRT",
		},
	},
	{
		Primary: "CONCEPT",
		Short:   "CNCPT",
		Alt: []string{
			"CNCPT", "CONCEPT",
		},
	},
	{
		Primary: "CONCESSION",
		Short:   "CONCSSN",
		Alt: []string{
			"CONCESSION", "CONCSSN",
		},
	},
	{
		Primary: "CONCOURSE",
		Short:   "CONCRS",
		Alt: []string{
			"CONCOURSE", "CONCRS",
		},
	},
	{
		Primary: "CONCRETE",
		Short:   "CONCRT",
		Alt: []string{
			"CON", "CONCRET", "CONCRETE", "CONCRT",
		},
	},
	{
		Primary: "CONDITIONING",
		Short:   "COND",
		Alt: []string{
			"CNDNTNG", "COND", "CONDITIONING",
		},
	},
	{
		Primary: "CONDOMINIUM",
		Short:   "CONDO",
		Alt: []string{
			"CNDMNM", "CONDO", "CONDOMINIUM",
		},
	},
	{
		Primary: "CONFECTIONERY",
		Short:   "CONF",
		Alt: []string{
			"CONF", "CONFECTIONERY",
		},
	},
	{
		Primary: "CONFEDERATED",
		Short:   "CONFDRTD",
		Alt: []string{
			"CONFDRTD", "CONFEDERATED",
		},
	},
	{
		Primary: "CONFEDERATION",
		Short:   "CONFDRTN",
		Alt: []string{
			"CONFDRTN", "CONFEDERATION",
		},
	},
	{
		Primary: "CONFER",
		Short:   "CNFR",
		Alt: []string{
			"CNFR", "CONFER",
		},
	},
	{
		Primary: "CONFERENCE",
		Short:   "CNFRNC",
		Alt: []string{
			"CNFRNC", "CONFERENCE", "CONFRENCE",
		},
	},
	{
		Primary: "CONGREGATION",
		Short:   "CONGREG",
		Alt: []string{
			"CONGREG", "CONGREGATION", "CONGRG",
		},
	},
	{
		Primary: "CONGREGATIONAL",
		Short:   "CONGREGTNL",
		Alt: []string{
			"CONGREGATIONAL", "CONGREGTNL",
		},
	},
	{
		Primary: "CONGRESS",
		Short:   "CNGRS",
		Alt: []string{
			"CNGRS", "CONGRESS",
		},
	},
	{
		Primary: "CONGRESSIONAL",
		Short:   "CNGRSNL",
		Alt: []string{
			"CNGRSNL", "CONGRESSIONAL",
		},
	},
	{
		Primary: "CONGRESSMAN",
		Short:   "CONGRSMAN",
		Alt: []string{
			"CONGRESSMAN", "CONGRSMAN",
		},
	},
	{
		Primary: "CONNECTION",
		Short:   "CONNECT",
		Alt: []string{
			"CONNECT", "CONNECTION",
		},
	},
	{
		Primary: "CONQUISTADOR",
		Short:   "CONQUISDR",
		Alt: []string{
			"CONQUISDR", "CONQUISTADOR",
		},
	},
	{
		Primary: "CONSERVATION",
		Short:   "CONSERVE",
		Alt: []string{
			"CNSRVTN", "CNSVTN", "CONSER", "CONSERV", "CONSERVATION",
			"CONSERVE",
		},
	},
	{
		Primary: "CONSERVATORY",
		Short:   "CONSRVTRY",
		Alt: []string{
			"CONSERVATORY", "CONSRVTRY",
		},
	},
	{
		Primary: "CONSOLATION",
		Short:   "CONSLTN",
		Alt: []string{
			"CONSLTN", "CONSOLATION",
		},
	},
	{
		Primary: "CONSOLIDATED",
		Short:   "CONS",
		Alt: []string{
			"CNSLD", "CNSLDTD", "CONS", "CONSOLIDATED",
		},
	},
	{
		Primary: "CONSOLIDATION",
		Short:   "CONSLDTN",
		Alt: []string{
			"CONSLDTN", "CONSOLIDATION",
		},
	},
	{
		Primary: "CONSOLIDATOR",
		Short:   "CONSLDTR",
		Alt: []string{
			"CONSLDTR", "CONSOLIDATOR",
		},
	},
	{
		Primary: "CONSORTIUM",
		Short:   "CNSRTM",
		Alt: []string{
			"CNSRTM", "CONSORTIUM",
		},
	},
	{
		Primary: "CONSTRUCT",
		Short:   "CONSTRCT",
		Alt: []string{
			"CONSTRCT", "CONSTRUCT",
		},
	},
	{
		Primary: "CONSTRUCTING",
		Short:   "CNSTRCTNG",
		Alt: []string{
			"CNSTRCTNG", "CONSTG", "CONSTRUCTING",
		},
	},
	{
		Primary: "CONSTRUCTION",
		Short:   "CONSTRCTN",
		Alt: []string{
			"CNST", "CNSTCONSTRCTN", "CNSTR", "CONSTN", "CONSTR",
			"CONSTRCTN", "CONSTRN", "CONSTRTN", "CONSTRUCTION", "CONSTRUCTN",
		},
	},
	{
		Primary: "CONSTRUCTOR",
		Short:   "CONSTR",
		Alt: []string{
			"CNSTR", "CONSTR", "CONSTRUCTOR",
		},
	},
	{
		Primary: "CONSULT",
		Short:   "CON",
		Alt: []string{
			"CON", "CONSULT",
		},
	},
	{
		Primary: "CONSULTANT",
		Short:   "CONSLNT",
		Alt: []string{
			"CNSLT", "CNSLTNT", "CON", "CONS", "CONSL",
			"CONSLTNT", "CONSULT", "CONSULTA", "CONSULTAN", "CONSULTANT",
			"CONSULTNT",
		},
	},
	{
		Primary: "CONSULTATION",
		Short:   "CNSLTN",
		Alt: []string{
			"CNSLTN", "CONSULTATION",
		},
	},
	{
		Primary: "CONSULTING",
		Short:   "CONSLTNG",
		Alt: []string{
			"CNSLTNG", "CONSLNTNG", "CONSLTG", "CONSLTNG", "CONSULTI",
			"CONSULTIN", "CONSULTING",
		},
	},
	{
		Primary: "CONSUMER",
		Short:   "CONSMR",
		Alt: []string{
			"CNSMR", "CONS", "CONSMR", "CONSUMER",
		},
	},
	{
		Primary: "CONTACT",
		Short:   "CONT",
		Alt: []string{
			"CONT", "CONTACT",
		},
	},
	{
		Primary: "CONTAIN",
		Short:   "CNTN",
		Alt: []string{
			"CNTN", "CONTAIN",
		},
	},
	{
		Primary: "CONTAINER",
		Short:   "CONTNR",
		Alt: []string{
			"CONTAINER", "CONTNR",
		},
	},
	{
		Primary: "CONTEMPORARY",
		Short:   "CONTEMP",
		Alt: []string{
			"CONTEMP", "CONTEMPO", "CONTEMPOR", "CONTEMPORAR", "CONTEMPORARY",
		},
	},
	{
		Primary: "CONTEST",
		Short:   "CNTST",
		Alt: []string{
			"CNTST", "CONTEST",
		},
	},
	{
		Primary: "CONTINENTAL",
		Short:   "CONTNTL",
		Alt: []string{
			"CNTNTL", "CONT", "CONTINENT", "CONTINENTAL", "CONTINENTL",
			"CONTNENTA", "CONTNTL",
		},
	},
	{
		Primary: "CONTINUING",
		Short:   "CONTNG",
		Alt: []string{
			"CONTINUING", "CONTNG",
		},
	},
	{
		Primary: "CONTINUOUS",
		Short:   "CONTNS",
		Alt: []string{
			"CONTINUOUS", "CONTNS",
		},
	},
	{
		Primary: "CONTRACT",
		Short:   "CNTRCT",
		Alt: []string{
			"CNTR", "CNTRCT", "CONTR", "CONTRAC", "CONTRACT",
		},
	},
	{
		Primary: "CONTRACTING",
		Short:   "CNTRCTNG",
		Alt: []string{
			"CNTRCTNG", "CONTG", "CONTR", "CONTRACTIN", "CONTRACTING",
			"CONTRG",
		},
	},
	{
		Primary: "CONTRACTOR",
		Short:   "CONTR",
		Alt: []string{
			"CNTRCTR", "CONTR", "CONTRACTOR", "COR",
		},
	},
	{
		Primary: "CONTRIBUTION",
		Short:   "CONTRBTN",
		Alt: []string{
			"CONTRBTN", "CONTRIBUTION",
		},
	},
	{
		Primary: "CONTROL",
		Short:   "CNTRL",
		Alt: []string{
			"CNTRL", "CONTRL", "CONTROL", "CTL", "CTRL",
		},
	},
	{
		Primary: "CONTROLLED",
		Short:   "CONTRLLD",
		Alt: []string{
			"CONTRLLD", "CONTROLLED",
		},
	},
	{
		Primary: "CONTROLLER",
		Short:   "CNTRLLR",
		Alt: []string{
			"CNTLR", "CNTR", "CNTRL", "CNTRLLR", "CNTRLR",
			"CONTLR", "CONTR", "CONTRLLR", "CONTRLR", "CONTROLER",
			"CONTROLL", "CONTROLLE", "CONTROLLER", "CONTROLLR", "CTL",
			"CTLR", "CTRLR",
		},
	},
	{
		Primary: "CONVALESCENT",
		Short:   "CONVAL",
		Alt: []string{
			"CONV", "CONVALESCEN", "CONVALESCENT",
		},
	},
	{
		Primary: "CONVENIENCE",
		Short:   "CONV",
		Alt: []string{
			"CONV", "CONVENIENCE",
		},
	},
	{
		Primary: "CONVENIENT",
		Short:   "CONVNT",
		Alt: []string{
			"CONVENIENT", "CONVNT",
		},
	},
	{
		Primary: "CONVENT",
		Short:   "CNVNT",
		Alt: []string{
			"CNVNT", "CONVENT", "CONVNT",
		},
	},
	{
		Primary: "CONVENTION",
		Short:   "CNVNTN",
		Alt: []string{
			"CNVNTN", "CONVENTION",
		},
	},
	{
		Primary: "CONVERSE",
		Short:   "CONVRS",
		Alt: []string{
			"CONVERSE", "CONVRS",
		},
	},
	{
		Primary: "CONVERSION",
		Short:   "CNVRSN",
		Alt: []string{
			"CNVRSN", "CONVERSION",
		},
	},
	{
		Primary: "CONVERTER",
		Short:   "CONVRTR",
		Alt: []string{
			"CONVERTER", "CONVRTR",
		},
	},
	{
		Primary: "CONVERTIBLE",
		Short:   "CONVRTBL",
		Alt: []string{
			"CONVERTIBLE", "CONVRTBL",
		},
	},
	{
		Primary: "CONVEYOR",
		Short:   "CONVYR",
		Alt: []string{
			"CONVEYOR", "CONVYR",
		},
	},
	{
		Primary: "COOKED",
		Short:   "CKD",
		Alt: []string{
			"CKD", "COOKED",
		},
	},
	{
		Primary: "COOKIE",
		Short:   "CK",
		Alt: []string{
			"CK", "COOKIE",
		},
	},
	{
		Primary: "COOLING",
		Short:   "COOL",
		Alt: []string{
			"COOL", "COOLG", "COOLING",
		},
	},
	{
		Primary: "COOPERATE",
		Short:   "COOP",
		Alt: []string{
			"COOP", "COOPERATE",
		},
	},
	{
		Primary: "COOPERATIVE",
		Short:   "COOPRTV",
		Alt: []string{
			"CO OP", "COOP", "COOPERATIVE", "COOPRTV",
		},
	},
	{
		Primary: "COORDINANT",
		Short:   "COORD",
		Alt: []string{
			"COORD", "COORDINANT",
		},
	},
	{
		Primary: "COORDINATE",
		Short:   "COORDNT",
		Alt: []string{
			"COORDINATE", "COORDNT",
		},
	},
	{
		Primary: "COORDINATOR",
		Short:   "COORDNTR",
		Alt: []string{
			"COOR", "COORD", "COORDINATOR", "COORDNTR",
		},
	},
	{
		Primary: "COPIER",
		Short:   "COPR",
		Alt: []string{
			"COPIER", "COPR",
		},
	},
	{
		Primary: "COPPER",
		Short:   "COP",
		Alt: []string{
			"COP", "COPPER",
		},
	},
	{
		Primary: "CORNER",
		Short:   "CORN",
		Alt: []string{
			"COR", "CORNER", "CORNR",
		},
	},
	{
		Primary: "CORONER",
		Short:   "COR",
		Alt: []string{
			"COR", "CORONER",
		},
	},
	{
		Primary: "CORPORAL",
		Short:   "CORPL",
		Alt: []string{
			"CORP", "CORPL", "CORPORAL", "CPL",
		},
	},
	{
		Primary: "CORPORATE",
		Short:   "CORPRT",
		Alt: []string{
			"CORP", "CORPORATE", "CORPORT", "CORPRT", "CRP",
		},
	},
	{
		Primary: "CORPORATION",
		Short:   "CORP",
		Alt: []string{
			"CORP", "CORPORATIN", "CORPORATIO", "CORPORATION",
		},
	},
	{
		Primary: "CORRECT",
		Short:   "CRRCT",
		Alt: []string{
			"CORRECT", "CRRCT",
		},
	},
	{
		Primary: "CORRECTION",
		Short:   "CRRCTN",
		Alt: []string{
			"CORRECTION", "CRRCTN",
		},
	},
	{
		Primary: "CORRECTIONAL",
		Short:   "CRRCTNL",
		Alt: []string{
			"CORCTNL", "CORRECTIONAL", "CRRCTNL",
		},
	},
	{
		Primary: "CORRESPONDENCE",
		Short:   "CORR",
		Alt: []string{
			"CORR", "CORRESPONDENCE",
		},
	},
	{
		Primary: "CORRESPONDENT",
		Short:   "CORRSPNDNT",
		Alt: []string{
			"COR", "CORR", "CORRESPONDENT", "CORRSPNDNT",
		},
	},
	{
		Primary: "CORRUGATED",
		Short:   "CORRGTD",
		Alt: []string{
			"CORRGTD", "CORRUGATED",
		},
	},
	{
		Primary: "COSMETIC",
		Short:   "CSMTC",
		Alt: []string{
			"COSMETIC", "COSMT", "CSMTC",
		},
	},
	{
		Primary: "COSMETOLOGIST",
		Short:   "CSMTLGST",
		Alt: []string{
			"COS", "COSMETOLOGIST", "CSMTLGST",
		},
	},
	{
		Primary: "COTTAGE",
		Short:   "CTG",
		Alt: []string{
			"COTTAGE", "CTG",
		},
	},
	{
		Primary: "COTTON",
		Short:   "COT",
		Alt: []string{
			"COT", "COTTON",
		},
	},
	{
		Primary: "COUNCIL",
		Short:   "CNCL",
		Alt: []string{
			"CL", "CNCL", "COUNCI", "COUNCIL",
		},
	},
	{
		Primary: "COUNCILING",
		Short:   "CNCLNG",
		Alt: []string{
			"CNCLNG", "COUNCILING",
		},
	},
	{
		Primary: "COUNSEL",
		Short:   "CNSL",
		Alt: []string{
			"CNSL", "COL", "COUNSEL",
		},
	},
	{
		Primary: "COUNSELING",
		Short:   "CNSLNG",
		Alt: []string{
			"CNSLNG", "COUNSELING",
		},
	},
	{
		Primary: "COUNSELLOR",
		Short:   "CNSLLR",
		Alt: []string{
			"CNSLLR", "CNSLR", "COUNSELLOR",
		},
	},
	{
		Primary: "COUNSELOR",
		Short:   "CNSLR",
		Alt: []string{
			"CNSLR", "COUNSELOR",
		},
	},
	{
		Primary: "COUNT",
		Short:   "CNT",
		Alt: []string{
			"CNT", "COUNT",
		},
	},
	{
		Primary: "COUNTER",
		Short:   "CNTR",
		Alt: []string{
			"CNTR", "COUNTER",
		},
	},
	{
		Primary: "COUNTRY",
		Short:   "CNTRY",
		Alt: []string{
			"CNTRY", "CO", "COUNTRY", "CTRY",
		},
	},
	{
		Primary: "COUNTRYSIDE",
		Short:   "CNTRYSD",
		Alt: []string{
			"CNTRYSD", "COUNTRYSIDE",
		},
	},
	{
		Primary: "COUNTY",
		Short:   "CNTY",
		Alt: []string{
			"CNTY", "CO", "COUNTY", "CTY",
		},
	},
	{
		Primary: "COUPE",
		Short:   "CP",
		Alt: []string{
			"COUPE", "CP",
		},
	},
	{
		Primary: "COURIER",
		Short:   "COUR",
		Alt: []string{
			"COUR", "COURIER", "COURIR",
		},
	},
	{
		Primary: "COURSE",
		Short:   "CRS",
		Alt: []string{
			"COURSE", "CRS", "CRSE",
		},
	},
	{
		Primary: "COURT",
		Short:   "CT",
		Alt: []string{
			"COURT", "CRT", "CT",
		},
	},
	{
		Primary: "COURTESY",
		Short:   "CRTSY",
		Alt: []string{
			"COURTESY", "CRTSY",
		},
	},
	{
		Primary: "COVENANT",
		Short:   "CVNNT",
		Alt: []string{
			"COVENANT", "CVNNT",
		},
	},
	{
		Primary: "COVERING",
		Short:   "COVER",
		Alt: []string{
			"COVER", "COVERING", "CVG", "CVRNG",
		},
	},
	{
		Primary: "COWBOY",
		Short:   "CWBY",
		Alt: []string{
			"COWBOY", "CWBY",
		},
	},
	{
		Primary: "CRAFT",
		Short:   "CRFT",
		Alt: []string{
			"CFT", "CRAFT", "CRFT",
		},
	},
	{
		Primary: "CRAFTER",
		Short:   "CFTR",
		Alt: []string{
			"CFTR", "CRAFTER",
		},
	},
	{
		Primary: "CRAFTSMAN",
		Short:   "CFT",
		Alt: []string{
			"CFT", "CRAFTSMAN",
		},
	},
	{
		Primary: "CRAFTSMEN",
		Short:   "CFTMN",
		Alt: []string{
			"CFTMN", "CRAFTSMEN",
		},
	},
	{
		Primary: "CRANBERRY",
		Short:   "CRNBRRY",
		Alt: []string{
			"CRANBERRY", "CRNBRRY",
		},
	},
	{
		Primary: "CRANE",
		Short:   "CRN",
		Alt: []string{
			"CRANE", "CRN",
		},
	},
	{
		Primary: "CRANKSHAFT",
		Short:   "CRNKSHFT",
		Alt: []string{
			"CRANKSHAFT", "CRNKSHFT",
		},
	},
	{
		Primary: "CRAZY",
		Short:   "CRZY",
		Alt: []string{
			"CRAZY", "CRZY",
		},
	},
	{
		Primary: "CREAM",
		Short:   "CRM",
		Alt: []string{
			"CREAM", "CRM",
		},
	},
	{
		Primary: "CREAMERY",
		Short:   "CRMRY",
		Alt: []string{
			"CREAMERY", "CRMRY",
		},
	},
	{
		Primary: "CREATION",
		Short:   "CREAT",
		Alt: []string{
			"CREAT", "CREATION",
		},
	},
	{
		Primary: "CREATIVE",
		Short:   "CREATV",
		Alt: []string{
			"CREAT", "CREATIVE", "CREATV", "CRTVE",
		},
	},
	{
		Primary: "CREDIT",
		Short:   "CRDT",
		Alt: []string{
			"CRDT", "CRED", "CREDIT",
		},
	},
	{
		Primary: "CREEK",
		Short:   "CRK",
		Alt: []string{
			"CREEK", "CRK",
		},
	},
	{
		Primary: "CREMATORY",
		Short:   "CRMTRY",
		Alt: []string{
			"CREMATORY", "CRMTRY",
		},
	},
	{
		Primary: "CREPE",
		Short:   "CRP",
		Alt: []string{
			"CREPE", "CRP",
		},
	},
	{
		Primary: "CRESCENT",
		Short:   "CRES",
		Alt: []string{
			"CRES", "CRESCENT",
		},
	},
	{
		Primary: "CREST",
		Short:   "CREST",
		Alt: []string{
			"CREST", "CRST",
		},
	},
	{
		Primary: "CRIMINAL",
		Short:   "CRMNL",
		Alt: []string{
			"CRIMINAL", "CRMNL",
		},
	},
	{
		Primary: "CROCKERY",
		Short:   "CKRY",
		Alt: []string{
			"CKRY", "CRK", "CROCKERY",
		},
	},
	{
		Primary: "CROSS",
		Short:   "CR",
		Alt: []string{
			"CR", "CROSS",
		},
	},
	{
		Primary: "CROSSING",
		Short:   "XING",
		Alt: []string{
			"CROSSING", "CRSSNG",
		},
	},
	{
		Primary: "CROSSROAD",
		Short:   "XROAD",
		Alt: []string{
			"CROSSRD", "CROSSROAD", "XRD", "XROAD",
		},
	},
	{
		Primary: "CROWN",
		Short:   "CRWN",
		Alt: []string{
			"CRN", "CROWN", "CRWN",
		},
	},
	{
		Primary: "CRUISE",
		Short:   "CRUS",
		Alt: []string{
			"CRS", "CRUISE", "CRUS",
		},
	},
	{
		Primary: "CRUSADE",
		Short:   "CRSD",
		Alt: []string{
			"CRSD", "CRUSADE",
		},
	},
	{
		Primary: "CRUSADER",
		Short:   "CRSDR",
		Alt: []string{
			"CRSDR", "CRUSADER",
		},
	},
	{
		Primary: "CRUST",
		Short:   "CRUST",
		Alt: []string{
			"CRST", "CRUST",
		},
	},
	{
		Primary: "CRYOGENIC",
		Short:   "CRYGNC",
		Alt: []string{
			"CRYGNC", "CRYOGENIC",
		},
	},
	{
		Primary: "CRYSTAL",
		Short:   "CRYSTL",
		Alt: []string{
			"CRYSTAL", "CRYSTL",
		},
	},
	{
		Primary: "CUISINE",
		Short:   "CSN",
		Alt: []string{
			"CSN", "CUISINE",
		},
	},
	{
		Primary: "CULTURAL",
		Short:   "CLTRL",
		Alt: []string{
			"CLTRL", "CULTURAL",
		},
	},
	{
		Primary: "CUPBOARD",
		Short:   "CPBRD",
		Alt: []string{
			"CPBRD", "CUPBOARD",
		},
	},
	{
		Primary: "CURATOR",
		Short:   "CUR",
		Alt: []string{
			"CUR", "CURATOR",
		},
	},
	{
		Primary: "CURRICULUM",
		Short:   "CURR",
		Alt: []string{
			"CURR", "CURRICULUM",
		},
	},
	{
		Primary: "CURTAIN",
		Short:   "CRTN",
		Alt: []string{
			"CRTN", "CURTAIN",
		},
	},
	{
		Primary: "CUSTODIAN",
		Short:   "CUSTDN",
		Alt: []string{
			"CUST", "CUSTDN", "CUSTODIAN",
		},
	},
	{
		Primary: "CUSTOM",
		Short:   "CSTM",
		Alt: []string{
			"CSTM", "CUST", "CUSTOM",
		},
	},
	{
		Primary: "CUSTOMER",
		Short:   "CUST",
		Alt: []string{
			"CUST", "CUSTOMER",
		},
	},
	{
		Primary: "CUTLERY",
		Short:   "CUTLY",
		Alt: []string{
			"CUTLERY", "CUTLY",
		},
	},
	{
		Primary: "CUTTING",
		Short:   "CUT",
		Alt: []string{
			"CUT", "CUTING", "CUTTING",
		},
	},
	{
		Primary: "CYBERNETIC",
		Short:   "CYBRNTC",
		Alt: []string{
			"CYBERNETIC", "CYBRNTC",
		},
	},
	{
		Primary: "CYCLE",
		Short:   "CYCL",
		Alt: []string{
			"CYCL", "CYCLE",
		},
	},
	{
		Primary: "DAILY",
		Short:   "DLY",
		Alt: []string{
			"DAILY", "DLY",
		},
	},
	{
		Primary: "DAIRY",
		Short:   "DRY",
		Alt: []string{
			"DAIRY", "DAR", "DRY",
		},
	},
	{
		Primary: "DAME",
		Short:   "DM",
		Alt: []string{
			"DAME", "DM",
		},
	},
	{
		Primary: "DANCE",
		Short:   "DNC",
		Alt: []string{
			"DANCE", "DNC",
		},
	},
	{
		Primary: "DATABASE",
		Short:   "DB",
		Alt: []string{
			"DATABASE", "DB",
		},
	},
	{
		Primary: "DATZUN",
		Short:   "DTZN",
		Alt: []string{
			"DATZUN", "DTZN",
		},
	},
	{
		Primary: "DAUGHTER",
		Short:   "DGHTR",
		Alt: []string{
			"DAUGHTER", "DGHTR",
		},
	},
	{
		Primary: "DEACON",
		Short:   "DCN",
		Alt: []string{
			"DCN", "DEACON",
		},
	},
	{
		Primary: "DEALER",
		Short:   "DLR",
		Alt: []string{
			"DEALER", "DLR",
		},
	},
	{
		Primary: "DEALING",
		Short:   "DLG",
		Alt: []string{
			"DEALING", "DLG",
		},
	},
	{
		Primary: "DECAL",
		Short:   "DEC",
		Alt: []string{
			"DEC", "DECAL",
		},
	},
	{
		Primary: "DECISION",
		Short:   "DCSN",
		Alt: []string{
			"DCSN", "DECISION",
		},
	},
	{
		Primary: "DECOR",
		Short:   "DCR",
		Alt: []string{
			"DCR", "DECOR",
		},
	},
	{
		Primary: "DECORATING",
		Short:   "DECOR",
		Alt: []string{
			"DCRTNG", "DCTG", "DECOR", "DECORATING",
		},
	},
	{
		Primary: "DECORATION",
		Short:   "DCTN",
		Alt: []string{
			"DCTN", "DECORATION",
		},
	},
	{
		Primary: "DECORATOR",
		Short:   "DCRTR",
		Alt: []string{
			"DCRTR", "DCTR", "DECORATOR",
		},
	},
	{
		Primary: "DEFENCE",
		Short:   "DEFNC",
		Alt: []string{
			"DEF", "DEFENCE", "DEFNC",
		},
	},
	{
		Primary: "DEFENSE",
		Short:   "DEFNS",
		Alt: []string{
			"DEFENSE", "DEFNS",
		},
	},
	{
		Primary: "DELICATESSEN",
		Short:   "DELI",
		Alt: []string{
			"DELI", "DELICATESSEN",
		},
	},
	{
		Primary: "DELIGHT",
		Short:   "DLGHT",
		Alt: []string{
			"DELIGHT", "DLGHT",
		},
	},
	{
		Primary: "DELINTING",
		Short:   "DLNTG",
		Alt: []string{
			"DELINTING", "DLNTG",
		},
	},
	{
		Primary: "DELIVERANCE",
		Short:   "DELVRNC",
		Alt: []string{
			"DELIVERANCE", "DELIVRANCE", "DELVRNC",
		},
	},
	{
		Primary: "DELIVERY",
		Short:   "DLVRY",
		Alt: []string{
			"DEL", "DELIVERY", "DLVRY",
		},
	},
	{
		Primary: "DELTA",
		Short:   "DLT",
		Alt: []string{
			"DELTA", "DLT",
		},
	},
	{
		Primary: "DEMOCRATIC",
		Short:   "DEM",
		Alt: []string{
			"DEM", "DEMOCRATIC",
		},
	},
	{
		Primary: "DEMOLITION",
		Short:   "DEMLTN",
		Alt: []string{
			"DEM", "DEMLTN", "DEMOLITION",
		},
	},
	{
		Primary: "DENTAL",
		Short:   "DNTL",
		Alt: []string{
			"DENTAL", "DNTL",
		},
	},
	{
		Primary: "DENTIST",
		Short:   "DDS",
		Alt: []string{
			"DDS", "DENT", "DENTIST",
		},
	},
	{
		Primary: "DENTISTRY",
		Short:   "DNTSTRY",
		Alt: []string{
			"DENTISTRY", "DNTSTRY",
		},
	},
	{
		Primary: "DENTURE",
		Short:   "DENTR",
		Alt: []string{
			"DENTR", "DENTURE", "DNTR",
		},
	},
	{
		Primary: "DEPARTMENT",
		Short:   "DEPT",
		Alt: []string{
			"DEP", "DEPART", "DEPARTM", "DEPARTMENT", "DEPARTMNT",
			"DEPT", "DPT",
		},
	},
	{
		Primary: "DEPENDABLE",
		Short:   "DPNDBL",
		Alt: []string{
			"DEPENDABLE", "DPNDBL",
		},
	},
	{
		Primary: "DEPOSIT",
		Short:   "DPST",
		Alt: []string{
			"DEPOSIT", "DPST",
		},
	},
	{
		Primary: "DEPOT",
		Short:   "DEP",
		Alt: []string{
			"DEP", "DEPOT", "DPT",
		},
	},
	{
		Primary: "DEPUTY",
		Short:   "DPTY",
		Alt: []string{
			"DEP", "DEPT", "DEPUTY", "DPTY",
		},
	},
	{
		Primary: "DERMATOLOGIST",
		Short:   "DERMTLGST",
		Alt: []string{
			"DERM", "DERMATOLOGIST", "DERMTLGST",
		},
	},
	{
		Primary: "DERMATOLOGY",
		Short:   "DERM",
		Alt: []string{
			"DERM", "DERMATOLOGY",
		},
	},
	{
		Primary: "DESERT",
		Short:   "DSRT",
		Alt: []string{
			"DESERT", "DSRT",
		},
	},
	{
		Primary: "DESIGN",
		Short:   "DSGN",
		Alt: []string{
			"DES", "DESIGN", "DSGN",
		},
	},
	{
		Primary: "DESIGNER",
		Short:   "DSGNR",
		Alt: []string{
			"DESGR", "DESIGNER", "DSGNR", "DSGR",
		},
	},
	{
		Primary: "DESIGNING",
		Short:   "DSGNG",
		Alt: []string{
			"DESIGNING", "DSGNG",
		},
	},
	{
		Primary: "DETAIL",
		Short:   "DTL",
		Alt: []string{
			"DETAIL", "DTL",
		},
	},
	{
		Primary: "DETECTIVE",
		Short:   "DET",
		Alt: []string{
			"DET", "DETECTIVE",
		},
	},
	{
		Primary: "DETENTION",
		Short:   "DETNTN",
		Alt:     []string{"DETENTION"},
	},
	{
		Primary: "DEVELOPER",
		Short:   "DVLPR",
		Alt: []string{
			"DEVELOPER", "DVLPR",
		},
	},
	{
		Primary: "DEVELOPMENT",
		Short:   "DEV",
		Alt: []string{
			"DEV", "DEVEL", "DEVELOP", "DEVELOPM", "DEVELOPMEN",
			"DEVELOPMENT", "DEVELOPMNT", "DEVELOPMT", "DEVELP", "DEVELPMT",
			"DEVLMNT", "DEVLPMNT", "DEVLPMT", "DEVMT", "DVLOPMT",
			"DVLPMNT", "DVLPMT",
		},
	},
	{
		Primary: "DEVELOPMENTAL",
		Short:   "DEVLPMNTL",
		Alt: []string{
			"DEVELOPMENTAL", "DEVLPMNTL",
		},
	},
	{
		Primary: "DEVICE",
		Short:   "DVC",
		Alt: []string{
			"DEVICE", "DVC",
		},
	},
	{
		Primary: "DIAGNOSTIC",
		Short:   "DGNSTC",
		Alt: []string{
			"DGNSTC", "DIAG", "DIAGNOSTIC",
		},
	},
	{
		Primary: "DIAMOND",
		Short:   "DMND",
		Alt: []string{
			"DIAMOND", "DMND",
		},
	},
	{
		Primary: "DIAPER",
		Short:   "DPR",
		Alt: []string{
			"DIAPER", "DPR",
		},
	},
	{
		Primary: "DICTATOR",
		Short:   "DICT",
		Alt: []string{
			"DICT", "DICTATOR",
		},
	},
	{
		Primary: "DIELECTRIC",
		Short:   "DLCTRC",
		Alt: []string{
			"DIELECTRIC", "DLCTRC",
		},
	},
	{
		Primary: "DIESEL",
		Short:   "DSL",
		Alt: []string{
			"DIESEL", "DSL",
		},
	},
	{
		Primary: "DIETARY",
		Short:   "DTRY",
		Alt: []string{
			"DIETARY", "DIETRY", "DTRY",
		},
	},
	{
		Primary: "DIETETIC",
		Short:   "DIETC",
		Alt: []string{
			"DIETC", "DIETEIC", "DIETETIC",
		},
	},
	{
		Primary: "DIFFERENT",
		Short:   "DIFF",
		Alt: []string{
			"DIFF", "DIFFERENT",
		},
	},
	{
		Primary: "DIFFUSION",
		Short:   "DIFFSN",
		Alt: []string{
			"DIFF", "DIFFSN", "DIFFUSION",
		},
	},
	{
		Primary: "DIGEST",
		Short:   "DGST",
		Alt: []string{
			"DGST", "DIGEST",
		},
	},
	{
		Primary: "DIGESTIVE",
		Short:   "DGSTV",
		Alt: []string{
			"DGSTV", "DIGESTIVE",
		},
	},
	{
		Primary: "DIGITAL",
		Short:   "DGTL",
		Alt: []string{
			"DGTL", "DIGITAL",
		},
	},
	{
		Primary: "DILIGENCE",
		Short:   "DLGNC",
		Alt: []string{
			"DILIGENCE", "DLGNC",
		},
	},
	{
		Primary: "DIMENSION",
		Short:   "DIM",
		Alt: []string{
			"DIM", "DIMENSION",
		},
	},
	{
		Primary: "DIMENSIONAL",
		Short:   "DIML",
		Alt: []string{
			"DIMENSIONAL", "DIML",
		},
	},
	{
		Primary: "DINER",
		Short:   "DNR",
		Alt: []string{
			"DIN", "DINER", "DNR",
		},
	},
	{
		Primary: "DIOCESE",
		Short:   "DIO",
		Alt: []string{
			"DIO", "DIOCESE",
		},
	},
	{
		Primary: "DIODE",
		Short:   "DIOD",
		Alt: []string{
			"DIOD", "DIODE",
		},
	},
	{
		Primary: "DIRECT",
		Short:   "DIRECT",
		Alt: []string{
			"DIR", "DIRECT",
		},
	},
	{
		Primary: "DIRECTION",
		Short:   "DIRCTN",
		Alt: []string{
			"DIRCTN", "DIRECTION",
		},
	},
	{
		Primary: "DIRECTIONAL",
		Short:   "DIRCTNL",
		Alt: []string{
			"DIRCTNL", "DIRECTIONAL",
		},
	},
	{
		Primary: "DIRECTOR",
		Short:   "DIR",
		Alt: []string{
			"DIR", "DIRCTR", "DIRECTOR",
		},
	},
	{
		Primary: "DIRECTORATE",
		Short:   "DIRCTRT",
		Alt: []string{
			"DIRCTRT", "DIRECTORATE",
		},
	},
	{
		Primary: "DIRECTORY",
		Short:   "DIRCTRY",
		Alt:     []string{"DIRECTORY"},
	},
	{
		Primary: "DISABILITY",
		Short:   "DSBLTY",
		Alt: []string{
			"DISABILITY", "DSBLTY",
		},
	},
	{
		Primary: "DISARMAMENT",
		Short:   "DSARMNT",
		Alt: []string{
			"DISARMAMENT", "DSARMNT",
		},
	},
	{
		Primary: "DISBURSEMENT",
		Short:   "DISBMT",
		Alt: []string{
			"DISBMT", "DISBURSEMENT",
		},
	},
	{
		Primary: "DISCOUNT",
		Short:   "DISC",
		Alt: []string{
			"DISC", "DISCOUNT",
		},
	},
	{
		Primary: "DISPATCH",
		Short:   "DISP",
		Alt: []string{
			"DISP", "DISPATCH", "DISPTCH",
		},
	},
	{
		Primary: "DISPATCHER",
		Short:   "DISPR",
		Alt: []string{
			"DISP", "DISPATCHER", "DISPR",
		},
	},
	{
		Primary: "DISPENSARY",
		Short:   "DSPN",
		Alt: []string{
			"DISPENSARY", "DSPN",
		},
	},
	{
		Primary: "DISPLAY",
		Short:   "DSPLY",
		Alt: []string{
			"DISP", "DISPLAY", "DSPLY",
		},
	},
	{
		Primary: "DISPOSAL",
		Short:   "DSPSL",
		Alt: []string{
			"DISPOSAL", "DSPSL",
		},
	},
	{
		Primary: "DISTILLER",
		Short:   "DISTLR",
		Alt: []string{
			"DIST", "DISTILLER", "DISTLR",
		},
	},
	{
		Primary: "DISTILLERY",
		Short:   "DISTLLRY",
		Alt: []string{
			"DIST", "DISTILLERY", "DISTLLRY",
		},
	},
	{
		Primary: "DISTINCTIVE",
		Short:   "DISTNCTV",
		Alt: []string{
			"DISTINCTIVE", "DISTNCTV",
		},
	},
	{
		Primary: "DISTRIBUTING",
		Short:   "DISTRG",
		Alt: []string{
			"DISTR", "DISTRG", "DISTRIB", "DISTRIBUTIN", "DISTRIBUTING",
		},
	},
	{
		Primary: "DISTRIBUTION",
		Short:   "DISTRB",
		Alt: []string{
			"DIST", "DISTR", "DISTRB", "DISTRIB", "DISTRIBUTIN",
			"DISTRIBUTION", "DSTRBTN",
		},
	},
	{
		Primary: "DISTRIBUTOR",
		Short:   "DISTR",
		Alt: []string{
			"DISTR", "DISTRIB", "DISTRIBTR", "DISTRIBUT", "DISTRIBUTOR",
			"DSTBTR",
		},
	},
	{
		Primary: "DISTRICT",
		Short:   "DIST",
		Alt: []string{
			"DIST", "DISTRICT", "DST",
		},
	},
	{
		Primary: "DIVERSIFIED",
		Short:   "DVSFD",
		Alt: []string{
			"DIVERSIFIED", "DVRSFD", "DVSFD",
		},
	},
	{
		Primary: "DIVIDE",
		Short:   "DV",
		Alt: []string{
			"DIV", "DIVIDE",
		},
	},
	{
		Primary: "DIVING",
		Short:   "DVNG",
		Alt: []string{
			"DIVING", "DVNG",
		},
	},
	{
		Primary: "DIVISION",
		Short:   "DIV",
		Alt: []string{
			"DIV", "DIVISION", "DIVSN",
		},
	},
	{
		Primary: "DIVISIONAL",
		Short:   "DIVSNL",
		Alt: []string{
			"DIV", "DIVISIONAL", "DIVSNL", "DVSNL",
		},
	},
	{
		Primary: "DOCTOR",
		Short:   "DR",
		Alt: []string{
			"DO", "DOCTOR", "DR", "M D", "MD",
			"PH D",
		},
	},
	{
		Primary: "DOCTRINE",
		Short:   "DOCTRN",
		Alt: []string{
			"DOCTRINE", "DOCTRN",
		},
	},
	{
		Primary: "DOCUMENTATION",
		Short:   "DCMNTN",
		Alt: []string{
			"DCMNTN", "DOCUMENTATION",
		},
	},
	{
		Primary: "DODGE",
		Short:   "DDG",
		Alt: []string{
			"DDG", "DODGE",
		},
	},
	{
		Primary: "DOLLAR",
		Short:   "DLLR",
		Alt: []string{
			"DLLR", "DLR", "DOLLAR",
		},
	},
	{
		Primary: "DOMESTIC",
		Short:   "DOM",
		Alt: []string{
			"DOM", "DOMESTIC",
		},
	},
	{
		Primary: "DOMINION",
		Short:   "DOMNN",
		Alt: []string{
			"DOMINION", "DOMNN",
		},
	},
	{
		Primary: "DONNEE",
		Short:   "DNN",
		Alt: []string{
			"DNN", "DONNEE",
		},
	},
	{
		Primary: "DOUBLE",
		Short:   "DBL",
		Alt: []string{
			"DBL", "DOUBLE",
		},
	},
	{
		Primary: "DOUGHNUT",
		Short:   "DONUT",
		Alt: []string{
			"DNT", "DONUT", "DOUGHNUT",
		},
	},
	{
		Primary: "DOWNTOWN",
		Short:   "DWNTN",
		Alt: []string{
			"DOWNTOWN", "DWNTN",
		},
	},
	{
		Primary: "DRAFTING",
		Short:   "DRFTNG",
		Alt: []string{
			"DRAFTING", "DRFTNG",
		},
	},
	{
		Primary: "DRAFTSMAN",
		Short:   "DFTSMAN",
		Alt: []string{
			"DFTSMAN", "DRAFTS", "DRAFTSMAN",
		},
	},
	{
		Primary: "DRAGON",
		Short:   "DRGN",
		Alt: []string{
			"DRAGON", "DRGN",
		},
	},
	{
		Primary: "DRAIN",
		Short:   "DRN",
		Alt: []string{
			"DRAIN", "DRN",
		},
	},
	{
		Primary: "DRAINAGE",
		Short:   "DRNG",
		Alt: []string{
			"DRAINAGE", "DRNG",
		},
	},
	{
		Primary: "DRAMA",
		Short:   "DRMA",
		Alt: []string{
			"DRAMA", "DRMA",
		},
	},
	{
		Primary: "DRAPERY",
		Short:   "DRAP",
		Alt: []string{
			"DRAP", "DRAPERIES", "DRAPERY",
		},
	},
	{
		Primary: "DREAM",
		Short:   "DRM",
		Alt: []string{
			"DREAM", "DRM",
		},
	},
	{
		Primary: "DRESS",
		Short:   "DRS",
		Alt: []string{
			"DRESS", "DRS",
		},
	},
	{
		Primary: "DRILL",
		Short:   "DRLL",
		Alt: []string{
			"DRILL", "DRLL",
		},
	},
	{
		Primary: "DRILLING",
		Short:   "DRILL",
		Alt: []string{
			"DRILL", "DRILLING", "DRLG",
		},
	},
	{
		Primary: "DRIVING",
		Short:   "DRG",
		Alt: []string{
			"DRIVING", "DRVG",
		},
	},
	{
		Primary: "DRYWALL",
		Short:   "DRYWL",
		Alt: []string{
			"DRYWALL", "DRYWL",
		},
	},
	{
		Primary: "DUCHESS",
		Short:   "DCHSS",
		Alt: []string{
			"DCHSS", "DUCHESS",
		},
	},
	{
		Primary: "DUPLICATING",
		Short:   "DUPNG",
		Alt: []string{
			"DUP", "DUPLICATING", "DUPNG",
		},
	},
	{
		Primary: "DUPLICATION",
		Short:   "DUP",
		Alt: []string{
			"DUP", "DUPLICATION",
		},
	},
	{
		Primary: "DUTCH",
		Short:   "DTCH",
		Alt: []string{
			"DTCH", "DUTCH",
		},
	},
	{
		Primary: "DWELLING",
		Short:   "DWLLNG",
		Alt: []string{
			"DWELLING", "DWLLNG",
		},
	},
	{
		Primary: "DYEING",
		Short:   "DYNG",
		Alt: []string{
			"DYEING", "DYG", "DYNG",
		},
	},
	{
		Primary: "DYING",
		Short:   "DYG",
		Alt: []string{
			"DYG", "DYING",
		},
	},
	{
		Primary: "DYNAMIC",
		Short:   "DYNMC",
		Alt: []string{
			"DYNA", "DYNAMIC", "DYNMC",
		},
	},
	{
		Primary: "EAGLE",
		Short:   "EGL",
		Alt: []string{
			"EAGLE", "EGL",
		},
	},
	{
		Primary: "EARLY",
		Short:   "ERLY",
		Alt: []string{
			"EARLY", "ERLY",
		},
	},
	{
		Primary: "EARTH",
		Short:   "ERTH",
		Alt: []string{
			"EARTH", "ERTH",
		},
	},
	{
		Primary: "EASTERN",
		Short:   "ESTRN",
		Alt: []string{
			"EASTERN", "ESTRN",
		},
	},
	{
		Primary: "EASTSIDE",
		Short:   "ESTSD",
		Alt: []string{
			"EASTSIDE", "ESTSD",
		},
	},
	{
		Primary: "EATERY",
		Short:   "ETRY",
		Alt: []string{
			"EATERY", "ETRY",
		},
	},
	{
		Primary: "ECOLOGY",
		Short:   "ECO",
		Alt: []string{
			"ECLGY", "ECO", "ECOLO", "ECOLOGY",
		},
	},
	{
		Primary: "ECONOMIC",
		Short:   "ECNMC",
		Alt: []string{
			"ECNMC", "ECON", "ECONOMIC",
		},
	},
	{
		Primary: "ECONOMIST",
		Short:   "ECONMST",
		Alt: []string{
			"ECOM", "ECON", "ECONMST", "ECONOMIST",
		},
	},
	{
		Primary: "ECONOMY",
		Short:   "ECON",
		Alt: []string{
			"ECON", "ECONOMY",
		},
	},
	{
		Primary: "EDIBLE",
		Short:   "EDBL",
		Alt: []string{
			"EDBL", "EDIBLE",
		},
	},
	{
		Primary: "EDIFICE",
		Short:   "EDFC",
		Alt: []string{
			"EDFC", "EDIFICE",
		},
	},
	{
		Primary: "EDITION",
		Short:   "ED",
		Alt: []string{
			"ED", "EDITION",
		},
	},
	{
		Primary: "EDITOR",
		Short:   "EDIT",
		Alt: []string{
			"EDIT", "EDITOR", "EDTR",
		},
	},
	{
		Primary: "EDUCATION",
		Short:   "EDUC",
		Alt: []string{
			"ED", "EDCT", "EDCTN", "EDUC", "EDUCATION",
		},
	},
	{
		Primary: "EDUCATIONAL",
		Short:   "EDUCL",
		Alt: []string{
			"EDUC", "EDUCATIONAL", "EDUCATIONL", "EDUCL", "EDUCTL",
		},
	},
	{
		Primary: "EIGHTH",
		Short:   "8TH",
		Alt: []string{
			"8TH", "EIGHTH", "VIII",
		},
	},
	{
		Primary: "ELDER",
		Short:   "ELDR",
		Alt: []string{
			"ELDER", "ELDR",
		},
	},
	{
		Primary: "ELDERLY",
		Short:   "ELDRLY",
		Alt: []string{
			"ELDERLY", "ELDRLY",
		},
	},
	{
		Primary: "ELECT",
		Short:   "ELEC",
		Alt: []string{
			"ELCT", "ELE", "ELEC", "ELECT",
		},
	},
	{
		Primary: "ELECTED",
		Short:   "ELCTD",
		Alt: []string{
			"ELCTD", "ELECT", "ELECTED",
		},
	},
	{
		Primary: "ELECTRIC",
		Short:   "ELECTR",
		Alt: []string{
			"ELC", "ELEC", "ELECT", "ELECTR", "ELECTRIC",
		},
	},
	{
		Primary: "ELECTRICAL",
		Short:   "ELECTRL",
		Alt: []string{
			"ELEC", "ELECT", "ELECTRICAL", "ELECTRL",
		},
	},
	{
		Primary: "ELECTRICIAN",
		Short:   "ELECTRCN",
		Alt: []string{
			"ELEC", "ELECT", "ELECTRCN", "ELECTRICIAN",
		},
	},
	{
		Primary: "ELECTRICITY",
		Short:   "ELECTRCTY",
		Alt: []string{
			"ELEC", "ELECT", "ELECTRCTY", "ELECTRICITY",
		},
	},
	{
		Primary: "ELECTROLOGIST",
		Short:   "ELCTRLGST",
		Alt: []string{
			"ELCTRLGST", "ELECTROLOGIST",
		},
	},
	{
		Primary: "ELECTROLYSIS",
		Short:   "ELCTRLYS",
		Alt: []string{
			"ELCTRLYS", "ELECTRLSIS", "ELECTRLYS", "ELECTROLYSIS",
		},
	},
	{
		Primary: "ELECTROMECHANICAL",
		Short:   "ELCTRMCHNCL",
		Alt: []string{
			"ELCTRMCHNCL", "ELECTROMECHANICAL",
		},
	},
	{
		Primary: "ELECTROMEDICAL",
		Short:   "ELCMED",
		Alt: []string{
			"ELCMED", "ELECTROMEDICAL",
		},
	},
	{
		Primary: "ELECTROMETALLURGICAL",
		Short:   "ELCMTLG",
		Alt: []string{
			"ELCMTLG", "ELECTROMETALLURGICAL",
		},
	},
	{
		Primary: "ELECTRON",
		Short:   "ELCTRN",
		Alt: []string{
			"ELCTRN", "ELECTRON",
		},
	},
	{
		Primary: "ELECTRONIC",
		Short:   "ELECT",
		Alt: []string{
			"ELEC", "ELECT", "ELECTRNC", "ELECTRONIC",
		},
	},
	{
		Primary: "ELECTROPLATING",
		Short:   "ELCPLTG",
		Alt: []string{
			"ELCPLTG", "ELECTROPLATING",
		},
	},
	{
		Primary: "ELEGANCE",
		Short:   "ELGNC",
		Alt: []string{
			"ELEGANCE", "ELGNC",
		},
	},
	{
		Primary: "ELEGANT",
		Short:   "ELGNT",
		Alt: []string{
			"ELEGANT", "ELGNT",
		},
	},
	{
		Primary: "ELEMENT",
		Short:   "ELMNT",
		Alt: []string{
			"ELEMENT", "ELMNT",
		},
	},
	{
		Primary: "ELEMENTARY",
		Short:   "ELEM",
		Alt: []string{
			"ELEM", "ELEMENTARY",
		},
	},
	{
		Primary: "ELEVATOR",
		Short:   "ELEV",
		Alt: []string{
			"ELEV", "ELEVATOR",
		},
	},
	{
		Primary: "ELEVENTH",
		Short:   "11TH",
		Alt: []string{
			"11", "11TH", "ELEVENTH", "XI",
		},
	},
	{
		Primary: "ELITE",
		Short:   "ELITE",
		Alt:     []string{"ELITE"},
	},
	{
		Primary: "EMBASSY",
		Short:   "EMBSSY",
		Alt: []string{
			"EMBASSY", "EMBSSY",
		},
	},
	{
		Primary: "EMBROIDERY",
		Short:   "EMB",
		Alt: []string{
			"EMB", "EMBROIDERY",
		},
	},
	{
		Primary: "EMERGENCY",
		Short:   "EMER",
		Alt: []string{
			"EMER", "EMERG", "EMERGENCY", "EMERGNCY",
		},
	},
	{
		Primary: "EMPIRE",
		Short:   "EMP",
		Alt: []string{
			"EMP", "EMPIRE",
		},
	},
	{
		Primary: "EMPLOYED",
		Short:   "EMPL",
		Alt: []string{
			"EMPL", "EMPLOY", "EMPLOYED",
		},
	},
	{
		Primary: "EMPLOYEE",
		Short:   "EMPLYE",
		Alt: []string{
			"EMPL", "EMPLOYEE", "EMPLYE",
		},
	},
	{
		Primary: "EMPLOYMENT",
		Short:   "EMPLMNT",
		Alt: []string{
			"EMPL", "EMPLMNT", "EMPLMT", "EMPLOYMENT",
		},
	},
	{
		Primary: "EMPORIUM",
		Short:   "EMPOR",
		Alt: []string{
			"EMPOR", "EMPORIUM", "EMPORM", "EMPRM",
		},
	},
	{
		Primary: "ENAMEL",
		Short:   "ENL",
		Alt: []string{
			"ENAMEL", "ENL",
		},
	},
	{
		Primary: "ENAMELING",
		Short:   "ENMLNG",
		Alt: []string{
			"ENAMELING", "ENMLNG",
		},
	},
	{
		Primary: "ENCYCLOPEDIA",
		Short:   "ENCY",
		Alt: []string{
			"ENCY", "ENCYCLOPEDIA",
		},
	},
	{
		Primary: "ENDEAVOR",
		Short:   "ENDVR",
		Alt: []string{
			"ENDEAVOR", "ENDVR",
		},
	},
	{
		Primary: "ENDOCRINOLOGIST",
		Short:   "ENDCRNLGST",
		Alt: []string{
			"ENDCRNLGST", "ENDOCRINOLOGIST",
		},
	},
	{
		Primary: "ENDODONTIC",
		Short:   "ENDDNTC",
		Alt: []string{
			"ENDDNTC", "ENDODONTIC",
		},
	},
	{
		Primary: "ENERGY",
		Short:   "ENGRY",
		Alt: []string{
			"ENERGY", "ENGRY", "ENGY", "ENRG",
		},
	},
	{
		Primary: "ENFORCEMENT",
		Short:   "ENFCMNT",
		Alt: []string{
			"ENFCMNT", "ENFORCEMENT",
		},
	},
	{
		Primary: "ENGINE",
		Short:   "ENG",
		Alt: []string{
			"ENG", "ENGINE",
		},
	},
	{
		Primary: "ENGINEER",
		Short:   "ENGR",
		Alt: []string{
			"ENG", "ENGINEER", "ENGR",
		},
	},
	{
		Primary: "ENGINEERED",
		Short:   "ENGRD",
		Alt: []string{
			"ENGINEERED", "ENGRD",
		},
	},
	{
		Primary: "ENGINEERING",
		Short:   "ENGRG",
		Alt: []string{
			"ENG", "ENGINEERING", "ENGINRNG", "ENGR", "ENGRG",
			"ENGRNG",
		},
	},
	{
		Primary: "ENGLAND",
		Short:   "ENGLD",
		Alt: []string{
			"ENG", "ENGL", "ENGLAND", "ENGLD",
		},
	},
	{
		Primary: "ENGLISH",
		Short:   "ENGL",
		Alt: []string{
			"ENGL", "ENGLISH", "ENGLSH",
		},
	},
	{
		Primary: "ENGRAVER",
		Short:   "ENGRVR",
		Alt: []string{
			"ENGRAVER", "ENGRVR",
		},
	},
	{
		Primary: "ENGRAVING",
		Short:   "ENGRV",
		Alt: []string{
			"ENGRAVING", "ENGRV",
		},
	},
	{
		Primary: "ENLARGE",
		Short:   "ENLRG",
		Alt: []string{
			"ENLARGE", "ENLRG",
		},
	},
	{
		Primary: "ENSIGN",
		Short:   "ENS",
		Alt: []string{
			"ENS", "ENSIGN",
		},
	},
	{
		Primary: "ENTERPRISE",
		Short:   "ENTRPRS",
		Alt: []string{
			"ENT", "ENTER", "ENTERP", "ENTERPRISE", "ENTERPRS",
			"ENTP", "ENTPR", "ENTPS", "ENTRPR", "ENTRPRS",
		},
	},
	{
		Primary: "ENTERTAINMENT",
		Short:   "ENTRTN",
		Alt: []string{
			"ENTERTAINMENT", "ENTRMT", "ENTRTN",
		},
	},
	{
		Primary: "ENTREPOT",
		Short:   "ENTRPT",
		Alt: []string{
			"ENTREPOT", "ENTRPT",
		},
	},
	{
		Primary: "ENTREPENEUR",
		Short:   "ENTRPRNR",
		Alt: []string{
			"ENTREPENEUR", "ENTRPRNR",
		},
	},
	{
		Primary: "ENTRY",
		Short:   "ENT",
		Alt: []string{
			"ENT", "ENTRY",
		},
	},
	{
		Primary: "ENVELOPE",
		Short:   "ENV",
		Alt: []string{
			"ENV", "ENVELOPE",
		},
	},
	{
		Primary: "ENVIRONMENT",
		Short:   "ENVIR",
		Alt: []string{
			"ENVIR", "ENVIRON", "ENVIRONMENT", "ENVRMT", "ENVRONMEN",
		},
	},
	{
		Primary: "ENVIRONMENTAL",
		Short:   "ENVIRON",
		Alt: []string{
			"ENVIRON", "ENVIRONMENTAL", "ENVRMTL", "ENVRNMTL",
		},
	},
	{
		Primary: "EPISCOPAL",
		Short:   "EPISCPL",
		Alt: []string{
			"EPIS", "EPISCOPAL", "EPISCPL", "EPSCP", "EPSCPL",
		},
	},
	{
		Primary: "EPSILON",
		Short:   "EPSLN",
		Alt: []string{
			"EPSILON", "EPSLN",
		},
	},
	{
		Primary: "EQUAL",
		Short:   "EQL",
		Alt: []string{
			"EQL", "EQUAL",
		},
	},
	{
		Primary: "EQUESTRIAN",
		Short:   "EQSTRN",
		Alt: []string{
			"EQSTRN", "EQUESTRIAN",
		},
	},
	{
		Primary: "EQUINE",
		Short:   "EQN",
		Alt: []string{
			"EQN", "EQUINE",
		},
	},
	{
		Primary: "EQUIPMENT",
		Short:   "EQUIP",
		Alt: []string{
			"EQIPMENT", "EQP", "EQPMNT", "EQPT", "EQUIP",
			"EQUIPMENT", "EQUIPT",
		},
	},
	{
		Primary: "EQUITABLE",
		Short:   "EQTBL",
		Alt: []string{
			"EQTBL", "EQUITABLE",
		},
	},
	{
		Primary: "EQUITY",
		Short:   "EQTY",
		Alt: []string{
			"EQTY", "EQUITY", "EQUTY",
		},
	},
	{
		Primary: "ERECTING",
		Short:   "ERCT",
		Alt: []string{
			"ERCT", "ERECTING",
		},
	},
	{
		Primary: "ERECTOR",
		Short:   "ERCTR",
		Alt: []string{
			"ERCTR", "ERECTOR",
		},
	},
	{
		Primary: "ESQUIRE",
		Short:   "ESQ",
		Alt: []string{
			"ESQ", "ESQUIRE",
		},
	},
	{
		Primary: "ESSENTIAL",
		Short:   "ESSNTL",
		Alt: []string{
			"ESSENTIAL", "ESSTNL",
		},
	},
	{
		Primary: "ESTABLISHMENT",
		Short:   "ESTAB",
		Alt: []string{
			"EST", "ESTAB", "ESTABLISHMENT",
		},
	},
	{
		Primary: "ESTATE",
		Short:   "EST",
		Alt: []string{
			"EST", "ESTATE",
		},
	},
	{
		Primary: "ESTIMATION",
		Short:   "ESTMTN",
		Alt: []string{
			"ESTIMATION", "ESTMTN",
		},
	},
	{
		Primary: "ESTIMATOR",
		Short:   "ESTMTR",
		Alt: []string{
			"EST", "ESTIMATOR", "ESTMTR",
		},
	},
	{
		Primary: "ETCETERA",
		Short:   "ETC",
		Alt: []string{
			"ETC", "ETCETERA",
		},
	},
	{
		Primary: "ETUDE",
		Short:   "ETD",
		Alt: []string{
			"ETD", "ETUDE",
		},
	},
	{
		Primary: "EUROPEAN",
		Short:   "ERPN",
		Alt: []string{
			"ERPN", "EUROPEAN",
		},
	},
	{
		Primary: "EVALUATION",
		Short:   "EVAL",
		Alt: []string{
			"EV", "EVAL", "EVALUATION",
		},
	},
	{
		Primary: "EVANGELICAL",
		Short:   "EVNGLCL",
		Alt: []string{
			"EVANGELICAL", "EVNGLCL",
		},
	},
	{
		Primary: "EVANGELIST",
		Short:   "EVNGLST",
		Alt: []string{
			"EVANGELIST", "EVNGLST",
		},
	},
	{
		Primary: "EVANGELISTIC",
		Short:   "EVNGLSTC",
		Alt: []string{
			"EVANGELISTIC", "EVNGLSTC",
		},
	},
	{
		Primary: "EVENING",
		Short:   "EVNNG",
		Alt: []string{
			"EVENING", "EVNNG",
		},
	},
	{
		Primary: "EVENT",
		Short:   "EVNT",
		Alt: []string{
			"EVENT", "EVNT",
		},
	},
	{
		Primary: "EVERGREEN",
		Short:   "EVRGRN",
		Alt: []string{
			"EVERGREEN", "EVRGRN",
		},
	},
	{
		Primary: "EXACT",
		Short:   "EXCT",
		Alt: []string{
			"EXACT", "EXCT",
		},
	},
	{
		Primary: "EXAMINATION",
		Short:   "EXMNTN",
		Alt: []string{
			"EXAMINATION", "EXMNTN",
		},
	},
	{
		Primary: "EXAMINE",
		Short:   "EXAM",
		Alt: []string{
			"EX", "EXAM", "EXAMINE", "EXMN",
		},
	},
	{
		Primary: "EXAMINER",
		Short:   "EXMNR",
		Alt: []string{
			"EXAMINER", "EXMNR",
		},
	},
	{
		Primary: "EXCAVATE",
		Short:   "EXCVT",
		Alt: []string{
			"EXCAVATE", "EXCVT",
		},
	},
	{
		Primary: "EXCAVATING",
		Short:   "EXCAVTG",
		Alt: []string{
			"EXCAVATING", "EXCAVATNG", "EXCAVTG", "EXCVTG",
		},
	},
	{
		Primary: "EXCAVATION",
		Short:   "EXCVTN",
		Alt: []string{
			"EXCAVATION", "EXCTVN",
		},
	},
	{
		Primary: "EXCAVATOR",
		Short:   "EXCVTR",
		Alt: []string{
			"EXCAVATOR", "EXCAVATR", "EXCVTR",
		},
	},
	{
		Primary: "EXCEL",
		Short:   "EXCL",
		Alt: []string{
			"EXCEL", "EXCL",
		},
	},
	{
		Primary: "EXCELSIOR",
		Short:   "EXCLSR",
		Alt: []string{
			"EXCEL", "EXCELSIOR", "EXCLSR",
		},
	},
	{
		Primary: "EXCEPTIONAL",
		Short:   "EXCPTNL",
		Alt: []string{
			"EXCEPTIONAL", "EXCPTNL",
		},
	},
	{
		Primary: "EXCESS",
		Short:   "EXCSS",
		Alt: []string{
			"EXCESS", "EXCSS",
		},
	},
	{
		Primary: "EXCHANGE",
		Short:   "EXCH",
		Alt: []string{
			"ECHANGE", "EXCH", "EXCHANGE",
		},
	},
	{
		Primary: "EXECUTIVE",
		Short:   "EXEC",
		Alt: []string{
			"EX", "EXC", "EXE", "EXEC", "EXECUTIVE",
		},
	},
	{
		Primary: "EXECUTOR",
		Short:   "EXTR",
		Alt: []string{
			"EXECUTOR", "EXTR",
		},
	},
	{
		Primary: "EXEMPT",
		Short:   "EXMPT",
		Alt: []string{
			"EXEMPT", "EXMPT",
		},
	},
	{
		Primary: "EXEMPTED",
		Short:   "EXMPTD",
		Alt: []string{
			"EXEMPTED", "EXMPTD",
		},
	},
	{
		Primary: "EXHIBIT",
		Short:   "EXHBT",
		Alt: []string{
			"EXHBT", "EXHIBIT",
		},
	},
	{
		Primary: "EXHIBITOR",
		Short:   "EXHBTR",
		Alt: []string{
			"EXHBTR", "EXHIBITOR",
		},
	},
	{
		Primary: "EXPEDITER",
		Short:   "EXPD",
		Alt: []string{
			"EXPD", "EXPEDITER",
		},
	},
	{
		Primary: "EXPEDITION",
		Short:   "EXPDTN",
		Alt: []string{
			"EXP", "EXPDTN", "EXPEDITION",
		},
	},
	{
		Primary: "EXPEDITOR",
		Short:   "EXPDTR",
		Alt: []string{
			"EXPDTR", "EXPEDITOR",
		},
	},
	{
		Primary: "EXPENSE",
		Short:   "EXP",
		Alt: []string{
			"EXP", "EXPENSE",
		},
	},
	{
		Primary: "EXPERIENCE",
		Short:   "EXPRNC",
		Alt: []string{
			"EXPERIENCE", "EXPRNC",
		},
	},
	{
		Primary: "EXPERIMENT",
		Short:   "EXPRMNT",
		Alt: []string{
			"EXPERIMENT", "EXPRMNT",
		},
	},
	{
		Primary: "EXPERT",
		Short:   "EXPR",
		Alt: []string{
			"EXPERT", "EXPR", "EXPRT",
		},
	},
	{
		Primary: "EXPLORATION",
		Short:   "EXPLRN",
		Alt: []string{
			"EXPLORATION", "EXPLRN", "EXPN",
		},
	},
	{
		Primary: "EXPLOSIVE",
		Short:   "EXPLSV",
		Alt: []string{
			"EXPL", "EXPLOSIVE", "EXPLSV",
		},
	},
	{
		Primary: "EXPORT",
		Short:   "EXPRT",
		Alt: []string{
			"EXP", "EXPORT", "EXPRT", "EXPT",
		},
	},
	{
		Primary: "EXPORTATION",
		Short:   "EXPN",
		Alt: []string{
			"EXPN", "EXPORTATION", "EXPRTTN",
		},
	},
	{
		Primary: "EXPORTER",
		Short:   "EXPRTR",
		Alt: []string{
			"EXP", "EXPORTER", "EXPRTR",
		},
	},
	{
		Primary: "EXPOSE",
		Short:   "EXPS",
		Alt: []string{
			"EXPOSE", "EXPS",
		},
	},
	{
		Primary: "EXPOSITION",
		Short:   "EXPO",
		Alt: []string{
			"EXPO", "EXPOSITION", "EXPSTN",
		},
	},
	{
		Primary: "EXPRESS",
		Short:   "EXPRSS",
		Alt: []string{
			"EX", "EXP", "EXPRESS", "EXPRSS",
		},
	},
	{
		Primary: "EXPRESSION",
		Short:   "EXPRSSN",
		Alt: []string{
			"EXPRESSION", "EXPRSSN",
		},
	},
	{
		Primary: "EXPRESSWAY",
		Short:   "EXPY",
		Alt: []string{
			"EXPRESSWAY", "EXPRSSWY", "EXPY",
		},
	},
	{
		Primary: "EXTENSION",
		Short:   "EXT",
		Alt: []string{
			"EXT", "EXTENSION", "EXTNSN",
		},
	},
	{
		Primary: "EXTERMINATING",
		Short:   "EXTERM",
		Alt: []string{
			"EXTERM", "EXTERMINATING", "EXTG", "EXTRMNTNG",
		},
	},
	{
		Primary: "EXTERMINATOR",
		Short:   "EXTRMNTR",
		Alt: []string{
			"EXTERMINATOR", "EXTRMNTR",
		},
	},
	{
		Primary: "EXTRACT",
		Short:   "EXTRCT",
		Alt: []string{
			"EXT", "EXTRACT", "EXTRCT",
		},
	},
	{
		Primary: "EXTRACTOR",
		Short:   "EXTRCTR",
		Alt: []string{
			"EXTRACTOR", "EXTRCTR",
		},
	},
	{
		Primary: "EXTRAORDINARY",
		Short:   "EXTRRDNRY",
		Alt: []string{
			"EXTRAORDINARY", "EXTRRDNRY",
		},
	},
	{
		Primary: "EXTREME",
		Short:   "EXTRM",
		Alt: []string{
			"EXTREME", "EXTRM",
		},
	},
	{
		Primary: "FABRIC",
		Short:   "FBRC",
		Alt: []string{
			"FABR", "FABRIC", "FBRC",
		},
	},
	{
		Primary: "FABRICATED",
		Short:   "FABD",
		Alt: []string{
			"FAB", "FABD", "FABRICATED",
		},
	},
	{
		Primary: "FABRICATING",
		Short:   "FABG",
		Alt: []string{
			"FABG", "FABRICATING",
		},
	},
	{
		Primary: "FABRICATION",
		Short:   "FBRCN",
		Alt: []string{
			"FABRICATION", "FBRCN",
		},
	},
	{
		Primary: "FABRICATOR",
		Short:   "FAB",
		Alt: []string{
			"FAB", "FABRICATOR", "FABRICTR", "FBRCTR",
		},
	},
	{
		Primary: "FACILITY",
		Short:   "FACLTY",
		Alt: []string{
			"FAC", "FACILITY", "FACLTY",
		},
	},
	{
		Primary: "FACTOR",
		Short:   "FCTR",
		Alt: []string{
			"FACTOR", "FCTR",
		},
	},
	{
		Primary: "FACTORY",
		Short:   "FCTRY",
		Alt: []string{
			"FAC", "FACTORY", "FCTRY",
		},
	},
	{
		Primary: "FACULTY",
		Short:   "FCLTY",
		Alt: []string{
			"FACULTY", "FCLTY",
		},
	},
	{
		Primary: "FAITH",
		Short:   "FTH",
		Alt: []string{
			"FAITH", "FTH",
		},
	},
	{
		Primary: "FALLS",
		Short:   "FLS",
		Alt: []string{
			"FALLS", "FLS",
		},
	},
	{
		Primary: "FAMILY",
		Short:   "FMLY",
		Alt: []string{
			"FAM", "FAMILY", "FMLY",
		},
	},
	{
		Primary: "FAMOUS",
		Short:   "FMS",
		Alt: []string{
			"FAMOUS", "FMS",
		},
	},
	{
		Primary: "FANCY",
		Short:   "FNCY",
		Alt: []string{
			"FANCY", "FNCY",
		},
	},
	{
		Primary: "FANTASTIC",
		Short:   "FNTSTIC",
		Alt: []string{
			"FANTASTIC", "FNTSTIC",
		},
	},
	{
		Primary: "FANTASY",
		Short:   "FNTSY",
		Alt: []string{
			"FANTASY", "FNTSY",
		},
	},
	{
		Primary: "FARM",
		Short:   "FRM",
		Alt: []string{
			"FARM", "FRM",
		},
	},
	{
		Primary: "FARMER",
		Short:   "FRMR",
		Alt: []string{
			"FARMER", "FRMR",
		},
	},
	{
		Primary: "FARMING",
		Short:   "FRMNG",
		Alt: []string{
			"FARMING", "FRMNG",
		},
	},
	{
		Primary: "FASHION",
		Short:   "FASHN",
		Alt: []string{
			"FASHION", "FASHN", "FSHN",
		},
	},
	{
		Primary: "FASTENER",
		Short:   "FAS",
		Alt: []string{
			"FAS", "FASTENER",
		},
	},
	{
		Primary: "FATHER",
		Short:   "FR",
		Alt: []string{
			"FATHER", "FR",
		},
	},
	{
		Primary: "FAUCET",
		Short:   "FCT",
		Alt: []string{
			"FAUCET", "FCT",
		},
	},
	{
		Primary: "FEATHER",
		Short:   "FE",
		Alt: []string{
			"FE", "FEATHER",
		},
	},
	{
		Primary: "FEDERAL",
		Short:   "FED",
		Alt: []string{
			"FDRL", "FED", "FEDERAL", "FEDL", "FEDRL",
		},
	},
	{
		Primary: "FEDERATED",
		Short:   "FDRTD",
		Alt: []string{
			"FDRTD", "FEDERATED",
		},
	},
	{
		Primary: "FEDERATION",
		Short:   "FEDRN",
		Alt: []string{
			"FEDERATION", "FEDRN",
		},
	},
	{
		Primary: "FELLOWSHIP",
		Short:   "FLLWSHP",
		Alt: []string{
			"FELLOWSHIP", "FELLOWSHP", "FLLWSHP", "FLWSHIP", "FLWSHP",
		},
	},
	{
		Primary: "FENCE",
		Short:   "FNC",
		Alt: []string{
			"FENCE", "FNC",
		},
	},
	{
		Primary: "FERROUS",
		Short:   "FER",
		Alt: []string{
			"FER", "FERROUS",
		},
	},
	{
		Primary: "FERTILIZER",
		Short:   "FERT",
		Alt: []string{
			"FERT", "FERTILIZER",
		},
	},
	{
		Primary: "FIBER",
		Short:   "FIBR",
		Alt: []string{
			"FIBER", "FIBR",
		},
	},
	{
		Primary: "FIBERGLASS",
		Short:   "FBRGLS",
		Alt: []string{
			"FBRGLS", "FIBERGLASS",
		},
	},
	{
		Primary: "FIBRE",
		Short:   "FBR",
		Alt: []string{
			"FBR", "FIBR", "FIBRE",
		},
	},
	{
		Primary: "FIDELITY",
		Short:   "FIDLTY",
		Alt: []string{
			"FDLTY", "FIDELITY", "FIDLTY",
		},
	},
	{
		Primary: "FIELD",
		Short:   "FLD",
		Alt: []string{
			"FIELD", "FLD",
		},
	},
	{
		Primary: "FIFTH",
		Short:   "5TH",
		Alt: []string{
			"5TH", "FIFTH", "V",
		},
	},
	{
		Primary: "FIGHT",
		Short:   "FGHT",
		Alt: []string{
			"FGHT", "FIGHT",
		},
	},
	{
		Primary: "FIGHTER",
		Short:   "FGHTR",
		Alt: []string{
			"FGHTR", "FIGHTER",
		},
	},
	{
		Primary: "FINANCE",
		Short:   "FIN",
		Alt: []string{
			"FIN", "FINANCE", "FNC",
		},
	},
	{
		Primary: "FINANCIAL",
		Short:   "FNCL",
		Alt: []string{
			"FINANCIAL", "FINL", "FNCL",
		},
	},
	{
		Primary: "FINANCIER",
		Short:   "FINR",
		Alt: []string{
			"FIN", "FINANCIER", "FINR",
		},
	},
	{
		Primary: "FINANCING",
		Short:   "FING",
		Alt: []string{
			"FINANCING", "FING",
		},
	},
	{
		Primary: "FINDING",
		Short:   "FNDG",
		Alt: []string{
			"FINDING", "FNDG",
		},
	},
	{
		Primary: "FINEST",
		Short:   "FNST",
		Alt: []string{
			"FINEST", "FNST",
		},
	},
	{
		Primary: "FINISH",
		Short:   "FNSH",
		Alt: []string{
			"FINISH", "FINSH", "FNSH",
		},
	},
	{
		Primary: "FINISHING",
		Short:   "FINISH",
		Alt: []string{
			"FINISH", "FINISHING", "FINSHG", "FNSHNG",
		},
	},
	{
		Primary: "FIREARM",
		Short:   "FRARM",
		Alt: []string{
			"FIREARM", "FRARM",
		},
	},
	{
		Primary: "FIREMAN",
		Short:   "FIRMN",
		Alt: []string{
			"FIREMAN", "FIRMN", "FRMN",
		},
	},
	{
		Primary: "FIREWORK",
		Short:   "FRWRK",
		Alt: []string{
			"FIREWORK", "FRWRK",
		},
	},
	{
		Primary: "FIRST",
		Short:   "1ST",
		Alt: []string{
			"1", "1ST", "FIRST", "I",
		},
	},
	{
		Primary: "FISCAL",
		Short:   "FISC",
		Alt: []string{
			"FISC", "FISCAL",
		},
	},
	{
		Primary: "FISHERY",
		Short:   "FSHRY",
		Alt: []string{
			"FISHERY", "FSHRY",
		},
	},
	{
		Primary: "FISHING",
		Short:   "FSHNG",
		Alt: []string{
			"FISHING", "FSHNG",
		},
	},
	{
		Primary: "FITNESS",
		Short:   "FITNS",
		Alt: []string{
			"FITNESS", "FITNS",
		},
	},
	{
		Primary: "FIXTURE",
		Short:   "FIX",
		Alt: []string{
			"FIX", "FIXTURE",
		},
	},
	{
		Primary: "FLAVOR",
		Short:   "FLVR",
		Alt: []string{
			"FL", "FLA", "FLAVOR", "FLVR",
		},
	},
	{
		Primary: "FLEET",
		Short:   "FLT",
		Alt: []string{
			"FLEET", "FLT",
		},
	},
	{
		Primary: "FLIGHT",
		Short:   "FLGT",
		Alt: []string{
			"FLGT", "FLIGHT", "FLT",
		},
	},
	{
		Primary: "FLOCK",
		Short:   "FLCK",
		Alt: []string{
			"FLCK", "FLOCK",
		},
	},
	{
		Primary: "FLOOR",
		Short:   "FL",
		Alt: []string{
			"FL", "FLOOR", "FLR",
		},
	},
	{
		Primary: "FLOORCOVERING",
		Short:   "FLRCVG",
		Alt: []string{
			"FLOORCOVERING", "FLRCVG",
		},
	},
	{
		Primary: "FLOORING",
		Short:   "FLRNG",
		Alt: []string{
			"FLOORING", "FLRG", "FLRNG",
		},
	},
	{
		Primary: "FLORAL",
		Short:   "FLRL",
		Alt: []string{
			"FLORAL", "FLRL",
		},
	},
	{
		Primary: "FLORIST",
		Short:   "FLRST",
		Alt: []string{
			"FLOR", "FLORIST", "FLRST",
		},
	},
	{
		Primary: "FLOWER",
		Short:   "FLWR",
		Alt: []string{
			"FLOWER", "FLWR",
		},
	},
	{
		Primary: "FLUID",
		Short:   "FLUD",
		Alt: []string{
			"FLD", "FLUD", "FLUID",
		},
	},
	{
		Primary: "FLYING",
		Short:   "FLY",
		Alt: []string{
			"FLY", "FLYING",
		},
	},
	{
		Primary: "FOCUS",
		Short:   "FCS",
		Alt: []string{
			"FCS", "FOCUS",
		},
	},
	{
		Primary: "FOOTBALL",
		Short:   "FTBLL",
		Alt: []string{
			"FOOTBALL", "FTBLL",
		},
	},
	{
		Primary: "FOOTWEAR",
		Short:   "FTWR",
		Alt: []string{
			"FOOTWEAR", "FTWR",
		},
	},
	{
		Primary: "FORCE",
		Short:   "FRC",
		Alt: []string{
			"FOR", "FORCE", "FRC",
		},
	},
	{
		Primary: "FORECASTING",
		Short:   "FRCSTNG",
		Alt: []string{
			"FORECASTING", "FRCSTNG",
		},
	},
	{
		Primary: "FOREIGN",
		Short:   "FRGN",
		Alt: []string{
			"FGN", "FOREIGN", "FRGN",
		},
	},
	{
		Primary: "FOREMAN",
		Short:   "FORMN",
		Alt: []string{
			"FOREMAN", "FORMN", "FRMN",
		},
	},
	{
		Primary: "FORESIGHT",
		Short:   "FORSGHT",
		Alt: []string{
			"FORESIGHT", "FORSGHT",
		},
	},
	{
		Primary: "FOREST",
		Short:   "FRST",
		Alt: []string{
			"FOREST", "FRST",
		},
	},
	{
		Primary: "FORESTRY",
		Short:   "FOR",
		Alt: []string{
			"FOR", "FORESTRY", "FRSTRY",
		},
	},
	{
		Primary: "FOREVER",
		Short:   "FORVR",
		Alt: []string{
			"FOREVER", "FORVR",
		},
	},
	{
		Primary: "FORGING",
		Short:   "FRG",
		Alt: []string{
			"FORGING", "FRG",
		},
	},
	{
		Primary: "FORGOING",
		Short:   "FORGNG",
		Alt: []string{
			"FORGOING", "FRGNG",
		},
	},
	{
		Primary: "FORKLIFT",
		Short:   "FRKLFT",
		Alt: []string{
			"FORKLIFT", "FRKLFT",
		},
	},
	{
		Primary: "FORMAL",
		Short:   "FRML",
		Alt: []string{
			"FORMAL", "FRML",
		},
	},
	{
		Primary: "FORMATION",
		Short:   "FRMTN",
		Alt: []string{
			"FORMATION", "FRMTN",
		},
	},
	{
		Primary: "FORTUNE",
		Short:   "FRTN",
		Alt: []string{
			"FORTUNE", "FRTN",
		},
	},
	{
		Primary: "FORUM",
		Short:   "FRUM",
		Alt: []string{
			"FORUM", "FRM", "FRUM",
		},
	},
	{
		Primary: "FORWARDING",
		Short:   "FWDG",
		Alt: []string{
			"FORWARDING", "FWDG",
		},
	},
	{
		Primary: "FOSTER",
		Short:   "FSTR",
		Alt: []string{
			"FOSTER", "FSTR",
		},
	},
	{
		Primary: "FOUND",
		Short:   "FND",
		Alt: []string{
			"FND", "FOUND",
		},
	},
	{
		Primary: "FOUNDATION",
		Short:   "FNDTN",
		Alt: []string{
			"FDN", "FNDTN", "FOUNDATION", "FOUNDTN",
		},
	},
	{
		Primary: "FOUNDRY",
		Short:   "FNDRY",
		Alt: []string{
			"FDRY", "FNDRY", "FOUNDRY",
		},
	},
	{
		Primary: "FOUNTAIN",
		Short:   "FTN",
		Alt: []string{
			"FOUNTAIN", "FTN",
		},
	},
	{
		Primary: "FOURGON",
		Short:   "FORGN",
		Alt: []string{
			"FORGN", "FOURGON",
		},
	},
	{
		Primary: "FOURTEENTH",
		Short:   "14TH",
		Alt: []string{
			"14", "14TH", "FOURTEENTH", "XIV",
		},
	},
	{
		Primary: "FOURTH",
		Short:   "4TH",
		Alt: []string{
			"4", "4TH", "FOURTH", "IV",
		},
	},
	{
		Primary: "FRAGRANCE",
		Short:   "FRGRNC",
		Alt: []string{
			"FRAGRANCE", "FRGRNC",
		},
	},
	{
		Primary: "FRAME",
		Short:   "FRAM",
		Alt: []string{
			"FRAM", "FRAME",
		},
	},
	{
		Primary: "FRAMEWORK",
		Short:   "FRMWRK",
		Alt: []string{
			"FRAMEWORK", "FRMWRK",
		},
	},
	{
		Primary: "FRAMING",
		Short:   "FRAMG",
		Alt: []string{
			"FRAMG", "FRAMING",
		},
	},
	{
		Primary: "FRANCHISE",
		Short:   "FRNCHS",
		Alt: []string{
			"FRANCHISE", "FRNCHS",
		},
	},
	{
		Primary: "FRANCHISING",
		Short:   "FRNCHSNG",
		Alt: []string{
			"FRANCHISING", "FRANCHSNG",
		},
	},
	{
		Primary: "FRATERNAL",
		Short:   "FRTRNL",
		Alt: []string{
			"FRATERNAL", "FRTRNL",
		},
	},
	{
		Primary: "FRATERNITY",
		Short:   "FRTRNTY",
		Alt: []string{
			"FRATERNITY", "FRTRNTY",
		},
	},
	{
		Primary: "FREEWAY",
		Short:   "FWY",
		Alt: []string{
			"FREEWAY", "FRWY", "FWY",
		},
	},
	{
		Primary: "FREEZE",
		Short:   "FREZ",
		Alt: []string{
			"FREEZE", "FREZ", "FRZ",
		},
	},
	{
		Primary: "FREEZER",
		Short:   "FRZR",
		Alt: []string{
			"FREEZER", "FRZR",
		},
	},
	{
		Primary: "FREIGHT",
		Short:   "FRGHT",
		Alt: []string{
			"FREIGHT", "FRGHT", "FRGT", "FRT",
		},
	},
	{
		Primary: "FRENCH",
		Short:   "FRNCH",
		Alt: []string{
			"FRENCH", "FRNCH",
		},
	},
	{
		Primary: "FRESH",
		Short:   "FRSH",
		Alt: []string{
			"FRESH", "FRSH",
		},
	},
	{
		Primary: "FRIARY",
		Short:   "FRY",
		Alt: []string{
			"FRIARY", "FRY",
		},
	},
	{
		Primary: "FRICTION",
		Short:   "FRCTN",
		Alt: []string{
			"FRCTN", "FRICTION",
		},
	},
	{
		Primary: "FRIED",
		Short:   "FRD",
		Alt: []string{
			"FRD", "FRIED",
		},
	},
	{
		Primary: "FRIEND",
		Short:   "FRND",
		Alt: []string{
			"FRIEND", "FRND",
		},
	},
	{
		Primary: "FRIENDLY",
		Short:   "FRNDLY",
		Alt: []string{
			"FRIENDLY", "FRNDLY",
		},
	},
	{
		Primary: "FRONTIER",
		Short:   "FRNTR",
		Alt: []string{
			"FRNTR", "FRONTIER",
		},
	},
	{
		Primary: "FROZEN",
		Short:   "FRZ",
		Alt: []string{
			"FROZEN", "FRZ", "FRZN",
		},
	},
	{
		Primary: "FRUIT",
		Short:   "FRT",
		Alt: []string{
			"FRT", "FRUIT",
		},
	},
	{
		Primary: "FUNCTIONAL",
		Short:   "FUNCTL",
		Alt: []string{
			"FUNCTIONAL", "FUNCTL",
		},
	},
	{
		Primary: "FUNCTIONARY",
		Short:   "FUNCTRY",
		Alt: []string{
			"FUNCTIONARY", "FUNCTRY",
		},
	},
	{
		Primary: "FUNDAMENTALIST",
		Short:   "FNDMNTLST",
		Alt: []string{
			"FNDMNTLST", "FUNDAMENTALIST",
		},
	},
	{
		Primary: "FUNDING",
		Short:   "FNDNG",
		Alt: []string{
			"FNDNG", "FUNDING",
		},
	},
	{
		Primary: "FUNERAL",
		Short:   "FNRL",
		Alt: []string{
			"FNRL", "FUNERAL",
		},
	},
	{
		Primary: "FURNACE",
		Short:   "FRNC",
		Alt: []string{
			"FRNC", "FURN", "FURNACE",
		},
	},
	{
		Primary: "FURNISHING",
		Short:   "FURNG",
		Alt: []string{
			"FURN", "FURNG", "FURNISHING",
		},
	},
	{
		Primary: "FURNITURE",
		Short:   "FURN",
		Alt: []string{
			"FURN", "FURNITURE",
		},
	},
	{
		Primary: "FURRIER",
		Short:   "FUR",
		Alt: []string{
			"FUR", "FURRIER",
		},
	},
	{
		Primary: "FUSIL",
		Short:   "FUSL",
		Alt: []string{
			"FUSIL", "FUSL",
		},
	},
	{
		Primary: "FUSION",
		Short:   "FUSN",
		Alt: []string{
			"FUSION", "FUSN",
		},
	},
	{
		Primary: "GALAXY",
		Short:   "GALXY",
		Alt: []string{
			"GALAXY", "GALXY",
		},
	},
	{
		Primary: "GALLERY",
		Short:   "GLLRY",
		Alt: []string{
			"GALLERY", "GLLRY",
		},
	},
	{
		Primary: "GALVANIZING",
		Short:   "GLVNZNG",
		Alt: []string{
			"GALVANIZING", "GLVNZNG",
		},
	},
	{
		Primary: "GARAGE",
		Short:   "GRGE",
		Alt: []string{
			"GAR", "GARAGE", "GRGE",
		},
	},
	{
		Primary: "GARDEN",
		Short:   "GDNS",
		Alt: []string{
			"GARDEN", "GDN", "GDNS", "GRDN",
		},
	},
	{
		Primary: "GARDENER",
		Short:   "GRDNR",
		Alt: []string{
			"GARDENER", "GRDNR",
		},
	},
	{
		Primary: "GARMENT",
		Short:   "GMT",
		Alt: []string{
			"GARMENT", "GMT",
		},
	},
	{
		Primary: "GASOLINE",
		Short:   "GAS",
		Alt: []string{
			"GAS", "GASOLINE",
		},
	},
	{
		Primary: "GASTROENTEROLOGIST",
		Short:   "GASTRNTRLGST",
		Alt: []string{
			"GAST", "GASTRNTRLGST", "GASTROENTEROLOGIST",
		},
	},
	{
		Primary: "GASTROENTEROLOGY",
		Short:   "GASTRNTRLGY",
		Alt: []string{
			"GAST", "GASTRNTRLGY", "GASTROENTEROLOGY",
		},
	},
	{
		Primary: "GATEWAY",
		Short:   "GTWY",
		Alt: []string{
			"GATEWAY", "GTWY",
		},
	},
	{
		Primary: "GATHERING",
		Short:   "GTHRNG",
		Alt: []string{
			"GATHERING", "GTHRNG",
		},
	},
	{
		Primary: "GAZETTE",
		Short:   "GAZ",
		Alt: []string{
			"GAZ", "GAZETTE",
		},
	},
	{
		Primary: "GENERAL",
		Short:   "GEN",
		Alt: []string{
			"GEN", "GENERAL", "GENL", "GN",
		},
	},
	{
		Primary: "GENERATING",
		Short:   "GNRTNG",
		Alt: []string{
			"GENERATING", "GNRTNG",
		},
	},
	{
		Primary: "GENERATION",
		Short:   "GNRTN",
		Alt: []string{
			"GENERATION", "GNRTN",
		},
	},
	{
		Primary: "GENERATOR",
		Short:   "GNRTR",
		Alt: []string{
			"GENERATOR", "GNRTR",
		},
	},
	{
		Primary: "GENESIS",
		Short:   "GNSS",
		Alt: []string{
			"GENESIS", "GNSS",
		},
	},
	{
		Primary: "GENTLEMEN",
		Short:   "GNTLMN",
		Alt: []string{
			"GENTLEMEN", "GNTLMN",
		},
	},
	{
		Primary: "GEODESIC",
		Short:   "GDSC",
		Alt: []string{
			"GDSC", "GEODESIC",
		},
	},
	{
		Primary: "GEOLOGICAL",
		Short:   "GEOLGCL",
		Alt: []string{
			"GEOLGCL", "GEOLOGICAL",
		},
	},
	{
		Primary: "GEOLOGIST",
		Short:   "GEOL",
		Alt: []string{
			"GEOL", "GEOLOGIST",
		},
	},
	{
		Primary: "GEOLOGY",
		Short:   "GEOLGY",
		Alt: []string{
			"GEOLGY", "GEOLOGY",
		},
	},
	{
		Primary: "GEOPHYSICAL",
		Short:   "GEOPHYS",
		Alt: []string{
			"GEOPHYS", "GEOPHYSICAL",
		},
	},
	{
		Primary: "GERIATRIC",
		Short:   "GERI",
		Alt: []string{
			"GERI", "GERIATRIC",
		},
	},
	{
		Primary: "GIANT",
		Short:   "GNT",
		Alt: []string{
			"GIANT", "GNT",
		},
	},
	{
		Primary: "GIFTWEAR",
		Short:   "GFTWR",
		Alt: []string{
			"GFTWR", "GIFTWEAR",
		},
	},
	{
		Primary: "GINGERBREAD",
		Short:   "GNGRBRD",
		Alt: []string{
			"GINGERBREAD", "GNGRBRD",
		},
	},
	{
		Primary: "GLACE",
		Short:   "GLC",
		Alt: []string{
			"GLACE", "GLC",
		},
	},
	{
		Primary: "GLADIATOR",
		Short:   "GLDTR",
		Alt: []string{
			"GLADIATOR", "GLDTR",
		},
	},
	{
		Primary: "GLASS",
		Short:   "GLS",
		Alt: []string{
			"GL", "GLASS", "GLS",
		},
	},
	{
		Primary: "GLASSWARE",
		Short:   "GLWR",
		Alt: []string{
			"GLASSWARE", "GLWR",
		},
	},
	{
		Primary: "GLAZE",
		Short:   "GLZ",
		Alt: []string{
			"GLAZE", "GLZ",
		},
	},
	{
		Primary: "GLOBAL",
		Short:   "GLBL",
		Alt: []string{
			"GLBL", "GLOBAL",
		},
	},
	{
		Primary: "GLOVE",
		Short:   "GLV",
		Alt: []string{
			"GLOVE", "GLV",
		},
	},
	{
		Primary: "GOLDEN",
		Short:   "GLDN",
		Alt: []string{
			"GLDN", "GOLDEN",
		},
	},
	{
		Primary: "GOSPEL",
		Short:   "GSPL",
		Alt: []string{
			"GOSPEL", "GSPL",
		},
	},
	{
		Primary: "GOURMET",
		Short:   "GRMT",
		Alt: []string{
			"GOURMET", "GRMT",
		},
	},
	{
		Primary: "GOVERNMENT",
		Short:   "GOVT",
		Alt: []string{
			"GOV", "GOVERMT", "GOVERNMENT", "GOVT",
		},
	},
	{
		Primary: "GOVERNMENTAL",
		Short:   "GVRNMNTL",
		Alt: []string{
			"GOVERNMENTAL", "GVRNMNTL",
		},
	},
	{
		Primary: "GOVERNOR",
		Short:   "GOV",
		Alt: []string{
			"GOV", "GOVERNOR", "GVRNR",
		},
	},
	{
		Primary: "GRACE",
		Short:   "GRC",
		Alt: []string{
			"GRACE", "GRC",
		},
	},
	{
		Primary: "GRADE",
		Short:   "GRDE",
		Alt: []string{
			"GRADE", "GRD", "GRDE",
		},
	},
	{
		Primary: "GRADUATE",
		Short:   "GRAD",
		Alt: []string{
			"GRAD", "GRADUATE",
		},
	},
	{
		Primary: "GRAIN",
		Short:   "GRAN",
		Alt: []string{
			"GRAIN", "GRAN", "GRN",
		},
	},
	{
		Primary: "GRAND",
		Short:   "GRND",
		Alt: []string{
			"GRAND", "GRD", "GRND",
		},
	},
	{
		Primary: "GRANDMA",
		Short:   "GRNDMA",
		Alt: []string{
			"GRANDMA", "GRNDMA",
		},
	},
	{
		Primary: "GRANDPA",
		Short:   "GRNDPA",
		Alt: []string{
			"GRANDPA", "GRNDPA",
		},
	},
	{
		Primary: "GRANITE",
		Short:   "GRNT",
		Alt: []string{
			"GRAN", "GRANITE", "GRNT",
		},
	},
	{
		Primary: "GRAPHIC",
		Short:   "GRPHC",
		Alt: []string{
			"GRAPHIC", "GRPHC",
		},
	},
	{
		Primary: "GRAVEL",
		Short:   "GRVL",
		Alt: []string{
			"GRAV", "GRAVEL", "GRAVL", "GRVL",
		},
	},
	{
		Primary: "GREAT",
		Short:   "GRT",
		Alt: []string{
			"GREAT", "GRT",
		},
	},
	{
		Primary: "GREATER",
		Short:   "GRTR",
		Alt: []string{
			"GREATER", "GRTR",
		},
	},
	{
		Primary: "GREEN",
		Short:   "GRN",
		Alt: []string{
			"GREEN", "GRN",
		},
	},
	{
		Primary: "GREENHOUSE",
		Short:   "GRNHS",
		Alt: []string{
			"GREENHOUSE", "GRNHS", "GRNHSE",
		},
	},
	{
		Primary: "GREETING",
		Short:   "GRTG",
		Alt: []string{
			"GREETING", "GRTG",
		},
	},
	{
		Primary: "GRILL",
		Short:   "GRL",
		Alt: []string{
			"GRILL", "GRL",
		},
	},
	{
		Primary: "GRINDER",
		Short:   "GRNDR",
		Alt: []string{
			"GRINDER", "GRNDR",
		},
	},
	{
		Primary: "GRINDING",
		Short:   "GRIND",
		Alt: []string{
			"GRIND", "GRINDING", "GRNDG",
		},
	},
	{
		Primary: "GROCER",
		Short:   "GROC",
		Alt: []string{
			"GROC", "GROCER",
		},
	},
	{
		Primary: "GROCERY",
		Short:   "GROCY",
		Alt: []string{
			"GROCERY", "GROCY",
		},
	},
	{
		Primary: "GROOMING",
		Short:   "GROOM",
		Alt: []string{
			"GROOM", "GROOMING",
		},
	},
	{
		Primary: "GROUP",
		Short:   "GRP",
		Alt: []string{
			"GP", "GROUP", "GRP",
		},
	},
	{
		Primary: "GROVE",
		Short:   "GRV",
		Alt: []string{
			"GROVE", "GRV",
		},
	},
	{
		Primary: "GUARANTEED",
		Short:   "GRNTD",
		Alt: []string{
			"GRNTD", "GUARANTEED",
		},
	},
	{
		Primary: "GUARD",
		Short:   "GRD",
		Alt: []string{
			"GRD", "GUARD",
		},
	},
	{
		Primary: "GUARDIAN",
		Short:   "GRDN",
		Alt: []string{
			"GRDN", "GUARDIAN",
		},
	},
	{
		Primary: "GUIDANCE",
		Short:   "GUIDNC",
		Alt: []string{
			"GUID", "GUIDANCE", "GUIDNC",
		},
	},
	{
		Primary: "GUIDE",
		Short:   "GUID",
		Alt: []string{
			"GUID", "GUIDE",
		},
	},
	{
		Primary: "GUILD",
		Short:   "GLD",
		Alt: []string{
			"GLD", "GUILD",
		},
	},
	{
		Primary: "GUNNERY",
		Short:   "GY",
		Alt: []string{
			"GNNRY", "GUNNERY", "GY",
		},
	},
	{
		Primary: "GUNSMITH",
		Short:   "GNSMTH",
		Alt: []string{
			"GNSMTH", "GUNSMITH",
		},
	},
	{
		Primary: "GYMNASTIC",
		Short:   "GYM",
		Alt: []string{
			"GYM", "GYMNASTIC",
		},
	},
	{
		Primary: "GYNECOLOGIST",
		Short:   "GYN",
		Alt: []string{
			"GYN", "GYNCLGST", "GYNECOLOGIST",
		},
	},
	{
		Primary: "GYNECOLOGY",
		Short:   "GYNCLGY",
		Alt: []string{
			"GYN", "GYNCLGY", "GYNECOLOGY",
		},
	},
	{
		Primary: "GYPSUM",
		Short:   "GYPS",
		Alt: []string{
			"GYPS", "GYPSUM",
		},
	},
	{
		Primary: "HABERDASHERY",
		Short:   "HDASHY",
		Alt: []string{
			"HABERDASHERY", "HDASHY",
		},
	},
	{
		Primary: "HAIRCUTTING",
		Short:   "HAIRCTTNG",
		Alt: []string{
			"HAIRCTTNG", "HAIRCUTTING",
		},
	},
	{
		Primary: "HAIRDRESSER",
		Short:   "HRDRSSR",
		Alt: []string{
			"HAIRDRESSER", "HRDRSSR",
		},
	},
	{
		Primary: "HAIRSTYLING",
		Short:   "HRSTYLNG",
		Alt: []string{
			"HAIRSTYLING", "HRSTYLNG",
		},
	},
	{
		Primary: "HAIRSTYLIST",
		Short:   "HRSTYLST",
		Alt: []string{
			"HAIRSTYLIST", "HAIRSTYLS", "HRSTYLST",
		},
	},
	{
		Primary: "HALLMARK",
		Short:   "HLLMRK",
		Alt: []string{
			"HALLMARK", "HLLMRK",
		},
	},
	{
		Primary: "HAMBURGER",
		Short:   "HAMBGR",
		Alt: []string{
			"HAMB", "HAMBURGER", "HB", "HMBG",
		},
	},
	{
		Primary: "HANDBAG",
		Short:   "HBAG",
		Alt: []string{
			"HANDBAG", "HBAG",
		},
	},
	{
		Primary: "HANDICAPPED",
		Short:   "HNDCPD",
		Alt: []string{
			"HANDICAPPED", "HNDCPD",
		},
	},
	{
		Primary: "HANDICRAFT",
		Short:   "HNDCRFT",
		Alt: []string{
			"HANDCRAFT", "HANDICRAFT", "HNDCRFT",
		},
	},
	{
		Primary: "HANDLER",
		Short:   "HNDLR",
		Alt: []string{
			"HANDLER", "HNDLR",
		},
	},
	{
		Primary: "HANDLING",
		Short:   "HNDLG",
		Alt: []string{
			"HANDLING", "HDLG", "HNDLING",
		},
	},
	{
		Primary: "HANDPRINT",
		Short:   "HNDPRNT",
		Alt: []string{
			"HANDPRINT", "HNDPRNT",
		},
	},
	{
		Primary: "HANDY",
		Short:   "HNDY",
		Alt: []string{
			"HANDY", "HNDY",
		},
	},
	{
		Primary: "HANDYMAN",
		Short:   "HNDYMN",
		Alt: []string{
			"HANDYMAN", "HNDYMN",
		},
	},
	{
		Primary: "HAPPY",
		Short:   "HAP",
		Alt: []string{
			"HAP", "HAPPY",
		},
	},
	{
		Primary: "HARBOR",
		Short:   "HBR",
		Alt: []string{
			"HARB", "HARBOR", "HARBR", "HBR", "HRBOR",
		},
	},
	{
		Primary: "HARDWARE",
		Short:   "HDWR",
		Alt: []string{
			"HARDWARE", "HDWR",
		},
	},
	{
		Primary: "HARNESS",
		Short:   "HARN",
		Alt: []string{
			"HARN", "HARNESS",
		},
	},
	{
		Primary: "HATCHERY",
		Short:   "HTCHY",
		Alt: []string{
			"HATCHERY", "HTCHY",
		},
	},
	{
		Primary: "HAULING",
		Short:   "HLG",
		Alt: []string{
			"HAULING", "HLG",
		},
	},
	{
		Primary: "HAVEN",
		Short:   "HVN",
		Alt: []string{
			"HAVEN", "HVN",
		},
	},
	{
		Primary: "HAYSTACK",
		Short:   "HYSTCK",
		Alt: []string{
			"HAYSTACK", "HYSTCK",
		},
	},
	{
		Primary: "HEADACHE",
		Short:   "HDCH",
		Alt: []string{
			"HDCH", "HEADACHE",
		},
	},
	{
		Primary: "HEADLINER",
		Short:   "HDLNR",
		Alt: []string{
			"HDLNR", "HEADLINER",
		},
	},
	{
		Primary: "HEADQUARTERS",
		Short:   "HDQTRS",
		Alt: []string{
			"HDQS", "HEADQUARTERS", "HQ", "HQS", "HQTS",
		},
	},
	{
		Primary: "HEALTH",
		Short:   "HLTH",
		Alt: []string{
			"HEALTH", "HLTH",
		},
	},
	{
		Primary: "HEARING",
		Short:   "HEAR",
		Alt: []string{
			"HEAR", "HEARING", "HRNG",
		},
	},
	{
		Primary: "HEART",
		Short:   "HRT",
		Alt: []string{
			"HEART", "HRT",
		},
	},
	{
		Primary: "HEATING",
		Short:   "HTG",
		Alt: []string{
			"HEATG", "HEATING", "HTG", "HTNG",
		},
	},
	{
		Primary: "HEAVY",
		Short:   "HVY",
		Alt: []string{
			"HEAVY", "HVY",
		},
	},
	{
		Primary: "HEIGHT",
		Short:   "HTS",
		Alt: []string{
			"HEIGHT", "HT",
		},
	},
	{
		Primary: "HELICOPTER",
		Short:   "HLCPTR",
		Alt: []string{
			"HELICOPTER", "HLCPTR",
		},
	},
	{
		Primary: "HELPER",
		Short:   "HLPR",
		Alt: []string{
			"HELPER", "HLPR",
		},
	},
	{
		Primary: "HEMATOLOGIST",
		Short:   "HEMATL",
		Alt: []string{
			"HEMATL", "HEMATOLOGIST",
		},
	},
	{
		Primary: "HEMATOLOGY",
		Short:   "HEMATLGY",
		Alt: []string{
			"HEMATL", "HEMATLGY", "HEMATOLOGY",
		},
	},
	{
		Primary: "HERITAGE",
		Short:   "HRTG",
		Alt: []string{
			"HERITAGE", "HRTG",
		},
	},
	{
		Primary: "HERMITAGE",
		Short:   "HRMTG",
		Alt: []string{
			"HERMITAGE", "HRMTG",
		},
	},
	{
		Primary: "HICKORY",
		Short:   "HCKRY",
		Alt: []string{
			"HCKRY", "HICKORY",
		},
	},
	{
		Primary: "HIDEAWAY",
		Short:   "HDWY",
		Alt: []string{
			"HDWY", "HIDEAWAY",
		},
	},
	{
		Primary: "HIGHER",
		Short:   "HGHR",
		Alt: []string{
			"HGHR", "HIGHER",
		},
	},
	{
		Primary: "HIGHLAND",
		Short:   "HGLND",
		Alt: []string{
			"HGLND", "HIGHLAND",
		},
	},
	{
		Primary: "HIGHWAY",
		Short:   "HWY",
		Alt: []string{
			"HIGHWAY", "HWY",
		},
	},
	{
		Primary: "HILLTOP",
		Short:   "HLTP",
		Alt: []string{
			"HILLTOP", "HLTP",
		},
	},
	{
		Primary: "HISTORICAL",
		Short:   "HISTRCL",
		Alt: []string{
			"HIST", "HISTORCL", "HISTORICAL", "HISTRCL",
		},
	},
	{
		Primary: "HITCHING",
		Short:   "HTCHNG",
		Alt: []string{
			"HITCHING", "HTCHNG",
		},
	},
	{
		Primary: "HOBBY",
		Short:   "HOB",
		Alt: []string{
			"HOB", "HOBBY",
		},
	},
	{
		Primary: "HOLDING",
		Short:   "HLDNG",
		Alt: []string{
			"HLDNG", "HOLDG", "HOLDING",
		},
	},
	{
		Primary: "HOLIDAY",
		Short:   "HLDY",
		Alt: []string{
			"HLDY", "HOLIDAY",
		},
	},
	{
		Primary: "HOLINESS",
		Short:   "HLNSS",
		Alt: []string{
			"HLNSS", "HOLINESS",
		},
	},
	{
		Primary: "HOMESTEAD",
		Short:   "HMSTD",
		Alt: []string{
			"HMSTD", "HOMESTEAD",
		},
	},
	{
		Primary: "HOMEWORK",
		Short:   "HMWRK",
		Alt: []string{
			"HMWRK", "HOMEWORK",
		},
	},
	{
		Primary: "HONEYBEE",
		Short:   "HNYB",
		Alt: []string{
			"HNYB", "HONEYBEE",
		},
	},
	{
		Primary: "HONORABLE",
		Short:   "HON",
		Alt: []string{
			"HON", "HONORABLE",
		},
	},
	{
		Primary: "HORIZON",
		Short:   "HRZN",
		Alt: []string{
			"HORIZON", "HRZN",
		},
	},
	{
		Primary: "HORSE",
		Short:   "HORSE",
		Alt: []string{
			"HORSE", "HRS",
		},
	},
	{
		Primary: "HORTICULTURAL",
		Short:   "HORTL",
		Alt: []string{
			"HORT", "HORTICULTURAL", "HORTL",
		},
	},
	{
		Primary: "HORTICULTURE",
		Short:   "HORT",
		Alt: []string{
			"HORT", "HORTICULTURE",
		},
	},
	{
		Primary: "HOSIERY",
		Short:   "HSY",
		Alt: []string{
			"HOS", "HOSIERY", "HSY",
		},
	},
	{
		Primary: "HOSPICE",
		Short:   "HSPC",
		Alt: []string{
			"HOSP", "HOSPI", "HOSPICE", "HSPC",
		},
	},
	{
		Primary: "HOSPITAL",
		Short:   "HOSP",
		Alt: []string{
			"HOSP", "HOSPIT", "HOSPITAL", "HSP", "HSPTL",
		},
	},
	{
		Primary: "HOSPITALITY",
		Short:   "HOSPTY",
		Alt: []string{
			"HOSPITALITY", "HOSPTY",
		},
	},
	{
		Primary: "HOTEL",
		Short:   "HTL",
		Alt: []string{
			"HOT", "HOTEL", "HT", "HTL",
		},
	},
	{
		Primary: "HOUSE",
		Short:   "HSE",
		Alt: []string{
			"HOUSE", "HS", "HSE",
		},
	},
	{
		Primary: "HOUSEHOLD",
		Short:   "HSEHLD",
		Alt: []string{
			"HHLD", "HOUSEHOLD", "HSEHLD",
		},
	},
	{
		Primary: "HOUSEWARES",
		Short:   "HSWRS",
		Alt: []string{
			"HOUSEWARES", "HSWRS",
		},
	},
	{
		Primary: "HOUSING",
		Short:   "HSNG",
		Alt: []string{
			"HOUSING", "HOUSNG", "HSNG",
		},
	},
	{
		Primary: "HUMAN",
		Short:   "HMN",
		Alt: []string{
			"HMN", "HUMAN",
		},
	},
	{
		Primary: "HUNGRY",
		Short:   "HNGRY",
		Alt: []string{
			"HNGRY", "HUNGRY",
		},
	},
	{
		Primary: "HUNTER",
		Short:   "HNTR",
		Alt: []string{
			"HNTR", "HUNTER",
		},
	},
	{
		Primary: "HYDRAULIC",
		Short:   "HYDRLC",
		Alt: []string{
			"HYDRAULIC", "HYDRLC",
		},
	},
	{
		Primary: "HYGIENE",
		Short:   "HYGN",
		Alt: []string{
			"HYGIENE", "HYGN",
		},
	},
	{
		Primary: "HYPNOSIS",
		Short:   "HYPNS",
		Alt: []string{
			"HYPNOSIS", "HYPNS",
		},
	},
	{
		Primary: "IDEAL",
		Short:   "IDL",
		Alt: []string{
			"IDEAL", "IDL",
		},
	},
	{
		Primary: "IGNITION",
		Short:   "IGN",
		Alt: []string{
			"IGN", "IGNITION",
		},
	},
	{
		Primary: "IMAGE",
		Short:   "IMG",
		Alt: []string{
			"IMAGE", "IMG",
		},
	},
	{
		Primary: "IMAGINATION",
		Short:   "IMGNTN",
		Alt: []string{
			"IMAGINATION", "IMGNTN",
		},
	},
	{
		Primary: "IMAGING",
		Short:   "IMGNG",
		Alt: []string{
			"IMAGING", "IMGNG",
		},
	},
	{
		Primary: "IMMACULATE",
		Short:   "IMMCLT",
		Alt: []string{
			"IMMACULATE", "IMMCLT",
		},
	},
	{
		Primary: "IMMEDIATE",
		Short:   "IMMDT",
		Alt: []string{
			"IMMDT", "IMMEDIATE",
		},
	},
	{
		Primary: "IMMIGRATION",
		Short:   "IMMGRTN",
		Alt: []string{
			"IMMGRTN", "IMMIGRATION",
		},
	},
	{
		Primary: "IMPACT",
		Short:   "IMP",
		Alt: []string{
			"IMP", "IMPACT",
		},
	},
	{
		Primary: "IMPAIRED",
		Short:   "IMPRD",
		Alt: []string{
			"IMPAIRED", "IMPRD",
		},
	},
	{
		Primary: "IMPEDIMENT",
		Short:   "IMPDMNT",
		Alt: []string{
			"IMPDMNT", "IMPEDIMENT",
		},
	},
	{
		Primary: "IMPERIAL",
		Short:   "IMPRL",
		Alt: []string{
			"IMPERIAL", "IMPRL",
		},
	},
	{
		Primary: "IMPLEMENT",
		Short:   "IMPL",
		Alt: []string{
			"IMPL", "IMPLEMENT", "IMPLMNT", "IMPT",
		},
	},
	{
		Primary: "IMPLEMENTATION",
		Short:   "IMPLNTN",
		Alt: []string{
			"IMPLEMENTATION", "IMPLNTN",
		},
	},
	{
		Primary: "IMPORT",
		Short:   "IMPRT",
		Alt: []string{
			"IMPORT", "IMPRT",
		},
	},
	{
		Primary: "IMPORTATION",
		Short:   "IMPN",
		Alt: []string{
			"IMPN", "IMPORTATION",
		},
	},
	{
		Primary: "IMPORTED",
		Short:   "IMPRTD",
		Alt: []string{
			"IMPORTED", "IMPRTD",
		},
	},
	{
		Primary: "IMPORTER",
		Short:   "IMPRTR",
		Alt: []string{
			"IMP", "IMPORTER", "IMPRTR",
		},
	},
	{
		Primary: "IMPORTING",
		Short:   "IMPRTNG",
		Alt: []string{
			"IMPORTING", "IMPRTNG",
		},
	},
	{
		Primary: "IMPRESSION",
		Short:   "IMPRESS",
		Alt: []string{
			"IMPRESS", "IMPRESSION",
		},
	},
	{
		Primary: "IMPROVEMENT",
		Short:   "IMPRVMT",
		Alt: []string{
			"IMPROVEMENT", "IMPRV", "IMPRVMNT", "IMPRVMT",
		},
	},
	{
		Primary: "INCARNATION",
		Short:   "INCRNTN",
		Alt: []string{
			"INCARNATION", "INCRNTN",
		},
	},
	{
		Primary: "INCOME",
		Short:   "INCM",
		Alt: []string{
			"INCM", "INCO", "INCOME",
		},
	},
	{
		Primary: "INCORPORATED",
		Short:   "INC",
		Alt: []string{
			"INC", "INCOR", "INCORP", "INCORPORATED",
		},
	},
	{
		Primary: "INCORPORATION",
		Short:   "INCTN",
		Alt: []string{
			"INCORPORATION", "INCTN",
		},
	},
	{
		Primary: "INDEMNITY",
		Short:   "INDMNTY",
		Alt: []string{
			"INDEMNITY", "INDMNTY",
		},
	},
	{
		Primary: "INDEPENDENCE",
		Short:   "INDPDNC",
		Alt: []string{
			"INDEP", "INDEPENDENCE", "INDPDNC",
		},
	},
	{
		Primary: "INDEPENDENT",
		Short:   "INDPNDNT",
		Alt: []string{
			"IND", "INDEPENDENT", "INDPDNT", "INDPNDNT",
		},
	},
	{
		Primary: "INDIAN",
		Short:   "INDN",
		Alt: []string{
			"INDIAN", "INDN",
		},
	},
	{
		Primary: "INDUSTRIAL",
		Short:   "IND",
		Alt: []string{
			"IND", "INDL", "INDSTRL", "INDUS", "INDUSTRIA",
			"INDUSTRIAL", "INDUSTRL",
		},
	},
	{
		Primary: "INDUSTRY",
		Short:   "INDUST",
		Alt: []string{
			"IND", "INDS", "INDTRY", "INDUS", "INDUST",
			"INDUSTR", "INDUSTRY",
		},
	},
	{
		Primary: "INFANT",
		Short:   "INFNT",
		Alt: []string{
			"INF", "INFANT", "INFNT",
		},
	},
	{
		Primary: "INFINITE",
		Short:   "INFINT",
		Alt: []string{
			"INFINITE", "INFINT",
		},
	},
	{
		Primary: "INFIRM",
		Short:   "INFRM",
		Alt: []string{
			"INFIRM", "INFRM",
		},
	},
	{
		Primary: "INFIRMARY",
		Short:   "INFRMRY",
		Alt: []string{
			"INFIRMARY", "INFRMRY",
		},
	},
	{
		Primary: "INFORM",
		Short:   "INF",
		Alt: []string{
			"INF", "INFORM",
		},
	},
	{
		Primary: "INFORMATICS",
		Short:   "INFRMTCS",
		Alt: []string{
			"INFORMATICS", "INFRMTCS",
		},
	},
	{
		Primary: "INFORMATION",
		Short:   "INFO",
		Alt: []string{
			"INF", "INFO", "INFOR", "INFORMATION",
		},
	},
	{
		Primary: "INGREDIENT",
		Short:   "INGRDNT",
		Alt: []string{
			"INGRDNT", "INGREDIENT",
		},
	},
	{
		Primary: "INITIAL",
		Short:   "INIT",
		Alt: []string{
			"INITIAL", "INTL",
		},
	},
	{
		Primary: "INJECTION",
		Short:   "INJCTN",
		Alt: []string{
			"INJCTN", "INJECTION",
		},
	},
	{
		Primary: "INLAND",
		Short:   "INLND",
		Alt: []string{
			"INLAND", "INLND",
		},
	},
	{
		Primary: "INNER",
		Short:   "INNR",
		Alt: []string{
			"INNER", "INNR",
		},
	},
	{
		Primary: "INNKEEPER",
		Short:   "INNKPR",
		Alt: []string{
			"INNKEEPER", "INNKPR",
		},
	},
	{
		Primary: "INNOCENT",
		Short:   "INNCNT",
		Alt: []string{
			"INNCNT", "INNOCENT",
		},
	},
	{
		Primary: "INNOVATION",
		Short:   "INNVTN",
		Alt: []string{
			"INNOVATION", "INNVTN",
		},
	},
	{
		Primary: "INNOVATIVE",
		Short:   "INNVTV",
		Alt: []string{
			"INNOVATIVE", "INNVTV",
		},
	},
	{
		Primary: "INQUISITIVE",
		Short:   "INQSTV",
		Alt: []string{
			"INQ", "INQSTV", "INQUISITIVE",
		},
	},
	{
		Primary: "INSCRIPTION",
		Short:   "INSCRPTN",
		Alt: []string{
			"INSCRIPTION", "INSCRPTN",
		},
	},
	{
		Primary: "INSECURE",
		Short:   "INSCR",
		Alt: []string{
			"INSCR", "INSECURE",
		},
	},
	{
		Primary: "INSPECTION",
		Short:   "INSPCTN",
		Alt: []string{
			"INSPCTN", "INSPECTION", "INSPTN",
		},
	},
	{
		Primary: "INSPECTOR",
		Short:   "INSPCTR",
		Alt: []string{
			"INS", "INSP", "INSPCTR", "INSPECTOR",
		},
	},
	{
		Primary: "INSTALLATION",
		Short:   "INSTLTN",
		Alt: []string{
			"INSTALLATION", "INSTLTN",
		},
	},
	{
		Primary: "INSTALLER",
		Short:   "INSTLLR",
		Alt: []string{
			"INSTALLER", "INSTLLR",
		},
	},
	{
		Primary: "INSTALLMENT",
		Short:   "INSTL",
		Alt: []string{
			"INSTALLMENT", "INSTL",
		},
	},
	{
		Primary: "INSTANT",
		Short:   "INSTNT",
		Alt: []string{
			"INSTANT", "INSTNT",
		},
	},
	{
		Primary: "INSTITUTE",
		Short:   "INST",
		Alt: []string{
			"INST", "INSTI", "INSTIT", "INSTITUE", "INSTITUT",
			"INSTITUTE",
		},
	},
	{
		Primary: "INSTITUTION",
		Short:   "INSTN",
		Alt: []string{
			"INSTITUTION", "INSTN",
		},
	},
	{
		Primary: "INSTITUTIONAL",
		Short:   "INSTNL",
		Alt: []string{
			"INSTITUTIONAL", "INSTNL",
		},
	},
	{
		Primary: "INSTRUCTOR",
		Short:   "INSTRCTR",
		Alt: []string{
			"INST", "INSTR", "INSTRCTR", "INSTRUCTOR",
		},
	},
	{
		Primary: "INSTRUMENT",
		Short:   "INSTR",
		Alt: []string{
			"INSTR", "INSTRUMENT",
		},
	},
	{
		Primary: "INSTRUMENTATION",
		Short:   "INSTRMNTN",
		Alt: []string{
			"INSTRMNTN", "INSTRUMENTA", "INSTRUMENTATION",
		},
	},
	{
		Primary: "INSULATED",
		Short:   "INSLTD",
		Alt: []string{
			"INSLTD", "INSULATED",
		},
	},
	{
		Primary: "INSULATING",
		Short:   "INSULG",
		Alt: []string{
			"INSULATING", "INSULG",
		},
	},
	{
		Primary: "INSULATION",
		Short:   "INSLTN",
		Alt: []string{
			"INSLTN", "INSUL", "INSULATION", "INSULATN",
		},
	},
	{
		Primary: "INSURANCE",
		Short:   "INS",
		Alt: []string{
			"INS", "INSUR", "INSURAN", "INSURANCE",
		},
	},
	{
		Primary: "INTEGRATED",
		Short:   "INTGRTD",
		Alt: []string{
			"INTEGRATED", "INTGRTD",
		},
	},
	{
		Primary: "INTELLIGENCE",
		Short:   "INTLLGNC",
		Alt: []string{
			"INTELLIGENCE", "INTLLGNC",
		},
	},
	{
		Primary: "INTENTIONAL",
		Short:   "INTNTNL",
		Alt: []string{
			"INTENTIONAL", "INTNTL",
		},
	},
	{
		Primary: "INTERACTION",
		Short:   "INTRCTN",
		Alt: []string{
			"INTER", "INTERACTION", "INTRCTN",
		},
	},
	{
		Primary: "INTERACTIVE",
		Short:   "INTRCTV",
		Alt: []string{
			"INTERACTIVE", "INTRCTV",
		},
	},
	{
		Primary: "INTERCHANGE",
		Short:   "INTRCHNG",
		Alt: []string{
			"INTERCHANGE", "INTRCHNG",
		},
	},
	{
		Primary: "INTERCONTINENTAL",
		Short:   "INTERCON",
		Alt: []string{
			"INTERCON", "INTERCONTINENTAL",
		},
	},
	{
		Primary: "INTEREST",
		Short:   "INTRST",
		Alt: []string{
			"INTEREST", "INTRST",
		},
	},
	{
		Primary: "INTERFAITH",
		Short:   "INTRFTH",
		Alt: []string{
			"INTERFAITH", "INTRFTH",
		},
	},
	{
		Primary: "INTERIOR",
		Short:   "INTR",
		Alt: []string{
			"INT", "INTERIOR", "INTR",
		},
	},
	{
		Primary: "INTERMEDIATE",
		Short:   "INTER",
		Alt: []string{
			"INTER", "INTERMED", "INTERMEDIATE",
		},
	},
	{
		Primary: "INTERMEDICS",
		Short:   "INTRMDCS",
		Alt: []string{
			"INTERMEDICS", "INTRMDCS",
		},
	},
	{
		Primary: "INTERNAL",
		Short:   "INTERNL",
		Alt: []string{
			"INTER", "INTERNAL", "INTERNL",
		},
	},
	{
		Primary: "INTERNATIONAL",
		Short:   "INTRNTL",
		Alt: []string{
			"INTERNATI", "INTERNATIO", "INTERNATION", "INTERNATIONA", "INTERNATIONAL",
			"INTERNATL", "INTL", "INTNL", "INTRNTL", "INTRNTNL",
		},
	},
	{
		Primary: "INTERNIST",
		Short:   "INTERNST",
		Alt: []string{
			"INTER", "INTERNIST", "INTERNST",
		},
	},
	{
		Primary: "INTERSTATE",
		Short:   "INTSTE",
		Alt: []string{
			"INTERSTATE", "INTRST", "INTSTE",
		},
	},
	{
		Primary: "INTERVIEWER",
		Short:   "INTERV",
		Alt: []string{
			"INTERV", "INTERVIEWER",
		},
	},
	{
		Primary: "INVENTORY",
		Short:   "INVTY",
		Alt: []string{
			"INVEN", "INVENTORY", "INVTY",
		},
	},
	{
		Primary: "INVEST",
		Short:   "INVST",
		Alt: []string{
			"INVEST", "INVST",
		},
	},
	{
		Primary: "INVESTED",
		Short:   "INVSTD",
		Alt: []string{
			"INVESTED", "INVSTD",
		},
	},
	{
		Primary: "INVESTIGATION",
		Short:   "INVSTGTN",
		Alt: []string{
			"INVESTIGATION", "INVSTGTN",
		},
	},
	{
		Primary: "INVESTIGATIVE",
		Short:   "INVSTGTV",
		Alt: []string{
			"INVESTIGATIVE", "INVSTGTV",
		},
	},
	{
		Primary: "INVESTIGATOR",
		Short:   "INVSTR",
		Alt: []string{
			"INVESTIGATOR", "INVSTR",
		},
	},
	{
		Primary: "INVESTMENT",
		Short:   "INVSTMNT",
		Alt: []string{
			"INV", "INVESTMENT", "INVESTMNT", "INVESTMT", "INVST",
			"INVSTMNT", "INVSTMT",
		},
	},
	{
		Primary: "INVITATIONAL",
		Short:   "INVTNL",
		Alt: []string{
			"INVITATIONAL", "INVTNL",
		},
	},
	{
		Primary: "INVOICE",
		Short:   "INV",
		Alt: []string{
			"INV", "INVOICE",
		},
	},
	{
		Primary: "IRONWORK",
		Short:   "IRNWRK",
		Alt: []string{
			"IRNWRK", "IRONWORK",
		},
	},
	{
		Primary: "IRRIGATION",
		Short:   "IRRGTN",
		Alt: []string{
			"IRRGTN", "IRRIG", "IRRIGAT", "IRRIGATION",
		},
	},
	{
		Primary: "ISLAND",
		Short:   "ISLE",
		Alt: []string{
			"IS", "ISL", "ISLAND", "ISLE",
		},
	},
	{
		Primary: "ISLANDER",
		Short:   "ISLER",
		Alt: []string{
			"ISLANDER", "ISLER",
		},
	},
	{
		Primary: "ISOLATION",
		Short:   "ISO",
		Alt: []string{
			"ISO", "ISOLATION",
		},
	},
	{
		Primary: "ISOTOPE",
		Short:   "ISTP",
		Alt: []string{
			"ISOTOPE", "ISTP",
		},
	},
	{
		Primary: "ITALIAN",
		Short:   "ITAL",
		Alt: []string{
			"IT", "ITAL", "ITALIAN", "ITLN",
		},
	},
	{
		Primary: "JAILER",
		Short:   "JLR",
		Alt: []string{
			"JAILER", "JLR",
		},
	},
	{
		Primary: "JANITOR",
		Short:   "JAN",
		Alt: []string{
			"JAN", "JANITOR",
		},
	},
	{
		Primary: "JANITORIAL",
		Short:   "JANTRL",
		Alt: []string{
			"JAN", "JANITOR", "JANITORIAL", "JNTRL",
		},
	},
	{
		Primary: "JEWELER",
		Short:   "JWLR",
		Alt: []string{
			"JEWELER", "JWLR",
		},
	},
	{
		Primary: "JEWELRY",
		Short:   "JWLRY",
		Alt: []string{
			"JEWELRY", "JEWLRY", "JWLRY", "JWLY",
		},
	},
	{
		Primary: "JEWISH",
		Short:   "JEW",
		Alt: []string{
			"JEW", "JEWISH",
		},
	},
	{
		Primary: "JOBBER",
		Short:   "JOB",
		Alt: []string{
			"JOB", "JOBBER",
		},
	},
	{
		Primary: "JOINT",
		Short:   "JNT",
		Alt: []string{
			"JNT", "JOINT",
		},
	},
	{
		Primary: "JOURNAL",
		Short:   "JRNL",
		Alt: []string{
			"JOURNAL", "JRNL",
		},
	},
	{
		Primary: "JOURNALIST",
		Short:   "JRNLST",
		Alt: []string{
			"JOURNALIST", "JRNLST",
		},
	},
	{
		Primary: "JOURNEY",
		Short:   "JRNY",
		Alt: []string{
			"JOURNEY", "JRNY",
		},
	},
	{
		Primary: "JUBILEE",
		Short:   "JBL",
		Alt: []string{
			"JBL", "JUBILEE",
		},
	},
	{
		Primary: "JUDGE",
		Short:   "JDG",
		Alt: []string{
			"JD", "JDG", "JUDGE",
		},
	},
	{
		Primary: "JUICE",
		Short:   "JC",
		Alt: []string{
			"JC", "JUICE",
		},
	},
	{
		Primary: "JUNCTION",
		Short:   "JCT",
		Alt: []string{
			"JC", "JCT", "JCTION", "JCTN", "JUNCTION",
			"JUNCTN", "JUNCTON",
		},
	},
	{
		Primary: "JUNIOR",
		Short:   "JR",
		Alt: []string{
			"JR", "JUNIOR",
		},
	},
	{
		Primary: "JUSTICE",
		Short:   "JSTC",
		Alt: []string{
			"JSTC", "JUSTICE",
		},
	},
	{
		Primary: "JUVENILE",
		Short:   "JVNL",
		Alt: []string{
			"JUVENILE", "JVNL",
		},
	},
	{
		Primary: "KARATE",
		Short:   "KRT",
		Alt: []string{
			"KARATE", "KRT",
		},
	},
	{
		Primary: "KENNEL",
		Short:   "KNL",
		Alt: []string{
			"KENNEL", "KNL",
		},
	},
	{
		Primary: "KEYBOARD",
		Short:   "KYBRD",
		Alt: []string{
			"KEYBOARD", "KYBRD",
		},
	},
	{
		Primary: "KEYSTONE",
		Short:   "KEYSTN",
		Alt: []string{
			"KEYSTN", "KEYSTONE",
		},
	},
	{
		Primary: "KIDDIE",
		Short:   "KID",
		Alt: []string{
			"KID", "KIDDIE",
		},
	},
	{
		Primary: "KINDERGARTEN",
		Short:   "KINDERGTN",
		Alt: []string{
			"KDRGRTN", "KINDERGARTEN", "KINDERGTN", "KNDGTRN", "KNDRGRTN",
		},
	},
	{
		Primary: "KINEMATICS",
		Short:   "KNMTCS",
		Alt: []string{
			"KINEMATICS", "KNMTCS",
		},
	},
	{
		Primary: "KINGDOM",
		Short:   "KNGDM",
		Alt: []string{
			"KINGDOM", "KNGDM",
		},
	},
	{
		Primary: "KITCHEN",
		Short:   "KTCHN",
		Alt: []string{
			"KIT", "KITCHEN", "KTCHN", "KTN",
		},
	},
	{
		Primary: "KNIGHT",
		Short:   "KNGHT",
		Alt: []string{
			"KNGHT", "KNIGHT", "KNT",
		},
	},
	{
		Primary: "KNITTED",
		Short:   "KNTTD",
		Alt: []string{
			"KNITTED", "KNTTD",
		},
	},
	{
		Primary: "KNITTING",
		Short:   "KNT",
		Alt: []string{
			"KNITTING", "KNT",
		},
	},
	{
		Primary: "KNITWEAR",
		Short:   "KNTWR",
		Alt: []string{
			"KNITWEAR", "KNTWR",
		},
	},
	{
		Primary: "KOSHER",
		Short:   "KSHR",
		Alt: []string{
			"KOSHER", "KSHR",
		},
	},
	{
		Primary: "LABEL",
		Short:   "LBL",
		Alt: []string{
			"LAB", "LABEL", "LBL",
		},
	},
	{
		Primary: "LABORATORY",
		Short:   "LAB",
		Alt: []string{
			"LAB", "LABORATORY",
		},
	},
	{
		Primary: "LABORER",
		Short:   "LBR",
		Alt: []string{
			"LABORER", "LBR",
		},
	},
	{
		Primary: "LACQUER",
		Short:   "LACQ",
		Alt: []string{
			"LACQ", "LACQUER",
		},
	},
	{
		Primary: "LAMINATE",
		Short:   "LMNT",
		Alt: []string{
			"LAMINATE", "LMNT",
		},
	},
	{
		Primary: "LAMINATING",
		Short:   "LMNTNG",
		Alt: []string{
			"LAMINATING", "LMNTNG",
		},
	},
	{
		Primary: "LANCE",
		Short:   "LNC",
		Alt: []string{
			"LANCE", "LNC",
		},
	},
	{
		Primary: "LANDFILL",
		Short:   "LNDFLL",
		Alt: []string{
			"LANDFILL", "LNDFLL",
		},
	},
	{
		Primary: "LANDMARK",
		Short:   "LNDMRK",
		Alt: []string{
			"LANDMARK", "LNDMRK",
		},
	},
	{
		Primary: "LANDSCAPE",
		Short:   "LNDSCP",
		Alt: []string{
			"LANDSCAPE", "LANDSCP", "LDSCP", "LNDSCP",
		},
	},
	{
		Primary: "LANDSCAPING",
		Short:   "LANDSCPG",
		Alt: []string{
			"LANDSCAPING", "LANDSCPG", "LDSCPG", "LNDSCPG",
		},
	},
	{
		Primary: "LANGUAGE",
		Short:   "LANG",
		Alt: []string{
			"LANG", "LANGUAGE",
		},
	},
	{
		Primary: "LAPIDARY",
		Short:   "LAPDRY",
		Alt: []string{
			"LAPDRY", "LAPIDARY",
		},
	},
	{
		Primary: "LARGE",
		Short:   "LRGE",
		Alt: []string{
			"LARGE", "LRGE",
		},
	},
	{
		Primary: "LARYNGOLOGIST",
		Short:   "LARYNGLGST",
		Alt: []string{
			"LAR", "LARYNGLGST", "LARYNGOLOGIST",
		},
	},
	{
		Primary: "LARYNGOLOGY",
		Short:   "LARYNGLGY",
		Alt: []string{
			"LAR", "LARYNGLGY", "LARYNGOLOGY",
		},
	},
	{
		Primary: "LASER",
		Short:   "LSR",
		Alt: []string{
			"LASER", "LSR",
		},
	},
	{
		Primary: "LASTING",
		Short:   "LSTNG",
		Alt: []string{
			"LASTING", "LSTNG",
		},
	},
	{
		Primary: "LATHING",
		Short:   "LTHG",
		Alt: []string{
			"LATHING", "LTHG",
		},
	},
	{
		Primary: "LATTER",
		Short:   "LTTR",
		Alt: []string{
			"LATTER", "LTTR",
		},
	},
	{
		Primary: "LAUNDERER",
		Short:   "LDRER",
		Alt: []string{
			"LAUNDERER", "LDRER",
		},
	},
	{
		Primary: "LAUNDROMAT",
		Short:   "LNDRMT",
		Alt: []string{
			"LAUNDROMAT", "LNDRMT",
		},
	},
	{
		Primary: "LAUNDRY",
		Short:   "LNDRY",
		Alt: []string{
			"LAUNDRY", "LDRY", "LNDRY",
		},
	},
	{
		Primary: "LAWYER",
		Short:   "LWYR",
		Alt: []string{
			"LAWYER", "LGL", "LWYR",
		},
	},
	{
		Primary: "LEADER",
		Short:   "LDR",
		Alt: []string{
			"LDR", "LEADER",
		},
	},
	{
		Primary: "LEAGUE",
		Short:   "LEA",
		Alt: []string{
			"LEA", "LEAG", "LEAGUE", "LGE",
		},
	},
	{
		Primary: "LEARNING",
		Short:   "LEARN",
		Alt: []string{
			"LEARN", "LEARNING", "LRNG",
		},
	},
	{
		Primary: "LEASE",
		Short:   "LS",
		Alt: []string{
			"LEAS", "LEASE", "LS",
		},
	},
	{
		Primary: "LEASING",
		Short:   "LEASE",
		Alt: []string{
			"LEASE", "LEASING", "LSG", "LSNG",
		},
	},
	{
		Primary: "LEATHER",
		Short:   "LTHR",
		Alt: []string{
			"LEA", "LEATHER", "LTHR",
		},
	},
	{
		Primary: "LECTURE",
		Short:   "LECT",
		Alt: []string{
			"LECT", "LECTURE",
		},
	},
	{
		Primary: "LECTURER",
		Short:   "LECTR",
		Alt: []string{
			"LEC", "LECT", "LECTR", "LECTURER",
		},
	},
	{
		Primary: "LEGAL",
		Short:   "LGL",
		Alt: []string{
			"LEG", "LEGAL", "LGL",
		},
	},
	{
		Primary: "LEGION",
		Short:   "LGN",
		Alt: []string{
			"LEGION", "LGN",
		},
	},
	{
		Primary: "LEISURE",
		Short:   "LSUR",
		Alt: []string{
			"LEISURE", "LSR", "LSUR",
		},
	},
	{
		Primary: "LENGTH",
		Short:   "LNGTH",
		Alt: []string{
			"LENGTH", "LNGTH",
		},
	},
	{
		Primary: "LESSOR",
		Short:   "LSSR",
		Alt: []string{
			"LESSOR", "LSSR",
		},
	},
	{
		Primary: "LETTER",
		Short:   "LTR",
		Alt: []string{
			"LETTER", "LTE", "LTR",
		},
	},
	{
		Primary: "LETTERPRESS",
		Short:   "LTRPRS",
		Alt: []string{
			"LETTERPRESS", "LTRPRS",
		},
	},
	{
		Primary: "LEVER",
		Short:   "LVR",
		Alt: []string{
			"LEVER", "LVR",
		},
	},
	{
		Primary: "LIABILITY",
		Short:   "LBLTY",
		Alt: []string{
			"LBLTY", "LIABILITY",
		},
	},
	{
		Primary: "LIBERTY",
		Short:   "LBRTY",
		Alt: []string{
			"LBRTY", "LIBERTY", "LIBTY",
		},
	},
	{
		Primary: "LIBRARIAN",
		Short:   "LIBRN",
		Alt: []string{
			"LBRN", "LIB", "LIBR", "LIBRARIAN", "LIBRN",
		},
	},
	{
		Primary: "LIBRARY",
		Short:   "LBRY",
		Alt: []string{
			"LBRRY", "LBRY", "LIB", "LIBRAR", "LIBRARY",
			"LIBRY",
		},
	},
	{
		Primary: "LICENSED",
		Short:   "LCNSD",
		Alt: []string{
			"LCNSD", "LICENSED",
		},
	},
	{
		Primary: "LIEUTENANT",
		Short:   "LT",
		Alt: []string{
			"LIEUTENANT", "LT",
		},
	},
	{
		Primary: "LIGHT",
		Short:   "LGT",
		Alt: []string{
			"LGT", "LIGHT", "LIT",
		},
	},
	{
		Primary: "LIGHTER",
		Short:   "LGHTR",
		Alt: []string{
			"LGHTR", "LIGHTER",
		},
	},
	{
		Primary: "LIGHTING",
		Short:   "LIGHT",
		Alt: []string{
			"LGHTG", "LIGHT", "LIGHTING", "LTG",
		},
	},
	{
		Primary: "LIMIT",
		Short:   "LMT",
		Alt: []string{
			"LIMIT", "LMT",
		},
	},
	{
		Primary: "LIMITED",
		Short:   "LTD",
		Alt: []string{
			"LIMITED", "LMTD", "LTD",
		},
	},
	{
		Primary: "LIMITLESS",
		Short:   "LMTLSS",
		Alt: []string{
			"LIMITLESS", "LMTLSS",
		},
	},
	{
		Primary: "LIMOUSINE",
		Short:   "LIMO",
		Alt: []string{
			"LIMO", "LIMOSINE", "LIMOUSINE", "LIMSNE",
		},
	},
	{
		Primary: "LINEN",
		Short:   "LIN",
		Alt: []string{
			"LIN", "LINEN",
		},
	},
	{
		Primary: "LINGERIE",
		Short:   "LNGR",
		Alt: []string{
			"LINGERIE", "LNGR",
		},
	},
	{
		Primary: "LINOLEUM",
		Short:   "LNLM",
		Alt: []string{
			"LINOLEUM", "LNLM",
		},
	},
	{
		Primary: "LIQUID",
		Short:   "LQD",
		Alt: []string{
			"LIQUID", "LQD",
		},
	},
	{
		Primary: "LIQUOR",
		Short:   "LQR",
		Alt: []string{
			"LIQUOR", "LQ", "LQR",
		},
	},
	{
		Primary: "LITHOGRAPH",
		Short:   "LITHO",
		Alt: []string{
			"LITHO", "LITHOGRAPH",
		},
	},
	{
		Primary: "LITHOGRAPHER",
		Short:   "LITHOR",
		Alt: []string{
			"LITHO", "LITHOGRAPHER", "LITHOR",
		},
	},
	{
		Primary: "LITHOGRAPHIC",
		Short:   "LITHOC",
		Alt: []string{
			"LITHOC", "LITHOGRAPHIC",
		},
	},
	{
		Primary: "LITHOGRAPHING",
		Short:   "LITHOG",
		Alt: []string{
			"LITHO", "LITHOG", "LITHOGRAPHING",
		},
	},
	{
		Primary: "LITHOGRAPHY",
		Short:   "LITHOY",
		Alt: []string{
			"LITHOGRAPHY", "LITHOY",
		},
	},
	{
		Primary: "LITTLE",
		Short:   "LTL",
		Alt: []string{
			"LITTLE", "LTL",
		},
	},
	{
		Primary: "LIVERY",
		Short:   "LV",
		Alt: []string{
			"LIVERY", "LV",
		},
	},
	{
		Primary: "LIVESTOCK",
		Short:   "LVSTCK",
		Alt: []string{
			"LIVESTOCK", "LVSTCK", "LVSTK",
		},
	},
	{
		Primary: "LIVING",
		Short:   "LVNG",
		Alt: []string{
			"LIVING", "LVNG",
		},
	},
	{
		Primary: "LOADER",
		Short:   "LODR",
		Alt: []string{
			"LDR", "LOADER", "LODR",
		},
	},
	{
		Primary: "LOADING",
		Short:   "LDNG",
		Alt: []string{
			"LDNG", "LOADING",
		},
	},
	{
		Primary: "LOBSTER",
		Short:   "LBSTR",
		Alt: []string{
			"LBSTR", "LOBSTER",
		},
	},
	{
		Primary: "LOCAL",
		Short:   "LCL",
		Alt: []string{
			"LCL", "LOC", "LOCAL",
		},
	},
	{
		Primary: "LOCATION",
		Short:   "LCTN",
		Alt: []string{
			"LCTN", "LOCATION",
		},
	},
	{
		Primary: "LOCKER",
		Short:   "LCKR",
		Alt: []string{
			"LCKR", "LOCKER",
		},
	},
	{
		Primary: "LOCKSMITH",
		Short:   "LOKSMTH",
		Alt: []string{
			"LCKSMTH", "LOCKSMITH", "LOCKSMTH", "LSMITH",
		},
	},
	{
		Primary: "LOCOMOTIVE",
		Short:   "LOCOM",
		Alt: []string{
			"LOCOM", "LOCOMOTIVE",
		},
	},
	{
		Primary: "LODGE",
		Short:   "LDG",
		Alt: []string{
			"LDG", "LDGE", "LODG", "LODGE",
		},
	},
	{
		Primary: "LOGGING",
		Short:   "LOG",
		Alt: []string{
			"LOG", "LOGGING",
		},
	},
	{
		Primary: "LOGIC",
		Short:   "LGC",
		Alt: []string{
			"LGC", "LOGIC",
		},
	},
	{
		Primary: "LOGICAL",
		Short:   "LGCL",
		Alt: []string{
			"LGCL", "LOGICAL",
		},
	},
	{
		Primary: "LOGISTIC",
		Short:   "LOGISTC",
		Alt: []string{
			"LOGISTC", "LOGISTIC", "LOGS",
		},
	},
	{
		Primary: "LOGISTICIAN",
		Short:   "LOGISTN",
		Alt:     []string{"LOGISTICIAN"},
	},
	{
		Primary: "LOUNGE",
		Short:   "LNG",
		Alt: []string{
			"LNG", "LOUNGE",
		},
	},
	{
		Primary: "LUBRICANT",
		Short:   "LUBR",
		Alt: []string{
			"LUBR", "LUBRICANT", "LUBRICNT",
		},
	},
	{
		Primary: "LUBRICATION",
		Short:   "LUBE",
		Alt: []string{
			"LUBE", "LUBRICATION",
		},
	},
	{
		Primary: "LUCKY",
		Short:   "LCKY",
		Alt: []string{
			"LCKY", "LUCKY",
		},
	},
	{
		Primary: "LUGGAGE",
		Short:   "LUG",
		Alt: []string{
			"LUG", "LUGGAGE",
		},
	},
	{
		Primary: "LUMBER",
		Short:   "LMBR",
		Alt: []string{
			"LBR", "LMBR", "LUMBER",
		},
	},
	{
		Primary: "LUTHERAN",
		Short:   "LUTH",
		Alt: []string{
			"LUTH", "LUTHERAN",
		},
	},
	{
		Primary: "MACARONI",
		Short:   "MCRN",
		Alt: []string{
			"MACARONI", "MCRN",
		},
	},
	{
		Primary: "MACHINE",
		Short:   "MACH",
		Alt: []string{
			"MACH", "MACHINE", "MCH", "MCHINE",
		},
	},
	{
		Primary: "MACHINER",
		Short:   "MACHR",
		Alt: []string{
			"MACH", "MACHINER", "MACHR",
		},
	},
	{
		Primary: "MACHINERY",
		Short:   "MACHY",
		Alt: []string{
			"MACH", "MACHINERY", "MACHY", "MCHY",
		},
	},
	{
		Primary: "MACHINING",
		Short:   "MACHG",
		Alt: []string{
			"MACH", "MACHG", "MACHINING",
		},
	},
	{
		Primary: "MACHINIST",
		Short:   "MACHST",
		Alt: []string{
			"MACH", "MACHINIST", "MACHST",
		},
	},
	{
		Primary: "MAGAZINE",
		Short:   "MAG",
		Alt: []string{
			"MAG", "MAGAZINE",
		},
	},
	{
		Primary: "MAGIC",
		Short:   "MGC",
		Alt: []string{
			"MAGIC", "MGC",
		},
	},
	{
		Primary: "MAGNETIC",
		Short:   "MGNTC",
		Alt: []string{
			"MAGNETIC", "MGNTC",
		},
	},
	{
		Primary: "MAGNETO",
		Short:   "MGNTO",
		Alt: []string{
			"MAGNETO", "MGNTO",
		},
	},
	{
		Primary: "MAILER",
		Short:   "MLR",
		Alt: []string{
			"MAILER", "MLR",
		},
	},
	{
		Primary: "MAILSTOP CODE",
		Short:   "MSC",
		Alt: []string{
			"MAILSTOP CODE", "MS", "MS#", "MSC",
		},
	},
	{
		Primary: "MAINSAIL",
		Short:   "MNSL",
		Alt: []string{
			"MAINSAIL", "MNSL",
		},
	},
	{
		Primary: "MAINTENANCE",
		Short:   "MNTNC",
		Alt: []string{
			"MAINT", "MAINTENANCE", "MNTNC", "MTNCE",
		},
	},
	{
		Primary: "MAJESTIC",
		Short:   "MJSTC",
		Alt: []string{
			"MAJESTIC", "MJSTC",
		},
	},
	{
		Primary: "MAJOR",
		Short:   "MJR",
		Alt: []string{
			"MAJ", "MAJOR", "MJR",
		},
	},
	{
		Primary: "MAMMOGRAPHY",
		Short:   "MAMGRAPHY",
		Alt: []string{
			"MAMGRPHY", "MAMMOGRAPHY",
		},
	},
	{
		Primary: "MANAGE",
		Short:   "MANAG",
		Alt: []string{
			"MANAG", "MANAGE", "MNG",
		},
	},
	{
		Primary: "MANAGEMENT",
		Short:   "MGMT",
		Alt: []string{
			"MANAGE", "MANAGEMENT", "MANGMNT", "MGMENT", "MGMT",
			"MGT", "MNGMNT", "MNGMT", "MNGN",
		},
	},
	{
		Primary: "MANAGER",
		Short:   "MGR",
		Alt: []string{
			"MANAGE", "MANAGER", "MG", "MGR", "MNAGER",
			"MNGR",
		},
	},
	{
		Primary: "MANAGERIAL",
		Short:   "MGRL",
		Alt: []string{
			"MANAGERIAL", "MGRL",
		},
	},
	{
		Primary: "MANAGING",
		Short:   "MNGNG",
		Alt: []string{
			"MANAGING", "MGNG", "MNG", "MNGNG",
		},
	},
	{
		Primary: "MANOR",
		Short:   "MNR",
		Alt: []string{
			"MANOR", "MNR",
		},
	},
	{
		Primary: "MANPOWER",
		Short:   "MNPWR",
		Alt: []string{
			"MANPOWER", "MNPWR",
		},
	},
	{
		Primary: "MANUFACTURE",
		Short:   "MFR",
		Alt: []string{
			"MANF", "MANUF", "MANUFACTURE", "MFR",
		},
	},
	{
		Primary: "MANUFACTURER",
		Short:   "MFGR",
		Alt: []string{
			"MANUFACTURER", "MFGR", "MFR",
		},
	},
	{
		Primary: "MANUFACTURING",
		Short:   "MFG",
		Alt: []string{
			"MANUFACTURI", "MANUFACTURING", "MFG", "MFGNG",
		},
	},
	{
		Primary: "MAPLE",
		Short:   "MPL",
		Alt: []string{
			"MAPLE", "MPL",
		},
	},
	{
		Primary: "MARATHON",
		Short:   "MRTHN",
		Alt: []string{
			"MARATHON", "MRTHN",
		},
	},
	{
		Primary: "MARBLE",
		Short:   "MRBL",
		Alt: []string{
			"MARBLE", "MBL", "MRBL",
		},
	},
	{
		Primary: "MARINA",
		Short:   "MRNA",
		Alt: []string{
			"MARINA", "MRNA",
		},
	},
	{
		Primary: "MARINE",
		Short:   "MRNE",
		Alt: []string{
			"MAR", "MARINE", "MRNE",
		},
	},
	{
		Primary: "MARITIME",
		Short:   "MRTM",
		Alt: []string{
			"MARITIME", "MRTM",
		},
	},
	{
		Primary: "MARKET",
		Short:   "MKT",
		Alt: []string{
			"MARKET", "MKT", "MRKT",
		},
	},
	{
		Primary: "MARKETER",
		Short:   "MRKTR",
		Alt: []string{
			"MARKETER", "MRKTR",
		},
	},
	{
		Primary: "MARKETING",
		Short:   "MKTG",
		Alt: []string{
			"MARKETING", "MKT", "MKTG", "MKTING", "MKTNG",
			"MRKT", "MRKTG",
		},
	},
	{
		Primary: "MARKETPLACE",
		Short:   "MRKTPLC",
		Alt: []string{
			"MARKETPLACE", "MRKTPLC",
		},
	},
	{
		Primary: "MARKING",
		Short:   "MKG",
		Alt: []string{
			"MARKING", "MKG",
		},
	},
	{
		Primary: "MARSHALL",
		Short:   "MRSHLL",
		Alt: []string{
			"MARSHALL", "MRSHLL",
		},
	},
	{
		Primary: "MASON",
		Short:   "MSN",
		Alt: []string{
			"MASON", "MSN",
		},
	},
	{
		Primary: "MASONIC",
		Short:   "MSNC",
		Alt: []string{
			"MASONIC", "MSNC",
		},
	},
	{
		Primary: "MASONRY",
		Short:   "MASON",
		Alt: []string{
			"MASON", "MASONRY", "MSN",
		},
	},
	{
		Primary: "MASTER",
		Short:   "MSTR",
		Alt: []string{
			"MASTER", "MSTR",
		},
	},
	{
		Primary: "MATERIAL",
		Short:   "MTRL",
		Alt: []string{
			"MATERIAL", "MATL", "MTL", "MTRL",
		},
	},
	{
		Primary: "MATERIEL",
		Short:   "MATL",
		Alt: []string{
			"MATERIEL", "MTREL",
		},
	},
	{
		Primary: "MATERNITY",
		Short:   "MTRNTY",
		Alt: []string{
			"MATERNITY", "MTRNTY",
		},
	},
	{
		Primary: "MATTRESS",
		Short:   "MATRS",
		Alt: []string{
			"MAT", "MATRS", "MATT", "MATTRESS", "MATTRS",
		},
	},
	{
		Primary: "MAYOR",
		Short:   "MAY",
		Alt: []string{
			"MAY", "MAYOR", "MYR",
		},
	},
	{
		Primary: "MEADOW",
		Short:   "MDWS",
		Alt: []string{
			"MDW", "MEADOW",
		},
	},
	{
		Primary: "MEASURE",
		Short:   "MSR",
		Alt: []string{
			"MEASURE", "MSR",
		},
	},
	{
		Primary: "MEASUREMENT",
		Short:   "MSRMNT",
		Alt: []string{
			"MEASUREMENT", "MEASUREMNT", "MSRMNT",
		},
	},
	{
		Primary: "MECHANIC",
		Short:   "MECH",
		Alt: []string{
			"MCHNC", "MECH", "MECHANIC",
		},
	},
	{
		Primary: "MECHANICAL",
		Short:   "MECHL",
		Alt: []string{
			"MECH", "MECHANICAL", "MECHL",
		},
	},
	{
		Primary: "MEDIA",
		Short:   "MEDIA",
		Alt: []string{
			"MED", "MEDIA",
		},
	},
	{
		Primary: "MEDICAL",
		Short:   "MEDCL",
		Alt: []string{
			"MDCL", "MED", "MEDCL", "MEDIC", "MEDICAL",
			"MEDL",
		},
	},
	{
		Primary: "MEDICAMENT",
		Short:   "MEDCMNT",
		Alt: []string{
			"MEDCMNT", "MEDICAMENT",
		},
	},
	{
		Primary: "MEDICINE",
		Short:   "MEDCN",
		Alt: []string{
			"MED", "MEDCN", "MEDICINE",
		},
	},
	{
		Primary: "MEDIUM",
		Short:   "MEDM",
		Alt: []string{
			"MED", "MEDIUM", "MEDM",
		},
	},
	{
		Primary: "MEETING",
		Short:   "MTG",
		Alt: []string{
			"MEETING", "MTG",
		},
	},
	{
		Primary: "MELANGE",
		Short:   "MLNG",
		Alt: []string{
			"MELANGE", "MLNG",
		},
	},
	{
		Primary: "MEMBER",
		Short:   "MBR",
		Alt: []string{
			"MBR", "MEMBER",
		},
	},
	{
		Primary: "MEMBERSHIP",
		Short:   "MBRSHP",
		Alt: []string{
			"MBRSHP", "MEMBERSHIP",
		},
	},
	{
		Primary: "MEMBRANE",
		Short:   "MBRM",
		Alt: []string{
			"MBRM", "MEMBRANE",
		},
	},
	{
		Primary: "MEMORANDUM",
		Short:   "MEMO",
		Alt: []string{
			"MEMO", "MEMORANDUM",
		},
	},
	{
		Primary: "MEMORIAL",
		Short:   "MEML",
		Alt: []string{
			"MEM", "MEML", "MEMORIAL", "MEMRL",
		},
	},
	{
		Primary: "MEMORY",
		Short:   "MEM",
		Alt: []string{
			"MEM", "MEMORY",
		},
	},
	{
		Primary: "MENNONITE",
		Short:   "MENIT",
		Alt: []string{
			"MENIT", "MENNONITE",
		},
	},
	{
		Primary: "MENTAL",
		Short:   "MNTL",
		Alt: []string{
			"MENT", "MENTAL", "MNTL",
		},
	},
	{
		Primary: "MERCANTILE",
		Short:   "MERCTL",
		Alt: []string{
			"MERC", "MERCANTILE", "MERCTL",
		},
	},
	{
		Primary: "MERCHANDISE",
		Short:   "MDSE",
		Alt: []string{
			"MDSE", "MERCHANDISE",
		},
	},
	{
		Primary: "MERCHANDISER",
		Short:   "MRCHNDSR",
		Alt: []string{
			"MERCHANDISER", "MRCHNDSR",
		},
	},
	{
		Primary: "MERCHANDISING",
		Short:   "MDSNG",
		Alt: []string{
			"MDSNG", "MERCH", "MERCHANDISING", "MHDSG",
		},
	},
	{
		Primary: "MERCHANT",
		Short:   "MRCHNT",
		Alt: []string{
			"MCHNT", "MERCHANT", "MRCHNT",
		},
	},
	{
		Primary: "MERCURY",
		Short:   "MERC",
		Alt: []string{
			"MERC", "MERCURY",
		},
	},
	{
		Primary: "MERIDIONAL",
		Short:   "MRDNL",
		Alt: []string{
			"MERIDIONAL", "MRDNL",
		},
	},
	{
		Primary: "METAL",
		Short:   "METL",
		Alt: []string{
			"MET", "METAL", "METL", "MTL",
		},
	},
	{
		Primary: "METALLIZING",
		Short:   "MTLNG",
		Alt: []string{
			"METALLIZING", "MTLNG",
		},
	},
	{
		Primary: "METALLURGICAL",
		Short:   "METLLRGCL",
		Alt: []string{
			"MET", "METALLURGICAL", "METLLRGCL",
		},
	},
	{
		Primary: "METALLURGIST",
		Short:   "METLLRGST",
		Alt: []string{
			"MET", "METALLURGIST", "METLLRGST",
		},
	},
	{
		Primary: "METALLURGY",
		Short:   "MTLGY",
		Alt: []string{
			"METALLURGY", "MTLGY",
		},
	},
	{
		Primary: "METEOROLOGIST",
		Short:   "METRLGST",
		Alt: []string{
			"MET", "METEOROLOGIST", "METRLGST",
		},
	},
	{
		Primary: "METHOD",
		Short:   "METH",
		Alt: []string{
			"METH", "METHOD",
		},
	},
	{
		Primary: "METHODIST",
		Short:   "METHDST",
		Alt: []string{
			"METH", "METHDST", "METHODIST",
		},
	},
	{
		Primary: "METRIC",
		Short:   "MTRC",
		Alt: []string{
			"METRIC", "MTRC",
		},
	},
	{
		Primary: "METROPOLITAN",
		Short:   "METRO",
		Alt: []string{
			"METRO", "METROPOLITAN",
		},
	},
	{
		Primary: "MEXICAN",
		Short:   "MEX",
		Alt: []string{
			"MEX", "MEXICAN",
		},
	},
	{
		Primary: "MICRO",
		Short:   "MCR",
		Alt: []string{
			"MCR", "MICRO",
		},
	},
	{
		Primary: "MICROBIOLOGY",
		Short:   "MCRBLGY",
		Alt: []string{
			"MCRBLGY", "MICROBIOLOGY",
		},
	},
	{
		Primary: "MICROCOMPUTER",
		Short:   "MCRCMPTR",
		Alt: []string{
			"MCRCMPTR", "MICRO", "MICROCOMPUTER",
		},
	},
	{
		Primary: "MICRODATA",
		Short:   "MCRDT",
		Alt: []string{
			"MCRDT", "MICRODATA",
		},
	},
	{
		Primary: "MICROELECTRONIC",
		Short:   "MCRELCTRNC",
		Alt: []string{
			"MCRELCTRNC", "MICROELECTRONIC",
		},
	},
	{
		Primary: "MICROFICHE",
		Short:   "MCRFCH",
		Alt: []string{
			"MCRFCH", "MICROFICHE",
		},
	},
	{
		Primary: "MICROWAVE",
		Short:   "MCRWV",
		Alt: []string{
			"MCRWV", "MICROWAVE",
		},
	},
	{
		Primary: "MIDDLE",
		Short:   "MID",
		Alt: []string{
			"MID", "MIDDLE", "MIDL",
		},
	},
	{
		Primary: "MIDLAND",
		Short:   "MDLND",
		Alt: []string{
			"MDLND", "MIDLAND",
		},
	},
	{
		Primary: "MIDSHIPMAN",
		Short:   "MDSHPMN",
		Alt: []string{
			"MDSHPMN", "MIDSHIPMAN",
		},
	},
	{
		Primary: "MIDTOWN",
		Short:   "MDTWN",
		Alt: []string{
			"MDTWN", "MIDTOWN",
		},
	},
	{
		Primary: "MIDWAY",
		Short:   "MDWY",
		Alt: []string{
			"MDWY", "MIDWAY",
		},
	},
	{
		Primary: "MIDWEST",
		Short:   "MDWST",
		Alt: []string{
			"MDWST", "MIDWEST", "MIDWST",
		},
	},
	{
		Primary: "MIDWESTERN",
		Short:   "MDWSTRN",
		Alt: []string{
			"MDWSTRN", "MIDWESTERN",
		},
	},
	{
		Primary: "MILIEU",
		Short:   "ML",
		Alt: []string{
			"MILIEU", "ML",
		},
	},
	{
		Primary: "MILITARY",
		Short:   "MLTRY",
		Alt: []string{
			"MILITARY", "MLTRY",
		},
	},
	{
		Primary: "MILLINERY",
		Short:   "MILNRY",
		Alt: []string{
			"MILLINERY", "MLY",
		},
	},
	{
		Primary: "MILLING",
		Short:   "MIL",
		Alt: []string{
			"MIL", "MILLING",
		},
	},
	{
		Primary: "MILLWORK",
		Short:   "MLLWK",
		Alt: []string{
			"MILLWORK", "MLLWK",
		},
	},
	{
		Primary: "MINERAL",
		Short:   "MNRL",
		Alt: []string{
			"MIN", "MINERAL", "MNRL",
		},
	},
	{
		Primary: "MINIATURE",
		Short:   "MINI",
		Alt: []string{
			"MINI", "MINIATURE",
		},
	},
	{
		Primary: "MINING",
		Short:   "MIN",
		Alt: []string{
			"MIN", "MINING", "MINNG",
		},
	},
	{
		Primary: "MINISTER",
		Short:   "MINSTR",
		Alt: []string{
			"MINISTER", "MNTR",
		},
	},
	{
		Primary: "MINISTRY",
		Short:   "MNSTRY",
		Alt: []string{
			"MINISTRY", "MNSTRY",
		},
	},
	{
		Primary: "MINISCULE",
		Short:   "MNSCL",
		Alt: []string{
			"MINISCULE", "MNSCL",
		},
	},
	{
		Primary: "MIRROR",
		Short:   "MIR",
		Alt: []string{
			"MIR", "MIRROR",
		},
	},
	{
		Primary: "MISCELLANEOUS",
		Short:   "MISC",
		Alt: []string{
			"MISC", "MISCELLANEOUS",
		},
	},
	{
		Primary: "MISSILE",
		Short:   "MIS",
		Alt: []string{
			"MIS", "MISSILE",
		},
	},
	{
		Primary: "MISSION",
		Short:   "MSSN",
		Alt: []string{
			"MISSION", "MSN", "MSSN",
		},
	},
	{
		Primary: "MISSIONARY",
		Short:   "MSSNRY",
		Alt: []string{
			"MISSIONARY", "MSSNRY",
		},
	},
	{
		Primary: "MISTER",
		Short:   "MR",
		Alt: []string{
			"MISTER", "MR",
		},
	},
	{
		Primary: "MIXED",
		Short:   "MXD",
		Alt: []string{
			"MIXED", "MXD",
		},
	},
	{
		Primary: "MIXING",
		Short:   "MIX",
		Alt: []string{
			"MIX", "MIXING",
		},
	},
	{
		Primary: "MOBILE",
		Short:   "MBL",
		Alt: []string{
			"MBL", "MO", "MOB", "MOBILE",
		},
	},
	{
		Primary: "MOCCASIN",
		Short:   "MOC",
		Alt: []string{
			"MOC", "MOCCASIN",
		},
	},
	{
		Primary: "MODEL",
		Short:   "MDL",
		Alt: []string{
			"MDL", "MODEL",
		},
	},
	{
		Primary: "MODERN",
		Short:   "MOD",
		Alt: []string{
			"MDRN", "MOD", "MODERN",
		},
	},
	{
		Primary: "MOLDED",
		Short:   "MLD",
		Alt: []string{
			"MLD", "MOLDED",
		},
	},
	{
		Primary: "MOLDING",
		Short:   "MLDG",
		Alt: []string{
			"MLDG", "MOLDING",
		},
	},
	{
		Primary: "MONASTERY",
		Short:   "MONSTRY",
		Alt: []string{
			"MONASTERY", "MONSTRY",
		},
	},
	{
		Primary: "MONEY",
		Short:   "MNY",
		Alt: []string{
			"MNY", "MONEY",
		},
	},
	{
		Primary: "MONITORING",
		Short:   "MNTRNG",
		Alt: []string{
			"MNTRNG", "MONITORING",
		},
	},
	{
		Primary: "MONOGRAM",
		Short:   "MNGRM",
		Alt: []string{
			"MNGRM", "MONOGRAM",
		},
	},
	{
		Primary: "MONTHLY",
		Short:   "MNTHLY",
		Alt: []string{
			"MNTHLY", "MONTHLY",
		},
	},
	{
		Primary: "MONUMENT",
		Short:   "MNMT",
		Alt: []string{
			"MNMT", "MONU", "MONUMENT",
		},
	},
	{
		Primary: "MOOSE",
		Short:   "MSE",
		Alt: []string{
			"MOOSE", "MSE",
		},
	},
	{
		Primary: "MORTGAGE",
		Short:   "MRTG",
		Alt: []string{
			"MORTG", "MORTGAGE", "MORTGE", "MRTG", "MRTGE",
			"MTG", "MTGE",
		},
	},
	{
		Primary: "MORTICIAN",
		Short:   "MORT",
		Alt: []string{
			"MORT", "MORTICIAN",
		},
	},
	{
		Primary: "MORTUARY",
		Short:   "MRTRY",
		Alt: []string{
			"MORTUARY", "MRTRY",
		},
	},
	{
		Primary: "MOSAIC",
		Short:   "MOSC",
		Alt: []string{
			"MOSAIC", "MSC",
		},
	},
	{
		Primary: "MOTEL",
		Short:   "MTL",
		Alt: []string{
			"MOTEL", "MTL",
		},
	},
	{
		Primary: "MOTHER",
		Short:   "MTHR",
		Alt: []string{
			"MOTHER", "MTHR",
		},
	},
	{
		Primary: "MOTIF",
		Short:   "MTF",
		Alt: []string{
			"MOTIF", "MTF",
		},
	},
	{
		Primary: "MOTION",
		Short:   "MOTN",
		Alt: []string{
			"MOTION", "MOTN", "MTN",
		},
	},
	{
		Primary: "MOTOR",
		Short:   "MTR",
		Alt: []string{
			"MOTOR", "MTR",
		},
	},
	{
		Primary: "MOTORCYCLE",
		Short:   "MTRCYL",
		Alt: []string{
			"MOTORCYCLE", "MTCYC",
		},
	},
	{
		Primary: "MOULAGE",
		Short:   "MLG",
		Alt: []string{
			"MLG", "MOULAGE",
		},
	},
	{
		Primary: "MOULDING",
		Short:   "MLDNG",
		Alt: []string{
			"MLDNG", "MOULDING",
		},
	},
	{
		Primary: "MOUNT",
		Short:   "MT",
		Alt: []string{
			"MOUNT", "MT",
		},
	},
	{
		Primary: "MOUNTAIN",
		Short:   "MTN",
		Alt: []string{
			"MNTN", "MOUNTAIN", "MOUNTIN", "MTN",
		},
	},
	{
		Primary: "MOVEMENT",
		Short:   "MVMNT",
		Alt: []string{
			"MOVEMENT", "MVMNT",
		},
	},
	{
		Primary: "MOVER",
		Short:   "MVR",
		Alt: []string{
			"MOVER", "MVR",
		},
	},
	{
		Primary: "MOVIE",
		Short:   "MOV",
		Alt: []string{
			"MOV", "MOVIE",
		},
	},
	{
		Primary: "MOVING",
		Short:   "MOVE",
		Alt: []string{
			"MOVE", "MOVING", "MVG",
		},
	},
	{
		Primary: "MOWER",
		Short:   "MWR",
		Alt: []string{
			"MOWER", "MWR",
		},
	},
	{
		Primary: "MUFFLER",
		Short:   "MUFLR",
		Alt: []string{
			"MFLR", "MUFFLER", "MUFLR",
		},
	},
	{
		Primary: "MUNICIPAL",
		Short:   "MNCPL",
		Alt: []string{
			"MNCPL", "MUNICIPAL",
		},
	},
	{
		Primary: "MUNICIPALITY",
		Short:   "MNCPLTY",
		Alt: []string{
			"MNCPLTY", "MUNICIPALITY",
		},
	},
	{
		Primary: "MUSEUM",
		Short:   "MUS",
		Alt: []string{
			"MUS", "MUSEUM",
		},
	},
	{
		Primary: "MUSIC",
		Short:   "MUSC",
		Alt: []string{
			"MUS", "MUSC", "MUSIC",
		},
	},
	{
		Primary: "MUSICAL",
		Short:   "MUSCL",
		Alt: []string{
			"MUSCL", "MUSICAL",
		},
	},
	{
		Primary: "MUTUAL",
		Short:   "MUTL",
		Alt: []string{
			"MTL", "MUTL", "MUTUAL",
		},
	},
	{
		Primary: "MYSTIC",
		Short:   "MYSTC",
		Alt: []string{
			"MYSTC", "MYSTIC",
		},
	},
	{
		Primary: "NATION",
		Short:   "NAT",
		Alt: []string{
			"NAT", "NATION",
		},
	},
	{
		Primary: "NATIONAL",
		Short:   "NATL",
		Alt: []string{
			"NATIONAL", "NATL", "NTL",
		},
	},
	{
		Primary: "NATIONWIDE",
		Short:   "NTNWD",
		Alt: []string{
			"NATIONWIDE", "NTNWD",
		},
	},
	{
		Primary: "NATURAL",
		Short:   "NTRL",
		Alt: []string{
			"NATURAL", "NTRL",
		},
	},
	{
		Primary: "NATURALLY",
		Short:   "NTRLLY",
		Alt: []string{
			"NATURALLY", "NTRLLY",
		},
	},
	{
		Primary: "NAUTICAL",
		Short:   "NTCL",
		Alt: []string{
			"NAUTICAL", "NTCL",
		},
	},
	{
		Primary: "NAVAL",
		Short:   "NVL",
		Alt: []string{
			"NAVAL", "NVL",
		},
	},
	{
		Primary: "NAVEL",
		Short:   "NVEL",
		Alt: []string{
			"NAVEL", "NVEL",
		},
	},
	{
		Primary: "NAVIGATION",
		Short:   "NVGTN",
		Alt: []string{
			"NAVIGATION", "NVGTN",
		},
	},
	{
		Primary: "NAZARENE",
		Short:   "NAZ",
		Alt: []string{
			"NAZ", "NAZARENE",
		},
	},
	{
		Primary: "NECESSITY",
		Short:   "NEC",
		Alt: []string{
			"NEC", "NECESSITY",
		},
	},
	{
		Primary: "NECKWEAR",
		Short:   "NCKWR",
		Alt: []string{
			"NCKWR", "NECKWEAR",
		},
	},
	{
		Primary: "NEIGHBORHOOD",
		Short:   "NGHBRHD",
		Alt: []string{
			"NEIGHBORHOOD", "NGHBRHG",
		},
	},
	{
		Primary: "NEPHROLOGY",
		Short:   "NEPH",
		Alt: []string{
			"NEPH", "NEPHROLOGY",
		},
	},
	{
		Primary: "NETWORK",
		Short:   "NTWRK",
		Alt: []string{
			"NET", "NETWK", "NETWORK", "NTK", "NTWK",
			"NTWRK",
		},
	},
	{
		Primary: "NETWORKING",
		Short:   "NTWRKNG",
		Alt: []string{
			"NETWORKING", "NTWRKNG",
		},
	},
	{
		Primary: "NEUROBIOLOGY",
		Short:   "NEUROBIOL",
		Alt: []string{
			"NEUROBIOL", "NEUROBIOLOGY",
		},
	},
	{
		Primary: "NEUROLOGIST",
		Short:   "NEUROLGST",
		Alt: []string{
			"NEUROLGST", "NEUROLOGIST",
		},
	},
	{
		Primary: "NEUROLOGY",
		Short:   "NRLGY",
		Alt: []string{
			"NEUROLOGY", "NRLGY",
		},
	},
	{
		Primary: "NEWSPAPER",
		Short:   "NWSPPR",
		Alt: []string{
			"NEWSPAPER", "NSWPPR",
		},
	},
	{
		Primary: "NINTH",
		Short:   "9TH",
		Alt: []string{
			"9TH", "IX", "NINTH",
		},
	},
	{
		Primary: "NONCOMMISSIONED",
		Short:   "NC",
		Alt: []string{
			"NC", "NONCOMMISSIONED",
		},
	},
	{
		Primary: "NONFERROUS",
		Short:   "NFER",
		Alt: []string{
			"NFER", "NONFERROUS",
		},
	},
	{
		Primary: "NORTHERN",
		Short:   "NTHRN",
		Alt: []string{
			"NORTHERN", "NTHRN",
		},
	},
	{
		Primary: "NORTHSIDE",
		Short:   "NRTHSD",
		Alt: []string{
			"NORTHSIDE", "NRTHSD",
		},
	},
	{
		Primary: "NORTHWESTERN",
		Short:   "NWN",
		Alt: []string{
			"NORTHWESTERN", "NWN",
		},
	},
	{
		Primary: "NOTION",
		Short:   "NOT",
		Alt: []string{
			"NOT", "NOTION",
		},
	},
	{
		Primary: "NOVELTY",
		Short:   "NOVLT",
		Alt: []string{
			"NOVELTY", "NOVLT",
		},
	},
	{
		Primary: "NUCLEAR",
		Short:   "NUC",
		Alt: []string{
			"NUC", "NUCLEAR",
		},
	},
	{
		Primary: "NURSE",
		Short:   "NUR",
		Alt: []string{
			"NUR", "NURSE",
		},
	},
	{
		Primary: "NURSERY",
		Short:   "NRSY",
		Alt: []string{
			"NRSY", "NURS", "NURSERY",
		},
	},
	{
		Primary: "NURSING",
		Short:   "NURSE",
		Alt: []string{
			"NURSE", "NURSING",
		},
	},
	{
		Primary: "NUTRITION",
		Short:   "NUTRI",
		Alt: []string{
			"NTRTN", "NUTRI", "NUTRITION",
		},
	},
	{
		Primary: "OBSERVATORY",
		Short:   "OBSRVTRY",
		Alt: []string{
			"OBSERVATORY", "OBSRVTRY",
		},
	},
	{
		Primary: "OBSTETRIC",
		Short:   "OBST",
		Alt: []string{
			"OBST", "OBSTETRIC",
		},
	},
	{
		Primary: "OBSTETRICIAN",
		Short:   "OB",
		Alt: []string{
			"OB", "OBSTETRICIAN", "OBSTRCN",
		},
	},
	{
		Primary: "OCCUPATION",
		Short:   "OCCUPTN",
		Alt: []string{
			"OCCUPATION", "OCCUPTN",
		},
	},
	{
		Primary: "OCCUPATIONAL",
		Short:   "OCCUPTNL",
		Alt: []string{
			"OCCUP", "OCCUPATIONAL", "OCCUPTNL",
		},
	},
	{
		Primary: "OCEAN",
		Short:   "OCN",
		Alt: []string{
			"OCEAN", "OCN",
		},
	},
	{
		Primary: "OFFICE",
		Short:   "OFC",
		Alt: []string{
			"OFC", "OFCE", "OFF", "OFFC", "OFFICE",
		},
	},
	{
		Primary: "OFFICER",
		Short:   "OFCR",
		Alt: []string{
			"OFFICER", "OFFICR", "OFFR",
		},
	},
	{
		Primary: "OFFICIAL",
		Short:   "OFCL",
		Alt: []string{
			"OFCL", "OFFICIAL",
		},
	},
	{
		Primary: "ONCOLOGIST",
		Short:   "ONCOL",
		Alt: []string{
			"ONCOL", "ONCOLOGIST",
		},
	},
	{
		Primary: "OPERATING",
		Short:   "OPG",
		Alt: []string{
			"OP", "OPERATING", "OPG", "OPRTNG",
		},
	},
	{
		Primary: "OPERATION",
		Short:   "OPRN",
		Alt: []string{
			"OP", "OPER", "OPERATION", "OPN", "OPR",
			"OPRN",
		},
	},
	{
		Primary: "OPERATIONAL",
		Short:   "OPRTNL",
		Alt: []string{
			"OP", "OPERATIONAL", "OPRTNL",
		},
	},
	{
		Primary: "OPERATIVE",
		Short:   "OPTV",
		Alt: []string{
			"OPER", "OPERATIVE", "OPTV",
		},
	},
	{
		Primary: "OPERATOR",
		Short:   "OPR",
		Alt: []string{
			"OP", "OPER", "OPERATOR", "OPR", "OPRTR",
		},
	},
	{
		Primary: "OPHTHALMIC",
		Short:   "OPHT",
		Alt: []string{
			"OPHT", "OPHTHALMIC",
		},
	},
	{
		Primary: "OPHTHALMOLOGIST",
		Short:   "OPH",
		Alt: []string{
			"OPH", "OPHTHALMOLOGIST",
		},
	},
	{
		Primary: "OPPORTUNITY",
		Short:   "OPRTNTY",
		Alt: []string{
			"OPPORTUNITY", "OPRTNTY",
		},
	},
	{
		Primary: "OPTICAL",
		Short:   "OPTIC",
		Alt: []string{
			"OPT", "OPTIC", "OPTICAL",
		},
	},
	{
		Primary: "OPTICIAN",
		Short:   "OPTCN",
		Alt: []string{
			"OPT", "OPTCN", "OPTICIAN",
		},
	},
	{
		Primary: "OPTOMETRIST",
		Short:   "OPTOM",
		Alt: []string{
			"OPTOM", "OPTOMETRIST",
		},
	},
	{
		Primary: "ORANGE",
		Short:   "ORNG",
		Alt: []string{
			"ORANGE", "ORNG",
		},
	},
	{
		Primary: "ORCHARD",
		Short:   "ORCH",
		Alt: []string{
			"ORCH", "ORCHARD", "ORCHRD",
		},
	},
	{
		Primary: "ORDER",
		Short:   "ORDR",
		Alt: []string{
			"ORD", "ORDER", "ORDR",
		},
	},
	{
		Primary: "ORDERING",
		Short:   "ORDNG",
		Alt: []string{
			"ORDERING", "ORDNG",
		},
	},
	{
		Primary: "ORDINATOR",
		Short:   "ORDNTR",
		Alt: []string{
			"ORDINATOR", "ORDNTR",
		},
	},
	{
		Primary: "ORDNANCE",
		Short:   "ORD",
		Alt: []string{
			"ORD", "ORDNANCE",
		},
	},
	{
		Primary: "ORGANIZATION",
		Short:   "ORGN",
		Alt: []string{
			"ORGANIZATION", "ORGN",
		},
	},
	{
		Primary: "ORGANIZATIONAL",
		Short:   "ORGNL",
		Alt: []string{
			"ORGANIZATIONAL", "ORGNL",
		},
	},
	{
		Primary: "ORIENTAL",
		Short:   "ORNTL",
		Alt: []string{
			"ORIENTAL", "ORNTL",
		},
	},
	{
		Primary: "ORNAMENTAL",
		Short:   "ORNMTL",
		Alt: []string{
			"ORNA", "ORNAMENTAL", "ORNMTL",
		},
	},
	{
		Primary: "ORTHOPEDIC",
		Short:   "ORTHO",
		Alt: []string{
			"ORTHO", "ORTHOPEDIC", "ORTHPD",
		},
	},
	{
		Primary: "ORTHOPTIST",
		Short:   "ORTHOPTST",
		Alt: []string{
			"ORTHOPTIST", "ORTHOPTST",
		},
	},
	{
		Primary: "OSTEOPATH",
		Short:   "OSTEOPTH",
		Alt: []string{
			"OSTEO", "OSTEOPATH", "OSTEOPTH",
		},
	},
	{
		Primary: "OSTEOPATHIC",
		Short:   "OSTEOPTHC",
		Alt: []string{
			"OSTEO", "OSTEOPATHIC", "OSTEOPTHC",
		},
	},
	{
		Primary: "OTOLOGY",
		Short:   "OTO",
		Alt: []string{
			"OTO", "OTOLOGY",
		},
	},
	{
		Primary: "OTORHINOLRYNGY",
		Short:   "OTRHNLRYNGY",
		Alt: []string{
			"OTORHINOLRYNGY", "OTRHNLRYNGY",
		},
	},
	{
		Primary: "OUTDOOR",
		Short:   "OTDR",
		Alt: []string{
			"OTDR", "OUTDOOR",
		},
	},
	{
		Primary: "OUTLET",
		Short:   "OUTLT",
		Alt: []string{
			"OTLT", "OUTL", "OUTLET", "OUTLT",
		},
	},
	{
		Primary: "OVERHEAD",
		Short:   "OVRHD",
		Alt: []string{
			"OVERHEAD", "OVRHD",
		},
	},
	{
		Primary: "OVERSIGHT",
		Short:   "OVRSGHT",
		Alt: []string{
			"OVERSIGHT", "OVRSGHT",
		},
	},
	{
		Primary: "OWNER",
		Short:   "OWNR",
		Alt: []string{
			"ONR", "OWN", "OWNE", "OWNER", "OWNR",
			"OWR",
		},
	},
	{
		Primary: "PACIFIC",
		Short:   "PAC",
		Alt: []string{
			"PAC", "PACIFIC", "PCF",
		},
	},
	{
		Primary: "PACKAGE",
		Short:   "PKG",
		Alt: []string{
			"PACKAGE", "PKG",
		},
	},
	{
		Primary: "PACKAGING",
		Short:   "PKGNG",
		Alt: []string{
			"PACKAGING", "PACKG", "PKG", "PKGNG",
		},
	},
	{
		Primary: "PACKER",
		Short:   "PKR",
		Alt: []string{
			"PACKER", "PKR",
		},
	},
	{
		Primary: "PACKING",
		Short:   "PCKG",
		Alt: []string{
			"PACKING", "PCKG", "PKG",
		},
	},
	{
		Primary: "PADDING",
		Short:   "PDG",
		Alt: []string{
			"PADDING", "PDG",
		},
	},
	{
		Primary: "PAINT",
		Short:   "PNT",
		Alt: []string{
			"PAINT", "PNT",
		},
	},
	{
		Primary: "PAINTER",
		Short:   "PNTR",
		Alt: []string{
			"PAINTER", "PNTR", "PTR",
		},
	},
	{
		Primary: "PAINTING",
		Short:   "PAINT",
		Alt: []string{
			"PAINT", "PAINTING", "PNT", "PNTG", "PNTNG",
		},
	},
	{
		Primary: "PALACE",
		Short:   "PALC",
		Alt: []string{
			"PALACE", "PALC", "PLC",
		},
	},
	{
		Primary: "PANCAKE",
		Short:   "PNCK",
		Alt: []string{
			"PANCAKE", "PNCK",
		},
	},
	{
		Primary: "PANHANDLE",
		Short:   "PNHDL",
		Alt: []string{
			"PANHANDLE", "PNHDL",
		},
	},
	{
		Primary: "PANTRY",
		Short:   "PNTRY",
		Alt: []string{
			"PANTRY", "PNTRY",
		},
	},
	{
		Primary: "PAPER",
		Short:   "PPR",
		Alt: []string{
			"PAPER", "PPR",
		},
	},
	{
		Primary: "PAPERBOARD",
		Short:   "PPRBD",
		Alt: []string{
			"PAPERBOARD", "PPRBD",
		},
	},
	{
		Primary: "PARADISE",
		Short:   "PRDS",
		Alt: []string{
			"PARADISE", "PRDS",
		},
	},
	{
		Primary: "PARKING",
		Short:   "PARK",
		Alt: []string{
			"PARK", "PARKING", "PRKG",
		},
	},
	{
		Primary: "PARKWAY",
		Short:   "PKWY",
		Alt: []string{
			"PARKWAY", "PKWY", "PKY",
		},
	},
	{
		Primary: "PARLOR",
		Short:   "PRLR",
		Alt: []string{
			"PARLOR", "PRLR",
		},
	},
	{
		Primary: "PARTICLEBOARD",
		Short:   "PTLBD",
		Alt: []string{
			"PARTICLEBOARD", "PTLBD",
		},
	},
	{
		Primary: "PARTNER",
		Short:   "PRTNR",
		Alt: []string{
			"PARTN", "PARTNER", "PARTNR", "PATNR", "PRT",
			"PRTNR", "PT", "PTNR", "PTR",
		},
	},
	{
		Primary: "PARTNERSHIP",
		Short:   "PRTNRSHP",
		Alt: []string{
			"PARTNERSHIP", "PRTNRSHP",
		},
	},
	{
		Primary: "PARTY",
		Short:   "PTY",
		Alt: []string{
			"PARTY", "PTY",
		},
	},
	{
		Primary: "PASSENGER",
		Short:   "PSSGR",
		Alt: []string{
			"PASS", "PASSENGER",
		},
	},
	{
		Primary: "PASTOR",
		Short:   "PSTR",
		Alt: []string{
			"PASTOR", "PST", "PSTR",
		},
	},
	{
		Primary: "PATCH",
		Short:   "PTCH",
		Alt: []string{
			"PATCH", "PTCH",
		},
	},
	{
		Primary: "PATENT",
		Short:   "PATNT",
		Alt: []string{
			"PAT", "PATENT", "PATNT",
		},
	},
	{
		Primary: "PATHOLOGIST",
		Short:   "PTHLGST",
		Alt: []string{
			"PATHOLOGIST", "PTHLGST",
		},
	},
	{
		Primary: "PATHOLOGY",
		Short:   "PATH",
		Alt: []string{
			"PATH", "PATHOLOGY",
		},
	},
	{
		Primary: "PATIO",
		Short:   "PAT",
		Alt: []string{
			"PAT", "PATIO",
		},
	},
	{
		Primary: "PATTERN",
		Short:   "PTTRN",
		Alt: []string{
			"PATTERN", "PTTRN",
		},
	},
	{
		Primary: "PAVING",
		Short:   "PAVE",
		Alt: []string{
			"PAV", "PAVE", "PAVING", "PVG",
		},
	},
	{
		Primary: "PAWNBROKER",
		Short:   "PWNBKR",
		Alt: []string{
			"PAWNBROKER", "PWNPKR",
		},
	},
	{
		Primary: "PAYABLE",
		Short:   "PAYABL",
		Alt: []string{
			"PAY", "PAYABL", "PAYABLE",
		},
	},
	{
		Primary: "PAYMENT",
		Short:   "PYMT",
		Alt: []string{
			"PAYMENT", "PYMT",
		},
	},
	{
		Primary: "PEDIATRIC",
		Short:   "PEDTRC",
		Alt: []string{
			"PED", "PEDIATRIC", "PEDTRC",
		},
	},
	{
		Primary: "PEDIATRICIAN",
		Short:   "PED",
		Alt: []string{
			"PED", "PEDIATRICIAN",
		},
	},
	{
		Primary: "PENNEY",
		Short:   "PNY",
		Alt: []string{
			"PENNEY", "PNY",
		},
	},
	{
		Primary: "PENINSULA",
		Short:   "PEN",
		Alt: []string{
			"PEN", "PENINSULA",
		},
	},
	{
		Primary: "PENSION",
		Short:   "PNSN",
		Alt: []string{
			"PENSION", "PNSN",
		},
	},
	{
		Primary: "PENTECOSTAL",
		Short:   "PENTE",
		Alt: []string{
			"PENT", "PENTE", "PENTECOSTAL", "PNTCSTL",
		},
	},
	{
		Primary: "PEOPLE",
		Short:   "PPL",
		Alt: []string{
			"PEOPLE", "PPL",
		},
	},
	{
		Primary: "PERFECT",
		Short:   "PERF",
		Alt: []string{
			"PERF", "PERFECT", "PRFCT",
		},
	},
	{
		Primary: "PERFORMANCE",
		Short:   "PERFORM",
		Alt: []string{
			"PERF", "PERFORM", "PERFORMANCE",
		},
	},
	{
		Primary: "PERIODICAL",
		Short:   "PERI",
		Alt: []string{
			"PERI", "PERIODICAL",
		},
	},
	{
		Primary: "PERIODONTIST",
		Short:   "PRDNTST",
		Alt: []string{
			"PERIODONTIST", "PRDNTST",
		},
	},
	{
		Primary: "PERSONAL",
		Short:   "PRSNL",
		Alt: []string{
			"PER", "PERS", "PERSONAL", "PRSNL",
		},
	},
	{
		Primary: "PERSONNEL",
		Short:   "PRSNNL",
		Alt: []string{
			"PERS", "PERSONNEL", "PRSNL", "PRSNNL",
		},
	},
	{
		Primary: "PESTICIDE",
		Short:   "PST",
		Alt: []string{
			"PESTICIDE", "PST",
		},
	},
	{
		Primary: "PETROLEUM",
		Short:   "PETRO",
		Alt: []string{
			"PETRO", "PETROLEUM",
		},
	},
	{
		Primary: "PETTY",
		Short:   "PTTY",
		Alt: []string{
			"PETTY", "PTTY",
		},
	},
	{
		Primary: "PHARMACEUTICAL",
		Short:   "PHARML",
		Alt: []string{
			"PHARMACEUTICAL", "PHARNL", "PHRM",
		},
	},
	{
		Primary: "PHARMACIST",
		Short:   "PHRMST",
		Alt: []string{
			"PHARM", "PHARMACIST", "PHRMST",
		},
	},
	{
		Primary: "PHARMACY",
		Short:   "PHARM",
		Alt: []string{
			"PHARM", "PHARMACY", "PHRM", "PHRMCY",
		},
	},
	{
		Primary: "PHONE",
		Short:   "PH",
		Alt: []string{
			"PHN", "PHONE",
		},
	},
	{
		Primary: "PHONOGRAPH",
		Short:   "PHONO",
		Alt: []string{
			"PHONO", "PHONOGRAPH",
		},
	},
	{
		Primary: "PHOTOGRAPH",
		Short:   "PHOTO",
		Alt: []string{
			"PHOTO", "PHOTOGRAPH",
		},
	},
	{
		Primary: "PHOTOGRAPHER",
		Short:   "PHOTOGR",
		Alt: []string{
			"PHOTOGR", "PHOTOGRAPHER",
		},
	},
	{
		Primary: "PHOTOGRAPHY",
		Short:   "PHOTO",
		Alt: []string{
			"PHOTO", "PHOTOGRAPHY",
		},
	},
	{
		Primary: "PHYSICAL",
		Short:   "PHYSCL",
		Alt: []string{
			"PHYS", "PHYSCL", "PHYSICAL",
		},
	},
	{
		Primary: "PHYSICIAN",
		Short:   "PHYS",
		Alt: []string{
			"PHYS", "PHYSCN", "PHYSICIAN",
		},
	},
	{
		Primary: "PHYSICIST",
		Short:   "PHYST",
		Alt: []string{
			"PHYS", "PHYSICIST", "PHYST",
		},
	},
	{
		Primary: "PIANO",
		Short:   "PNO",
		Alt: []string{
			"PIANO", "PNO",
		},
	},
	{
		Primary: "PICTURE",
		Short:   "PIC",
		Alt: []string{
			"PCTR", "PIC", "PICTURE",
		},
	},
	{
		Primary: "PIEDMONT",
		Short:   "PDMNT",
		Alt: []string{
			"PDMNT", "PIEDMONT",
		},
	},
	{
		Primary: "PIONEER",
		Short:   "PNR",
		Alt: []string{
			"PIONEER", "PNR",
		},
	},
	{
		Primary: "PIZZA",
		Short:   "PZ",
		Alt: []string{
			"PIZZA", "PZ", "PZA",
		},
	},
	{
		Primary: "PIZZERIA",
		Short:   "PZA",
		Alt: []string{
			"PIZZERIA", "PZ", "PZA",
		},
	},
	{
		Primary: "PLACE",
		Short:   "PL",
		Alt: []string{
			"PL", "PLACE",
		},
	},
	{
		Primary: "PLAIN",
		Short:   "PLN",
		Alt: []string{
			"PLAIN", "PLN",
		},
	},
	{
		Primary: "PLANNER",
		Short:   "PLNR",
		Alt: []string{
			"PLANNER", "PLNR",
		},
	},
	{
		Primary: "PLANNING",
		Short:   "PLAN",
		Alt: []string{
			"PLAN", "PLANNING", "PLG", "PLN", "PLNG",
			"PLNNG",
		},
	},
	{
		Primary: "PLANT",
		Short:   "PLNT",
		Alt: []string{
			"PLANT", "PLNT", "PLT",
		},
	},
	{
		Primary: "PLASTERING",
		Short:   "PLST",
		Alt: []string{
			"PLASTERING", "PLST",
		},
	},
	{
		Primary: "PLASTIC",
		Short:   "PLAS",
		Alt: []string{
			"PLAS", "PLASTIC", "PLST",
		},
	},
	{
		Primary: "PLATING",
		Short:   "PLTG",
		Alt: []string{
			"PLATING", "PLTG",
		},
	},
	{
		Primary: "PLATOON",
		Short:   "PLTN",
		Alt: []string{
			"PLATOON", "PLTN",
		},
	},
	{
		Primary: "PLAZA",
		Short:   "PLZ",
		Alt: []string{
			"PLAZA", "PLZ",
		},
	},
	{
		Primary: "PLEASANT",
		Short:   "PLSNT",
		Alt: []string{
			"PLEASANT", "PLSNT",
		},
	},
	{
		Primary: "PLUMBER",
		Short:   "PLMBR",
		Alt: []string{
			"PLMBR", "PLUMBER",
		},
	},
	{
		Primary: "PLUMBING",
		Short:   "PLBG",
		Alt: []string{
			"PLUMB", "PLUMBING",
		},
	},
	{
		Primary: "PLYWOOD",
		Short:   "PLYWD",
		Alt: []string{
			"PLYWD", "PLYWOOD",
		},
	},
	{
		Primary: "PODIATRIST",
		Short:   "PDTRST",
		Alt: []string{
			"PDTRST", "PODIATRIST",
		},
	},
	{
		Primary: "POINT",
		Short:   "PT",
		Alt: []string{
			"POINT", "PT",
		},
	},
	{
		Primary: "POLICE",
		Short:   "PLC",
		Alt: []string{
			"PLC", "POL", "POLICE",
		},
	},
	{
		Primary: "POLICY",
		Short:   "PLCY",
		Alt: []string{
			"PLCY", "POLICY",
		},
	},
	{
		Primary: "POLISHING",
		Short:   "POLSG",
		Alt: []string{
			"POLISHING", "POLSG",
		},
	},
	{
		Primary: "POLLUTION",
		Short:   "POLTN",
		Alt: []string{
			"POLLUTION", "POLTN",
		},
	},
	{
		Primary: "PORTER",
		Short:   "PRTR",
		Alt: []string{
			"PORTER", "PRTR", "PTR",
		},
	},
	{
		Primary: "POSITION",
		Short:   "PSTN",
		Alt: []string{
			"POSITION", "PSTN",
		},
	},
	{
		Primary: "POSTAL",
		Short:   "PSTL",
		Alt: []string{
			"POSTAL", "PSTL",
		},
	},
	{
		Primary: "POSTMASTER",
		Short:   "PM",
		Alt: []string{
			"PM", "POSTMASTER",
		},
	},
	{
		Primary: "POTTERY",
		Short:   "POT",
		Alt: []string{
			"POT", "POTTERY",
		},
	},
	{
		Primary: "POULTRY",
		Short:   "PLTY",
		Alt: []string{
			"PLTY", "POULTRY",
		},
	},
	{
		Primary: "POWER",
		Short:   "PWR",
		Alt: []string{
			"POWER", "PWR",
		},
	},
	{
		Primary: "PRACTICAL",
		Short:   "PRACL",
		Alt: []string{
			"PRAC", "PRACL", "PRACTICAL",
		},
	},
	{
		Primary: "PRACTICE",
		Short:   "PRAC",
		Alt: []string{
			"PRAC", "PRACTICE", "PRCTC",
		},
	},
	{
		Primary: "PRACTITIONER",
		Short:   "PRACTNR",
		Alt: []string{
			"PRAC", "PRACTITIONER", "PRACTNR", "PRCTTNR",
		},
	},
	{
		Primary: "PRAIRIE",
		Short:   "PR",
		Alt: []string{
			"PR", "PRAIRIE",
		},
	},
	{
		Primary: "PRECISION",
		Short:   "PRCSN",
		Alt: []string{
			"PRCSN", "PRECISION",
		},
	},
	{
		Primary: "PREFABRICATED",
		Short:   "PFAB",
		Alt: []string{
			"PFAB", "PREFABRICATED",
		},
	},
	{
		Primary: "PREFERRED",
		Short:   "PREF",
		Alt: []string{
			"PREF", "PREFERRED",
		},
	},
	{
		Primary: "PREMIER",
		Short:   "PREM",
		Alt: []string{
			"PREM", "PREMIER",
		},
	},
	{
		Primary: "PREPARATION",
		Short:   "PREP",
		Alt: []string{
			"PREP", "PREPARATION",
		},
	},
	{
		Primary: "PREPARER",
		Short:   "PRPRR",
		Alt: []string{
			"PREPARER", "PRPRR",
		},
	},
	{
		Primary: "PRESBYTERIAN",
		Short:   "PRESBY",
		Alt: []string{
			"PRES", "PRESBY", "PRESBYTERIAN", "PRSBY",
		},
	},
	{
		Primary: "PRESCHOOL",
		Short:   "PRSCHL",
		Alt: []string{
			"PRESCHOOL", "PRSCHL",
		},
	},
	{
		Primary: "PRESCRIPTION",
		Short:   "PRESCR",
		Alt: []string{
			"PRESCR", "PRESCRIPTION",
		},
	},
	{
		Primary: "PRESERVING",
		Short:   "PRSV",
		Alt: []string{
			"PRESERVING", "PRSV",
		},
	},
	{
		Primary: "PRESIDENT",
		Short:   "PRES",
		Alt: []string{
			"PR", "PRES", "PRESIDENT", "PRS",
		},
	},
	{
		Primary: "PRESS",
		Short:   "PRS",
		Alt: []string{
			"PRESS", "PRS",
		},
	},
	{
		Primary: "PRESSING",
		Short:   "PRSG",
		Alt: []string{
			"PRESSING", "PRSG",
		},
	},
	{
		Primary: "PRESTIGE",
		Short:   "PRSTG",
		Alt: []string{
			"PRESTIGE", "PRSTG",
		},
	},
	{
		Primary: "PREVENTION",
		Short:   "PRVNTN",
		Alt: []string{
			"PREVENTION", "PRVNTN",
		},
	},
	{
		Primary: "PRICE",
		Short:   "PRC",
		Alt: []string{
			"PRC", "PRICE",
		},
	},
	{
		Primary: "PRIDE",
		Short:   "PRD",
		Alt: []string{
			"PRD", "PRIDE",
		},
	},
	{
		Primary: "PRIEST",
		Short:   "PRST",
		Alt: []string{
			"PR", "PRIEST", "PRST",
		},
	},
	{
		Primary: "PRIME",
		Short:   "PRM",
		Alt: []string{
			"PRIME", "PRM",
		},
	},
	{
		Primary: "PRINCE",
		Short:   "PRNC",
		Alt: []string{
			"PR", "PRINCE", "PRNC",
		},
	},
	{
		Primary: "PRINCIPAL",
		Short:   "PRIN",
		Alt: []string{
			"PRIN", "PRINC", "PRINCIPAL", "PRN", "PRNCPL",
		},
	},
	{
		Primary: "PRINT",
		Short:   "PRT",
		Alt: []string{
			"PRINT", "PRT",
		},
	},
	{
		Primary: "PRINTER",
		Short:   "PRINTR",
		Alt: []string{
			"PRINT", "PRINTER", "PRINTR", "PRTR",
		},
	},
	{
		Primary: "PRINTING",
		Short:   "PRINTG",
		Alt: []string{
			"PRINT", "PRINTG", "PRINTING", "PRNTNG", "PRTG",
			"PTG",
		},
	},
	{
		Primary: "PRIVATE",
		Short:   "PVT",
		Alt: []string{
			"PRIVATE", "PVT",
		},
	},
	{
		Primary: "PROCESS",
		Short:   "PRCS",
		Alt: []string{
			"PRCS", "PROCES", "PROCESS",
		},
	},
	{
		Primary: "PROCESSING",
		Short:   "PRCSG",
		Alt: []string{
			"PRCS", "PRCSG", "PRCSNG", "PROC", "PROCESSING",
		},
	},
	{
		Primary: "PROCESSOR",
		Short:   "PRCSR",
		Alt: []string{
			"PRCSR", "PROCESSOR",
		},
	},
	{
		Primary: "PROCUREMENT",
		Short:   "PRCMNT",
		Alt: []string{
			"PRCMNT", "PROCU", "PROCUREMENT",
		},
	},
	{
		Primary: "PRODUCE",
		Short:   "PROD",
		Alt: []string{
			"PROD", "PRODUCE",
		},
	},
	{
		Primary: "PRODUCER",
		Short:   "PRODR",
		Alt: []string{
			"PROD", "PRODR", "PRODUCER",
		},
	},
	{
		Primary: "PRODUCING",
		Short:   "PRDCNG",
		Alt: []string{
			"PRDCNG", "PRODUCING",
		},
	},
	{
		Primary: "PRODUCT",
		Short:   "PRODT",
		Alt: []string{
			"PRO", "PROD", "PRODT", "PRODUCT",
		},
	},
	{
		Primary: "PRODUCTION",
		Short:   "PRODN",
		Alt: []string{
			"PRD", "PRDTN", "PROD", "PRODCTN", "PRODN",
			"PRODT", "PRODUCTION",
		},
	},
	{
		Primary: "PRODUCTIVITY",
		Short:   "PRDCTVTY",
		Alt: []string{
			"PRDCTVTY", "PRODUCTIVITY",
		},
	},
	{
		Primary: "PROFESSIONAL",
		Short:   "PRO",
		Alt: []string{
			"PRO", "PROF", "PROFESSIONAL", "PROFL",
		},
	},
	{
		Primary: "PROFESSOR",
		Short:   "PROF",
		Alt: []string{
			"PROF", "PROFESSOR",
		},
	},
	{
		Primary: "PROGRAM",
		Short:   "PRGM",
		Alt: []string{
			"PRGM", "PROG", "PROGRAM",
		},
	},
	{
		Primary: "PROGRAMMER",
		Short:   "PRGRMR",
		Alt: []string{
			"PRGMR", "PRGRMR", "PROG", "PROGR", "PROGRAMER",
			"PROGRAMMER", "PROGRMMR",
		},
	},
	{
		Primary: "PROGRAMMING",
		Short:   "PRGMNG",
		Alt: []string{
			"PRGMNG", "PROGRAMMING",
		},
	},
	{
		Primary: "PROGRESSIVE",
		Short:   "PROGS",
		Alt: []string{
			"PROG", "PROGRESSIVE", "PROGS",
		},
	},
	{
		Primary: "PROJECT",
		Short:   "PROJ",
		Alt: []string{
			"PRJ", "PROJ", "PROJECT",
		},
	},
	{
		Primary: "PROMOTION",
		Short:   "PROM",
		Alt: []string{
			"PROM", "PROMOTION",
		},
	},
	{
		Primary: "PROPANE",
		Short:   "PROPN",
		Alt: []string{
			"LPG", "PROPANE", "PROPN", "PRPN",
		},
	},
	{
		Primary: "PROPERTY",
		Short:   "PROP",
		Alt: []string{
			"PROP", "PROPERTY", "PRPTY",
		},
	},
	{
		Primary: "PROPRIETARY",
		Short:   "PROPTY",
		Alt: []string{
			"PROPRIETARY", "PROPTY",
		},
	},
	{
		Primary: "PROTECTION",
		Short:   "PROTECT",
		Alt: []string{
			"PROTCTN", "PROTECT", "PROTECTION", "PRTCTN",
		},
	},
	{
		Primary: "PROTECTIVE",
		Short:   "PRTCTV",
		Alt: []string{
			"PROTECTIVE", "PRTCTV",
		},
	},
	{
		Primary: "PROTESTANT",
		Short:   "PRTSTNT",
		Alt: []string{
			"PROTESTANT", "PRTSTNT",
		},
	},
	{
		Primary: "PROVIDENCE",
		Short:   "PRVDNCE",
		Alt:     []string{"PROVIDENCE"},
	},
	{
		Primary: "PRVDNC",
		Short:   "PRVDNC",
		Alt:     []string{"PRVDNC"},
	},
	{
		Primary: "PROVINCE",
		Short:   "PROVNC",
		Alt: []string{
			"PROV", "PROVINCE", "PROVNC",
		},
	},
	{
		Primary: "PROVISION",
		Short:   "PROVSN",
		Alt: []string{
			"PROV", "PROVISION", "PROVSN",
		},
	},
	{
		Primary: "PSYCHIATRIC",
		Short:   "PSYCHC",
		Alt: []string{
			"PSYCH", "PSYCHC", "PSYCHIATRIC",
		},
	},
	{
		Primary: "PSYCHIATRIST",
		Short:   "PSYCH",
		Alt: []string{
			"PSYCH", "PSYCHIATRIST",
		},
	},
	{
		Primary: "PSYCHIATRY",
		Short:   "PSYCHY",
		Alt: []string{
			"PSHYCHY", "PSYCH", "PSYCHIATRY",
		},
	},
	{
		Primary: "PSYCHOLOGICAL",
		Short:   "PSYCHL",
		Alt: []string{
			"PSYCH", "PSYCHL", "PSYCHOLOGICAL",
		},
	},
	{
		Primary: "PSYCHOLOGIST",
		Short:   "PSYC",
		Alt: []string{
			"PSYC", "PSYCHOLOGIST",
		},
	},
	{
		Primary: "PSYCHOLOGY",
		Short:   "PSYCY",
		Alt: []string{
			"PSYC", "PSYCH", "PSYCHOLOGY", "PSYCLGY",
		},
	},
	{
		Primary: "PUBLIC",
		Short:   "PUB",
		Alt: []string{
			"PBLC", "PUB", "PUBLIC",
		},
	},
	{
		Primary: "PUBLICATION",
		Short:   "PUBLCTN",
		Alt: []string{
			"PBLCNTN", "PUBL", "PUBLCTN", "PUBLICATION",
		},
	},
	{
		Primary: "PUBLISHER",
		Short:   "PUBLR",
		Alt: []string{
			"PBLSHR", "PUB", "PUBL", "PUBLISHER", "PUBLR",
			"PUBLSHR",
		},
	},
	{
		Primary: "PUBLISHING",
		Short:   "PBLSHNG",
		Alt: []string{
			"PBLSHNG", "PUB", "PUBG", "PUBLISHING",
		},
	},
	{
		Primary: "PUMPING",
		Short:   "PMPG",
		Alt: []string{
			"PMPG", "PUMPING",
		},
	},
	{
		Primary: "PUNCH",
		Short:   "PNCH",
		Alt: []string{
			"PNCH", "PUNCH",
		},
	},
	{
		Primary: "PURCHASE",
		Short:   "PURCH",
		Alt: []string{
			"PUR", "PURCH", "PURCHASE",
		},
	},
	{
		Primary: "PURCHASER",
		Short:   "PURCHR",
		Alt: []string{
			"PUR", "PURCHASER", "PURCHR",
		},
	},
	{
		Primary: "PURCHASING",
		Short:   "PRCHNG",
		Alt: []string{
			"PRCHNG", "PURCH", "PURCHASING",
		},
	},
	{
		Primary: "QUADRANGLE",
		Short:   "QUAD",
		Alt: []string{
			"QUAD", "QUADRANGLE",
		},
	},
	{
		Primary: "QUALITY",
		Short:   "QLTY",
		Alt: []string{
			"QLTY", "QUAL", "QUALITY", "QULTY",
		},
	},
	{
		Primary: "QUANTITY",
		Short:   "QTY",
		Alt: []string{
			"QTY", "QUANTITY",
		},
	},
	{
		Primary: "QUARRY",
		Short:   "QUAR",
		Alt: []string{
			"QUAR", "QUARRY",
		},
	},
	{
		Primary: "QUARTER",
		Short:   "QTR",
		Alt: []string{
			"QTR", "QUARTER",
		},
	},
	{
		Primary: "QUEEN",
		Short:   "QN",
		Alt: []string{
			"QN", "QUEEN",
		},
	},
	{
		Primary: "QUICK",
		Short:   "QCK",
		Alt: []string{
			"QCK", "QUICK",
		},
	},
	{
		Primary: "RABBI",
		Short:   "RBB",
		Alt: []string{
			"RABBI", "RBB",
		},
	},
	{
		Primary: "RACING",
		Short:   "RACG",
		Alt: []string{
			"RACG", "RACING",
		},
	},
	{
		Primary: "RADIATOR",
		Short:   "RADTR",
		Alt: []string{
			"RAD", "RADIATOR", "RADTR",
		},
	},
	{
		Primary: "RADIO",
		Short:   "RDO",
		Alt: []string{
			"RADIO", "RDO",
		},
	},
	{
		Primary: "RADIOLOGIST",
		Short:   "RAD",
		Alt: []string{
			"RAD", "RADIOLOGIST",
		},
	},
	{
		Primary: "RADIOLOGY",
		Short:   "RADY",
		Alt: []string{
			"RAD", "RADIOLOGY", "RADY",
		},
	},
	{
		Primary: "RAILROAD",
		Short:   "RR",
		Alt: []string{
			"R R", "RAILROAD", "RR",
		},
	},
	{
		Primary: "RAILWAY",
		Short:   "RLWY",
		Alt: []string{
			"RAILWAY", "RLWY",
		},
	},
	{
		Primary: "RAINBOW",
		Short:   "RNBW",
		Alt: []string{
			"RAINBOW", "RNBW",
		},
	},
	{
		Primary: "RANCH",
		Short:   "RNCH",
		Alt: []string{
			"RANCH", "RNCH",
		},
	},
	{
		Primary: "READABLE",
		Short:   "RDBL",
		Alt: []string{
			"RDBL", "READABLE",
		},
	},
	{
		Primary: "READY",
		Short:   "RDY",
		Alt: []string{
			"RDY", "READY",
		},
	},
	{
		Primary: "REALTOR",
		Short:   "RLTR",
		Alt: []string{
			"REALTOR", "RLTR",
		},
	},
	{
		Primary: "REALTY",
		Short:   "RLTY",
		Alt: []string{
			"REALTY", "RLTY",
		},
	},
	{
		Primary: "REBUILDER",
		Short:   "RBLDR",
		Alt: []string{
			"RBLDR", "REBUILDER",
		},
	},
	{
		Primary: "RECEIPT",
		Short:   "RECPT",
		Alt: []string{
			"REC", "RECEIPT", "RECP", "RECPT",
		},
	},
	{
		Primary: "RECEIVABLE",
		Short:   "RCVBL",
		Alt: []string{
			"RCV", "RECEIVABLE",
		},
	},
	{
		Primary: "RECEIVE",
		Short:   "RCV",
		Alt: []string{
			"RCV", "RECEIVE",
		},
	},
	{
		Primary: "RECEIVED",
		Short:   "RCVD",
		Alt: []string{
			"RCVD", "RECEIVED",
		},
	},
	{
		Primary: "RECEIVING",
		Short:   "RCVNG",
		Alt: []string{
			"RCVNG", "RECEIVING",
		},
	},
	{
		Primary: "RECONSTRUCTIVE",
		Short:   "RECNSTRCTV",
		Alt: []string{
			"RECNSTRCTV", "RECONSTRUCTIVE",
		},
	},
	{
		Primary: "RECORD",
		Short:   "REC",
		Alt: []string{
			"REC", "RECORD",
		},
	},
	{
		Primary: "RECOVERY",
		Short:   "RECVY",
		Alt: []string{
			"RECOVERY", "RECVY",
		},
	},
	{
		Primary: "RECREATION",
		Short:   "RCRTN",
		Alt: []string{
			"RCRTN", "REC", "RECREATION",
		},
	},
	{
		Primary: "RECREATIONAL",
		Short:   "RCRTNL",
		Alt: []string{
			"RCRTNL", "RECREATIONAL", "RECRTL",
		},
	},
	{
		Primary: "RECRUITER",
		Short:   "RCRTR",
		Alt: []string{
			"RCRTR", "RECRUITER",
		},
	},
	{
		Primary: "RECRUITING",
		Short:   "RECRUIT",
		Alt: []string{
			"RECRUIT", "RECRUITING",
		},
	},
	{
		Primary: "RECYCLING",
		Short:   "RECYCLE",
		Alt: []string{
			"RCYCLNG", "RECYCLE", "RECYCLING",
		},
	},
	{
		Primary: "REDUCTION",
		Short:   "RDCTN",
		Alt: []string{
			"RDCTN", "REDUCTION",
		},
	},
	{
		Primary: "REFERENCE",
		Short:   "REF",
		Alt: []string{
			"REF", "REFERENCE",
		},
	},
	{
		Primary: "REFINERY",
		Short:   "RFNRY",
		Alt: []string{
			"REFINERY", "RFNRY",
		},
	},
	{
		Primary: "REFINING",
		Short:   "RFNG",
		Alt: []string{
			"REF", "REFINING", "RFNG",
		},
	},
	{
		Primary: "REFRACTORY",
		Short:   "REFR",
		Alt: []string{
			"REFR", "REFRACTORY",
		},
	},
	{
		Primary: "REFRIGERATION",
		Short:   "REFRIG",
		Alt: []string{
			"REFRIG", "REFRIGERATION", "RFRGRTN",
		},
	},
	{
		Primary: "REFRIGERATOR",
		Short:   "RFRG",
		Alt: []string{
			"REFRIGERATOR", "RFRG",
		},
	},
	{
		Primary: "REGION",
		Short:   "REGN",
		Alt: []string{
			"REG", "REGION", "REGN",
		},
	},
	{
		Primary: "REGIONAL",
		Short:   "REGL",
		Alt: []string{
			"REG", "REGIONAL", "REGL", "REGNL",
		},
	},
	{
		Primary: "REGISTER",
		Short:   "REG",
		Alt: []string{
			"REG", "REGISTER", "RGSTR",
		},
	},
	{
		Primary: "REGISTERED",
		Short:   "REGD",
		Alt: []string{
			"REG", "REGD", "REGISTERED",
		},
	},
	{
		Primary: "REGISTRAR",
		Short:   "REGR",
		Alt: []string{
			"REG", "REGISTRAR", "REGR",
		},
	},
	{
		Primary: "REGISTRY",
		Short:   "RGSTY",
		Alt: []string{
			"REGISTRY", "RGSTY",
		},
	},
	{
		Primary: "REGULATORY",
		Short:   "RGLTRY",
		Alt: []string{
			"REGULATORY", "RGLTRY",
		},
	},
	{
		Primary: "REHABILITATION",
		Short:   "REHAB",
		Alt: []string{
			"REHAB", "REHABILITATION",
		},
	},
	{
		Primary: "RELATED",
		Short:   "RLTD",
		Alt: []string{
			"RELATED", "RLTD",
		},
	},
	{
		Primary: "RELATION",
		Short:   "REL",
		Alt: []string{
			"REL", "RELA", "RELATION",
		},
	},
	{
		Primary: "RELIABLE",
		Short:   "RELI",
		Alt: []string{
			"RELI", "RELIABLE",
		},
	},
	{
		Primary: "RELOCATION",
		Short:   "RLCTN",
		Alt: []string{
			"RELOCATION", "RLCTN",
		},
	},
	{
		Primary: "REMEDIAL",
		Short:   "RMDL",
		Alt: []string{
			"REMEDIAL", "RMDL",
		},
	},
	{
		Primary: "REMODELING",
		Short:   "REMOD",
		Alt: []string{
			"REMOD", "REMODELING", "RMDLG",
		},
	},
	{
		Primary: "RENTAL",
		Short:   "RENT",
		Alt: []string{
			"RENT", "RENTAL", "RNT", "RNTL",
		},
	},
	{
		Primary: "REPAIR",
		Short:   "RPR",
		Alt: []string{
			"REPAIR", "REPR", "RPR",
		},
	},
	{
		Primary: "REPORT",
		Short:   "REPT",
		Alt: []string{
			"REP", "REPORT", "REPT",
		},
	},
	{
		Primary: "REPORTER",
		Short:   "REPTR",
		Alt: []string{
			"REP", "REPORTER", "REPTR",
		},
	},
	{
		Primary: "REPRESENTATIVE",
		Short:   "REP",
		Alt: []string{
			"REP", "REPRESENTATIVE",
		},
	},
	{
		Primary: "REPUBLIC",
		Short:   "REPB",
		Alt: []string{
			"REPB", "REPUBLIC",
		},
	},
	{
		Primary: "REPUBLICAN",
		Short:   "REPUB",
		Alt: []string{
			"REPUB", "REPUBLICAN",
		},
	},
	{
		Primary: "REQUIREMENT",
		Short:   "RQRMNT",
		Alt: []string{
			"REQUIREMENT", "RQRMNT",
		},
	},
	{
		Primary: "RESEARCH",
		Short:   "RSRCH",
		Alt: []string{
			"RES", "RESEARCH", "RSCH", "RSRCH",
		},
	},
	{
		Primary: "RESERVE",
		Short:   "RESV",
		Alt: []string{
			"RESERVE", "RESV",
		},
	},
	{
		Primary: "RESIDENCE",
		Short:   "RSDNC",
		Alt: []string{
			"RESIDENCE", "RSDNC",
		},
	},
	{
		Primary: "RESIDENT",
		Short:   "RES",
		Alt: []string{
			"RES", "RESIDENT", "RSDNT",
		},
	},
	{
		Primary: "RESORT",
		Short:   "RESRT",
		Alt: []string{
			"RESORT", "RESRT",
		},
	},
	{
		Primary: "RESOURCE",
		Short:   "RESRC",
		Alt: []string{
			"RES", "RESOURCE", "RESRC", "RSCE", "RSRC",
		},
	},
	{
		Primary: "RESPONSIBLE",
		Short:   "RESP",
		Alt: []string{
			"RESP", "RESPONSIBLE",
		},
	},
	{
		Primary: "RESTAURANT",
		Short:   "RSTRNT",
		Alt: []string{
			"RESTAURANT", "RSTRNT",
		},
	},
	{
		Primary: "RESTORATION",
		Short:   "RESTOR",
		Alt: []string{
			"RESTOR", "RESTORATION", "RSTRTN",
		},
	},
	{
		Primary: "RETAIL",
		Short:   "RTL",
		Alt: []string{
			"RETAIL", "RTL",
		},
	},
	{
		Primary: "RETAILER",
		Short:   "RET",
		Alt: []string{
			"RET", "RETAILER",
		},
	},
	{
		Primary: "RETARDATION",
		Short:   "RTRDTN",
		Alt: []string{
			"RETARDATION", "RTRDTN",
		},
	},
	{
		Primary: "RETIRED",
		Short:   "RTRD",
		Alt: []string{
			"RET", "RETIRED", "RTRD",
		},
	},
	{
		Primary: "RETIREMENT",
		Short:   "RTRMNT",
		Alt: []string{
			"RETIREMENT", "RTRMNT",
		},
	},
	{
		Primary: "RETRAINING",
		Short:   "RETRNG",
		Alt: []string{
			"RETRAINING", "RETRNG",
		},
	},
	{
		Primary: "REVEREND",
		Short:   "REV",
		Alt: []string{
			"REV", "REVEREND",
		},
	},
	{
		Primary: "RIDGE",
		Short:   "RDG",
		Alt: []string{
			"RDG", "RIDGE",
		},
	},
	{
		Primary: "RIVER",
		Short:   "RIV",
		Alt: []string{
			"RIV", "RIVER", "RIVR", "RVR",
		},
	},
	{
		Primary: "ROADWAY",
		Short:   "RDWY",
		Alt: []string{
			"RDWY", "ROADWAY",
		},
	},
	{
		Primary: "ROCKY",
		Short:   "RCKY",
		Alt: []string{
			"RCKY", "ROCKY",
		},
	},
	{
		Primary: "ROOFING",
		Short:   "ROOF",
		Alt: []string{
			"ROOF", "ROOFG", "ROOFING",
		},
	},
	{
		Primary: "ROUND",
		Short:   "RND",
		Alt: []string{
			"RND", "ROUND",
		},
	},
	{
		Primary: "ROUTE",
		Short:   "RT",
		Alt: []string{
			"ROUTE", "RT", "RTE",
		},
	},
	{
		Primary: "ROYAL",
		Short:   "RYL",
		Alt: []string{
			"ROYAL", "RYL",
		},
	},
	{
		Primary: "ROYALTY",
		Short:   "ROY",
		Alt: []string{
			"ROY", "ROYALTY",
		},
	},
	{
		Primary: "RUBBER",
		Short:   "RBR",
		Alt: []string{
			"RBR", "RUBBER",
		},
	},
	{
		Primary: "RURAL",
		Short:   "RUR",
		Alt: []string{
			"RUR", "RURAL",
		},
	},
	{
		Primary: "SADDLERY",
		Short:   "SAD",
		Alt: []string{
			"SAD", "SADDLERY",
		},
	},
	{
		Primary: "SAFETY",
		Short:   "SFTY",
		Alt: []string{
			"SAFETY", "SFTY",
		},
	},
	{
		Primary: "SAINT",
		Short:   "ST",
		Alt: []string{
			"SAINT", "ST",
		},
	},
	{
		Primary: "SALES",
		Short:   "SLS",
		Alt: []string{
			"SALES", "SLS",
		},
	},
	{
		Primary: "SALESMAN",
		Short:   "SLSMN",
		Alt: []string{
			"SALESMAN", "SLSMAN", "SLSMN",
		},
	},
	{
		Primary: "SALON",
		Short:   "SLN",
		Alt: []string{
			"SALON", "SLN",
		},
	},
	{
		Primary: "SALOON",
		Short:   "SLON",
		Alt: []string{
			"SALOON", "SLN", "SLON",
		},
	},
	{
		Primary: "SALVAGE",
		Short:   "SLVG",
		Alt: []string{
			"SALV", "SALVAGE", "SLVG",
		},
	},
	{
		Primary: "SALVATION",
		Short:   "SLVTN",
		Alt: []string{
			"SALVATION", "SLVTN",
		},
	},
	{
		Primary: "SANDWICH",
		Short:   "SNDWCH",
		Alt: []string{
			"SAND", "SANDWICH", "SNDWCH",
		},
	},
	{
		Primary: "SANITARY",
		Short:   "SANI",
		Alt: []string{
			"SANI", "SANITARY",
		},
	},
	{
		Primary: "SANITATION",
		Short:   "SANITN",
		Alt: []string{
			"SANI", "SANITATION", "SANITN",
		},
	},
	{
		Primary: "SATELLITE",
		Short:   "SAT",
		Alt: []string{
			"SAT", "SATELLITE",
		},
	},
	{
		Primary: "SATISFACTION",
		Short:   "STSFCTN",
		Alt: []string{
			"SATISFACTION", "STSFCTN",
		},
	},
	{
		Primary: "SAVINGS",
		Short:   "SVNGS",
		Alt: []string{
			"SAV", "SAVE", "SAVINGS", "SVNGS",
		},
	},
	{
		Primary: "SCHOOL",
		Short:   "SCHL",
		Alt: []string{
			"SCH", "SCHL", "SCHOOL",
		},
	},
	{
		Primary: "SCIENCE",
		Short:   "SCI",
		Alt: []string{
			"SC", "SCI", "SCIENCE",
		},
	},
	{
		Primary: "SCIENTIFIC",
		Short:   "SCNTFC",
		Alt: []string{
			"SCI", "SCIENTIFIC", "SCNTFC",
		},
	},
	{
		Primary: "SCIENTIST",
		Short:   "SCNTST",
		Alt: []string{
			"SCIENTIST", "SCNTST",
		},
	},
	{
		Primary: "SCREEN",
		Short:   "SCRN",
		Alt: []string{
			"SCREEN", "SCRN",
		},
	},
	{
		Primary: "SEAFOOD",
		Short:   "SEAFD",
		Alt: []string{
			"SEAFD", "SEAFOOD",
		},
	},
	{
		Primary: "SEAMAN",
		Short:   "SMN",
		Alt: []string{
			"SEAMAN", "SMN",
		},
	},
	{
		Primary: "SEASON",
		Short:   "SN",
		Alt: []string{
			"SEASON", "SN",
		},
	},
	{
		Primary: "SECOND",
		Short:   "2ND",
		Alt: []string{
			"2", "2ND", "II", "SEC", "SECOND",
		},
	},
	{
		Primary: "SECRETARIAL",
		Short:   "SECL",
		Alt: []string{
			"SEC", "SECL", "SECRETARIAL",
		},
	},
	{
		Primary: "SECRETARY",
		Short:   "SECY",
		Alt: []string{
			"SEC", "SECR", "SECRETARY", "SECT", "SECTY",
			"SECY",
		},
	},
	{
		Primary: "SECTION",
		Short:   "SECT",
		Alt: []string{
			"SCTN", "SECT", "SECTION",
		},
	},
	{
		Primary: "SECTIONAL",
		Short:   "SECTL",
		Alt: []string{
			"SECT", "SECTIONAL", "SECTL",
		},
	},
	{
		Primary: "SECURITY",
		Short:   "SEC",
		Alt: []string{
			"SCRTY", "SEC", "SECURITY",
		},
	},
	{
		Primary: "SEMINARY",
		Short:   "SMRY",
		Alt: []string{
			"SEMINARY", "SMRY",
		},
	},
	{
		Primary: "SENATOR",
		Short:   "SEN",
		Alt: []string{
			"SEN", "SENATOR",
		},
	},
	{
		Primary: "SENIOR",
		Short:   "SR",
		Alt: []string{
			"SENIOR", "SR",
		},
	},
	{
		Primary: "SENSORY",
		Short:   "SNSRY",
		Alt: []string{
			"SENSORY", "SNSRY",
		},
	},
	{
		Primary: "SEPTIC",
		Short:   "SPTC",
		Alt: []string{
			"SEPTIC", "SPTC",
		},
	},
	{
		Primary: "SERGEANT",
		Short:   "SGT",
		Alt: []string{
			"SEGT", "SERGEANT", "SERGNT", "SG", "SGT",
		},
	},
	{
		Primary: "SERIAL",
		Short:   "SER",
		Alt: []string{
			"SER", "SERIAL",
		},
	},
	{
		Primary: "SERVICE",
		Short:   "SVC",
		Alt: []string{
			"SER", "SERV", "SERVIC", "SERVICE", "SRV",
			"SV", "SVC", "SVCE",
		},
	},
	{
		Primary: "SEVENTH",
		Short:   "7TH",
		Alt: []string{
			"7TH", "SEVENTH", "VII",
		},
	},
	{
		Primary: "SEWER",
		Short:   "SWR",
		Alt: []string{
			"SEWER", "SWR",
		},
	},
	{
		Primary: "SEWING",
		Short:   "SEW",
		Alt: []string{
			"SEW", "SEWING",
		},
	},
	{
		Primary: "SHADE",
		Short:   "SHD",
		Alt: []string{
			"SHADE", "SHD",
		},
	},
	{
		Primary: "SHEAR",
		Short:   "SHR",
		Alt: []string{
			"SHEAR", "SHR",
		},
	},
	{
		Primary: "SHEET",
		Short:   "SHT",
		Alt: []string{
			"SHEET", "SHT",
		},
	},
	{
		Primary: "SHELL",
		Short:   "SHL",
		Alt: []string{
			"SHELL", "SHL",
		},
	},
	{
		Primary: "SHERIFF",
		Short:   "SHER",
		Alt: []string{
			"SH", "SHER", "SHERIF", "SHERIFF",
		},
	},
	{
		Primary: "SHIELD",
		Short:   "SHLD",
		Alt: []string{
			"SHIELD", "SHLD",
		},
	},
	{
		Primary: "SHIFT",
		Short:   "SHFT",
		Alt: []string{
			"SHFT", "SHIFT",
		},
	},
	{
		Primary: "SHIPBUILDING",
		Short:   "SHIPBLDG",
		Alt: []string{
			"SHIPBLDG", "SHIPBUILDING",
		},
	},
	{
		Primary: "SHIPPING",
		Short:   "SHIPG",
		Alt: []string{
			"SHIPG", "SHIPPING", "SHPNG",
		},
	},
	{
		Primary: "SHOPPE",
		Short:   "SHP",
		Alt: []string{
			"SHOPPE", "SHP",
		},
	},
	{
		Primary: "SHOPPING",
		Short:   "SHPG",
		Alt: []string{
			"SHOPG", "SHOPPING",
		},
	},
	{
		Primary: "SHORE",
		Short:   "SHOR",
		Alt: []string{
			"SHOR", "SHORE", "SHR",
		},
	},
	{
		Primary: "SHOWCASE",
		Short:   "SHWCS",
		Alt: []string{
			"SHOWCASE", "SHWCS",
		},
	},
	{
		Primary: "SIDING",
		Short:   "SIDE",
		Alt: []string{
			"SIDE", "SIDING",
		},
	},
	{
		Primary: "SILVER",
		Short:   "SLVR",
		Alt: []string{
			"SILVER", "SLVR",
		},
	},
	{
		Primary: "SILVERPLATING",
		Short:   "SILPLTG",
		Alt: []string{
			"SILPLTG", "SILVERPLATING",
		},
	},
	{
		Primary: "SILVERWARE",
		Short:   "SILWR",
		Alt: []string{
			"SILVERWARE", "SILWR",
		},
	},
	{
		Primary: "SISTER",
		Short:   "SIS",
		Alt: []string{
			"SIS", "SISTER", "SR",
		},
	},
	{
		Primary: "SIXTH",
		Short:   "6TH",
		Alt: []string{
			"6TH", "SIXTH", "VI",
		},
	},
	{
		Primary: "SKILL",
		Short:   "SKLL",
		Alt: []string{
			"SKILL", "SKLL",
		},
	},
	{
		Primary: "SMALL",
		Short:   "SM",
		Alt: []string{
			"SM", "SMALL", "SML",
		},
	},
	{
		Primary: "SMELTING",
		Short:   "SMELT",
		Alt: []string{
			"SMELT", "SMELTING",
		},
	},
	{
		Primary: "SOCIAL",
		Short:   "SCL",
		Alt: []string{
			"SCL", "SOC", "SOCIAL",
		},
	},
	{
		Primary: "SOCIETY",
		Short:   "SCTY",
		Alt: []string{
			"SCTY", "SOC", "SOCIETY",
		},
	},
	{
		Primary: "SOFTWARE",
		Short:   "SFTWR",
		Alt: []string{
			"SFTWE", "SFTWR", "SOFT", "SOFTWARE",
		},
	},
	{
		Primary: "SOLAR",
		Short:   "SLR",
		Alt: []string{
			"SLR", "SOLAR",
		},
	},
	{
		Primary: "SOLICITOR",
		Short:   "SOLCR",
		Alt: []string{
			"SOLCR", "SOLICITOR",
		},
	},
	{
		Primary: "SOLID",
		Short:   "SLD",
		Alt: []string{
			"SLD", "SOLID",
		},
	},
	{
		Primary: "SOLUTION",
		Short:   "SLTN",
		Alt: []string{
			"SLTN", "SOLUTION",
		},
	},
	{
		Primary: "SOUND",
		Short:   "SND",
		Alt: []string{
			"SND", "SOUND",
		},
	},
	{
		Primary: "SOURCE",
		Short:   "SRC",
		Alt: []string{
			"SOURCE", "SRC",
		},
	},
	{
		Primary: "SOUTHERN",
		Short:   "STHRN",
		Alt: []string{
			"SOUTHERN", "STHRN",
		},
	},
	{
		Primary: "SOUTHSIDE",
		Short:   "STHSD",
		Alt: []string{
			"SOUTHSIDE", "STHSD",
		},
	},
	{
		Primary: "SOUVENIR",
		Short:   "SUV",
		Alt: []string{
			"SOUVENIR", "SUV",
		},
	},
	{
		Primary: "SPACE",
		Short:   "SP",
		Alt: []string{
			"SP", "SPACE", "SPC",
		},
	},
	{
		Primary: "SPECIAL",
		Short:   "SPEC",
		Alt: []string{
			"SPCL", "SPEC", "SPECIAL",
		},
	},
	{
		Primary: "SPECIALIST",
		Short:   "SPCLST",
		Alt: []string{
			"SPCLST", "SPEC", "SPECIALIST", "SPECIALIT",
		},
	},
	{
		Primary: "SPECIALTY",
		Short:   "SPCLTY",
		Alt: []string{
			"SPC", "SPCLT", "SPCLTY", "SPEC", "SPECIALTY",
		},
	},
	{
		Primary: "SPECIFICATION",
		Short:   "SPCFCTN",
		Alt: []string{
			"SPCFCTN", "SPECIFICATION",
		},
	},
	{
		Primary: "SPECTRUM",
		Short:   "SPECT",
		Alt: []string{
			"SPECT", "SPECTRUM",
		},
	},
	{
		Primary: "SPEED",
		Short:   "SPD",
		Alt: []string{
			"SPD", "SPEED",
		},
	},
	{
		Primary: "SPEEDOMETER",
		Short:   "SPDMTR",
		Alt: []string{
			"SPDMTR", "SPEEDOMETER",
		},
	},
	{
		Primary: "SPEEDY",
		Short:   "SPDY",
		Alt: []string{
			"SPDY", "SPEEDY",
		},
	},
	{
		Primary: "SPONSOR",
		Short:   "SPON",
		Alt: []string{
			"SPONG", "SPONSOR",
		},
	},
	{
		Primary: "SPONSORING",
		Short:   "SPONG",
		Alt:     []string{"SPONSORING"},
	},
	{
		Primary: "SPORT",
		Short:   "SPRT",
		Alt: []string{
			"SPORT", "SPRT", "SPT",
		},
	},
	{
		Primary: "SPORTING",
		Short:   "SPORT",
		Alt: []string{
			"SPORT", "SPORTING", "SPRTG", "SPTG",
		},
	},
	{
		Primary: "SPORTSWEAR",
		Short:   "SPORTSWR",
		Alt: []string{
			"SPORTSWEAR", "SPORTSWR",
		},
	},
	{
		Primary: "SPRING",
		Short:   "SPG",
		Alt: []string{
			"SPG", "SPNG", "SPRING", "SPRNG",
		},
	},
	{
		Primary: "SPRINKLER",
		Short:   "SPRINK",
		Alt: []string{
			"SPRINK", "SPRINKLER",
		},
	},
	{
		Primary: "SQUARE",
		Short:   "SQ",
		Alt: []string{
			"SQ", "SQUARE",
		},
	},
	{
		Primary: "STABLE",
		Short:   "STBL",
		Alt: []string{
			"STABLE", "STBL",
		},
	},
	{
		Primary: "STAFF",
		Short:   "STAF",
		Alt: []string{
			"STAF", "STAFF",
		},
	},
	{
		Primary: "STAINLESS",
		Short:   "STNLS",
		Alt: []string{
			"STAINLESS", "STNLS",
		},
	},
	{
		Primary: "STAMP",
		Short:   "STMP",
		Alt: []string{
			"STAMP", "STMP",
		},
	},
	{
		Primary: "STAMPING",
		Short:   "STAMPG",
		Alt: []string{
			"STAMPG", "STAMPING",
		},
	},
	{
		Primary: "STANDARD",
		Short:   "STAND",
		Alt: []string{
			"STAND", "STANDARD", "STD",
		},
	},
	{
		Primary: "START",
		Short:   "STRT",
		Alt: []string{
			"START", "STRT",
		},
	},
	{
		Primary: "STATE",
		Short:   "STAT",
		Alt: []string{
			"ST", "STAT", "STATE",
		},
	},
	{
		Primary: "STATION",
		Short:   "STA",
		Alt: []string{
			"STA", "STATION", "STATN", "STN",
		},
	},
	{
		Primary: "STATIONER",
		Short:   "STATNR",
		Alt: []string{
			"STATIONER", "STATNR",
		},
	},
	{
		Primary: "STATIONARY",
		Short:   "STATNRY",
		Alt: []string{
			"STATIONARY", "STATNRY", "STY",
		},
	},
	{
		Primary: "STEAK",
		Short:   "STK",
		Alt: []string{
			"STEAK", "STK",
		},
	},
	{
		Primary: "STEAM",
		Short:   "STM",
		Alt: []string{
			"STEAM", "STM",
		},
	},
	{
		Primary: "STEEL",
		Short:   "STL",
		Alt: []string{
			"STEEL", "STL",
		},
	},
	{
		Primary: "STEREO",
		Short:   "STER",
		Alt: []string{
			"STER", "STEREO", "STR",
		},
	},
	{
		Primary: "STERLING",
		Short:   "STRLNG",
		Alt: []string{
			"STERLING", "STRLNG",
		},
	},
	{
		Primary: "STOCK",
		Short:   "STCK",
		Alt: []string{
			"STCK", "STOCK",
		},
	},
	{
		Primary: "STOCKHOLDER",
		Short:   "STCKHLDR",
		Alt: []string{
			"STCKHLDR", "STOCKHOLDER",
		},
	},
	{
		Primary: "STOCKYARD",
		Short:   "STKYD",
		Alt: []string{
			"STKYD", "STOCKYARD",
		},
	},
	{
		Primary: "STONE",
		Short:   "STN",
		Alt: []string{
			"STN", "STONE",
		},
	},
	{
		Primary: "STORAGE",
		Short:   "STGE",
		Alt: []string{
			"STGE", "STOR", "STORAGE", "STRGE",
		},
	},
	{
		Primary: "STORE",
		Short:   "STR",
		Alt: []string{
			"STORE", "STR",
		},
	},
	{
		Primary: "STOREKEEPER",
		Short:   "STRKP",
		Alt: []string{
			"STOREKEEPER", "STRKP",
		},
	},
	{
		Primary: "STRATEGIC",
		Short:   "STRTGC",
		Alt: []string{
			"STRATEGIC", "STRTGC",
		},
	},
	{
		Primary: "STREET",
		Short:   "STRET",
		Alt: []string{
			"ST", "STREET", "STRET", "STRT",
		},
	},
	{
		Primary: "STRUCTURAL",
		Short:   "STRL",
		Alt: []string{
			"STRL", "STRUCTURAL",
		},
	},
	{
		Primary: "STRUCTURED",
		Short:   "STRCTRD",
		Alt: []string{
			"STRCTRD", "STRUCTURED",
		},
	},
	{
		Primary: "STUDENT",
		Short:   "STDNT",
		Alt: []string{
			"STDNT", "STU", "STUDENT",
		},
	},
	{
		Primary: "STUDIO",
		Short:   "STD",
		Alt: []string{
			"STD", "STUDIO",
		},
	},
	{
		Primary: "STUDY",
		Short:   "STUD",
		Alt: []string{
			"STUD", "STUDY",
		},
	},
	{
		Primary: "STUFF",
		Short:   "STFF",
		Alt: []string{
			"STFF", "STUFF",
		},
	},
	{
		Primary: "STYLE",
		Short:   "STYL",
		Alt: []string{
			"STYL", "STYLE",
		},
	},
	{
		Primary: "STYLING",
		Short:   "STYLG",
		Alt: []string{
			"STYL", "STYLG", "STYLING",
		},
	},
	{
		Primary: "STYLIST",
		Short:   "STYLST",
		Alt: []string{
			"STYL", "STYLIST", "STYLST",
		},
	},
	{
		Primary: "SUBSCRIPTION",
		Short:   "SUBSCR",
		Alt: []string{
			"SUB", "SUBSC", "SUBSCR", "SUBSCRIPTION", "SUBSCRON",
		},
	},
	{
		Primary: "SUBSIDIARY",
		Short:   "SUBY",
		Alt: []string{
			"SUB", "SUBSIDIARY", "SUBY",
		},
	},
	{
		Primary: "SUBSTANCE",
		Short:   "SBSTNC",
		Alt: []string{
			"SBSTNC", "SUBSTANCE",
		},
	},
	{
		Primary: "SUBSTITUTE",
		Short:   "SUB",
		Alt: []string{
			"SUB", "SUBSTITUTE",
		},
	},
	{
		Primary: "SUBURBAN",
		Short:   "SUBN",
		Alt: []string{
			"SUB", "SUBN", "SUBURBAN",
		},
	},
	{
		Primary: "SUBWAY",
		Short:   "SBWY",
		Alt: []string{
			"SBWY", "SUBWAY",
		},
	},
	{
		Primary: "SUGAR",
		Short:   "SUG",
		Alt: []string{
			"SUG", "SUGAR",
		},
	},
	{
		Primary: "SUITE",
		Short:   "STE",
		Alt: []string{
			"STE", "SUITE",
		},
	},
	{
		Primary: "SUMMIT",
		Short:   "SMT",
		Alt: []string{
			"SMT", "SUMMIT",
		},
	},
	{
		Primary: "SUNDRY",
		Short:   "SNDRY",
		Alt: []string{
			"SND", "SNDRY", "SUNDRY",
		},
	},
	{
		Primary: "SUNRISE",
		Short:   "SNRS",
		Alt: []string{
			"SNRS", "SUNRISE",
		},
	},
	{
		Primary: "SUNSET",
		Short:   "SNST",
		Alt: []string{
			"SNST", "SUNSET",
		},
	},
	{
		Primary: "SUNSHINE",
		Short:   "SNSHN",
		Alt: []string{
			"SNSHN", "SUNSHINE",
		},
	},
	{
		Primary: "SUPER",
		Short:   "SPR",
		Alt: []string{
			"SPR", "SUPER",
		},
	},
	{
		Primary: "SUPERINTENDENT",
		Short:   "SUPT",
		Alt: []string{
			"SUPERINTENDENT", "SUPT",
		},
	},
	{
		Primary: "SUPERIOR",
		Short:   "SUPER",
		Alt: []string{
			"SPR", "SUP", "SUPER", "SUPERIOR",
		},
	},
	{
		Primary: "SUPERMARKET",
		Short:   "SPRMRKT",
		Alt: []string{
			"SPRMKT", "SPRMRKT", "SUPERMARKET",
		},
	},
	{
		Primary: "SUPERVISING",
		Short:   "SUPVG",
		Alt: []string{
			"SPVNG", "SUPERVISING", "SUPVG",
		},
	},
	{
		Primary: "SUPERVISION",
		Short:   "SUPRVSN",
		Alt: []string{
			"SUPERVISION", "SUPRVSN",
		},
	},
	{
		Primary: "SUPERVISOR",
		Short:   "SUPVSR",
		Alt: []string{
			"SPV", "SPVR", "SPVSR", "SUPER", "SUPERVISOR",
			"SUPV", "SUPVR", "SUPVSR",
		},
	},
	{
		Primary: "SUPERVISORY",
		Short:   "SUPVRY",
		Alt: []string{
			"SUPERVISORY", "SUPVRY",
		},
	},
	{
		Primary: "SUPPLY",
		Short:   "SUPL",
		Alt: []string{
			"SPLY", "SUP", "SUPL", "SUPLY", "SUPPLY",
		},
	},
	{
		Primary: "SUPPORT",
		Short:   "SPPRT",
		Alt: []string{
			"SPPRT", "SPRT", "SUPPORT",
		},
	},
	{
		Primary: "SUPREME",
		Short:   "SPRM",
		Alt: []string{
			"SPRM", "SUPREME",
		},
	},
	{
		Primary: "SURFACE",
		Short:   "SURFC",
		Alt: []string{
			"SRFC", "SURFACE", "SURFC",
		},
	},
	{
		Primary: "SURGEON",
		Short:   "SRGN",
		Alt: []string{
			"SRGN", "SURGEON",
		},
	},
	{
		Primary: "SURGERY",
		Short:   "SURG",
		Alt: []string{
			"SRGRY", "SURG", "SURGERY", "SURGY",
		},
	},
	{
		Primary: "SURGICAL",
		Short:   "SURGCL",
		Alt: []string{
			"SURGCL", "SURGICAL",
		},
	},
	{
		Primary: "SURPLUS",
		Short:   "SURPL",
		Alt: []string{
			"SRPLS", "SURPL", "SURPLUS",
		},
	},
	{
		Primary: "SURVEY",
		Short:   "SRVY",
		Alt: []string{
			"SRVY", "SURVEY",
		},
	},
	{
		Primary: "SURVEYOR",
		Short:   "SURVYR",
		Alt: []string{
			"SURVEYOR", "SURVYR",
		},
	},
	{
		Primary: "SUSPENSION",
		Short:   "SUSPNSN",
		Alt: []string{
			"SUSPENSION", "SUSPNSN",
		},
	},
	{
		Primary: "SWEEP",
		Short:   "SWP",
		Alt: []string{
			"SWEEP", "SWP",
		},
	},
	{
		Primary: "SWEET",
		Short:   "SWT",
		Alt: []string{
			"SWEET", "SWT",
		},
	},
	{
		Primary: "SYNDICATE",
		Short:   "SYND",
		Alt: []string{
			"SINDICATE", "SYNDICATE",
		},
	},
	{
		Primary: "SYNTHETIC",
		Short:   "SYNT",
		Alt: []string{
			"SYNT", "SYNTHETIC",
		},
	},
	{
		Primary: "SYSTEM",
		Short:   "SYST",
		Alt: []string{
			"SYS", "SYST", "SYSTEM",
		},
	},
	{
		Primary: "TABLE",
		Short:   "TBL",
		Alt: []string{
			"TABLE", "TBL",
		},
	},
	{
		Primary: "TACKLE",
		Short:   "TCKL",
		Alt: []string{
			"TACKLE", "TCKL",
		},
	},
	{
		Primary: "TAILOR",
		Short:   "TLR",
		Alt: []string{
			"TAILOR", "TLR",
		},
	},
	{
		Primary: "TAILORING",
		Short:   "TLRG",
		Alt: []string{
			"TAILORING", "TLRG",
		},
	},
	{
		Primary: "TANNING",
		Short:   "TAN",
		Alt: []string{
			"TAN", "TANNING",
		},
	},
	{
		Primary: "TAVERN",
		Short:   "TRVN",
		Alt: []string{
			"TAV", "TAVERN", "TRVN",
		},
	},
	{
		Primary: "TAXIDERMY",
		Short:   "TXDRMY",
		Alt: []string{
			"TAXIDERMY", "TXDRMY",
		},
	},
	{
		Primary: "TEACHER",
		Short:   "TEACH",
		Alt: []string{
			"TEACH", "TEACHER",
		},
	},
	{
		Primary: "TECHNICAL",
		Short:   "TECHL",
		Alt: []string{
			"TECH", "TECHL", "TECHNICAL",
		},
	},
	{
		Primary: "TECHNICIAN",
		Short:   "TECHN",
		Alt: []string{
			"TECH", "TECHN", "TECHNICIAN",
		},
	},
	{
		Primary: "TECHNOLOGICAL",
		Short:   "TCHNLGCL",
		Alt: []string{
			"TCHNLGCL", "TECHNOLOGICAL",
		},
	},
	{
		Primary: "TECHNOLOGIST",
		Short:   "TECH",
		Alt: []string{
			"TECH", "TECHNOLOGIST",
		},
	},
	{
		Primary: "TECHNOLOGY",
		Short:   "TECHLGY",
		Alt: []string{
			"TCHNLGY", "TECH", "TECHLGY", "TECHNOL", "TECHNOLOGY",
		},
	},
	{
		Primary: "TELECOMMUNICATION",
		Short:   "TELECOM",
		Alt: []string{
			"TELCOMMN", "TELECOM", "TELECOMM", "TELECOMMUNICATION",
		},
	},
	{
		Primary: "TELEGRAPH",
		Short:   "TELG",
		Alt: []string{
			"TELEGRAPH", "TELG",
		},
	},
	{
		Primary: "TELEMARKETING",
		Short:   "TELMKTG",
		Alt: []string{
			"TELEMARKETING", "TELMKTG",
		},
	},
	{
		Primary: "TELEPHONE",
		Short:   "TEL",
		Alt: []string{
			"PHONE", "TELE", "TELEPHONE",
		},
	},
	{
		Primary: "TELETYPE",
		Short:   "TLTYP",
		Alt: []string{
			"TELETYPE", "TLTYP",
		},
	},
	{
		Primary: "TELEVISION",
		Short:   "TV",
		Alt: []string{
			"T V", "TELEVISION",
		},
	},
	{
		Primary: "TELEX",
		Short:   "TLX",
		Alt: []string{
			"TELEX", "TLX",
		},
	},
	{
		Primary: "TEMPERATURE",
		Short:   "TEMP",
		Alt: []string{
			"TEMP", "TEMPERATURE",
		},
	},
	{
		Primary: "TEMPLE",
		Short:   "TMPL",
		Alt: []string{
			"TEMPLE", "TMPL",
		},
	},
	{
		Primary: "TEMPORARY",
		Short:   "TEMPY",
		Alt: []string{
			"TEMP", "TEMPORARY", "TEMPY",
		},
	},
	{
		Primary: "TENNIS",
		Short:   "TEN",
		Alt: []string{
			"TEN", "TENNIS",
		},
	},
	{
		Primary: "TENTH",
		Short:   "10TH",
		Alt: []string{
			"10TH", "TENTH", "X",
		},
	},
	{
		Primary: "TERMINAL",
		Short:   "TRMNL",
		Alt: []string{
			"TERMINAL", "TRML", "TRMNL",
		},
	},
	{
		Primary: "TERMITE",
		Short:   "TRMT",
		Alt: []string{
			"TERMITE", "TRMT",
		},
	},
	{
		Primary: "TERRACE",
		Short:   "TER",
		Alt: []string{
			"TER", "TERR", "TERRACE",
		},
	},
	{
		Primary: "TESTING",
		Short:   "TEST",
		Alt: []string{
			"TEST", "TESTING", "TSTG",
		},
	},
	{
		Primary: "TEXTILE",
		Short:   "TXTL",
		Alt: []string{
			"TEX", "TEXTILE", "TXTL",
		},
	},
	{
		Primary: "THEATRE",
		Short:   "THTR",
		Alt: []string{
			"THEATRE", "THTR",
		},
	},
	{
		Primary: "THEATRICAL",
		Short:   "THEA",
		Alt: []string{
			"THEA", "THEATRICAL", "THTRCL",
		},
	},
	{
		Primary: "THERAPIST",
		Short:   "THRPST",
		Alt: []string{
			"THERAPIST", "THRPST",
		},
	},
	{
		Primary: "THERAPY",
		Short:   "THRPY",
		Alt: []string{
			"THERAPY", "THRPY",
		},
	},
	{
		Primary: "THING",
		Short:   "THNG",
		Alt: []string{
			"THING", "THNG",
		},
	},
	{
		Primary: "THIRD",
		Short:   "3RD",
		Alt: []string{
			"3", "3RD", "III", "THIRD",
		},
	},
	{
		Primary: "THREAD",
		Short:   "THD",
		Alt: []string{
			"THD", "THREAD",
		},
	},
	{
		Primary: "THRIFT",
		Short:   "THRFT",
		Alt: []string{
			"THRFT", "THRIFT",
		},
	},
	{
		Primary: "THRIFTY",
		Short:   "THRFTY",
		Alt: []string{
			"THRFT", "THRFTY", "THRIFTY",
		},
	},
	{
		Primary: "THRUWAY",
		Short:   "THRWY",
		Alt: []string{
			"THRUWAY", "THRWY",
		},
	},
	{
		Primary: "TIMBER",
		Short:   "TMBR",
		Alt: []string{
			"TIMBER", "TMBR",
		},
	},
	{
		Primary: "TITLE",
		Short:   "TITL",
		Alt: []string{
			"TITL", "TITLE", "TTL",
		},
	},
	{
		Primary: "TOBACCO",
		Short:   "TOB",
		Alt: []string{
			"TOB", "TOBACCO",
		},
	},
	{
		Primary: "TOILET",
		Short:   "TOIL",
		Alt: []string{
			"TOIL", "TOILET",
		},
	},
	{
		Primary: "TOTAL",
		Short:   "TTL",
		Alt: []string{
			"TOTAL", "TTL",
		},
	},
	{
		Primary: "TOUCH",
		Short:   "TCH",
		Alt: []string{
			"TCH", "TOUCH",
		},
	},
	{
		Primary: "TOWER",
		Short:   "TWR",
		Alt: []string{
			"TOWER", "TWR",
		},
	},
	{
		Primary: "TOWING",
		Short:   "TOW",
		Alt: []string{
			"TOW", "TOWING",
		},
	},
	{
		Primary: "TOWN",
		Short:   "TWN",
		Alt: []string{
			"TOWN", "TWN",
		},
	},
	{
		Primary: "TOWNE",
		Short:   "TWNE",
		Alt: []string{
			"TOWNE", "TWN", "TWNE",
		},
	},
	{
		Primary: "TOWNSHIP",
		Short:   "TWP",
		Alt: []string{
			"TOWNSHIP", "TWNSHP", "TWP",
		},
	},
	{
		Primary: "TRACTOR",
		Short:   "TRCTR",
		Alt: []string{
			"TRACTOR", "TRCTR",
		},
	},
	{
		Primary: "TRADE",
		Short:   "TRD",
		Alt: []string{
			"TRADE", "TRD",
		},
	},
	{
		Primary: "TRADESMAN",
		Short:   "TRDSMN",
		Alt: []string{
			"TRADESMAN", "TRDSMN",
		},
	},
	{
		Primary: "TRADING",
		Short:   "TRADE",
		Alt: []string{
			"TRADE", "TRADING", "TRDG",
		},
	},
	{
		Primary: "TRAFFIC",
		Short:   "TRFC",
		Alt: []string{
			"TRAFFIC", "TRFC",
		},
	},
	{
		Primary: "TRAIL",
		Short:   "TRL",
		Alt: []string{
			"TRAIL", "TRL",
		},
	},
	{
		Primary: "TRAILER",
		Short:   "TRLR",
		Alt: []string{
			"TRAILER", "TRLR",
		},
	},
	{
		Primary: "TRAINEE",
		Short:   "TRN",
		Alt: []string{
			"TRAINEE", "TRN",
		},
	},
	{
		Primary: "TRAINER",
		Short:   "TRNR",
		Alt: []string{
			"TRAINER", "TRNR",
		},
	},
	{
		Primary: "TRAINING",
		Short:   "TRAIN",
		Alt: []string{
			"TRAIN", "TRAINING", "TRNG",
		},
	},
	{
		Primary: "TRANSFER",
		Short:   "TRNSFR",
		Alt: []string{
			"TRANSF", "TRANSFER", "TRNSFR",
		},
	},
	{
		Primary: "TRANSFORMER",
		Short:   "TRANSFRMR",
		Alt: []string{
			"TRANS", "TRANSFORMER", "TRANSFRMR",
		},
	},
	{
		Primary: "TRANSIT",
		Short:   "TRAN",
		Alt: []string{
			"TRAN", "TRANSIT", "TRNST",
		},
	},
	{
		Primary: "TRANSMISSION",
		Short:   "TRANS",
		Alt: []string{
			"TRANS", "TRANSM", "TRANSMISSION", "TRANSMSSN",
		},
	},
	{
		Primary: "TRANSPORT",
		Short:   "TRNSPRT",
		Alt: []string{
			"TRANS", "TRANSPORT", "TRNSPRT", "TRNSPT",
		},
	},
	{
		Primary: "TRANSPORTATION",
		Short:   "TRNSPRTN",
		Alt: []string{
			"TRANSP", "TRANSPORTATION", "TRNSP", "TRNSPRTN", "TRNSPTN",
		},
	},
	{
		Primary: "TRAVEL",
		Short:   "TRVL",
		Alt: []string{
			"TRAVEL", "TRVL",
		},
	},
	{
		Primary: "TREASURE",
		Short:   "TREAS",
		Alt: []string{
			"TREAS", "TREASURE",
		},
	},
	{
		Primary: "TREASURER",
		Short:   "TRES",
		Alt: []string{
			"TR", "TREA", "TREAS", "TREASURER", "TRES",
			"TRS",
		},
	},
	{
		Primary: "TREASURY",
		Short:   "TRSRY",
		Alt: []string{
			"TREASURY", "TRSRY",
		},
	},
	{
		Primary: "TREATMENT",
		Short:   "TRTMNT",
		Alt: []string{
			"TREATMENT", "TRTMNT",
		},
	},
	{
		Primary: "TRIANGLE",
		Short:   "TRI",
		Alt: []string{
			"TRI", "TRIANGLE",
		},
	},
	{
		Primary: "TRINITY",
		Short:   "TRNTY",
		Alt: []string{
			"TRINITY", "TRNTY",
		},
	},
	{
		Primary: "TRIPLE",
		Short:   "TRPL",
		Alt: []string{
			"TRIPLE", "TRPL",
		},
	},
	{
		Primary: "TROOPER",
		Short:   "TRPR",
		Alt: []string{
			"TROOPER", "TRPR",
		},
	},
	{
		Primary: "TROPHY",
		Short:   "TROPH",
		Alt: []string{
			"TROPH", "TROPHY",
		},
	},
	{
		Primary: "TROPICAL",
		Short:   "TRPCL",
		Alt: []string{
			"TROPICAL", "TRPCL",
		},
	},
	{
		Primary: "TRUCK",
		Short:   "TRCK",
		Alt: []string{
			"TRCK", "TRUCK",
		},
	},
	{
		Primary: "TRUCKING",
		Short:   "TRCKNG",
		Alt: []string{
			"TRCKG", "TRCKNG", "TRUCKING",
		},
	},
	{
		Primary: "TRUST",
		Short:   "TRST",
		Alt: []string{
			"TR", "TRST", "TRUST",
		},
	},
	{
		Primary: "TRUSTEE",
		Short:   "TR",
		Alt: []string{
			"TR", "TRSTE", "TRUSTEE",
		},
	},
	{
		Primary: "TURNPIKE",
		Short:   "TPKE",
		Alt: []string{
			"TPK", "TPKE", "TURNPIKE",
		},
	},
	{
		Primary: "TYPESETTING",
		Short:   "TYPSG",
		Alt: []string{
			"TYPESETTING", "TYPSG",
		},
	},
	{
		Primary: "TYPEWRITER",
		Short:   "TYPWRTR",
		Alt: []string{
			"TYPEWRITER", "TYPTR", "TYPWRTR",
		},
	},
	{
		Primary: "UNDERGRADUATE",
		Short:   "UNDGRAD",
		Alt: []string{
			"UNDERGRADUATE", "UNDGRAD",
		},
	},
	{
		Primary: "UNDERGROUND",
		Short:   "UNDGRD",
		Alt: []string{
			"UNDERGROUND", "UNDGRD",
		},
	},
	{
		Primary: "UNDERWEAR",
		Short:   "UNDWR",
		Alt: []string{
			"UNDERWEAR", "UNDWR",
		},
	},
	{
		Primary: "UNDERWRITER",
		Short:   "UNDERWRTR",
		Alt: []string{
			"UNDERWRITER", "UNDERWRTR", "UNDRWRTR",
		},
	},
	{
		Primary: "UNDERWRITING",
		Short:   "UNDERWRTNG",
		Alt: []string{
			"UNDERWRITING", "UNDERWRTNG",
		},
	},
	{
		Primary: "UNIFORM",
		Short:   "UNFRM",
		Alt: []string{
			"UNF", "UNFRM", "UNIF", "UNIFORM",
		},
	},
	{
		Primary: "UNION",
		Short:   "UN",
		Alt: []string{
			"UN", "UNION",
		},
	},
	{
		Primary: "UNIQUE",
		Short:   "UNQ",
		Alt: []string{
			"UNIQUE", "UNQ",
		},
	},
	{
		Primary: "UNITED",
		Short:   "UNTD",
		Alt: []string{
			"UNITED", "UNTD",
		},
	},
	{
		Primary: "UNITED STATES",
		Short:   "US",
		Alt: []string{
			"UNITED STATES", "US",
		},
	},
	{
		Primary: "UNITED STATES OF AMERICA",
		Short:   "USA",
		Alt:     []string{"UNITED STATES OF AMERICA"},
	},
	{
		Primary: "UNIVERSAL",
		Short:   "UNIVRSL",
		Alt: []string{
			"UNIV", "UNIVERSAL", "UNIVRSL",
		},
	},
	{
		Primary: "UNIVERSITY",
		Short:   "UNIV",
		Alt: []string{
			"UNIV", "UNIVERSITY",
		},
	},
	{
		Primary: "UNLIMITED",
		Short:   "UNLTD",
		Alt: []string{
			"UNLIMITED", "UNLTD",
		},
	},
	{
		Primary: "UPHOLSTERER",
		Short:   "UPHLR",
		Alt: []string{
			"UPHLR", "UPHOLSTERER",
		},
	},
	{
		Primary: "UPHOLSTERING",
		Short:   "UPHLSTRNG",
		Alt: []string{
			"UPHLSTR", "UPHLSTRNG", "UPHOL", "UPHOLSTERING",
		},
	},
	{
		Primary: "UPHOLSTERY",
		Short:   "UPHLSTRY",
		Alt: []string{
			"UPHL", "UPHLSTRY", "UPHOL", "UPHOLSTERY",
		},
	},
	{
		Primary: "URANIUM",
		Short:   "URNM",
		Alt: []string{
			"URANIUM", "URNM",
		},
	},
	{
		Primary: "UROLOGY",
		Short:   "URO",
		Alt: []string{
			"URO", "UROLOGY",
		},
	},
	{
		Primary: "UTILITY",
		Short:   "UTLTY",
		Alt: []string{
			"UTILITY", "UTLTY",
		},
	},
	{
		Primary: "UTILIZATION",
		Short:   "UTLZTN",
		Alt: []string{
			"UTILIZATION", "UTLZTN",
		},
	},
	{
		Primary: "VACUUM",
		Short:   "VCM",
		Alt: []string{
			"VAC", "VACUUM", "VCM",
		},
	},
	{
		Primary: "VALLEY",
		Short:   "VLY",
		Alt: []string{
			"VALLEY", "VALLY", "VLLY", "VLY",
		},
	},
	{
		Primary: "VALUE",
		Short:   "VAL",
		Alt: []string{
			"VAL", "VALUE",
		},
	},
	{
		Primary: "VARIETY",
		Short:   "VRTY",
		Alt: []string{
			"VAR", "VARIETY", "VRTY",
		},
	},
	{
		Primary: "VAULT",
		Short:   "VLT",
		Alt: []string{
			"VAULT", "VLT",
		},
	},
	{
		Primary: "VEGETABLE",
		Short:   "VEG",
		Alt: []string{
			"VEG", "VEGETABLE",
		},
	},
	{
		Primary: "VEHICLE",
		Short:   "VEHIC",
		Alt: []string{
			"VEHIC", "VEHICLE", "VEHK",
		},
	},
	{
		Primary: "VENDING",
		Short:   "VNDNG",
		Alt: []string{
			"VEND", "VENDING", "VNDNG",
		},
	},
	{
		Primary: "VENTILATING",
		Short:   "VENT",
		Alt: []string{
			"VENT", "VENTILATING",
		},
	},
	{
		Primary: "VETERAN",
		Short:   "VETRN",
		Alt: []string{
			"VET", "VETERAN", "VETRN",
		},
	},
	{
		Primary: "VETERINARIAN",
		Short:   "VET",
		Alt: []string{
			"VET", "VETERINARIAN", "VETRN",
		},
	},
	{
		Primary: "VETERINARY",
		Short:   "VETRNRY",
		Alt: []string{
			"VET", "VETERINARY", "VETRNRY",
		},
	},
	{
		Primary: "VIADUCT",
		Short:   "VIA",
		Alt: []string{
			"VIA", "VIADUCT",
		},
	},
	{
		Primary: "VICE",
		Short:   "V",
		Alt: []string{
			"V", "VICE",
		},
	},
	{
		Primary: "VICTORY",
		Short:   "VCTRY",
		Alt: []string{
			"VCTRY", "VICTORY",
		},
	},
	{
		Primary: "VIDEO",
		Short:   "VID",
		Alt: []string{
			"VID", "VIDEO",
		},
	},
	{
		Primary: "VIKING",
		Short:   "VKG",
		Alt: []string{
			"VIKING", "VKG",
		},
	},
	{
		Primary: "VILLAGE",
		Short:   "VLG",
		Alt: []string{
			"VILLAGE", "VLG",
		},
	},
	{
		Primary: "VISION",
		Short:   "VSN",
		Alt: []string{
			"VISION", "VSN",
		},
	},
	{
		Primary: "VISITING",
		Short:   "VSTNG",
		Alt: []string{
			"VISITING", "VSTNG",
		},
	},
	{
		Primary: "VISITOR",
		Short:   "VSTR",
		Alt: []string{
			"VISITOR", "VSTR",
		},
	},
	{
		Primary: "VISTA",
		Short:   "VIS",
		Alt: []string{
			"VIS", "VISTA",
		},
	},
	{
		Primary: "VISUAL",
		Short:   "VISL",
		Alt: []string{
			"VIS", "VISL", "VISUAL",
		},
	},
	{
		Primary: "VOCATION",
		Short:   "VOCN",
		Alt: []string{
			"VOCATION", "VOCN",
		},
	},
	{
		Primary: "VOCATIONAL",
		Short:   "VOCNL",
		Alt: []string{
			"VOCATIONAL", "VOCNL",
		},
	},
	{
		Primary: "VOLUME",
		Short:   "VOL",
		Alt: []string{
			"VOL", "VOLUME",
		},
	},
	{
		Primary: "VOLUNTARY",
		Short:   "VOLNTRY",
		Alt: []string{
			"VOL", "VOLNTRY", "VOLUNTARY",
		},
	},
	{
		Primary: "VOLUNTEER",
		Short:   "VOLNTR",
		Alt:     []string{"VOLUNTEER"},
	},
	{
		Primary: "VULCANIZATION",
		Short:   "VULCN",
		Alt: []string{
			"VULCANIZATION", "VULCN",
		},
	},
	{
		Primary: "VUCANIZING",
		Short:   "VULC",
		Alt: []string{
			"VUCANIZING", "VULC",
		},
	},
	{
		Primary: "WALKWAY",
		Short:   "WLKWY",
		Alt: []string{
			"WALKWAY", "WLKWY",
		},
	},
	{
		Primary: "WALLPAPER",
		Short:   "WLPAPER",
		Alt: []string{
			"PAPER", "WALLPAPER", "WLPAPER", "WLPR",
		},
	},
	{
		Primary: "WARDEN",
		Short:   "WRDN",
		Alt: []string{
			"WARDEN", "WRDN",
		},
	},
	{
		Primary: "WAREHOUSE",
		Short:   "WRHSE",
		Alt: []string{
			"WAREHOUSE", "WHSE", "WRHSE",
		},
	},
	{
		Primary: "WAREHOUSING",
		Short:   "WHSNG",
		Alt: []string{
			"WAREHOUSING", "WHSNG",
		},
	},
	{
		Primary: "WARRANT",
		Short:   "WRRNT",
		Alt: []string{
			"WARRANT", "WRRNT",
		},
	},
	{
		Primary: "WASHING",
		Short:   "WSHG",
		Alt: []string{
			"WASHING", "WSHG",
		},
	},
	{
		Primary: "WASTE",
		Short:   "WST",
		Alt: []string{
			"WASTE", "WST",
		},
	},
	{
		Primary: "WASTEWATER",
		Short:   "WSTWTR",
		Alt: []string{
			"WASTEWATER", "WSTWTR",
		},
	},
	{
		Primary: "WATER",
		Short:   "WTR",
		Alt: []string{
			"WATER", "WTR",
		},
	},
	{
		Primary: "WEBER",
		Short:   "WBR",
		Alt: []string{
			"WBR", "WEBER",
		},
	},
	{
		Primary: "WEIGHT",
		Short:   "WGHT",
		Alt: []string{
			"WEIGHT", "WGHT", "WT",
		},
	},
	{
		Primary: "WELDING",
		Short:   "WELD",
		Alt: []string{
			"WELD", "WELDING", "WLDG",
		},
	},
	{
		Primary: "WESTERN",
		Short:   "WSTRN",
		Alt: []string{
			"WESTERN", "WSTRN",
		},
	},
	{
		Primary: "WESTSIDE",
		Short:   "WSTSD",
		Alt: []string{
			"WESTSIDE", "WSTSD",
		},
	},
	{
		Primary: "WHEEL",
		Short:   "WHL",
		Alt: []string{
			"WHEEL", "WHL",
		},
	},
	{
		Primary: "WHEELER",
		Short:   "WHLR",
		Alt: []string{
			"WHEELER", "WHLR",
		},
	},
	{
		Primary: "WHITE",
		Short:   "WHT",
		Alt: []string{
			"WHITE", "WHT",
		},
	},
	{
		Primary: "WHOLESALE",
		Short:   "WHOL",
		Alt: []string{
			"WHLSE", "WHOL", "WHOLESALE", "WHS", "WHSE",
			"WHSL",
		},
	},
	{
		Primary: "WHOLESALER",
		Short:   "WHSLR",
		Alt: []string{
			"WHOLESALER", "WHSLR",
		},
	},
	{
		Primary: "WINDOW",
		Short:   "WNDW",
		Alt: []string{
			"WIN", "WINDOW", "WNDW",
		},
	},
	{
		Primary: "WIRING",
		Short:   "WIRG",
		Alt: []string{
			"WIRG", "WIRING",
		},
	},
	{
		Primary: "WITNESS",
		Short:   "WTNS",
		Alt: []string{
			"WITNESS", "WTNS",
		},
	},
	{
		Primary: "WOMEN",
		Short:   "WMN",
		Alt: []string{
			"WM", "WMN", "WOMEN",
		},
	},
	{
		Primary: "WOODWORK",
		Short:   "WOODWK",
		Alt: []string{
			"WOODWK", "WOODWORK",
		},
	},
	{
		Primary: "WOODWORKING",
		Short:   "WOODWKG",
		Alt: []string{
			"WDWKG", "WOODWKG", "WOODWORKING",
		},
	},
	{
		Primary: "WOOLEN",
		Short:   "WOOL",
		Alt: []string{
			"WOOL", "WOOLEN",
		},
	},
	{
		Primary: "WORKER",
		Short:   "WRKR",
		Alt: []string{
			"WKR", "WORKER", "WRKR",
		},
	},
	{
		Primary: "WORKING",
		Short:   "WKG",
		Alt: []string{
			"WKG", "WORKING",
		},
	},
	{
		Primary: "WORKSHOP",
		Short:   "WRKSHP",
		Alt: []string{
			"WORKSHOP", "WRKSHP",
		},
	},
	{
		Primary: "WORLD",
		Short:   "WLD",
		Alt: []string{
			"WLD", "WORLD", "WRLD",
		},
	},
	{
		Primary: "WORLDWIDE",
		Short:   "WRLDWD",
		Alt: []string{
			"WORLDWIDE", "WRLDWD",
		},
	},
	{
		Primary: "WRECKER",
		Short:   "WRCKR",
		Alt: []string{
			"WRCKR", "WRECKER",
		},
	},
	{
		Primary: "WRECKING",
		Short:   "WRCKG",
		Alt: []string{
			"WRCKG", "WRECKING",
		},
	},
	{
		Primary: "WRITER",
		Short:   "WRTR",
		Alt: []string{
			"WRITER", "WRTR",
		},
	},
	{
		Primary: "YACHT",
		Short:   "YCHT",
		Alt: []string{
			"YACHT", "YCHT",
		},
	},
	{
		Primary: "YELLOW",
		Short:   "YLW",
		Alt: []string{
			"YELLOW", "YLW",
		},
	},
	{
		Primary: "YOGURT",
		Short:   "YGRT",
		Alt: []string{
			"YGRT", "YOGURT",
		},
	},
	{
		Primary: "YOUNG",
		Short:   "YNG",
		Alt: []string{
			"YNG", "YOUNG",
		},
	},
	{
		Primary: "YOUTH",
		Short:   "YTH",
		Alt: []string{
			"YOUTH", "YTH",
		},
	},
}

// All yields every row of Appendix G, in the order the standard prints them.
//
// The rows are yielded rather than returned as a slice so a consumer cannot
// mutate the table underneath another one. Alt is a slice and is still
// shared; a consumer that needs to hold on to it should copy it.
func All() iter.Seq[BusinessWord] {
	return slices.Values(businessWords)
}

// Len reports how many rows All will yield.
func Len() int {
	return len(businessWords)
}
