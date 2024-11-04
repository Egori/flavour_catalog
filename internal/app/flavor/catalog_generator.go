package app

import (
	"flavor/internal/domain/entity"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createProduct() entity.Product {
	product := new(entity.Product)

	//err := faker.FakeData(&product)
	err := gofakeit.Struct(&product)

	if err != nil {
		fmt.Println(err)
	}

	product.ID = primitive.NewObjectID()

	for i := range product.CategoryIDs {
		product.CategoryIDs[i] = primitive.NewObjectID()
	}

	return *product
}
