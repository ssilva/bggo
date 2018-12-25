package bggo

import (
	"encoding/xml"
)

// HotResponse represents the /hot API response
type HotResponse struct {
	XMLName    xml.Name  `xml:"items"`
	TermsOfUse string    `xml:"termsofuse,attr"`
	Items      []hotitem `xml:"item"`
}

type hotitem struct {
	ID            string      `xml:"id,attr"`
	Rank          int         `xml:"rank,attr"`
	Thumbnail     stringvalue `xml:"thumbnail"`
	Name          stringvalue `xml:"name"`
	YearPublished stringvalue `xml:"yearpublished"`
}
