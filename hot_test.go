package bggo

import (
	"encoding/xml"
	"testing"
)

func TestHotResponse(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<items termsofuse="https://example.com">
		<item id="1234" rank="1">
			<thumbnail value="thumbnail1.jpg"/>
			<name value="Game 1"/>
			<yearpublished value="2019" />
		</item>
		<item id="5678" rank="2">
			<thumbnail value="thumbnail2.jpg"/>
			<name value="Game 2"/>
			<yearpublished value="2018" />
		</item>
	</items>`

	hot := &HotResponse{}
	err := xml.Unmarshal([]byte(data), hot)

	assertNil(t, err)

	assertEqual(t, hot.TermsOfUse, "https://example.com")
	assertEqual(t, len(hot.Items), 2)
	assertEqual(t, hot.Items[0].ID, "1234")
	assertEqual(t, hot.Items[0].Rank, 1)
	assertEqual(t, hot.Items[0].Thumbnail.Value, "thumbnail1.jpg")
	assertEqual(t, hot.Items[0].Name.Value, "Game 1")
	assertEqual(t, hot.Items[0].YearPublished.Value, "2019")
	assertEqual(t, hot.Items[1].ID, "5678")
	assertEqual(t, hot.Items[1].Rank, 2)
	assertEqual(t, hot.Items[1].Thumbnail.Value, "thumbnail2.jpg")
	assertEqual(t, hot.Items[1].Name.Value, "Game 2")
	assertEqual(t, hot.Items[1].YearPublished.Value, "2018")

}
