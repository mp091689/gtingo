package country

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Country struct {
	code string
	name string
}

func NewCountry() Country {
	rand.Seed(time.Now().UnixNano())

	return countries[rand.Intn(len(countries)-1)]
}

func (c Country) GetCode() string {
	match, _ := regexp.MatchString("\\d{3}-\\d{3}", c.code)

	code := c.code

	if match {
		rangeCode := strings.Split(c.code, "-")

		max, _ := strconv.Atoi(rangeCode[1])
		min, _ := strconv.Atoi(rangeCode[0])

		code = strconv.Itoa(rand.Intn(max-min) + min)
	}

	for len(code) < 3 {
		code = "0" + code
	}

	return code
}

func (c Country) GetName() string {
	return c.name
}

var countries = []Country{
	{
		code: "000-019",
		name: "UPC-A compatible - United States and Canada",
	},
	{
		code: "020-029",
		name: "UPC-A compatible - Used to issue restricted circulation numbers within a geographic region[m]",
	},
	{
		code: "030-039",
		name: "UPC-A compatible - United States drugs (see United States National Drug Code)",
	},
	{
		code: "040-049",
		name: "PC-A compatible - Used to issue restricted circulation numbers within a geographic region[m]",
	},
	{
		code: "050-059",
		name: "PC-A compatible - GS1 US reserved for future use",
	},
	{
		code: "060-099",
		name: "PC-A compatible - United States and Canada",
	},
	{
		code: "100-139",
		name: "United States",
	},
	{
		code: "200-299",
		name: "Used to issue GS1 restricted circulation number within a geographic region[m]",
	},
	{
		code: "275",
		name: "Palestine",
	},
	{
		code: "300-379",
		name: "France and Monaco",
	},
	{
		code: "380",
		name: "Bulgaria",
	},
	{
		code: "383",
		name: "Slovenia",
	},
	{
		code: "385",
		name: "Croatia",
	},
	{
		code: "387",
		name: "Bosnia and Herzegovina l",
	},
	{
		code: "389",
		name: "Montenegro",
	},
	{
		code: "390",
		name: "Kosovo",
	},
	{
		code: "400-440",
		name: "Germany (440 code inherited from old East Germany on reunification, 1990)",
	},
	{
		code: "450-459",
		name: "Japan (new Japanese Article Number range)",
	},
	{
		code: "460-469",
		name: "Russia (barcodes inherited from the Soviet Union)",
	},
	{
		code: "470",
		name: "Kyrgyzstan",
	},
	{
		code: "471",
		name: "Republic of China (Taiwan)",
	},
	{
		code: "474",
		name: "Estonia",
	},
	{
		code: "475",
		name: "Latvia",
	},
	{
		code: "476",
		name: "Azerbaijan",
	},
	{
		code: "477",
		name: "Lithuania",
	},
	{
		code: "478",
		name: "Uzbekistan",
	},
	{
		code: "479",
		name: "Sri Lanka",
	},
	{
		code: "480",
		name: "Philippines",
	},
	{
		code: "481",
		name: "Belarus",
	},
	{
		code: "482",
		name: "Ukraine",
	},
	{
		code: "483",
		name: "Turkmenistan",
	},
	{
		code: "484",
		name: "Moldova",
	},
	{
		code: "485",
		name: "Armenia",
	},
	{
		code: "486",
		name: "Georgia",
	},
	{
		code: "487",
		name: "Kazakhstan",
	},
	{
		code: "488",
		name: "Tajikistan",
	},
	{
		code: "489",
		name: "Hong Kong",
	},
	{
		code: "490-499",
		name: "Japan (original Japanese Article Number range)",
	},
	{
		code: "500-509",
		name: "United Kingdom",
	},
	{
		code: "520-521",
		name: "Greece",
	},
	{
		code: "528",
		name: "Lebanon",
	},
	{
		code: "529",
		name: "Cyprus",
	},
	{
		code: "530",
		name: "Albania",
	},
	{
		code: "531",
		name: "Macedonia",
	},
	{
		code: "535",
		name: "Malta",
	},
	{
		code: "539",
		name: "Ireland",
	},
	{
		code: "540-549",
		name: "Belgium and Luxembourg",
	},
	{
		code: "560",
		name: "Portugal",
	},
	{
		code: "569",
		name: "Iceland",
	},
	{
		code: "570-579",
		name: "Denmark , Faroe Islands and Greenland",
	},
	{
		code: "590",
		name: "Poland",
	},
	{
		code: "594",
		name: "Romania",
	},
	{
		code: "599",
		name: "Hungary",
	},
	{
		code: "600-601",
		name: "South Africa",
	},
	{
		code: "603",
		name: "Ghana",
	},
	{
		code: "604",
		name: "Senegal",
	},
	{
		code: "608",
		name: "Bahrain",
	},
	{
		code: "609",
		name: "Mauritius",
	},
	{
		code: "611",
		name: "Morocco",
	},
	{
		code: "613",
		name: "Algeria",
	},
	{
		code: "615",
		name: "Nigeria",
	},
	{
		code: "616",
		name: "Kenya",
	},
	{
		code: "618",
		name: "Ivory Coast",
	},
	{
		code: "619",
		name: "Tunisia",
	},
	{
		code: "620",
		name: "Tanzania",
	},
	{
		code: "621",
		name: "Syria",
	},
	{
		code: "622",
		name: "Egypt",
	},
	{
		code: "623",
		name: "Brunei",
	},
	{
		code: "624",
		name: "Libya",
	},
	{
		code: "625",
		name: "Jordan",
	},
	{
		code: "626",
		name: "Iran",
	},
	{
		code: "627",
		name: "Kuwait",
	},
	{
		code: "628",
		name: "Saudi Arabia",
	},
	{
		code: "629",
		name: "United Arab Emirates",
	},
	{
		code: "640-649",
		name: "Finland",
	},
	{
		code: "690-699",
		name: "People's Republic of China",
	},
	{
		code: "700-709",
		name: "Norway",
	},
	{
		code: "729",
		name: "Israel",
	},
	{
		code: "730-739",
		name: "Sweden : EAN/GS1 Sweden",
	},
	{
		code: "740",
		name: "Guatemala",
	},
	{
		code: "741",
		name: "El Salvador",
	},
	{
		code: "742",
		name: "Honduras",
	},
	{
		code: "743",
		name: "Nicaragua",
	},
	{
		code: "744",
		name: "Costa Rica",
	},
	{
		code: "745",
		name: "Panama",
	},
	{
		code: "746",
		name: "Dominican Republic",
	},
	{
		code: "750",
		name: "Mexico",
	},
	{
		code: "754-755",
		name: "Canada",
	},
	{
		code: "759",
		name: "Venezuela",
	},
	{
		code: "760-769",
		name: "Switzerland and Liechtenstein",
	},
	{
		code: "770-771",
		name: "Colombia",
	},
	{
		code: "773",
		name: "Uruguay",
	},
	{
		code: "775",
		name: "Peru",
	},
	{
		code: "777",
		name: "Bolivia",
	},
	{
		code: "778-779",
		name: "Argentina",
	},
	{
		code: "780",
		name: "Chile",
	},
	{
		code: "784",
		name: "Paraguay",
	},
	{
		code: "786",
		name: "Ecuador",
	},
	{
		code: "789-790",
		name: "Brazil",
	},
	{
		code: "800-839",
		name: "Italy, San Marino and Vatican City",
	},
	{
		code: "840-849",
		name: "Spain and Andorra",
	},
	{
		code: "850",
		name: "Cuba",
	},
	{
		code: "858",
		name: "Slovakia",
	},
	{
		code: "859",
		name: "Czech Republic",
	},
	{
		code: "860",
		name: "Serbia",
	},
	{
		code: "865",
		name: "Mongolia",
	},
	{
		code: "867",
		name: "North Korea",
	},
	{
		code: "868-869",
		name: "Turkey",
	},
	{
		code: "870-879",
		name: "Netherlands",
	},
	{
		code: "880",
		name: "South Korea",
	},
	{
		code: "884",
		name: "Cambodia",
	},
	{
		code: "885",
		name: "Thailand",
	},
	{
		code: "888",
		name: "Singapore",
	},
	{
		code: "890",
		name: "India",
	},
	{
		code: "893",
		name: "Vietnam (previously used by North Vietnam and South Vietnam before 1975)",
	},
	{
		code: "894",
		name: "Bangladesh",
	},
	{
		code: "896",
		name: "Pakistan",
	},
	{
		code: "899",
		name: "Indonesia",
	},
	{
		code: "900-919",
		name: "Austria",
	},
	{
		code: "930-939",
		name: "Australia",
	},
	{
		code: "940-949",
		name: "New Zealand",
	},
	{
		code: "950",
		name: "GS1 Global Office: Special applications",
	},
	{
		code: "951",
		name: "Used to issue General Manager Numbers for the EPC General Identifier (GID) scheme as defined by the EPC Tag Data Standard",
	},
	{
		code: "955",
		name: "Malaysia",
	},
	{
		code: "958",
		name: "Macau",
	},
	{
		code: "960-961",
		name: "GS1 UK Office: GTIN-8 allocations",
	},
	{
		code: "962-969",
		name: "GS1 Global Office: GTIN-8 allocations",
	},
	{
		code: "977",
		name: "Serial publications (ISSN)",
	},
	{
		code: "978-979",
		name: "\"Bookland\"(ISBN) - 979-0 used for sheet music (ISMN-13, replaces deprecated ISMN M- numbers)",
	},
	{
		code: "980",
		name: "Refund receipts",
	},
	{
		code: "981-984",
		name: "GS1 coupon identification for common currency areas",
	},
	{
		code: "990-999",
		name: "GS1 coupon identification",
	},
}
