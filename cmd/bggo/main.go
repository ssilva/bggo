package main

import (
	"flag"
	"fmt"
	"math"
	"net/url"
	"os"

	"github.com/ssilva/bggo"
)

const (
	bggurlplays      string = "https://www.boardgamegeek.com/xmlapi2/plays?username="
	bggurlsearch     string = "https://www.boardgamegeek.com/xmlapi2/search?type=boardgame&query="
	bggurlthing      string = "https://www.boardgamegeek.com/xmlapi2/thing?stats=1&id="
	bggurlhot        string = "https://www.boardgamegeek.com/xmlapi2/hot?type=boardgame"
	bggurlcollection string = "https://www.boardgamegeek.com/xmlapi2/collection?own=1&stats=1&username="
)

func printHelp() {
	fmt.Println("bggo: Get statistics from BoardGameGeek.com")
	fmt.Println()
	fmt.Println("To get the rating of a board game:")
	fmt.Println("  bggo GAMENAME")
	fmt.Println()
	fmt.Println("To get the rating of a board game, using exact search:")
	fmt.Println("  bggo -exact GAMENAME")
	fmt.Println()
	fmt.Println("To get a user's plays:")
	fmt.Println("  bggo -plays USERNAME")
	fmt.Println()
	fmt.Println("To get statistcs on a user's collection of owned games:")
	fmt.Println("  bggo -collection USERNAME")
	fmt.Println()
	fmt.Println("To get the list of most active games:")
	fmt.Println("  bggo -hot")
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

// `gameIDs` is a comma-separated list of game IDs
func retrieveGames(gameIDs string) (resp *bggo.ThingResponse) {
	xmldata := httpGetAndReadAll(bggurlthing + gameIDs)
	resp = &bggo.ThingResponse{}
	unmarshalOrDie(xmldata, resp)

	return
}

func printGames(resp *bggo.ThingResponse) {
	for _, item := range resp.Items {
		fmt.Printf("[%.1f] (%5d votes) %s\n",
			item.Ratings.Average.Value,
			item.Ratings.UsersRated.Value,
			item.PrimaryName(),
		)
	}
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
		game := retrieveGames(item.ID)
		printGames(game)
	}
}

func retrieveAndPrintHotGames() {
	xmldata := httpGetAndReadAll(bggurlhot)
	resp := &bggo.HotResponse{}
	unmarshalOrDie(xmldata, resp)

	for _, item := range resp.Items {
		fmt.Printf("[%2d] %s (%s)\n", item.Rank, item.Name.Value, item.YearPublished.Value)
	}

	return
}

type customstats struct {
	// Stats based on CollectionResponse
	mostPlayedName  string
	mostPlayedCount int

	// Stats based on ThingResponmse
	designers  map[string]int
	mechanics  map[string]int
	categories map[string]int

	mostPopularName  string
	mostPopularCount int

	leastPopularName  string
	leastPopularCount int

	highestRatedName  string
	highestRatedAvg   float32
	highestRatedVotes int

	lowestRatedName  string
	lowestRatedAvg   float32
	lowestRatedVotes int
}

func makeCustomStats() *customstats {
	return &customstats{
		designers:         make(map[string]int),
		mechanics:         make(map[string]int),
		categories:        make(map[string]int),
		leastPopularCount: math.MaxUint32,
		lowestRatedAvg:    math.MaxFloat32,
	}
}

func collectStatsOnCollection(stats *customstats, coll *bggo.CollectionResponse) {
	for _, item := range coll.Items {
		if item.NumPlays >= stats.mostPlayedCount {
			stats.mostPlayedName = item.Name.Value
			stats.mostPlayedCount = item.NumPlays
		}
	}
}

func collectStatsOnGames(stats *customstats, games *bggo.ThingResponse) {
	for _, g := range games.Items {
		for _, link := range g.Links {
			switch link.Type {
			case "boardgamedesigner":
				stats.designers[link.Value]++
			case "boardgamemechanic":
				stats.mechanics[link.Value]++
			case "boardgamecategory":
				stats.categories[link.Value]++
			}
		}

		if g.Ratings.Owned.Value >= stats.mostPopularCount {
			stats.mostPopularName = g.PrimaryName()
			stats.mostPopularCount = g.Ratings.Owned.Value
		}

		if g.Ratings.Owned.Value < stats.leastPopularCount {
			stats.leastPopularName = g.PrimaryName()
			stats.leastPopularCount = g.Ratings.Owned.Value
		}

		if g.Ratings.Average.Value >= stats.highestRatedAvg {
			stats.highestRatedName = g.PrimaryName()
			stats.highestRatedAvg = g.Ratings.Average.Value
			stats.highestRatedVotes = g.Ratings.UsersRated.Value
		}

		if g.Ratings.Average.Value < stats.lowestRatedAvg {
			stats.lowestRatedName = g.PrimaryName()
			stats.lowestRatedAvg = g.Ratings.Average.Value
			stats.lowestRatedVotes = g.Ratings.UsersRated.Value
		}
	}
}

func printStats(username string, stats *customstats) {
	fmt.Println()
	fmt.Printf("Stats for %s's Collection\n", username)
	fmt.Println()
	fmt.Println("Owned Games")
	fmt.Printf("\tMost played:   %s (%d plays by %s)\n", stats.mostPlayedName, stats.mostPlayedCount, username)
	fmt.Printf("\tMost popular:  %s (%d owners)\n", stats.mostPopularName, stats.mostPopularCount)
	fmt.Printf("\tLeast popular: %s (%d owners)\n", stats.leastPopularName, stats.leastPopularCount)
	fmt.Printf("\tHighest rated: %s (%.1f average, %d votes)\n", stats.highestRatedName, stats.highestRatedAvg, stats.highestRatedVotes)
	fmt.Printf("\tLowest rated:  %s (%.1f average, %d votes)\n", stats.lowestRatedName, stats.lowestRatedAvg, stats.lowestRatedVotes)

	fmt.Println()
	printCollectionStats(10, &stats.designers, "Designers")
	fmt.Println()
	printCollectionStats(10, &stats.mechanics, "Mechanics")
	fmt.Println()
	printCollectionStats(10, &stats.categories, "Categories")
}

func retrieveCollection(username string) (coll *bggo.CollectionResponse) {
	xmldata := httpGetAndReadAll(bggurlcollection + username)
	coll = &bggo.CollectionResponse{}
	unmarshalOrDie(xmldata, coll)
	return
}

func retrieveAndPrintCollectionStats(username string) {
	collection := retrieveCollection(username)
	games := retrieveGames(collection.JoinObjectIDs())

	stats := makeCustomStats()
	collectStatsOnCollection(stats, collection)
	collectStatsOnGames(stats, games)

	printStats(username, stats)
}

// printCollectionStats prints the top `limit` items of a map holding collection statistics of a
// particular type (e.g. number of games owned, grouped by boardgame designers)
func printCollectionStats(limit int, collectionStats *map[string]int, collectionStatsLabel string) {
	keyValues := sortMapByValue(collectionStats)

	fmt.Printf("Top %d %s\n", limit, collectionStatsLabel)
	for _, keyvalue := range keyValues[:limit] {
		fmt.Printf("\t%s [%d]\n", keyvalue.key, keyvalue.value)
	}
}

func main() {
	help := flag.Bool("help", false, "print usage")
	plays := flag.String("plays", "", "get plays for USER")
	hot := flag.Bool("hot", false, "get the list of most active games")
	collection := flag.String("collection", "", "get collection stats for USER")
	exactSearch := flag.Bool("exact", false, "exact search")
	flag.Parse()
	gameName := flag.Arg(0)

	if *help {
		printHelp()

	} else if *plays != "" {
		response := retrievePlays(*plays)
		printPlays(response)

	} else if *hot {
		retrieveAndPrintHotGames()

	} else if *collection != "" {
		retrieveAndPrintCollectionStats(*collection)

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
