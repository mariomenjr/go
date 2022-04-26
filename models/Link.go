package models

import (
	"github.com/Kamva/mgm"
)

type Link struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"name" bson:"title"`
	Url              string `json:"url" bson:"url"`
	Id               string `json:"id" bson:"id"`
}

func NewLink(title string, url string) *Link {
	return &Link{
		Title: title,
		Url:   url,
	}
}

func (model *Link) Created() error {

	model.Id = model.ID.Hex() // Gotta convert this one
	return mgm.Coll(model).Update(model)
}
