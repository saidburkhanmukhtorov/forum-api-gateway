package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/Forum-service/Forum-api-gateway/genproto/category"
	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with given information
// @Tags category
// @Accept json
// @Produce json
// @Param category body category.CreateCategoryRequest true "Category information"
// @Success 201 {object} category.Category
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /categories [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var req category.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CategoryService.CreateCategory(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetCategoryById godoc
// @Summary Get a category by its ID
// @Description Retrieve a category by its unique ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} category.Category
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Category item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /categories/{id} [get]
func (h *Handler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	log.Print(id)
	resp, err := h.CategoryService.GetCategory(c.Request.Context(), &category.GetCategoryRequest{
		Id: id,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCategory godoc
// @Summary Update a category by its ID
// @Description Update a category with the given information
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body category.UpdateCategoryRequest true "Category information"
// @Success 200 {object} category.Category
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Category item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /categories [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	var req category.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CategoryService.UpdateCategory(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to update category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCategory godoc
// @Summary Delete a category by its ID
// @Description Delete a category by its unique ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} category.DeleteCategoryResponse
// @Failure 404 {object} string "Category item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /categories/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.CategoryService.DeleteCategory(c.Request.Context(), &category.DeleteCategoryRequest{Id: id})
	if err != nil {
		log.Error().Err(err).Msg("failed to delete category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories with optional pagination
// @Tags category
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of categories per page"
// @Success 200 {object} category.GetAllCategoriesResponse
// @Failure 500 {object} string "Internal server error"
// @Router /categories [get]
func (h *Handler) GetAllCategories(c *gin.Context) {
	var (
		req category.GetAllCategoriesRequest
		err error
	)
	req.Page, req.Limit, err = ReadPageLimit(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to read page limit")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CategoryService.GetAllCategories(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get categories")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp1, err := h.CategoryService.GetCategory(c.Request.Context(), &category.GetCategoryRequest{
		Id: resp.Categories[0].Id,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Logger.Println(resp1)
	c.JSON(http.StatusOK, resp)
}
