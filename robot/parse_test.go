package robot

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("./example.html")
	if err != nil {
		t.Fatal(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	pets, err := ParsePetsHTML(content)
	if err != nil {
		t.Fatal(err)
	}
	if len(pets) != 26 {
		t.Fatalf("Unexpected number of ParsePetsHTML results: %d", len(pets))
	}

	tippy := pets[0]
	if tippy.ID != 10049878 {
		t.Errorf("Unexpected ID: %d", tippy.ID)
	}
	if tippy.Name != "Tippy" {
		t.Errorf("Unexpected name: %s", tippy.Name)
	}
	if tippy.Type != "small anim" {
		t.Errorf("Unexpected type: %s", tippy.Type)
	}
	if tippy.Gender != "Female" {
		t.Errorf("Unexpected gender: %s", tippy.Gender)
	}
	if tippy.Color != "" {
		t.Errorf("Unexpected color: %s", tippy.Color)
	}
	if tippy.Breed != "Rat" {
		t.Errorf("Unexpected breed: %s", tippy.Breed)
	}
	if tippy.Age != "1 year, 8 months old" {
		t.Errorf("Unexpected age: %s", tippy.Age)
	}
	if tippy.Location != "Apple Wood Rescue, Inc." {
		t.Errorf("Unexpected location: %s", tippy.Location)
	}
	if tippy.DetailURL != "http://www.petharbor.com/detail.asp?ID=10049878&LOCATION=82294&searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=100&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist='83615','80454','79367','82294','77298','84657','69972','84715','79780','77803','76338','85330','76065','78397','86214','85252','74805','73867','82242','81793','72856','73086','82431','86406','74867','83241','72907','74328','86813','71436','82755','82206','76134'&atype=&where=type_OO" {
		t.Errorf("Unexpected detail URL: %s", tippy.DetailURL)
	}
	if tippy.ImageURL != "http://www.petharbor.com/get_image.asp?ID=10049878&LOCATION=82294&RES=Detail" {
		t.Errorf("Unexpected image URL: %s", tippy.ImageURL)
	}
}

func TestUpdateImageURL(t *testing.T) {
	orig := "get_image.asp?RES=thumb&ID=9979224&LOCATION=74328"
	// Golang's Values Encode() will sort the keys!
	// expected :=  "get_image.asp?RES=Detail&ID=9979224&LOCATION=74328"
	expected := "get_image.asp?ID=9979224&LOCATION=74328&RES=Detail"
	result := UpdateImageURL(orig)
	if expected != result {
		t.Fatalf("Unexpected image URL: %s", result)
	}
}
