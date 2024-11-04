package storage

import (
	"context"
	"flavor/internal/domain/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductStorager interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Product, error)
	GetByID(ctx context.Context, id string) (entity.Product, error)
	GetByCategoryID(ctx context.Context, id primitive.ObjectID, limit int, offset int) ([]entity.Product, error)
	GetByOwnCategoryName(ctx context.Context, catName string, limit int, offset int) ([]entity.Product, int, error)
	GetByPath(ctx context.Context, path string) (entity.Product, error)
	Create(ctx context.Context, m entity.Product) error
	Update(ctx context.Context, m entity.Product) error
	Delete(ctx context.Context, id string) error
	GetCount(ctx context.Context) (int64, error)
}

type ProductStorage struct {
	C        *mongo.Collection
	countEls *int64
}

func NewproductStorage(collection *mongo.Collection) *ProductStorage {

	return &ProductStorage{C: collection}
}

func (s *ProductStorage) GetAll(ctx context.Context, limit int, offset int) ([]entity.Product, error) {

	var findoptions = new(options.FindOptions)

	if limit > 0 {
		findoptions.SetLimit(int64(limit))
		findoptions.SetSkip(int64(limit * offset))
	}

	cur, err := s.C.Find(ctx, bson.M{}, findoptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []entity.Product

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		//	elem := bson.D{}
		var elem entity.Product
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *ProductStorage) GetByCategoryID(ctx context.Context, id primitive.ObjectID, limit int, offset int) ([]entity.Product, error) {

	var findoptions = new(options.FindOptions)

	// if limit > 0 {
	// 	findoptions.SetLimit(int64(limit))
	// 	findoptions.SetSkip(int64(limit * offset))
	// }

	cur, err := s.C.Find(ctx, bson.M{"categoryIds": id}, findoptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []entity.Product

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		//	elem := bson.D{}
		var elem entity.Product
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *ProductStorage) GetByOwnCategoryName(ctx context.Context, catName string, limit int, offset int) ([]entity.Product, int, error) {

	var findoptions = new(options.FindOptions)

	if limit > 0 {
		findoptions.SetLimit(int64(limit))
		findoptions.SetSkip(int64(limit * offset))
	}

	cur, err := s.C.Find(ctx, bson.M{"categoryProps.path": catName}, findoptions)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	count64, _ := s.C.CountDocuments(ctx, bson.M{"categoryProps.path": catName})
	count := int(count64)

	var results []entity.Product

	for cur.Next(ctx) {
		if err = cur.Err(); err != nil {
			return nil, 0, err
		}

		//	elem := bson.D{}
		var elem entity.Product
		err = cur.Decode(&elem)
		if err != nil {
			return nil, 0, err
		}

		results = append(results, elem)
	}

	return results, count, nil
}

func (s *ProductStorage) GetByID(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product
	filter, err := matchID(id)
	if err != nil {
		return product, err
	}

	err = s.C.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return product, ErrNotFound
	}
	return product, err
}

func (s *ProductStorage) GetByPath(ctx context.Context, path string) (entity.Product, error) {
	var product entity.Product
	filter := bson.M{"path": path}

	err := s.C.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return product, ErrNotFound
	}
	return product, err
}

func (s *ProductStorage) GetCount(ctx context.Context) (int64, error) {

	if s.countEls != nil {
		return *s.countEls, nil
	}
	filter := bson.M{}

	count, err := s.C.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	s.countEls = &count

	return count, err
}

func (s *ProductStorage) Create(ctx context.Context, m entity.Product) error {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}

	_, err := s.C.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStorage) Update(ctx context.Context, m entity.Product) error {
	// filter, err := matchID(id)
	// if err != nil {
	// 	return err
	// }

	filter := bson.D{{"_id", m.ID}}

	elem := bson.D{}

	// if m.Name != "" {
	// 	elem = append(elem, bson.E{Key: "name", Value: m.Name})
	// }

	// if m.Description != "" {
	// 	elem = append(elem, bson.E{Key: "description", Value: m.Description})
	// }

	// if len(m.CategoryProperties) > 0 {
	// 	elem = append(elem, bson.E{Key: "categoryProps", Value: m.CategoryProperties})
	// }

	if m.Path != "" {
		elem = append(elem, bson.E{Key: "path", Value: m.Path})
	}

	// TODO add other rows

	update := bson.D{
		{Key: "$set", Value: elem},
	}

	_, err := s.C.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (s *ProductStorage) Delete(ctx context.Context, id string) error {
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
