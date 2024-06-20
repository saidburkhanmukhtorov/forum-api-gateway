package handler

import (
	"net/http"

	"github.com/Forum-service/Forum-api-gateway/genproto/post"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with given information
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param post body post.CreatePostRequest true "Post information"
// @Success 201 {object} post.Post
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /posts [post]
func (h *Handler) CreatePost(c *gin.Context) {
	var req post.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostService.CreatePost(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetPostById godoc
// @Summary Get a post by its ID
// @Description Retrieve a post by its unique ID
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Success 200 {object} post.Post
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Post item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /posts/{id} [get]
func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.PostService.GetPost(c.Request.Context(), &post.GetPostRequest{
		Id: id,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePost godoc
// @Summary Update a post by its ID
// @Description Update a post with the given information
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Param post body post.UpdatePostRequest true "Post information"
// @Success 200 {object} post.Post
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Post item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /posts/{id} [put]
func (h *Handler) UpdatePost(c *gin.Context) {
	var req post.UpdatePostRequest
	id := c.Param("id")
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostService.UpdatePost(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to update post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePost godoc
// @Summary Delete a post by its ID
// @Description Delete a post by its unique ID
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Success 200 {object} post.DeletePostResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Post item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /posts/{id} [delete]
func (h *Handler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.PostService.DeletePost(c.Request.Context(), &post.DeletePostRequest{Id: id})
	if err != nil {
		log.Error().Err(err).Msg("failed to delete post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllPosts godoc
// @Summary Get all posts
// @Description Retrieve a list of all posts with optional filtering and pagination
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Number of posts per page"
// @Param user_id query string false "Filter by user ID"
// @Param title query string false "Filter by title"
// @Param category_id query string false "Filter by category ID"
// @Param body query string false "Filter by body content (partial match)"
// @Success 200 {object} post.GetAllPostsResponse
// @Failure 500 {object} string "Internal server error"
// @Router /posts [get]
func (h *Handler) GetAllPosts(c *gin.Context) {
	var (
		req post.GetAllPostsRequest
		err error
	)

	req.UserId = c.Query("user_id")
	req.Title = c.Query("title")
	req.CategoryId = c.Query("category_id")
	req.Body = c.Query("body")
	req.Page, req.Limit, err = ReadPageLimit(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to read page limit")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.PostService.GetAllPosts(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get posts")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
