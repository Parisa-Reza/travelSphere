package models

type CountryInfo struct {
	Name       NameData                  `json:"name"`
	Capital    []string                  `json:"capital"`
	Region     string                    `json:"region"`
	Population int                       `json:"population"`
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

type RestCountriesV5Response struct {
	Data struct {
		Objects []RestCountriesV5Country `json:"objects"`
		Meta    struct {
			More bool `json:"more"`
		} `json:"meta"`
	} `json:"data"`
}

type RestCountriesV5Country struct {
	Names struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"names"`
	Capitals []struct {
		Name string `json:"name"`
	} `json:"capitals"`
	Flag struct {
		Png string `json:"url_png"`
		Svg string `json:"url_svg"`
	} `json:"flag"`
	Region     string `json:"region"`
	Population int    `json:"population"`
	Languages  []struct {
		Code string `json:"iso639_3"`
		Name string `json:"name"`
	} `json:"languages"`
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}
