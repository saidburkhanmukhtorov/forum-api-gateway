package api

import (
	_ "github.com/Forum-service/Forum-api-gateway/api/docs"
	"github.com/Forum-service/Forum-api-gateway/api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Forum API Gateway
// @version 1.0
// @description This is a forum API gateway service.
// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewEngine() *gin.Engine {
	h, err := handler.NewHandler()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Use(middlewares.Auth)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Grouping API routes under /v1
	v1 := r.Group("/v1")
	{
		// Categories
		v1.POST("/categories", h.CreateCategory)
		v1.GET("/categories/:id", h.GetCategoryById)
		v1.PUT("/categories", h.UpdateCategory)
		v1.DELETE("/categories/:id", h.DeleteCategory)
		v1.GET("/categories", h.GetAllCategories)

		// Tags
		v1.POST("/tags", h.CreateTag)
		v1.GET("/tags/:id", h.GetTagById)
		v1.PUT("/tags/:id", h.UpdateTag)
		v1.DELETE("/tags/:id", h.DeleteTag)
		v1.GET("/tags", h.GetAllTags)

		// Posts
		v1.POST("/posts", h.CreatePost)
		v1.GET("/posts/:id", h.GetPostById)
		v1.PUT("/posts/:id", h.UpdatePost)
		v1.DELETE("/posts/:id", h.DeletePost)
		v1.GET("/posts", h.GetAllPosts)

		// Comments
		v1.POST("/comments", h.CreateComment)
		v1.GET("/comments/:id", h.GetCommentById)
		v1.PUT("/comments/:id", h.UpdateComment)
		v1.DELETE("/comments/:id", h.DeleteComment)
		v1.GET("/comments", h.GetAllComments)

		// PostTags
		v1.POST("/posttags", h.CreatePostTag)
		v1.DELETE("/posttags/:postid/:tagid", h.DeletePostTag)
		v1.GET("/posttags", h.GetAllPostTags)
		v1.GET("/posttags/:tag_id/posts", h.GetPostsByTag)
		v1.GET("/tags/popular", h.GetFamousTags)
	}

	return r
}
