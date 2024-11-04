package app

import (
	"context"
	"flag"
	"flavor/internal/config"
	"flavor/internal/domain/entity"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatalogGenerator struct {
	serviceProvider *serviceProvider
}

func (cg *CatalogGenerator) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		cg.initConfig,
		cg.initServiceProvider,
		cg.GenerateProducts,
		//a.UpdateRandProd,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cg *CatalogGenerator) initServiceProvider(_ context.Context) error {
	cg.serviceProvider = newServiceProvider()
	cg.serviceProvider.InitServices()
	return nil
}

func (cg *CatalogGenerator) initConfig(_ context.Context) error {
	envFileName := ".env"

	flagset := flag.CommandLine
	flagset.StringVar(&envFileName, "env", envFileName, "the env file which web app will use to extract its environment variables")
	flagset.Parse(os.Args[1:])

	config.Load(envFileName)
	return nil
}

func (cg *CatalogGenerator) createProduct() entity.Product {
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

func (cg *CatalogGenerator) UpdateRandProd(ctx context.Context) error {
	prodStor := cg.serviceProvider.productStorage

	for offset := 0; offset < 1000; offset++ {
		products, err := prodStor.GetAll(ctx, 1000, offset)
		if len(products) == 0 || err != nil {
			return err
		}
		for _, prod := range products {
			newProd := cg.createProduct()
			prod.Path = newProd.Path
			prodStor.Update(ctx, prod)
		}
	}

	return nil

}

func (cg *CatalogGenerator) GenerateProducts(ctx context.Context) error {

	prodStor := cg.serviceProvider.productStorage

	println("generate started...")

	startTime := time.Now()

	for j := 0; j < 100000; j++ {
		product := cg.createProduct()
		err := prodStor.Create(ctx, product)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	// Засекаем конечное время
	endTime := time.Now()

	// Вычисляем разницу между начальным и конечным временем
	elapsedTime := endTime.Sub(startTime)

	// Выводим результат
	fmt.Printf("Операция заняла %s\n", elapsedTime)

	println("generate success")

	return nil

	//fmt.Printf("%+v\n", product)
}
