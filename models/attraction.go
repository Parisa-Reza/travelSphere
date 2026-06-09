package models

type Attraction struct {
	Xid   string          `json:"xid"`
	Name  string          `json:"name"`
	Dist  float64         `json:"dist"`
	Rate  int             `json:"rate"`
	Kinds string          `json:"kinds"`
	Point AttractionPoint `json:"point"`

	
	DisplayKinds string `json:"DisplayKinds"`
}

type AttractionPoint struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}