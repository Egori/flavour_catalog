package app

import (
	"context"
	storage "flavor/internal/adapters/db/mongo/storage"
	"flavor/internal/config"
	service "flavor/internal/domain/service/catalog"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type serviceProvider struct {
	catalogservice *service.CatalogService

	productStorage storage.ProductStorager

	categoryStorage storage.CategoryStorager

	mongoDB *mongo.Database
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) MongoDB() *mongo.Database {
	dbName := config.DBName
	clientOptions := options.Client().ApplyURI(config.DSN)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	if s.mongoDB == nil {
		s.mongoDB = client.Database(dbName)
	}

	return s.mongoDB
}

func (s *serviceProvider) ProductStorage() storage.ProductStorager {

	db := s.MongoDB()
	productsCollection := db.Collection("products")
	if s.productStorage == nil {
		s.productStorage = storage.NewproductStorage(productsCollection)
	}

	return s.productStorage
}

func (s *serviceProvider) CategoryStorage() storage.CategoryStorager {

	db := s.MongoDB()
	productsCollection := db.Collection("categories")
	if s.categoryStorage == nil {
		s.categoryStorage = storage.NewcategoryStorage(productsCollection)
	}

	return s.categoryStorage
}

func (s *serviceProvider) ProductService() service.CatalogService {

	if s.catalogservice == nil {
		s.catalogservice = service.NewCatalogService(s.ProductStorage(), s.CategoryStorage())
	}

	return *s.catalogservice
}

func (s *serviceProvider) InitServices() {
	//_ = s.MovieService()
	_ = s.ProductService()
}
