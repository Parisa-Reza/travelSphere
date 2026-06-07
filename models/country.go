package models

type CountryInfo struct {
	Flags     FlagData            `json:"flags"`
	Name      NameData            `json:"name"`
	Languages map[string]string   `json:"languages"`
	Capital   []string            `json:"capital"`
	Region    string              `json:"region"`
}

type FlagData struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
	Alt string `json:"alt"`
}

type NameData struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}