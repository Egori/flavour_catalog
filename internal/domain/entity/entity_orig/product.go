package entity_orig

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID                    primitive.ObjectID             `json:"_id" bson:"_id"` /* you need the bson:"_id" to be able to retrieve with ID filled */
	Name                  string                         `json:"name" bson:"name"`
	Description           string                         `json:"description" bson:"description"`
	CategoryIDs           []primitive.ObjectID           `json:"categoryIds" bson:"categoryIds"`
	CategoryProperties    []ProductCategoryProperty      `json:"categoryProps" bson:"categoryProps"`
	StringProperties      []ProductStringProperty        `json:"strProps" bson:"strProps"`
	MultipleRangeProperty []ProductMultipleRangeProperty `json:"multiRangeProps" bson:"multiRangeProps"`
	HighNotes             []ProductNoteRangeProperty     `json:"highNotes" bson:"highNotes"`
	MiddleNotes           []ProductNoteRangeProperty     `json:"middleNotes" bson:"middleNotes"`
	LowNotes              []ProductNoteRangeProperty     `json:"lowNotes" bson:"lowNotes"`
	BasicNotes            []ProductNoteRangeProperty     `json:"basicNotes" bson:"basicNotes"`

	Brand       ProductCategoryProperty `json:"brand" bson:"brand"`
	Author      ProductCategoryProperty `json:"author" bson:"author"`
	Range       ProductRange            `json:"range" bson:"range"`
	Sex         SexType                 `json:"sex" bson:"sex"`
	ProductType string                  `json:"prodType" bson:"prodType"`
}

type ProductCategoryProperty struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
	Path  string `json:"path" bson:"path"`
	Image string `json:"image" bson:"image"`
}

type ProductStringProperty struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type ProductMultipleRangeProperty struct {
	Name   string           `json:"name" bson:"name"`
	Values map[string]int32 `json:"values" bson:"values"`
}

type ProductNoteRangeProperty struct {
	Name  string  `json:"name" bson:"name"`
	Range float32 `json:"range" bson:"range"`
	Path  string  `json:"path" bson:"path"`
	Image string  `json:"image" bson:"image"`
}

type ProductRange struct {
	Value  float32 `json:"value" bson:"value"`
	Weight int32   `json:"weight" bson:"weight"`
}

type SexType string

const (
	Men    SexType = "man"
	Women  SexType = "woman"
	Unisex SexType = "unisex"
)
