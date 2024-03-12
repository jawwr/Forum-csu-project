package mapper

import (
	"post-service/internal/core/model"
	"post-service/internal/core/repository/dbModel"
)

func FromPostDb(post *dbModel.Post) *model.Post {
	return &model.Post{
		Id:     post.Id,
		Title:  post.Title,
		Text:   post.Text,
		UserId: post.UserId,
	}
}

func ToHandlerPost(post *model.Post, user *model.User) *model.HandlePost {
	return &model.HandlePost{
		Id:    post.Id,
		Title: post.Title,
		Text:  post.Text,
		User: model.HandleUser{
			Id:    user.Id,
			Login: user.Login,
		},
	}
}
