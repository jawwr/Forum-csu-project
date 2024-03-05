package service

import (
	"context"
	"errors"
	"log/slog"
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
)

type _postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) service.PostService {
	return _postService{repo: repo}
}

func (postService _postService) CreatePost(ctx context.Context, post model.Post) (int, error) {
	id, err := postService.repo.CreatePost(ctx, post)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("create post error")
	}

	return id, nil
}

func (postService _postService) GetPost(ctx context.Context, postId int) (model.Post, error) {
	return postService.repo.GetPost(ctx, postId)
}

func (postService _postService) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	return postService.repo.GetAllPosts(ctx)
}
