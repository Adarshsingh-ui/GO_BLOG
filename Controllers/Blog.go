package controllers

import (
	"net/http"
	"time"
	"myapp/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var blogs []*models.Blog

func init() {
	blogs = []*models.Blog{
		{
			ID:          uuid.New().String(),
			Title:       "Go Basics",
			Content:     "An introduction to Go programming language.",
			Author:      "John Doe",
			Tags:        []string{"golang", "programming", "basics"},
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now().Add(-24 * time.Hour),
			ReadTime:    5,
			IsPublished: true,
		},
		{
			ID:          uuid.New().String(),
			Title:       "Building APIs with Gin",
			Content:     "How to create RESTful APIs using Gin framework.",
			Author:      "Jane Smith",
			Tags:        []string{"golang", "gin", "api", "web"},
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-48 * time.Hour),
			ReadTime:    8,
			IsPublished: true,
		},
	}
}

func GetBlogs(c *gin.Context) {
	if len(blogs) == 0 {
		c.JSON(http.StatusNotFound, models.BlogListResponse{
			Success: false,
			Message: "No blogs found",
		})
		return
	}
	c.JSON(http.StatusOK, models.BlogListResponse{
		Success: true,
		Data:    blogs,
	})
}

func GetBlogByID(c *gin.Context) {
	blogID := c.Param("id")
	for _, blog := range blogs {
		if blog.ID == blogID {
			c.JSON(http.StatusOK, models.BlogResponse{
				Success: true,
				Data:    blog,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.BlogResponse{
		Success: false,
		Message: "Blog not found",
	})
}

func CreateBlog(c *gin.Context) {
	var newBlog models.Blog
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, models.BlogResponse{
			Success: false,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	newBlog.ID = uuid.New().String()
	newBlog.CreatedAt = time.Now()
	newBlog.UpdatedAt = time.Now()
	newBlog.IsPublished = false 

	newBlog.ReadTime = len(newBlog.Content) / 200 

	blogs = append(blogs, &newBlog)
	c.JSON(http.StatusCreated, models.BlogResponse{
		Success: true,
		Data:    &newBlog,
	})
}

func UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")
	var updateData models.Blog
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, models.BlogResponse{
			Success: false,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	for i, blog := range blogs {
		if blog.ID == blogID {
			blogs[i].Title = updateData.Title
			blogs[i].Content = updateData.Content
			blogs[i].Author = updateData.Author
			blogs[i].Tags = updateData.Tags
			blogs[i].UpdatedAt = time.Now()
			blogs[i].ReadTime = len(updateData.Content) / 200

			c.JSON(http.StatusOK, models.BlogResponse{
				Success: true,
				Data:    blogs[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, models.BlogResponse{
		Success: false,
		Message: "Blog not found",
	})
}

func DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")
	for i, blog := range blogs {
		if blog.ID == blogID {
			blogs = append(blogs[:i], blogs[i+1:]...)
			c.JSON(http.StatusOK, models.BlogResponse{
				Success: true,
				Message: "Blog deleted successfully",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.BlogResponse{
		Success: false,
		Message: "Blog not found",
	})
}

func PublishBlog(c *gin.Context) {
	blogID := c.Param("id")
	for i, blog := range blogs {
		if blog.ID == blogID {
			blogs[i].IsPublished = true
			blogs[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, models.BlogResponse{
				Success: true,
				Data:    blogs[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.BlogResponse{
		Success: false,
		Message: "Blog not found",
	})
}