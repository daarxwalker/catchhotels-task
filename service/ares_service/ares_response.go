package ares_service

type Response struct {
	Address Address `json:"sidlo"`
	Cin     string  `json:"ico"`
	Vat     string  `json:"dic"`
	Name    string  `json:"obchodniJmeno"`
}

type Address struct {
	Full                   string `json:"textovaAdresa"`
	CountryCode            string `json:"kodStatu"`
	City                   string `json:"nazevObce"`
	Street                 string `json:"nazevUlice"`
	StreetHouseNumber      int    `json:"cisloDomovni"`
	StreetIndicativeNumber int    `json:"cisloOrientacni"`
	Zip                    int    `json:"psc"`
}
