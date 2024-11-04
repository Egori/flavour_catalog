package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID                    primitive.ObjectID             `json:"_id" bson:"_id" fake:"-"`
	Path                  string                         `json:"path" bson:"path" `
	Name                  string                         `json:"name" bson:"name" fake:"{sentence:3}"`
	Description           string                         `json:"description" bson:"description" fake:"{paragraph:2,3,10,\n}"`
	CategoryIDs           []primitive.ObjectID           `json:"categoryIds" bson:"categoryIds" `
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
	Sex         string                  `json:"sex" bson:"sex" fake:"{randomstring:[men,women,unisex]}"`
	ProductType string                  `json:"prodType" bson:"prodType" fake:"{randomstring:[Духи/Экстракт,Туалетная вода,Парфюмерная вода,Одеколон, Масляные духи, Ароматная дымка, Твердые духи, Дневные духи]}"`
	BrandStr    string                  `json:"brandStr" bson:"brandStr" fake:"{carmaker}"`
}

type ProductCategoryProperty struct {
	Name  string `json:"name" bson:"name" fake:"{beerstyle}"`
	Value string `json:"value" bson:"value" fake:"{beername}"`
	Path  string `json:"path" bson:"path" fake:"{color}"`
	Image string `json:"image" bson:"image"`
}

type ProductStringProperty struct {
	Name  string `json:"name" bson:"name" fake:"{carmaker}"`
	Value string `json:"value" bson:"value" fake:"{carmodel}"`
}

type ProductMultipleRangeProperty struct {
	Name   string           `json:"name" bson:"name" fake:"{carmaker}"`
	Values map[string]int32 `json:"values" bson:"values"`
}

type ProductNoteRangeProperty struct {
	Name  string  `json:"name" bson:"name" fake:"{carmaker}"`
	Range float32 `json:"range" bson:"range" fake:"{float32range:0,1}"`
	Path  string  `json:"path" bson:"path" fake:"{year}/{month}"`
	Image string  `json:"image" bson:"image" `
}

type ProductRange struct {
	Value  float32 `json:"value" bson:"value" fake:"{float32range:1,5}"`
	Weight int32   `json:"weight" bson:"weight" fake:"{number:0,1000}"`
}

type SexType string

const (
	Men    SexType = "man"
	Women  SexType = "woman"
	Unisex SexType = "unisex"
)
