package service

import (
	"context"
	"post-service/internal/core/model"
)

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
}

type UserService interface {
	GetUserByToken(ctx context.Context, token string) (model.User, error)
}
