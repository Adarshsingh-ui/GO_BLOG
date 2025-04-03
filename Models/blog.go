package models

import "time"

type Blog struct {
	ID          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required,min=3,max=100"`
	Content     string    `json:"content" binding:"required,min=10"`
	Author      string    `json:"author" binding:"required"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ReadTime    int       `json:"read_time"`
	IsPublished bool      `json:"is_published"`
}

type BlogResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    *Blog  `json:"data,omitempty"`
}

type BlogListResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []*Blog `json:"data,omitempty"`
} 