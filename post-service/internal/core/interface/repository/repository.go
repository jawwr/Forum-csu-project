package repository

import (
	"context"
	"post-service/internal/core/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
}
