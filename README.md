# bggo

bggo is
- a command line tool for retrieving stats from BoardGameGeek.com
- a Go package for unmarshalling [BGG XML API2](https://boardgamegeek.com/wiki/page/BGG_XML_API2) responses

[![Go Report Card](https://goreportcard.com/badge/github.com/ssilva/bggo)](https://goreportcard.com/report/github.com/ssilva/bggo)
![Go](https://github.com/ssilva/bggo/workflows/Go/badge.svg)

## The command line tool

### Installation
```bash
$ git clone https://github.com/ssilva/bggo.git
$ cd bggo
$ go build ./cmd/bggo
# Optionally, add `bggo` to the path
```

### Examples

For the moment, the following use cases are available.

1. Get the rating of a board game:
    ```
    $ bggo "Terra Mystica"
    [8.2] (31304 votes, rank  10) Terra Mystica
    [8.1] (  758 votes, rank n/a) Terra Mystica: 4 Town Tiles
    [8.7] (  187 votes, rank n/a) Terra Mystica: Big Box
    [7.8] (  536 votes, rank n/a) Terra Mystica: Bonus Card Shipping Value
    [8.2] (  493 votes, rank n/a) Terra Mystica: Erweiterungsbogen
    [8.5] ( 3063 votes, rank n/a) Terra Mystica: Fire & Ice
    [8.6] ( 7603 votes, rank   8) Gaia Project
    ```
1. Get the rating of a board game, using exact search:
    ```
    $ bggo -exact "Puerto Rico"
    [8.0] (54936 votes, rank  16) Puerto Rico
    ```
1. Get a user's plays:
    ```
    $ bggo -plays Silvast
    Last 100 plays for Silvast
            2018-10-23: BANG! The Dice Game
            2018-10-13: Gloomhaven
            2018-09-22: Gloomhaven
            2018-09-02: Hero Realms
            2018-08-23: Gloomhaven
            2018-03-28: 6 nimmt!
            ⋮
    ```

1. Get stats on a user's collection:
    ```
    $ bggo -collection Silvast
    
    Stats for Silvast's Collection

    Owned Games
            Most played:   Love Letter (16 plays by Silvast)
            Most popular:  Pandemic (116203 owners)
            Least popular: Sphinx (523 owners)
            Highest rated: Terraforming Mars (8.4 average, 33384 votes)
            Lowest rated:  Sphinx (5.3 average, 190 votes)

    Top 10 Designers
            Robert Dougherty [3]
            Darwin Kastle [3]
            Uwe Rosenberg [3]
            Thomas Lehmann [2]
            Vlaada Chvátil [2]
            Donald X. Vaccarino [2]
            Scott Almes [2]
            Stefan Feld [2]
            Matt Leacock [1]
            Tom Cleaver [1]

    Top 10 Mechanics
            Hand Management [22]
            Card Drafting [17]
            Set Collection [17]
            Variable Player Powers [9]
            Simultaneous Action Selection [7]
            Deck / Pool Building [7]
            Press Your Luck [6]
            Take That [6]
            Variable Phase Order [5]
            Worker Placement [5]

    Top 10 Categories
            Card Game [21]
            Economic [8]
            Fantasy [7]
            Medieval [6]
            Science Fiction [5]
            Civilization [5]
            Fighting [5]
            Party Game [5]
            Territory Building [5]
            Deduction [4]
    ```

1. Get the list of most active games:
    ```
    $ bggo -hot
    [ 1] Tainted Grail: Fall of Avalon (2019)
    [ 2] New Frontiers (2018)
    [ 3] Mage Knight: Ultimate Edition (2018)
    [ 4] Gloomhaven (2017)
    [ 5] Beta Colony (2018)
    [ 6] Nemesis (2018)
    [ 7] Terraforming Mars (2016)
    [ 8] KeyForge: Call of the Archons (2018)
    [ 9] Azul (2017)
    [10] Architects of the West Kingdom (2018)
    ⋮
    ```

## The Go package

### Installation
```bash
$ go get github.com/ssilva/bggo
```

### Example
```go
package main

import (
	"fmt"
	"encoding/xml"

	"github.com/ssilva/bggo"
)

func main() {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<plays username="Silvast" userid="1234567" total="100" page="1" termsofuse="https://example.com">
		<play id="12345" date="1999-09-09" quantity="1" length="10" incomplete="1" nowinstats="1" location="Montreal">
			<item name="Terra Mystica" objecttype="thing" objectid="123">
				<subtypes>
					<subtype value="boardgame" />
					<subtype value="boardgameimplementation" />
				</subtypes>
			</item>
		</play>
	</plays>`

	plays := bggo.PlaysResponse{}
	err := xml.Unmarshal([]byte(data), &plays)
	if err != nil {
		panic(err)
	}
	fmt.Println(plays.Plays[0].Items[0].Name) // Terra Mystica
}
```
For more examples, see the `*_test.go` files.

### API coverage

Implemented:

- Thing Items (`/thing`)
- Plays (`/plays`)
- Search (`/search`)
- Hot Items (`/hot`)
- Collection (`/collection`)

To be implemented (in order of priority):

1. Users (`/user`)
1. Guilds (`/guild`)
1. Family Items (`/family`)
1. Forum Lists (`/forumlist`)
1. Forums (`/forum`)
1. Threads (`/thread`)

## Contributing

Please [file an issue](https://github.com/ssilva/bggo/issues) for bugs and feature suggestions.
