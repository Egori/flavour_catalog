package storage

import (
	"context"
	"flavor/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryStorager interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Category, error)
	GetByID(ctx context.Context, id string) (entity.Category, error)
	GetByPath(ctx context.Context, path string) (entity.Category, error)
	GetByParentPath(ctx context.Context, path string) ([]entity.Category, error)

	GetAllMain(ctx context.Context) ([]entity.Category, error)

	Create(ctx context.Context, m entity.Category) error
	Update(ctx context.Context, id string, m entity.Category) error
	Delete(ctx context.Context, id string) error
}

type CategoryStorage struct {
	C *mongo.Collection
}

func NewcategoryStorage(collection *mongo.Collection) *CategoryStorage {

	return &CategoryStorage{C: collection}
}
func (s *CategoryStorage) GetAllMain(ctx context.Context) ([]entity.Category, error) {

	filter := bson.M{
		// "parentId": bson.M{
		// 	"$eq":   nil,
		// 	"$type": 10, // Код типа null в BSON
		// },

		"categoryIds": bson.M{
			"$eq": primitive.NilObjectID,
			//"$type": 7, // Код типа null в BSON
		},
	}

	cur, err := s.C.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []entity.Category

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem entity.Category
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}
func (s *CategoryStorage) GetAll(ctx context.Context, limit int, offset int) ([]entity.Category, error) {

	var findoptions = new(options.FindOptions)

	if limit > 0 {
		findoptions.SetLimit(int64(limit))
		findoptions.SetSkip(int64(limit * offset))
	}
	// Note:
	// The mongodb's go-driver's docs says that you can pass `nil` to "find all" but this gives NilDocument error,
	// probably it's a bug or a documentation's mistake, you have to pass `bson.D{}` instead.
	cur, err := s.C.Find(ctx, bson.M{}, findoptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []entity.Category

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		//	elem := bson.D{}
		var elem entity.Category
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *CategoryStorage) GetByPath(ctx context.Context, path string) (entity.Category, error) {

	filter := bson.M{"path": path}

	var category entity.Category

	err := s.C.FindOne(ctx, filter).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return category, ErrNotFound
	}
	return category, err
}

func (s *CategoryStorage) GetByParentPath(ctx context.Context, path string) ([]entity.Category, error) {

	filter := bson.M{"parentPath": path}

	cur, err := s.C.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []entity.Category

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem entity.Category
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *CategoryStorage) GetByID(ctx context.Context, id string) (entity.Category, error) {
	var category entity.Category
	filter, err := matchID(id)
	if err != nil {
		return category, err
	}

	err = s.C.FindOne(ctx, filter).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return category, ErrNotFound
	}
	return category, err
}

func (s *CategoryStorage) Create(ctx context.Context, m entity.Category) error {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}

	_, err := s.C.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStorage) Update(ctx context.Context, id string, m entity.Category) error {
	filter, err := matchID(id)
	if err != nil {
		return err
	}

	elem := bson.D{}

	if m.Name != "" {
		elem = append(elem, bson.E{Key: "name", Value: m.Name})
	}

	if m.Description != "" {
		elem = append(elem, bson.E{Key: "description", Value: m.Description})
	}

	// TODO add other rows

	update := bson.D{
		{Key: "$set", Value: elem},
	}

	_, err = s.C.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (s *CategoryStorage) Delete(ctx context.Context, id string) error {
	filter, err := matchID(id)
	if err != nil {
		return err
	}
	_, err = s.C.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
		return err
	}

	return nil
}
