package Routes

import (
	"myapp/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		blogs := v1.Group("/blogs")
		{
			blogs.GET("", Controllers.GetBlogs)
			blogs.POST("", Controllers.CreateBlog)
			blogs.GET("/:id", Controllers.GetBlogByID)
			blogs.PUT("/:id", Controllers.UpdateBlog)
			blogs.DELETE("/:id", Controllers.DeleteBlog)
			blogs.PATCH("/:id/publish", Controllers.PublishBlog)
		}
	}

	return r
}