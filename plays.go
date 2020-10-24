package bggo

import (
	"encoding/xml"
)

// PlaysResponse represents the /plays API response
type PlaysResponse struct {
	XMLName    xml.Name `xml:"plays"`
	Username   string   `xml:"username,attr"`
	UserID     string   `xml:"userid,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	TermsOfUse string   `xml:"termsofuse,attr"`
	Plays      []play   `xml:"play"`
}

type play struct {
	ID         string     `xml:"id,attr"`
	Date       string     `xml:"date,attr"` // TODO Make this an actual `date` (not a `string`)
	Quantity   int        `xml:"quantity,attr"`
	Length     int        `xml:"length,attr"`
	Incomplete bool       `xml:"incomplete,attr"`
	NowInStats bool       `xml:"nowinstats,attr"`
	Location   string     `xml:"location,attr"`
	Items      []playitem `xml:"item"`
	Players    []player   `xml:"players>player"`
}

type playitem struct {
	Name       string    `xml:"name,attr"`
	ObjectType string    `xml:"objecttype,attr"`
	ObjectID   string    `xml:"objectid,attr"`
	Subtypes   []subtype `xml:"subtypes>subtype"`
}

type player struct {
	XMLName       xml.Name `xml:"player"`
	Username      string   `xml:"username,attr"`
	UserID        string   `xml:"userid,attr"`
	Name          string   `xml:"name,attr"`
	StartPosition int      `xml:"startposition,attr"`
	Color         string   `xml:"color,attr"`
	Score         int      `xml:"score,attr"`
	New           bool     `xml:"new,attr"`
	Rating        int      `xml:"rating,attr"`
	Win           bool     `xml:"win,attr"`
}

type subtype struct {
	Name string `xml:"value,attr"`
}
