# bggo

bggo is
- a command line tool for retrieving stats from BoardGameGeek.com
- a Go package for unmarshalling [BGG XML API2](https://boardgamegeek.com/wiki/page/BGG_XML_API2) responses

[![Go Report Card](https://goreportcard.com/badge/github.com/ssilva/bggo)](https://goreportcard.com/report/github.com/ssilva/bggo)
[![Build Status](https://travis-ci.org/ssilva/bggo.svg?branch=master)](https://travis-ci.org/ssilva/bggo)

## The command line tool

### Installation
```bash
$ git clone https://github.com/ssilva/bggo.git
$ cd bggo
$ go build ./cmd/bggo
# Optionally, add `bggo` to the path
```

### Examples

For the moment, the following use cases are availabe.

1. Get the rating of a board game:
    ```
    $ bggo "Terra Mystica"
    [8.2] (31152 votes) Terra Mystica
    [8.1] (  759 votes) Terra Mystica: 4 Town Tiles
    [8.7] (  183 votes) Terra Mystica: Big Box
    [7.8] (  537 votes) Terra Mystica: Bonus Card Shipping Value
    [8.2] (  491 votes) Terra Mystica: Erweiterungsbogen
    [8.5] ( 3054 votes) Terra Mystica: Fire & Ice
    [8.6] ( 7426 votes) Gaia Project
    ```
2. Get the rating of a board game, using exact search:
    ```
    $ bggo -exact "Puerto Rico"
    [8.0] (54791 votes) Puerto Rico
    ```
2. Get stats for a user:
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

To be implemented (in order of priority):

1. Hot Items (`/hot`)
1. Users (`/user`)
1. Collection (`/collection`)
1. Guilds (`/guild`)
1. Family Items (`/family`)
1. Forum Lists (`/forumlist`)
1. Forums (`/forum`)
1. Threads (`/thread`)

## Contributing

Please [file an issue](https://github.com/ssilva/bggo/issues) for bugs and feature suggestions.
