package bggo

import (
	"encoding/xml"
	"testing"
)

func TestCollectionResponse(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8" standalone="yes"?>
	<items totalitems="2" termsofuse="https://example.com" pubdate="Tue, 25 Dec 2018 21:33:02 +0000">
		<item objecttype="thing" objectid="12345" subtype="boardgame" collid="33431844">
			<name sortindex="1">Game 1</name>
			<yearpublished>2010</yearpublished>
			<image>image1.jpg</image>
			<thumbnail>thumbnail1.jpg</thumbnail>
			<stats minplayers="2" maxplayers="10" minplaytime="45" maxplaytime="45" playingtime="45" numowned="19973">
				<rating value="6">
					<usersrated value="14741" />
					<average value="6.88335" />
					<bayesaverage value="6.75541" />
					<stddev value="1.24755" />
					<median value="0" />
					<ranks>
						<rank type="subtype" id="1" name="boardgame" friendlyname="Board Game Rank" value="595" bayesaverage="6.75541" />
						<rank type="family" id="5499" name="familygames" friendlyname="Family Game Rank" value="137" bayesaverage="6.78126" />
					</ranks>
				</rating>
			</stats>
			<status own="1" prevowned="1" fortrade="1" want="1" wanttoplay="1" wanttobuy="1" wishlist="1" preordered="1" lastmodified="2016-06-18 11:59:52" />
			<numplays>3</numplays>
			<comment>Comment 1</comment>
		</item>
		<item objecttype="thing" objectid="678901" subtype="boardgame" collid="41702762">
			<name sortindex="1">Game 2</name>
			<yearpublished>2016</yearpublished>
			<image>image2.jpg</image>
			<thumbnail>thumbnail2.jpg</thumbnail>
			<stats minplayers="2" maxplayers="7" minplaytime="30" maxplaytime="30" playingtime="30" numowned="86099">
				<rating value="N/A">
					<usersrated value="65664" />
					<average value="7.78963" />
					<bayesaverage value="7.69947" />
					<stddev value="1.27475" />
					<median value="0" />
					<ranks>
						<rank type="subtype" id="1" name="boardgame" friendlyname="Board Game Rank" value="44" bayesaverage="7.69947" />
						<rank type="family" id="5497" name="strategygames" friendlyname="Strategy Game Rank" value="44" bayesaverage="7.66743" />
					</ranks>
				</rating>
			</stats>
			<status own="1" prevowned="0" fortrade="0" want="0" wanttoplay="0" wanttobuy="0" wishlist="0" preordered="0" lastmodified="2017-04-23 14:30:49" />
			<numplays>2</numplays>
			<comment>Comment 2</comment>
		</item>
	</items>`

	coll := &CollectionResponse{}
	err := xml.Unmarshal([]byte(data), coll)

	assertNil(t, err)

	assertEqual(t, coll.TotalItems, 2)
	assertEqual(t, coll.TermsOfUse, "https://example.com")
	assertEqual(t, coll.PubDate, "Tue, 25 Dec 2018 21:33:02 +0000")

	assertEqual(t, len(coll.Items), 2)

	assertEqual(t, coll.Items[0].ObjectType, "thing")
	assertEqual(t, coll.Items[0].ObjectID, "12345")
	assertEqual(t, coll.Items[0].SubType, "boardgame")
	assertEqual(t, coll.Items[0].CollID, "33431844")
	assertEqual(t, coll.Items[0].Name.SortIndex, 1)
	assertEqual(t, coll.Items[0].Name.Value, "Game 1")
	assertEqual(t, coll.Items[0].YearPublished, "2010")
	assertEqual(t, coll.Items[0].Image, "image1.jpg")
	assertEqual(t, coll.Items[0].Thumbnail, "thumbnail1.jpg")

	assertEqual(t, coll.Items[0].Stats.MinPlayers, 2)
	assertEqual(t, coll.Items[0].Stats.MaxPlayers, 10)
	assertEqual(t, coll.Items[0].Stats.MinPlayTime, 45)
	assertEqual(t, coll.Items[0].Stats.MaxPlayTime, 45)
	assertEqual(t, coll.Items[0].Stats.PlayingTime, 45)
	assertEqual(t, coll.Items[0].Stats.NumOwned, 19973)
	assertEqual(t, coll.Items[0].Stats.Rating.Value, "6")
	assertEqual(t, coll.Items[0].Stats.Rating.UsersRated.Value, 14741)
	assertEqual(t, coll.Items[0].Stats.Rating.Average.Value, float32(6.88335))
	assertEqual(t, coll.Items[0].Stats.Rating.BayesAverage.Value, float32(6.75541))
	assertEqual(t, coll.Items[0].Stats.Rating.StdDev.Value, float32(1.24755))
	assertEqual(t, coll.Items[0].Stats.Rating.Median.Value, 0)

	assertEqual(t, len(coll.Items[0].Stats.Rating.Ranks), 2)
	assertEqual(t, coll.Items[0].Stats.Rating.BoardGameRank(), "595")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].Type, "subtype")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].ID, "1")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].Name, "boardgame")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].FriendlyName, "Board Game Rank")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].Value, "595")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[0].BayesAverage, "6.75541")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].Type, "family")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].ID, "5499")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].Name, "familygames")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].FriendlyName, "Family Game Rank")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].Value, "137")
	assertEqual(t, coll.Items[0].Stats.Rating.Ranks[1].BayesAverage, "6.78126")

	assertEqual(t, coll.Items[0].Status.Own, true)
	assertEqual(t, coll.Items[0].Status.PrevOwned, true)
	assertEqual(t, coll.Items[0].Status.ForTrade, true)
	assertEqual(t, coll.Items[0].Status.Want, true)
	assertEqual(t, coll.Items[0].Status.WantToPlay, true)
	assertEqual(t, coll.Items[0].Status.WantToBuy, true)
	assertEqual(t, coll.Items[0].Status.Wishlist, true)
	assertEqual(t, coll.Items[0].Status.Preordered, true)
	assertEqual(t, coll.Items[0].Status.LastModified, "2016-06-18 11:59:52")
	assertEqual(t, coll.Items[0].NumPlays, 3)
	assertEqual(t, coll.Items[0].Comment, "Comment 1")

	assertEqual(t, coll.Items[1].ObjectType, "thing")
	assertEqual(t, coll.Items[1].ObjectID, "678901")
	assertEqual(t, coll.Items[1].SubType, "boardgame")
	assertEqual(t, coll.Items[1].CollID, "41702762")
	assertEqual(t, coll.Items[1].Name.SortIndex, 1)
	assertEqual(t, coll.Items[1].Name.Value, "Game 2")
	assertEqual(t, coll.Items[1].YearPublished, "2016")
	assertEqual(t, coll.Items[1].Image, "image2.jpg")
	assertEqual(t, coll.Items[1].Thumbnail, "thumbnail2.jpg")

	assertEqual(t, coll.Items[1].Stats.MinPlayers, 2)
	assertEqual(t, coll.Items[1].Stats.MaxPlayers, 7)
	assertEqual(t, coll.Items[1].Stats.MinPlayTime, 30)
	assertEqual(t, coll.Items[1].Stats.MaxPlayTime, 30)
	assertEqual(t, coll.Items[1].Stats.PlayingTime, 30)
	assertEqual(t, coll.Items[1].Stats.NumOwned, 86099)
	assertEqual(t, coll.Items[1].Stats.Rating.Value, "N/A")
	assertEqual(t, coll.Items[1].Stats.Rating.UsersRated.Value, 65664)
	assertEqual(t, coll.Items[1].Stats.Rating.Average.Value, float32(7.78963))
	assertEqual(t, coll.Items[1].Stats.Rating.BayesAverage.Value, float32(7.69947))
	assertEqual(t, coll.Items[1].Stats.Rating.StdDev.Value, float32(1.27475))
	assertEqual(t, coll.Items[1].Stats.Rating.Median.Value, 0)

	assertEqual(t, len(coll.Items[1].Stats.Rating.Ranks), 2)
	assertEqual(t, coll.Items[1].Stats.Rating.BoardGameRank(), "44")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].Type, "subtype")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].ID, "1")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].Name, "boardgame")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].FriendlyName, "Board Game Rank")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].Value, "44")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[0].BayesAverage, "7.69947")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].Type, "family")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].ID, "5497")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].Name, "strategygames")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].FriendlyName, "Strategy Game Rank")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].Value, "44")
	assertEqual(t, coll.Items[1].Stats.Rating.Ranks[1].BayesAverage, "7.66743")

	assertEqual(t, coll.Items[1].Status.Own, true)
	assertEqual(t, coll.Items[1].Status.PrevOwned, false)
	assertEqual(t, coll.Items[1].Status.ForTrade, false)
	assertEqual(t, coll.Items[1].Status.Want, false)
	assertEqual(t, coll.Items[1].Status.WantToPlay, false)
	assertEqual(t, coll.Items[1].Status.WantToBuy, false)
	assertEqual(t, coll.Items[1].Status.Wishlist, false)
	assertEqual(t, coll.Items[1].Status.Preordered, false)
	assertEqual(t, coll.Items[1].Status.LastModified, "2017-04-23 14:30:49")
	assertEqual(t, coll.Items[1].NumPlays, 2)
	assertEqual(t, coll.Items[1].Comment, "Comment 2")
}
