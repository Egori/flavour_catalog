package apihandler

import (
	"flavor/internal/domain/entity"
	service "flavor/internal/domain/service/catalog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service service.CatalogService
}

func NewCategoryHandler(service service.CatalogService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAllMain(ctx echo.Context) error {

	category, err := h.service.GetCategoriesAllMain(ctx.Request().Context())
	if err != nil {
		ctx.Error(err)
		return err
	}

	if category == nil {
		// will return "null" if empty, with this "trick" we return "[]" json.
		category = make([]entity.Category, 0)
	}

	err = ctx.JSON(http.StatusOK, category)
	if err != nil {
		ctx.Error(err)
		return err
	}

	return err
}

func (h *CategoryHandler) Add(ctx echo.Context) error {

	var category entity.Category
	if err := ctx.Bind(&category); err != nil {
		ctx.Error(err)
		return err
	}
	// if err := ctx.Validate(&category); err != nil {
	// 	ctx.Error(err)
	// 	return err
	// }

	err := h.service.AddCategory(ctx.Request().Context(), category)
	if err != nil {
		ctx.Error(err)
		return err
	}

	err = ctx.JSON(http.StatusOK, category)
	if err != nil {
		ctx.Error(err)
		return err
	}
	ctx.JSON(http.StatusOK, "Ok")
	return err
}
