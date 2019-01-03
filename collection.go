package bggo

import (
	"encoding/xml"
	"math"
)

// CollectionResponse represents the /collection API response
type CollectionResponse struct {
	XMLName    xml.Name   `xml:"items"`
	TotalItems int        `xml:"totalitems,attr"`
	TermsOfUse string     `xml:"termsofuse,attr"`
	PubDate    string     `xml:"pubdate,attr"` // TODO Make this an actual `time.Time` (not a `string`)
	Items      []collitem `xml:"item"`
}

// MostPopularGameOwned returns the game from the collection with highest number of owners
func (c *CollectionResponse) MostPopularGameOwned() (name string, count int) {
	for _, item := range c.Items {
		if item.Status.Own && item.Stats.NumOwned >= count {
			name = item.Name.Value
			count = item.Stats.NumOwned
		}
	}
	return
}

// LeastPopularGameOwned returns the game from the collection with lowest number of owners
func (c *CollectionResponse) LeastPopularGameOwned() (name string, count int) {
	count = math.MaxUint32

	for _, item := range c.Items {
		if item.Status.Own && item.Stats.NumOwned < count {
			name = item.Name.Value
			count = item.Stats.NumOwned
		}
	}
	return
}

// MostPlayedGameOwned returns the game from the collection with highest number of plays by the user
func (c *CollectionResponse) MostPlayedGameOwned() (name string, count int) {
	for _, item := range c.Items {
		if item.Status.Own && item.NumPlays >= count {
			name = item.Name.Value
			count = item.NumPlays
		}
	}
	return
}

// LeastPlayedGameOwned returns the game from the collection with lowest number of plays by the user
func (c *CollectionResponse) LeastPlayedGameOwned() (name string, count int) {
	count = math.MaxUint32

	for _, item := range c.Items {
		if item.Status.Own && item.NumPlays < count {
			name = item.Name.Value
			count = item.NumPlays
		}
	}
	return
}

// JoinObjectIDs joins the collection's object IDs, separated by commas
func (c *CollectionResponse) JoinObjectIDs() (joinedIDs string) {
	for i, collitem := range c.Items {
		joinedIDs += collitem.ObjectID
		if i < len(c.Items) {
			joinedIDs += ","
		}
	}
	return
}

type collitem struct {
	ObjectType string `xml:"objecttype,attr"`
	ObjectID   string `xml:"objectid,attr"`
	SubType    string `xml:"subtype,attr"`
	CollID     string `xml:"collid,attr"`
	Name       struct {
		Value     string `xml:",chardata"`
		SortIndex int    `xml:"sortindex,attr"`
	} `xml:"name"`
	YearPublished string `xml:"yearpublished"`
	Image         string `xml:"image"`
	Thumbnail     string `xml:"thumbnail"`
	Stats         stats  `xml:"stats"`
	Status        status `xml:"status"`
	NumPlays      int    `xml:"numplays"`
	Comment       string `xml:"comment"`
}

type stats struct {
	MinPlayers  int    `xml:"minplayers,attr"`
	MaxPlayers  int    `xml:"maxplayers,attr"`
	MinPlayTime int    `xml:"minplaytime,attr"`
	MaxPlayTime int    `xml:"maxplaytime,attr"`
	PlayingTime int    `xml:"playingtime,attr"`
	NumOwned    int    `xml:"numowned,attr"`
	Rating      rating `xml:"rating"`
}

type rating struct {
	Value        string     `xml:"value,attr"`
	UsersRated   intvalue   `xml:"usersrated"`
	Average      floatvalue `xml:"average"`
	BayesAverage floatvalue `xml:"bayesaverage"`
	StdDev       floatvalue `xml:"stddev"`
	Median       intvalue   `xml:"median"`
	Ranks        []rank     `xml:"ranks>rank"`
}

// BoardGameRank returns the rank that is the board game one
func (r *rating) BoardGameRank() (boardGameRank string) {
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

type status struct {
	Own          bool   `xml:"own,attr"`
	PrevOwned    bool   `xml:"prevowned,attr"`
	ForTrade     bool   `xml:"fortrade,attr"`
	Want         bool   `xml:"want,attr"`
	WantToPlay   bool   `xml:"wanttoplay,attr"`
	WantToBuy    bool   `xml:"wanttobuy,attr"`
	Wishlist     bool   `xml:"wishlist,attr"`
	Preordered   bool   `xml:"preordered,attr"`
	LastModified string `xml:"lastmodified,attr"` // TODO Make this `time.Time` (instead of `string`)
}
