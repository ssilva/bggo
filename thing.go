package bggo

import (
	"encoding/xml"
)

// ThingResponse represents the /thing API response
type ThingResponse struct {
	XMLName    xml.Name    `xml:"items"`
	TermsOfUse string      `xml:"termsofuse,attr"`
	Items      []thingitem `xml:"item"`
}

type thingitem struct {
	Type          string      `xml:"type,attr"`
	ID            string      `xml:"id,attr"`
	Thumbnail     string      `xml:"thumbnail"`
	Image         string      `xml:"image"`
	Names         []name      `xml:"name"`
	Description   string      `xml:"description"`
	YearPublished stringvalue `xml:"yearpublished"`
	MinPlayers    intvalue    `xml:"minplayers"`
	MaxPlayers    intvalue    `xml:"maxplayers"`
	Polls         []poll      `xml:"poll"`
	PlayingTime   intvalue    `xml:"playingtime"`
	MinPlayTime   intvalue    `xml:"minplaytime"`
	MaxPlayTime   intvalue    `xml:"maxplaytime"`
	MinAge        intvalue    `xml:"minage"`
	Links         []link      `xml:"link"`
	Ratings       ratings     `xml:"statistics>ratings"`
}

// PrimaryName returns the name that is the primary one
func (t *thingitem) PrimaryName() (primaryName string) {
	for _, name := range t.Names {
		if name.Type == "primary" {
			primaryName = name.Value
			break
		}
	}
	return
}

type name struct {
	Type      string `xml:"type,attr"`
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:"value,attr"`
}

type poll struct {
	Name       string    `xml:"name,attr"`
	Title      string    `xml:"title,attr"`
	TotalVotes int       `xml:"totalvotes,attr"`
	Results    []results `xml:"results"`
}

type results struct {
	NumPlayers string   `xml:"numplayers,attr"` // Set when poll.Name == 'suggested_numplayers'
	Results    []result `xml:"result"`
}

type result struct {
	Level    int    `xml:"level,attr"` // Set when poll.Name == 'language_dependence'
	Value    string `xml:"value,attr"`
	NumVotes int    `xml:"numvotes,attr"`
}

type link struct {
	Type  string `xml:"type,attr"`
	ID    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type ratings struct {
	UsersRated    intvalue   `xml:"usersrated"`
	Average       floatvalue `xml:"average"`
	BayesAverage  floatvalue `xml:"bayesaverage"`
	Ranks         []rank     `xml:"ranks>rank"`
	StdDev        floatvalue `xml:"stddev"`
	Median        intvalue   `xml:"median"`
	Owned         intvalue   `xml:"owned"`
	Trading       intvalue   `xml:"trading"`
	Wanting       intvalue   `xml:"wanting"`
	Wishing       intvalue   `xml:"wishing"`
	NumComments   intvalue   `xml:"numcomments"`
	NumWeights    intvalue   `xml:"numweights"`
	AverageWeight floatvalue `xml:"averageweight"`
}

// BoardGameRank returns the rank that is the board game one
func (r *ratings) BoardGameRank() (boardGameRank string) {
	for _, rank := range r.Ranks {
		if rank.Name == "boardgame" {
			if rank.Value == "Not Ranked" {
				boardGameRank = "n/a"
			} else {
				boardGameRank = rank.Value
			}
			break
		}
	}
	return
}

// TODO Consider moving this struct to common.go as it is used by ThingResponse and CollectionResponse
type rank struct {
	Type         string `xml:"type,attr"`
	ID           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	FriendlyName string `xml:"friendlyname,attr"`
	Value        string `xml:"value,attr"`        // Cannot be an `int` becasuse "Not Ranked" is a possible value
	BayesAverage string `xml:"bayesaverage,attr"` // Cannot be a `float32` becasuse becasuse "Not Ranked" is a possible value
}
