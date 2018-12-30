package bggo

import (
	"encoding/xml"
	"testing"
)

func TestThingResponse(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<items termsofuse="https://example.com">
		<item type="boardgame" id="1234">
			<thumbnail>thumbnail.jpg</thumbnail>
			<image>image.jpg</image>
			<name type="primary" sortindex="1" value="Terra Mystica" />
			<name type="alternate" sortindex="1" value="テラミスティカ" />
			<description>Description</description>
			<yearpublished value="2012" />
			<minplayers value="2" />
			<maxplayers value="5" />
			<poll name="suggested_numplayers" title="User Suggested Number of Players" totalvotes="720">
				<results numplayers="1">
					<result value="Best" numvotes="0" />
					<result value="Recommended" numvotes="14" />
				</results>
				<results numplayers="5+">
					<result value="Best" numvotes="6" />
					<result value="Recommended" numvotes="13" />
				</results>
			</poll>
			<playingtime value="150" />
			<minplaytime value="60" />
			<maxplaytime value="150" />
			<minage value="12" />
			<poll name="suggested_playerage" title="User Suggested Player Age" totalvotes="160">
				<results>
					<result value="18" numvotes="2" />
					<result value="21 and up" numvotes="3" />
				</results>
			</poll>
			<poll name="language_dependence" title="Language Dependence" totalvotes="170">
				<results>
					<result level="21" value="No necessary in-game text" numvotes="161" />
					<result level="22" value="Some necessary text - easily memorized or small crib sheet" numvotes="9" />
				</results>
			</poll>
			<link type="boardgamecategory" id="1234" value="Civilization" />
			<link type="boardgamemechanic" id="5678" value="Variable Player Powers" />
			<statistics page="1">
				<ratings>
					<usersrated value="31115" />
					<average value="8.20954" />
					<bayesaverage value="8.06335" />
					<ranks>
						<rank type="subtype" id="1" name="boardgame" friendlyname="Board Game Rank" value="10" bayesaverage="8.06335" />
						<rank type="family" id="5497" name="strategygames" friendlyname="Strategy Game Rank" value="11" bayesaverage="8.07124" />
					</ranks>
					<stddev value="1.46947" />
					<median value="1" />
					<owned value="34484" />
					<trading value="487" />
					<wanting value="1740" />
					<wishing value="10493" />
					<numcomments value="4786" />
					<numweights value="2232" />
					<averageweight value="3.9507" />
				</ratings>
			</statistics>
		</item>
	</items>`

	thing := &ThingResponse{}
	err := xml.Unmarshal([]byte(data), thing)

	assertNil(t, err)

	assertEqual(t, thing.TermsOfUse, "https://example.com")
	assertEqual(t, thing.Items[0].Type, "boardgame")
	assertEqual(t, thing.Items[0].ID, "1234")
	assertEqual(t, thing.Items[0].Thumbnail, "thumbnail.jpg")
	assertEqual(t, thing.Items[0].Image, "image.jpg")

	assertEqual(t, len(thing.Items[0].Names), 2)
	assertEqual(t, thing.Items[0].PrimaryName(), "Terra Mystica")
	assertEqual(t, thing.Items[0].Names[0].Type, "primary")
	assertEqual(t, thing.Items[0].Names[0].SortIndex, 1)
	assertEqual(t, thing.Items[0].Names[0].Value, "Terra Mystica")
	assertEqual(t, thing.Items[0].Names[1].Type, "alternate")
	assertEqual(t, thing.Items[0].Names[1].SortIndex, 1)
	assertEqual(t, thing.Items[0].Names[1].Value, "テラミスティカ")

	assertEqual(t, thing.Items[0].Description, "Description")
	assertEqual(t, thing.Items[0].YearPublished.Value, "2012")
	assertEqual(t, thing.Items[0].MinPlayers.Value, 2)
	assertEqual(t, thing.Items[0].MaxPlayers.Value, 5)

	assertEqual(t, len(thing.Items[0].Polls), 3)
	assertEqual(t, thing.Items[0].Polls[0].Name, "suggested_numplayers")
	assertEqual(t, thing.Items[0].Polls[0].Title, "User Suggested Number of Players")
	assertEqual(t, thing.Items[0].Polls[0].TotalVotes, 720)
	assertEqual(t, len(thing.Items[0].Polls[0].Results), 2)
	assertEqual(t, thing.Items[0].Polls[0].Results[0].NumPlayers, "1")
	assertEqual(t, len(thing.Items[0].Polls[0].Results[0].Results), 2)
	assertEqual(t, thing.Items[0].Polls[0].Results[0].Results[0].Value, "Best")
	assertEqual(t, thing.Items[0].Polls[0].Results[0].Results[0].NumVotes, 0)
	assertEqual(t, thing.Items[0].Polls[0].Results[0].Results[1].Value, "Recommended")
	assertEqual(t, thing.Items[0].Polls[0].Results[0].Results[1].NumVotes, 14)
	assertEqual(t, thing.Items[0].Polls[0].Results[1].NumPlayers, "5+")
	assertEqual(t, len(thing.Items[0].Polls[0].Results[1].Results), 2)
	assertEqual(t, thing.Items[0].Polls[0].Results[1].Results[0].Value, "Best")
	assertEqual(t, thing.Items[0].Polls[0].Results[1].Results[0].NumVotes, 6)
	assertEqual(t, thing.Items[0].Polls[0].Results[1].Results[1].Value, "Recommended")
	assertEqual(t, thing.Items[0].Polls[0].Results[1].Results[1].NumVotes, 13)

	assertEqual(t, thing.Items[0].Polls[1].Name, "suggested_playerage")
	assertEqual(t, thing.Items[0].Polls[1].Title, "User Suggested Player Age")
	assertEqual(t, thing.Items[0].Polls[1].TotalVotes, 160)
	assertEqual(t, len(thing.Items[0].Polls[1].Results), 1)
	assertEqual(t, thing.Items[0].Polls[1].Results[0].Results[0].Value, "18")
	assertEqual(t, thing.Items[0].Polls[1].Results[0].Results[0].NumVotes, 2)
	assertEqual(t, thing.Items[0].Polls[1].Results[0].Results[1].Value, "21 and up")
	assertEqual(t, thing.Items[0].Polls[1].Results[0].Results[1].NumVotes, 3)

	assertEqual(t, thing.Items[0].Polls[2].Name, "language_dependence")
	assertEqual(t, thing.Items[0].Polls[2].Title, "Language Dependence")
	assertEqual(t, thing.Items[0].Polls[2].TotalVotes, 170)
	assertEqual(t, len(thing.Items[0].Polls[2].Results), 1)
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[0].Level, 21)
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[0].Value, "No necessary in-game text")
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[0].NumVotes, 161)
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[1].Level, 22)
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[1].Value, "Some necessary text - easily memorized or small crib sheet")
	assertEqual(t, thing.Items[0].Polls[2].Results[0].Results[1].NumVotes, 9)

	assertEqual(t, thing.Items[0].PlayingTime.Value, 150)
	assertEqual(t, thing.Items[0].MinPlayTime.Value, 60)
	assertEqual(t, thing.Items[0].MaxPlayTime.Value, 150)
	assertEqual(t, thing.Items[0].MinAge.Value, 12)

	assertEqual(t, len(thing.Items[0].Links), 2)
	assertEqual(t, thing.Items[0].Links[0].Type, "boardgamecategory")
	assertEqual(t, thing.Items[0].Links[0].ID, "1234")
	assertEqual(t, thing.Items[0].Links[0].Value, "Civilization")
	assertEqual(t, thing.Items[0].Links[1].Type, "boardgamemechanic")
	assertEqual(t, thing.Items[0].Links[1].ID, "5678")
	assertEqual(t, thing.Items[0].Links[1].Value, "Variable Player Powers")

	assertEqual(t, thing.Items[0].Ratings.UsersRated.Value, 31115)
	assertEqual(t, thing.Items[0].Ratings.Average.Value, float32(8.20954))
	assertEqual(t, thing.Items[0].Ratings.BayesAverage.Value, float32(8.06335))

	assertEqual(t, len(thing.Items[0].Ratings.Ranks), 2)
	assertEqual(t, thing.Items[0].Ratings.BoardGameRank(), "10")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].Type, "subtype")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].ID, "1")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].Name, "boardgame")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].FriendlyName, "Board Game Rank")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].Value, "10")
	assertEqual(t, thing.Items[0].Ratings.Ranks[0].BayesAverage, "8.06335")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].Type, "family")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].ID, "5497")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].Name, "strategygames")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].FriendlyName, "Strategy Game Rank")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].Value, "11")
	assertEqual(t, thing.Items[0].Ratings.Ranks[1].BayesAverage, "8.07124")

	assertEqual(t, thing.Items[0].Ratings.StdDev.Value, float32(1.46947))
	assertEqual(t, thing.Items[0].Ratings.Median.Value, 1)
	assertEqual(t, thing.Items[0].Ratings.Owned.Value, 34484)
	assertEqual(t, thing.Items[0].Ratings.Trading.Value, 487)
	assertEqual(t, thing.Items[0].Ratings.Wanting.Value, 1740)
	assertEqual(t, thing.Items[0].Ratings.Wishing.Value, 10493)
	assertEqual(t, thing.Items[0].Ratings.NumComments.Value, 4786)
	assertEqual(t, thing.Items[0].Ratings.NumWeights.Value, 2232)
	assertEqual(t, thing.Items[0].Ratings.AverageWeight.Value, float32(3.9507))
}
