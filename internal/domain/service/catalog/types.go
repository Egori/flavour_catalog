package service

import "flavor/internal/domain/entity"

type paginationData struct {
	CurrentPage int
	CountPages  int
	CountEls    int
	PrevPage    int
	NextPage    int

	URLNext    string
	URLPrev    string
	URLCurrent string
	URLFirst   string
	URLLast    string

	HasPrev bool
	HasNext bool
}

type ProdListTData struct {
	Products   *[]entity.Product
	Pagination *paginationData
	URI        *string
}

type CategoriesTData struct {
	Category      *entity.Category
	SubCategories *[]entity.Category
	Products      ProdListTData
}

type ProductTData struct {
	Product *entity.Product
}
