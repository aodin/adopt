package robot

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

const baseURL = "http://www.petharbor.com/"

type Pet struct {
	ID        string
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

func UpdateParameter(input, key, value string) (output string) {
	// Parse the input URL
	u, err := url.Parse(input)
	if err != nil {
		return
	}
	// Update the query parameter
	updated := u.Query()
	updated.Set(key, value)
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

		// Separate the name and the id
		raw := strings.TrimSpace(cells[2].Content())
		parts := strings.Split(raw, "(")

		var id string
		var name string

		// If there are no parentheses, then assume the animal has an id
		// and no name
		if len(parts) == 1 {
			id = parts[0]
		} else {
			// combine all but the last part to form the name
			id = parts[len(parts)-1]
			name = strings.TrimSpace(parts[0])

			// Parse the ID
			if id == "" {
				id = strings.TrimSpace(parts[0])
			} else {
				// Remove the trailing parentheses
				id = strings.TrimSpace(id[:len(id)-1])
			}
		}

		u := UpdateParameter(img.Attributes()["src"].String(), "RES", "Detail")
		pet := Pet{
			ID:        id,
			ImageURL:  baseURL + u,
			DetailURL: baseURL + link.Attributes()["href"].String(),
			Type:      strings.TrimSpace(cells[1].Content()),
			Name:      name,
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

func ParsePetsFile(file string) ([]Pet, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	pets, err := ParsePetsHTML(content)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
