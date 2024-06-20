package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/Forum-service/Forum-api-gateway/genproto/comment"
	"github.com/gin-gonic/gin"
)

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment with given information
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param comment body comment.CreateCommentRequest true "Comment information"
// @Success 201 {object} comment.Comment
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /comments [post]
func (h *Handler) CreateComment(c *gin.Context) {
	var req comment.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CommentService.CreateComment(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetCommentById godoc
// @Summary Get a comment by its ID
// @Description Retrieve a comment by its unique ID
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Success 200 {object} comment.Comment
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Comment item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /comments/{id} [get]
func (h *Handler) GetCommentById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.CommentService.GetComment(c.Request.Context(), &comment.GetCommentRequest{
		Id: id,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get comment")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateComment godoc
// @Summary Update a comment by its ID
// @Description Update a comment with the given information
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Param comment body comment.UpdateCommentRequest true "Comment information"
// @Success 200 {object} comment.Comment
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Comment item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /comments/{id} [put]
func (h *Handler) UpdateComment(c *gin.Context) {
	var req comment.UpdateCommentRequest
	id := c.Param("id")
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CommentService.UpdateComment(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to update comment")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteComment godoc
// @Summary Delete a comment by its ID
// @Description Delete a comment by its unique ID
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Success 200 {object} comment.DeleteCommentResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Comment item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /comments/{id} [delete]
func (h *Handler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.CommentService.DeleteComment(c.Request.Context(), &comment.DeleteCommentRequest{Id: id})
	if err != nil {
		log.Error().Err(err).Msg("failed to delete comment")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Retrieve a list of all comments with optional filtering and pagination
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Number of comments per page"
// @Param post_id query string false "Filter by post ID"
// @Param user_id query string false "Filter by user ID"
// @Success 200 {object} comment.GetAllCommentsResponse
// @Failure 500 {object} string "Internal server error"
// @Router /comments [get]
func (h *Handler) GetAllComments(c *gin.Context) {
	var (
		req comment.GetAllCommentsRequest
		err error
	)
	req.PostId = c.Query("post_id")
	req.UserId = c.Query("user_id")

	req.Page, req.Limit, err = ReadPageLimit(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to read page limit")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.CommentService.GetAllComments(c.Request.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get comments")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
