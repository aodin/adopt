package robot

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "http://www.petharbor.com/"

type Pet struct {
	ID        int64
	Name      string
	Type      string
	Gender    string
	Color     string
	Breed     string
	Age       string
	Location  string
	DetailURL string // TODO Or URL?
	ImageURL  string // TODO Or URL?
}

func UpdateImageURL(input string) (output string) {
	// Parse the input URL
	u, err := url.Parse(input)
	if err != nil {
		return
	}
	// Update the query parameter RES
	updated := u.Query()
	updated.Set("RES", "Detail")
	u.RawQuery = updated.Encode()
	return u.String()
}

func ParsePetsHTML(content []byte) (pets []Pet, err error) {
	dom, err := gokogiri.ParseHtml(content)
	if err != nil {
		err = fmt.Errorf("Error parsing HTML: %s")
		return
	}

	// TODO This will work if there is only one class!
	q := "//table[@class='ResultsTable']//tr"
	rows, err := dom.Search(q)
	if err != nil {
		return
	}

	// TODO For multiple classes:
	// q := "//table[contains(@class, 'Test')]//tr"

	if len(rows) < 2 {
		err = fmt.Errorf("Insufficient number of rows: %d", len(rows))
		return
	}

	// Skip the first row in the table body - it is a header
	var cells []xml.Node
	for i, row := range rows[1:] {
		// Select the cells from the row
		cells, err = row.Search("./td")
		if err != nil {
			err = fmt.Errorf("Error finding cells in row %d: %s", i+1, err)
			return
		}

		if len(cells) < 8 {
			err = fmt.Errorf("Row %d does not have at least 8 cells", i)
			return
		}

		// Get the link
		link := cells[0].FirstChild()
		img := link.FirstChild()

		// TODO separate the name from the id
		raw := strings.TrimSpace(cells[2].Content())
		parts := strings.SplitN(raw, "(", 2)

		// Parse the ID
		rawID := parts[1]
		if rawID != "" {
			rawID = rawID[:len(rawID)-1]
		}

		var id int64
		id, err = strconv.ParseInt(rawID, 10, 64)
		if err != nil {
			err = fmt.Errorf("Cannot parse the ID in row %d: %s", i, rawID)
			return
		}

		pet := Pet{
			ID:        id,
			ImageURL:  UpdateImageURL(img.Attributes()["src"].String()),
			DetailURL: link.Attributes()["href"].String(),
			Type:      strings.TrimSpace(cells[1].Content()),
			Name:      strings.TrimSpace(parts[0]),
			Gender:    strings.TrimSpace(cells[3].Content()),
			Color:     strings.TrimSpace(cells[4].Content()),
			Breed:     strings.TrimSpace(cells[5].Content()),
			Age:       strings.TrimSpace(cells[6].Content()),
			Location:  strings.TrimSpace(cells[7].Content()),
		}
		pets = append(pets, pet)
	}
	return
}