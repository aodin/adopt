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

// TODO add URL value "where"
const denverURL = `http://www.petharbor.com/results.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=1000&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist=%27ARAP%27,%27AURO%27,%27DNVR%27,%27DDFL%27,%2783615%27,%2780454%27,%2779367%27,%2782294%27,%2777298%27,%2784657%27,%2769972%27,%2784715%27,%2779780%27,%2777803%27,%2776338%27,%2785330%27,%2776065%27,%2778397%27,%2786214%27,%2785252%27,%2774805%27,%2773867%27,%2782242%27,%2781793%27,%2772856%27,%2773086%27,%2782431%27,%2786406%27,%2774867%27,%2783241%27,%2772907%27,%2774328%27,%2786813%27,%2771436%27,%2782755%27,%2782206%27,%2776134%27&atype=&PAGE=1`

var whereTypes = []string{"type_OO", "type_CAT", "type_DOG"}

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
			ImageURL:  baseURL + UpdateImageURL(img.Attributes()["src"].String()),
			DetailURL: baseURL + link.Attributes()["href"].String(),
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
