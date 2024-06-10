package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCategories(ctx *gin.Context) {
	categoriesWithSubCategories, err := h.ds.GetCategoriesWithSubCategories(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, categoriesWithSubCategories)
}

func (h *Handler) getLanguages(ctx *gin.Context) {
	languages, err := h.ds.GetLanguages(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, languages)
}
func (h *Handler) getMetasCount(ctx *gin.Context) {
	metasCount, err := h.ds.GetMetasCount(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, metasCount)
}

func (h *Handler) filterByCategoryAndSubcategory(ctx *gin.Context) {
}
