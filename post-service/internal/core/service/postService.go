package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"log/slog"
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
)

type _postService struct {
	repo  repository.PostRepository
	kafka repository.EventRepository
}

func NewPostService(repo repository.PostRepository, kafka repository.EventRepository) service.PostService {
	return _postService{repo: repo, kafka: kafka}
}

func (postService _postService) CreatePost(ctx context.Context, post model.Post) (int, error) {
	id, err := postService.repo.CreatePost(ctx, post)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("create post error")
	}

	requestId, err := uuid.GenerateUUID()

	if err != nil {
		slog.Error("generate uuid error: ", err.Error())
		return id, nil
	}

	event := model.Event{
		Id:    requestId,
		Key:   requestId,
		Value: id,
	}

	err = postService.kafka.SendNewPostEvent(ctx, event)

	if err != nil {
		slog.Error(fmt.Sprint("send msg error: ", err.Error()))
		return id, nil
	}

	return id, nil
}

func (postService _postService) GetPost(ctx context.Context, postId int) (model.Post, error) {
	return postService.repo.GetPost(ctx, postId)
}

func (postService _postService) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	return postService.repo.GetAllPosts(ctx)
}
