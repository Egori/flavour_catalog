package apihandler

import (
	service "flavor/internal/domain/service/catalog"
)

type ProductHandler struct {
	service service.CatalogService
}

func NewProductHandler(service service.CatalogService) *ProductHandler {
	return &ProductHandler{service: service}
}

// func (h *ProductHandler) GetAll(ctx iris.Context) {
// 	limit := 10
// 	offsetStr := ctx.Params().Get("offset")
// 	offset, err := strconv.Atoi(offsetStr)
// 	if err != nil {
// 		httputil.InternalServerErrorJSON(ctx, err, "Offset param is incorrect")
// 		return
// 	}
// 	products, err := h.service.GetAll(ctx, limit, offset)
// 	if err != nil {
// 		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to retrieve all products")
// 		return
// 	}

// 	if products == nil {
// 		// will return "null" if empty, with this "trick" we return "[]" json.
// 		products = make([]entity.Product, 0)
// 	}

// 	ctx.JSON(products)
// }

// func (h *ProductHandler) Get(ctx iris.Context) {
// 	id := ctx.Params().Get("id")

// 	m, err := h.service.GetByID(ctx, id)
// 	if err != nil {
// 		if err == storage.ErrNotFound {
// 			ctx.NotFound()
// 		} else {
// 			httputil.InternalServerErrorJSON(ctx, err, "Server was unable to retrieve product [%s]", id)
// 		}
// 		return
// 	}

// 	ctx.JSON(m)
// }

// func (h *ProductHandler) Add(ctx iris.Context) {
// 	m := new(entity.Product)

// 	err := ctx.ReadJSON(m)
// 	if err != nil {
// 		httputil.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
// 		return
// 	}

// 	err = h.service.Create(ctx, *m)
// 	if err != nil {
// 		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to create a product")
// 		return
// 	}

// 	ctx.StatusCode(iris.StatusCreated)
// 	ctx.JSON(m)
// }

// func (h *ProductHandler) Update(ctx iris.Context) {
// 	id := ctx.Params().Get("id")

// 	var m storage.Product
// 	err := ctx.ReadJSON(&m)
// 	if err != nil {
// 		httputil.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
// 		return
// 	}

// 	err = h.service.Update(nil, id, m)
// 	if err != nil {
// 		if err == store.ErrNotFound {
// 			ctx.NotFound()
// 			return
// 		}
// 		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to update product [%s]", id)
// 		return
// 	}
// }

// func (h *ProductHandler) Delete(ctx iris.Context) {
// 	id := ctx.Params().Get("id")

// 	err := h.service.Delete(nil, id)
// 	if err != nil {
// 		if err == store.ErrNotFound {
// 			ctx.NotFound()
// 			return
// 		}
// 		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to delete product [%s]", id)
// 		return
// 	}
// }
