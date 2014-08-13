package robot

import (
	"github.com/aodin/aspect"
	"time"
)

var Pets = aspect.Table("pets",
	aspect.Column("id", aspect.Integer{PrimaryKey: true}),
	aspect.Column("name", aspect.String{}),
	aspect.Column("type", aspect.String{}),
	aspect.Column("gender", aspect.String{}),
	aspect.Column("color", aspect.String{}),
	aspect.Column("breed", aspect.String{}),
	aspect.Column("age", aspect.String{}),
	aspect.Column("location", aspect.String{}),
	aspect.Column("detail_url", aspect.String{}),
	aspect.Column("image_url", aspect.String{}),
	aspect.Column("added", aspect.Timestamp{}),
	aspect.Column("removed", aspect.Timestamp{}),
)

type PetWithTimestamp struct {
	ID        int64
	Name      string
	Type      string
	Gender    string
	Color     string
	Breed     string
	Age       string
	Location  string
	DetailURL string
	ImageURL  string
	Added     time.Time
	Removed   *time.Time // A pointer is used so the zero init is nil
}

func PWTFromPet(i Pet) (o PetWithTimestamp) {
	o.ID = i.ID
	o.Name = i.Name
	o.Type = i.Type
	o.Gender = i.Gender
	o.Color = i.Color
	o.Breed = i.Breed
	o.Age = i.Age
	o.Location = i.Location
	o.DetailURL = i.DetailURL
	o.ImageURL = i.ImageURL
	return
}
