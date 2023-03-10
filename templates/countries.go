package templates

// Maps country codes to English names
var countryMap = map[string]string{"ad": "Andorra",
	"ae":    "United Arab Emirates",
	"af":    "Afghanistan",
	"ag":    "Antigua and Barbuda",
	"ai":    "Anguilla",
	"al":    "Albania",
	"am":    "Armenia",
	"ao":    "Angola",
	"aq":    "Antarctica",
	"ar":    "Argentina",
	"as":    "American Samoa",
	"at":    "Austria",
	"au":    "Australia",
	"aw":    "Aruba",
	"ax":    "Aland Islands !",
	"az":    "Azerbaijan",
	"ba":    "Bosnia and Herzegovina",
	"bb":    "Barbados",
	"bd":    "Bangladesh",
	"be":    "Belgium",
	"bf":    "Burkina Faso",
	"bg":    "Bulgaria",
	"bh":    "Bahrain",
	"bi":    "Burundi",
	"bj":    "Benin",
	"bl":    "Saint Barthélemy",
	"bm":    "Bermuda",
	"bn":    "Brunei Darussalam",
	"bo":    "Bolivia",
	"bq":    "Bonaire",
	"br":    "Brazil",
	"bs":    "Bahamas",
	"bt":    "Bhutan",
	"bv":    "Bouvet Island",
	"bw":    "Botswana",
	"by":    "Belarus",
	"bz":    "Belize",
	"ca":    "Canada",
	"cc":    "Cocos (Keeling) Islands",
	"cd":    "Congo",
	"cf":    "Central African Republic",
	"cg":    "Congo",
	"ch":    "Switzerland",
	"ci":    "Cote d'Ivoire !",
	"ck":    "Cook Islands",
	"cl":    "Chile",
	"cm":    "Cameroon",
	"cn":    "China",
	"co":    "Colombia",
	"cr":    "Costa Rica",
	"cu":    "Cuba",
	"cv":    "Cabo Verde",
	"cw":    "Curaçao",
	"cx":    "Christmas Island",
	"cy":    "Cyprus",
	"cz":    "Czechia",
	"de":    "Germany",
	"dj":    "Djibouti",
	"dk":    "Denmark",
	"dm":    "Dominica",
	"do":    "Dominican Republic",
	"dz":    "Algeria",
	"ec":    "Ecuador",
	"ee":    "Estonia",
	"eg":    "Egypt",
	"eh":    "Western Sahara",
	"er":    "Eritrea",
	"es":    "Spain",
	"et":    "Ethiopia",
	"fi":    "Finland",
	"fj":    "Fiji",
	"fk":    "Falkland Islands (Malvinas)",
	"fm":    "Micronesia",
	"fo":    "Faroe Islands",
	"fr":    "France",
	"ga":    "Gabon",
	"gb":    "United Kingdom",
	"gd":    "Grenada",
	"ge":    "Georgia",
	"gf":    "French Guiana",
	"gg":    "Guernsey",
	"gh":    "Ghana",
	"gi":    "Gibraltar",
	"gl":    "Greenland",
	"gm":    "Gambia",
	"gn":    "Guinea",
	"gp":    "Guadeloupe",
	"gq":    "Equatorial Guinea",
	"gr":    "Greece",
	"gs":    "South Georgia and the South Sandwich Islands",
	"gt":    "Guatemala",
	"gu":    "Guam",
	"gw":    "Guinea-Bissau",
	"gy":    "Guyana",
	"hk":    "Hong Kong",
	"hm":    "Heard Island and McDonald Islands",
	"hn":    "Honduras",
	"hr":    "Croatia",
	"ht":    "Haiti",
	"hu":    "Hungary",
	"id":    "Indonesia",
	"ie":    "Ireland",
	"il":    "Israel",
	"im":    "Isle of Man",
	"in":    "India",
	"io":    "British Indian Ocean Territory",
	"iq":    "Iraq",
	"ir":    "Iran",
	"is":    "Iceland",
	"it":    "Italy",
	"je":    "Jersey",
	"jm":    "Jamaica",
	"jo":    "Jordan",
	"jp":    "Japan",
	"ke":    "Kenya",
	"kg":    "Kyrgyzstan",
	"kh":    "Cambodia",
	"ki":    "Kiribati",
	"km":    "Comoros",
	"kn":    "Saint Kitts and Nevis",
	"kp":    "Democratic People's Republic of Korea",
	"kr":    "Republic of Korea",
	"kw":    "Kuwait",
	"ky":    "Cayman Islands",
	"kz":    "Kazakhstan",
	"la":    "Lao People's Democratic Republic",
	"lb":    "Lebanon",
	"lc":    "Saint Lucia",
	"li":    "Liechtenstein",
	"lk":    "Sri Lanka",
	"lr":    "Liberia",
	"ls":    "Lesotho",
	"lt":    "Lithuania",
	"lu":    "Luxembourg",
	"lv":    "Latvia",
	"ly":    "Libya",
	"ma":    "Morocco",
	"mc":    "Monaco",
	"md":    "Moldova",
	"me":    "Montenegro",
	"mf":    "Saint Martin (French part)",
	"mg":    "Madagascar",
	"mh":    "Marshall Islands",
	"mk":    "Macedonia",
	"ml":    "Mali",
	"mm":    "Myanmar",
	"mn":    "Mongolia",
	"mo":    "Macao",
	"mp":    "Northern Mariana Islands",
	"mq":    "Martinique",
	"mr":    "Mauritania",
	"ms":    "Montserrat",
	"mt":    "Malta",
	"mu":    "Mauritius",
	"mv":    "Maldives",
	"mw":    "Malawi",
	"mx":    "Mexico",
	"my":    "Malaysia",
	"mz":    "Mozambique",
	"na":    "Namibia",
	"nc":    "New Caledonia",
	"ne":    "Niger",
	"nf":    "Norfolk Island",
	"ng":    "Nigeria",
	"ni":    "Nicaragua",
	"nl":    "Netherlands",
	"no":    "Norway",
	"np":    "Nepal",
	"nr":    "Nauru",
	"nu":    "Niue",
	"nz":    "New Zealand",
	"om":    "Oman",
	"pa":    "Panama",
	"pe":    "Peru",
	"pf":    "French Polynesia",
	"pg":    "Papua New Guinea",
	"ph":    "Philippines",
	"pk":    "Pakistan",
	"pl":    "Poland",
	"pm":    "Saint Pierre and Miquelon",
	"pn":    "Pitcairn",
	"pr":    "Puerto Rico",
	"ps":    "Palestine",
	"pt":    "Portugal",
	"pw":    "Palau",
	"py":    "Paraguay",
	"qa":    "Qatar",
	"re":    "Reunion !",
	"ro":    "Romania",
	"rs":    "Serbia",
	"ru":    "Russian Federation",
	"rw":    "Rwanda",
	"sa":    "Saudi Arabia",
	"sb":    "Solomon Islands",
	"sc":    "Seychelles",
	"sd":    "Sudan",
	"se":    "Sweden",
	"sg":    "Singapore",
	"sh":    "Saint Helena",
	"si":    "Slovenia",
	"sj":    "Svalbard and Jan Mayen",
	"sk":    "Slovakia",
	"sl":    "Sierra Leone",
	"sm":    "San Marino",
	"sn":    "Senegal",
	"so":    "Somalia",
	"sr":    "Suriname",
	"ss":    "South Sudan",
	"st":    "Sao Tome and Principe",
	"sv":    "El Salvador",
	"sx":    "Sint Maarten (Dutch part)",
	"sy":    "Syrian Arab Republic",
	"sz":    "Swaziland",
	"tc":    "Turks and Caicos Islands",
	"td":    "Chad",
	"tf":    "French Southern Territories",
	"tg":    "Togo",
	"th":    "Thailand",
	"tj":    "Tajikistan",
	"tk":    "Tokelau",
	"tl":    "Timor-Leste",
	"tm":    "Turkmenistan",
	"tn":    "Tunisia",
	"to":    "Tonga",
	"tr":    "Turkey",
	"tt":    "Trinidad and Tobago",
	"tv":    "Tuvalu",
	"tw":    "Taiwan",
	"tz":    "Tanzania",
	"ua":    "Ukraine",
	"ug":    "Uganda",
	"um":    "United States Minor Outlying Islands",
	"us":    "United States of America",
	"us-al": "Alabama",
	"us-ak": "Alaska",
	"us-az": "Arizona",
	"us-ar": "Arkansas",
	"us-ca": "California",
	"us-co": "Colorado",
	"us-ct": "Connecticut",
	"us-de": "Delaware",
	"us-dc": "District of Columbia",
	"us-fl": "Florida",
	"us-ga": "Georgia",
	"us-hi": "Hawaii",
	"us-id": "Idaho",
	"us-il": "Illinois",
	"us-in": "Indiana",
	"us-ia": "Iowa",
	"us-ks": "Kansas",
	"us-ky": "Kentucky",
	"us-la": "Louisiana",
	"us-me": "Maine",
	"us-md": "Maryland",
	"us-ma": "Massachusetts",
	"us-mi": "Michigan",
	"us-mn": "Minnesota",
	"us-ms": "Mississippi",
	"us-mo": "Missouri",
	"us-mt": "Montana",
	"us-ne": "Nebraska",
	"us-nv": "Nevada",
	"us-nh": "New Hampshire",
	"us-nj": "New Jersey",
	"us-nm": "New Mexico",
	"us-ny": "New York",
	"us-nc": "North Carolina",
	"us-nd": "North Dakota",
	"us-oh": "Ohio",
	"us-ok": "Oklahoma",
	"us-or": "Oregon",
	"us-pa": "Pennsylvania",
	"us-ri": "Rhode Island",
	"us-sc": "South Carolina",
	"us-sd": "South Dakota",
	"us-tn": "Tennessee",
	"us-tx": "Texas",
	"us-ut": "Utah",
	"us-vt": "Vermont",
	"us-va": "Virginia",
	"us-wa": "Washington",
	"us-wv": "West Virginia",
	"us-wi": "Wisconsin",
	"us-wy": "Wyoming",
	"uy":    "Uruguay",
	"uz":    "Uzbekistan",
	"va":    "Holy See",
	"vc":    "Saint Vincent and the Grenadines",
	"ve":    "Venezuela",
	"vg":    "British Virgin Islands",
	"vi":    "U.S. Virgin Islands",
	"vn":    "Viet Nam",
	"vu":    "Vanuatu",
	"wf":    "Wallis and Futuna",
	"ws":    "Samoa",
	"ye":    "Yemen",
	"yt":    "Mayotte",
	"za":    "South Africa",
	"zm":    "Zambia",
	"zw":    "Zimbabwe",
}
