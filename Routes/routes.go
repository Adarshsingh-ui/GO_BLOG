package routes

import (
	"myapp/controllers"
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
			blogs.GET("", controllers.GetBlogs)
			blogs.POST("", controllers.CreateBlog)
			blogs.GET("/:id", controllers.GetBlogByID)
			blogs.PUT("/:id", controllers.UpdateBlog)
			blogs.DELETE("/:id", controllers.DeleteBlog)
			blogs.PATCH("/:id/publish", controllers.PublishBlog)
		}
	}

	return r
}