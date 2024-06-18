package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/Forum-service/Forum-api-gateway/genproto/posttag"
	"github.com/gin-gonic/gin"
)

// CreatePostTag godoc
// @Summary Create a new post-tag relationship
// @Description Create a new post-tag relationship with given information
// @Tags posttag
// @Accept json
// @Produce json
// @Param posttag body posttag.CreatePostTagRequest true "PostTag information"
// @Success 201 {object} posttag.PostTag
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /posttags [post]
func (h *Handler) CreatePostTag(c *gin.Context) {
	var req posttag.CreatePostTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostTagService.CreatePostTag(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post-tag relationship")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// DeletePostTag godoc
// @Summary Delete a post-tag relationship
// @Description Delete a post-tag relationship with given information
// @Tags posttag
// @Accept json
// @Produce json
// @Param posttag body posttag.DeletePostTagRequest true "PostTag information"
// @Success 200 {object} posttag.DeletePostTagResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Post_tag item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /posttags [delete]
func (h *Handler) DeletePostTag(c *gin.Context) {
	var req posttag.DeletePostTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostTagService.DeletePostTag(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete post-tag relationship")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllPostTags godoc
// @Summary Get all post-tag relationships
// @Description Retrieve a list of all post-tag relationships with optional filtering and pagination
// @Tags posttag
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of post-tag relationships per page"
// @Param post_id query string false "Filter by post ID"
// @Param tag_id query string false "Filter by tag ID"
// @Success 200 {object} posttag.GetAllPostTagsResponse
// @Failure 500 {object} string "Internal server error"
// @Router /posttags [get]
func (h *Handler) GetAllPostTags(c *gin.Context) {
	var (
		req posttag.GetAllPostTagsRequest
		err error
	)
	req.PostId = c.Query("post_id")
	req.TagId = c.Query("tag_id")

	req.Page, req.Limit, err = ReadPageLimit(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to read page limit")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostTagService.GetAllPostTags(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get post-tag relationships")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetPostsByTag godoc
// @Summary Get posts by tag ID
// @Description Retrieve a list of posts associated with a given tag ID
// @Tags posttag
// @Accept json
// @Produce json
// @Param tag_id path string true "Tag ID"
// @Param page query int false "Page number"
// @Param limit query int false "Number of posts per page"
// @Success 200 {object} posttag.GetPostsByTagResponse
// / @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Posttag item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /posttags/{tag_id}/posts [get]
func (h *Handler) GetPostsByTag(c *gin.Context) {
	var (
		req posttag.GetPostsByTagRequest
		err error
	)
	req.TagId = c.Param("tag_id")

	req.Page, req.Limit, err = ReadPageLimit(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to read page limit")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostTagService.GetPostsByTag(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get posts by tag")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
