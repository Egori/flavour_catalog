package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID               primitive.ObjectID       `json:"_id" bson:"_id" fake:"-"`
	Name             string                   `json:"name" bson:"name" fake:"{sentence:3}"`
	Description      string                   `json:"description" bson:"description" fake:"{paragraph:2,3,10,\n}"`
	Path             string                   `json:"path" bson:"path" fake:"{year}/{month}"`
	ParentPath       string                   `json:"parentPath" bson:"parentPath" fake:"{year}/{month}"`
	PathName         string                   `json:"pathName" bson:"pathName" fake:"{year}/{month}"`
	Parentid         primitive.ObjectID       `json:"parentId" bson:"parentId" `
	Properties       []CategoryProperty       `json:"properties" bson:"properties"`
	Stringproperties []CategoryStringProperty `json:"strProps" bson:"strProps"`
	Image            string                   `json:"image" bson:"image" fake:"{year}/{month}"`
}

type CategoryProperty struct {
	Name  string `json:"name" bson:"name" fake:"{beerstyle}"`
	Value string `json:"value" bson:"value" fake:"{beername}"`
	Path  string `json:"path" bson:"path" fake:"{year}/{month}"`
	Image string `json:"image" bson:"image"`
}

type CategoryStringProperty struct {
	Name  string `json:"name" bson:"name" fake:"{carmaker}"`
	Value string `json:"value" bson:"value" fake:"{carmodel}"`
}

type CategoryRange struct {
	Value  float32 `json:"value" bson:"value" fake:"{float32range:1,5}"`
	Weight int32   `json:"weight" bson:"weight" fake:"{number:0,1000}"`
}
