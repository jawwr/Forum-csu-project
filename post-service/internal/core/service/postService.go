package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
)

type _postService struct {
	repo              repository.PostRepository
	kafka             repository.EventRepository
	subscriberService service.SubscriberService
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

	subscribers, err := postService.subscriberService.GetAllUserSubscribers(ctx, post.UserId)
	if err != nil {
		return 0, err
	}

	var subscribersIds []int

	for _, subscriber := range subscribers {
		subscribersIds = append(subscribersIds, subscriber.Id)
	}
	event := model.PostEvent{
		Title:         "Новый пост: " + post.Title,
		PostId:        post.Id,
		SubscriberIds: subscribersIds,
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
