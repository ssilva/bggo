package bggo

import (
	"encoding/xml"
	"testing"
)

func TestPlayResponse(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<plays username="Silvast" userid="1234567" total="100" page="1" termsofuse="https://example.com">
		<play id="12345" date="1999-09-09" quantity="1" length="10" incomplete="1" nowinstats="1" location="Montreal">
			<item name="Any Game" objecttype="thing" objectid="123">
				<subtypes>
					<subtype value="boardgame" />
					<subtype value="boardgameimplementation" />
				</subtypes>
			</item>
			<players>
				<player username="aaa" userid="1234" startposition="2" color="red" score="999" new="1" rating="5" win="1"/>
			</players>
		</play>
	</plays>`

	plays := &PlaysResponse{}
	err := xml.Unmarshal([]byte(data), plays)

	assertNil(t, err)

	assertEqual(t, plays.Username, "Silvast")
	assertEqual(t, plays.UserID, "1234567")
	assertEqual(t, plays.Total, 100)
	assertEqual(t, plays.Page, 1)
	assertEqual(t, plays.TermsOfUse, "https://example.com")
	assertEqual(t, len(plays.Plays), 1)

	assertEqual(t, plays.Plays[0].ID, "12345")
	assertEqual(t, plays.Plays[0].Date, "1999-09-09")
	assertEqual(t, plays.Plays[0].Quantity, 1)
	assertEqual(t, plays.Plays[0].Length, 10)
	assertEqual(t, plays.Plays[0].Incomplete, true)
	assertEqual(t, plays.Plays[0].NowInStats, true)
	assertEqual(t, plays.Plays[0].Location, "Montreal")
	assertEqual(t, len(plays.Plays[0].Items), 1)

	assertEqual(t, plays.Plays[0].Items[0].Name, "Any Game")
	assertEqual(t, plays.Plays[0].Items[0].ObjectType, "thing")
	assertEqual(t, plays.Plays[0].Items[0].ObjectID, "123")
	assertEqual(t, len(plays.Plays[0].Items[0].Subtypes), 2)

	assertEqual(t, plays.Plays[0].Players[0].Username, "aaa")
	assertEqual(t, plays.Plays[0].Players[0].UserID, "1234")
	assertEqual(t, plays.Plays[0].Players[0].StartPosition, 2)
	assertEqual(t, plays.Plays[0].Players[0].Color, "red")
	assertEqual(t, plays.Plays[0].Players[0].Score, 999)
	assertEqual(t, plays.Plays[0].Players[0].New, true)

	assertEqual(t, plays.Plays[0].Items[0].Subtypes[0].Name, "boardgame")
	assertEqual(t, plays.Plays[0].Items[0].Subtypes[1].Name, "boardgameimplementation")
}
