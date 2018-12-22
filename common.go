package bggo

type intvalue struct {
	Value int `xml:"value,attr"`
}

type floatvalue struct {
	Value float32 `xml:"value,attr"`
}

type stringvalue struct {
	Value string `xml:"value,attr"`
}
