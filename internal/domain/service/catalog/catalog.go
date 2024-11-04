package service

import (
	"context"
	storage "flavor/internal/adapters/db/mongo/storage"
	"flavor/internal/domain/entity"
	"fmt"
	"net/url"
	"strconv"
)

type CatalogService struct {
	prodStorage storage.ProductStorager

	categoryStorage storage.CategoryStorager
}

func NewCatalogService(productStorage storage.ProductStorager, categoryStorage storage.CategoryStorager) *CatalogService {
	return &CatalogService{
		prodStorage:     productStorage,
		categoryStorage: categoryStorage,
	}
}

func (catalog CatalogService) GetProduct(ctx context.Context, id string) (entity.Product, error) {

	return catalog.prodStorage.GetByID(ctx, id)

}

func (catalog CatalogService) GetProdAllTData(ctx context.Context, limit int, page int, url url.URL) (ProdListTData, error) {

	countEls, err := catalog.GetProdCount(ctx)

	if err != nil {
		return ProdListTData{}, fmt.Errorf("can`t get prodCount: %w", err)
	}

	paginationData := newPaginationData(limit, page, int(countEls), url)

	products, err := catalog.prodStorage.GetAll(ctx, limit, page)
	if err != nil {
		return ProdListTData{}, fmt.Errorf("can`t get products: %w", err)
	}

	return ProdListTData{Products: &products, Pagination: paginationData}, err

}

func (catalog CatalogService) GetProdCount(ctx context.Context) (int64, error) {
	return catalog.prodStorage.GetCount(ctx)
}

func newPaginationData(limit int, page int, countEls int, url url.URL) *paginationData {

	countPages := countEls / limit

	hasPrev := false
	hasNext := false

	if page > 1 {
		hasPrev = true
	}

	if page < countPages {
		hasNext = true
	}

	urlPrev := url
	q := urlPrev.Query()
	q.Del("page")
	q.Add("page", strconv.Itoa(page-1))
	urlPrev.RawQuery = q.Encode()
	urlPrevStr := urlPrev.String()

	urlNext := url
	q = urlNext.Query()
	q.Del("page")
	q.Add("page", strconv.Itoa(page+1))
	urlNext.RawQuery = q.Encode()
	urlNextStr := urlNext.String()

	urlFirst := url
	q = urlFirst.Query()
	q.Del("page")
	q.Add("page", strconv.Itoa(1))
	urlFirst.RawQuery = q.Encode()
	urlFirstStr := urlFirst.String()

	urlLast := url
	q = urlLast.Query()
	q.Del("page")
	q.Add("page", strconv.Itoa(countPages))
	urlLast.RawQuery = q.Encode()
	urlLastStr := urlLast.String()

	urlCurrentStr := url.String()

	pd := &paginationData{
		CurrentPage: page,
		CountEls:    countEls,
		CountPages:  countPages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
		HasPrev:     hasPrev,
		HasNext:     hasNext,
		URLNext:     urlNextStr,
		URLPrev:     urlPrevStr,
		URLCurrent:  urlCurrentStr,
		URLFirst:    urlFirstStr,
		URLLast:     urlLastStr,
	}
	return pd
}

// categories
func (catalog CatalogService) AddCategory(ctx context.Context, category entity.Category) error {
	return catalog.categoryStorage.Create(ctx, category)
}

func (catalog CatalogService) GetCategoriesByPathTData(ctx context.Context, path string) (CategoriesTData, error) {
	category, err := catalog.categoryStorage.GetByPath(ctx, path)
	if err != nil {
		return CategoriesTData{}, err
	}
	subcategories, err := catalog.categoryStorage.GetByParentPath(ctx, path)
	result := CategoriesTData{Category: &category, SubCategories: &subcategories}

	var products = []entity.Product{}
	if len(subcategories) == 0 {
		products, err = catalog.prodStorage.GetByCategoryID(ctx, category.ID, 5, 1)
	}

	result.Products = ProdListTData{Products: &products, Pagination: &paginationData{}}

	return result, err
}

func (catalog CatalogService) GetCategoriesAllMain(ctx context.Context) ([]entity.Category, error) {
	return catalog.categoryStorage.GetAllMain(ctx)
}

func (catalog CatalogService) GetProdTData(ctx context.Context, path string) (ProductTData, error) {

	product, err := catalog.prodStorage.GetByPath(ctx, path)
	if err != nil {
		return ProductTData{}, fmt.Errorf("can`t get products: %w", err)
	}

	return ProductTData{Product: &product}, err

}

func (catalog CatalogService) GetProdByOwnCategoryPathTData(ctx context.Context, catName string, limit int, page int, url url.URL) (ProdListTData, error) {

	products, count, err := catalog.prodStorage.GetByOwnCategoryName(ctx, catName, limit, page)
	if err != nil {
		return ProdListTData{}, fmt.Errorf("can`t get products: %w", err)
	}

	paginationData := newPaginationData(limit, page, count, url)

	return ProdListTData{Products: &products, Pagination: paginationData, URI: &url.Path}, err

}
