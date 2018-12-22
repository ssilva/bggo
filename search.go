package bggo

import (
	"encoding/xml"
)

// SearchResponse represents the /search API response
type SearchResponse struct {
	XMLName    xml.Name      `xml:"items"`
	Total      int           `xml:"total,attr"`
	TermsOfUse string        `xml:"termsofuse,attr"`
	Items      []searchitems `xml:"item"`
}

type searchitems struct {
	Type          string         `xml:"type,attr"`
	ID            string         `xml:"id,attr"`
	Name          searchitemname `xml:"name"`
	YearPublished stringvalue    `xml:"yearpublished"`
}

type searchitemname struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}
