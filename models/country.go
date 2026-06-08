package models

type CountryInfo struct {
	Name       NameData                  `json:"name"`
	Capital    []string                  `json:"capital"`
	Region     string                    `json:"region"`
	Population int    					 `json:"population"`
	Languages  map[string]string         `json:"languages"`
	Currencies map[string]CurrencyDetail `json:"currencies"`
	Flags      FlagData                  `json:"flags"`

	DisplayCapital    string `json:"DisplayCapital"`
	DisplayLanguages  string `json:"DisplayLanguages"`
	DisplayCurrencies string `json:"DisplayCurrencies"`
	Slug              string `json:"Slug"`
}

type NameData struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type CurrencyDetail struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type FlagData struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}