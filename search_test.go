package bggo

import (
	"encoding/xml"
	"testing"
)

func TestSearchResponse(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<items total="6" termsofuse="https://example.com">
		<item type="boardgame" id="1234">
			<name type="primary" value="Game 1"/>			
			<yearpublished value="2012" />
		</item>
		<item type="boardgame" id="5678">
			<name type="primary" value="Game 2"/>			
			<yearpublished value="2013" />
		</item>
	</items>`

	search := &SearchResponse{}
	err := xml.Unmarshal([]byte(data), search)

	assertNil(t, err)

	assertEqual(t, search.Total, 6)
	assertEqual(t, search.TermsOfUse, "https://example.com")
	assertEqual(t, len(search.Items), 2)

	assertEqual(t, search.Items[0].Type, "boardgame")
	assertEqual(t, search.Items[0].ID, "1234")
	assertEqual(t, search.Items[0].Name.Type, "primary")
	assertEqual(t, search.Items[0].Name.Value, "Game 1")
	assertEqual(t, search.Items[0].YearPublished.Value, "2012")

	assertEqual(t, search.Items[1].Type, "boardgame")
	assertEqual(t, search.Items[1].ID, "5678")
	assertEqual(t, search.Items[1].Name.Type, "primary")
	assertEqual(t, search.Items[1].Name.Value, "Game 2")
	assertEqual(t, search.Items[1].YearPublished.Value, "2013")
}
