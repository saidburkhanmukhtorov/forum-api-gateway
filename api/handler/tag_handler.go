package handler

import (
	"net/http"

	"github.com/Forum-service/Forum-api-gateway/genproto/tag"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// CreateTag godoc
// @Summary Create a new tag
// @Description Create a new tag with given information
// @Tags tag
// @Accept json
// @Produce json
// @Param tag body tag.CreateTagRequest true "Tag information"
// @Success 201 {object} tag.Tag
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /tags [post]
func (h *Handler) CreateTag(c *gin.Context) {
	var req tag.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.TagService.CreateTag(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create tag")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetTagById godoc
// @Summary Get a tag by its ID
// @Description Retrieve a tag by its unique ID
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} tag.Tag
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Tag item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /tags/{id} [get]
func (h *Handler) GetTagById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.TagService.GetTag(c.Request.Context(), &tag.GetTagRequest{
		Id: id,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get tag")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTag godoc
// @Summary Update a tag by its ID
// @Description Update a tag with the given information
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Param tag body tag.UpdateTagRequest true "Tag information"
// @Success 200 {object} tag.Tag
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Tag item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /tags/{id} [put]
func (h *Handler) UpdateTag(c *gin.Context) {
	var req tag.UpdateTagRequest
	id := c.Param("id")
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.TagService.UpdateTag(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to update tag")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteTag godoc
// @Summary Delete a tag by its ID
// @Description Delete a tag by its unique ID
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} tag.DeleteTagResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Tag item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /tags/{id} [delete]
func (h *Handler) DeleteTag(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.TagService.DeleteTag(c.Request.Context(), &tag.DeleteTagRequest{Id: id})
	if err != nil {
		log.Error().Err(err).Msg("failed to delete tag")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllTags godoc
// @Summary Get all tags
// @Description Retrieve a list of all tags with optional pagination
// @Tags tag
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of tags per page"
// @Success 200 {object} tag.GetAllTagsResponse
// @Failure 500 {object} string"Internal server error"
// @Router /tags [get]
func (h *Handler) GetAllTags(c *gin.Context) {
	var (
		req tag.GetAllTagsRequest
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
	resp, err := h.TagService.GetAllTags(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get tags")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
