package htmlHandler

import (
	"context"
	catalogservice "flavor/internal/domain/service/catalog"
	"net/http"
	"net/url"
	"strconv"

	//"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

type catalogService interface {
	GetProdAllTData(ctx context.Context, limit int, offset int, url url.URL) (catalogservice.ProdListTData, error)
	GetProdCount(ctx context.Context) (int64, error)
	GetCategoriesByPathTData(ctx context.Context, path string) (catalogservice.CategoriesTData, error)
	GetProdTData(ctx context.Context, path string) (catalogservice.ProductTData, error)
	GetProdByOwnCategoryPathTData(ctx context.Context, catName string, limit int, page int, url url.URL) (catalogservice.ProdListTData, error)
}

type CatalogHandler struct {
	catalogService catalogService
}

func NewCatalogHandler(catalog catalogService) *CatalogHandler {
	return &CatalogHandler{catalogService: catalog}
}

func (h *CatalogHandler) GetTest(ctx echo.Context) error {

	err := ctx.Render(http.StatusOK, "main-test.html", "")

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetMain(ctx echo.Context) error {
	limit := 3
	pageStr := ctx.Param("page")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	templateData, err := h.catalogService.GetProdAllTData(ctx.Request().Context(), limit, pageNum, *ctx.Request().URL)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "prod_page_list.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetProdListAll(ctx echo.Context) error {
	limit := 3
	pageStr := ctx.Param("page")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	templateData, err := h.catalogService.GetProdAllTData(ctx.Request().Context(), limit, pageNum, *ctx.Request().URL)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "prod_list.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetProdPage(ctx echo.Context) error {

	path := ctx.Param("path")
	templateData, err := h.catalogService.GetProdTData(ctx.Request().Context(), path)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "prod_layout.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetCategoryPage(ctx echo.Context) error {

	path := ctx.Param("path")

	templateData, err := h.catalogService.GetCategoriesByPathTData(ctx.Request().Context(), path)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "category_page.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetCatalogPageByCategoryOwn(ctx echo.Context) error {
	limit := 10

	pageNum := 1

	path := ctx.Param("path")

	pageStr := ctx.QueryParam("page")

	if pageStr != "" {
		return h.GetCatalogListByCategoryOwn(ctx)
	}

	templateData, err := h.catalogService.GetProdByOwnCategoryPathTData(ctx.Request().Context(), path, limit, pageNum, *ctx.Request().URL)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "prod_list_page.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}

func (h *CatalogHandler) GetCatalogListByCategoryOwn(ctx echo.Context) error {

	limit := 10
	pageStr := ctx.QueryParam("page")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	path := ctx.Param("path")

	templateData, err := h.catalogService.GetProdByOwnCategoryPathTData(ctx.Request().Context(), path, limit, pageNum, *ctx.Request().URL)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	err = ctx.Render(http.StatusOK, "prod_list.html", templateData)

	if err != nil {
		ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
		return err
	}

	return err
}
