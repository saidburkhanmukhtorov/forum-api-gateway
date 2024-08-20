package handler

import (
	"fmt"
	"strconv"

	"github.com/Forum-service/Forum-api-gateway/genproto/category"
	"github.com/Forum-service/Forum-api-gateway/genproto/comment"
	"github.com/Forum-service/Forum-api-gateway/genproto/post"
	"github.com/Forum-service/Forum-api-gateway/genproto/posttag"
	"github.com/Forum-service/Forum-api-gateway/genproto/tag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Handler struct holds gRPC client connections.
type Handler struct {
	CategoryService category.CategoryServiceClient
	TagService      tag.TagServiceClient
	PostService     post.PostServiceClient
	CommentService  comment.CommentServiceClient
	PostTagService  posttag.PostTagServiceClient
}

// NewHandler establishes gRPC connections and returns a Handler struct.
func NewHandler() (*Handler, error) {
	conn, err := grpc.NewClient(
		"forum_service:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("error connecting to gRPC server: %v", err)
	}

	return &Handler{
		CategoryService: category.NewCategoryServiceClient(conn),
		TagService:      tag.NewTagServiceClient(conn),
		PostService:     post.NewPostServiceClient(conn),
		CommentService:  comment.NewCommentServiceClient(conn),
		PostTagService:  posttag.NewPostTagServiceClient(conn),
	}, nil
}

func ReadPageLimit(c *gin.Context) (int32, int32, error) {
	var (
		page  int64
		limit int64
		err   error
	)

	// Read page query parameter
	pageStr := c.Query("page")
	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			return 0, 0, err
		}
	} else {
		page = 1 // Default page
	}

	// Read limit query parameter
	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 32)
		if err != nil {
			return 0, 0, err
		}
	} else {
		limit = 10
	}

	return int32(page), int32(limit), nil
}
