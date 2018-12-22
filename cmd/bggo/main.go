package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ssilva/bggo"
)

const (
	bggurlplays  string = "https://www.boardgamegeek.com/xmlapi2/plays?username="
	bggurlsearch string = "https://www.boardgamegeek.com/xmlapi2/search?type=boardgame&query="
	bggurlthing  string = "https://www.boardgamegeek.com/xmlapi2/thing?stats=1&id="
)

func printHelp() {
	fmt.Println("bggo: Get statistics from BoardGameGeek.com")
	fmt.Println("To get statistcs on a board game:")
	fmt.Println(" bggo GAMENAME")
	fmt.Println("To get play statistcs on a user:")
	fmt.Println(" bggo -plays USERNAME")
}

func httpGetAndReadAll(url string) (xmldata []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	defer resp.Body.Close()

	xmldata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	return
}

func unmarshalOrDie(xmldata []byte, object interface{}) {
	err := xml.Unmarshal(xmldata, object)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	return
}

func retrievePlays(username string) (resp *bggo.PlaysResponse) {
	xmldata := httpGetAndReadAll(bggurlplays + url.QueryEscape(username))
	resp = &bggo.PlaysResponse{}
	unmarshalOrDie(xmldata, resp)

	return
}

func printPlays(resp *bggo.PlaysResponse) {
	fmt.Printf("Last %d plays for %s\n", len(resp.Plays), resp.Username)
	for _, play := range resp.Plays {
		fmt.Printf("\t%s: ", play.Date)
		for i, item := range play.Items {
			fmt.Printf("%s", item.Name)
			if i < (len(play.Items) - 1) {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

func retrieveGame(gameID string) (resp *bggo.ThingResponse) {
	xmldata := httpGetAndReadAll(bggurlthing + gameID)
	resp = &bggo.ThingResponse{}
	unmarshalOrDie(xmldata, resp)

	return
}

func printGame(resp *bggo.ThingResponse) {
	fmt.Printf("[%.1f] (%5d votes) %s\n",
		resp.Item.Ratings.Average.Value,
		resp.Item.Ratings.UsersRated.Value,
		resp.Item.PrimaryName(),
	)
}

func searchGame(gameName string, exactSearch bool) (resp *bggo.SearchResponse) {
	url := bggurlsearch + url.QueryEscape(gameName)
	if exactSearch {
		url += "&exact=1"
	}

	xmldata := httpGetAndReadAll(url)
	resp = &bggo.SearchResponse{}
	unmarshalOrDie(xmldata, resp)

	return
}

func retrieveAndPrintGameRating(gameName string, exactSearch bool) {
	results := searchGame(gameName, exactSearch)

	if results.Total == 0 {
		fmt.Println("Search returned no items")
		return
	}

	for _, item := range results.Items {
		game := retrieveGame(item.ID)
		printGame(game)
	}
}

func main() {
	help := flag.Bool("help", false, "print usage")
	plays := flag.String("plays", "", "get plays for USER")
	exactSearch := flag.Bool("exact", false, "exact search")
	flag.Parse()
	gameName := flag.Arg(0)

	if *help == true {
		printHelp()

	} else if *plays != "" {
		response := retrievePlays(*plays)
		printPlays(response)

	} else if gameName != "" {
		if len(gameName) < 3 {
			fmt.Println("The game name must be longer than 2 characters.")
			os.Exit(1)
		}
		retrieveAndPrintGameRating(gameName, *exactSearch)

	} else {
		printHelp()
	}
}
