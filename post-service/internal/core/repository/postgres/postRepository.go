package postgres

import (
	"context"
	"fmt"
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/mapper"
	"post-service/internal/core/model"
	"post-service/internal/core/repository/dbModel"
	"post-service/internal/lib/db"
)

type _postRepository struct {
	db *db.Db
}

func NewPostRepo(db *db.Db) repository.PostRepository {
	return _postRepository{db}
}

func (postRepository _postRepository) CreatePost(ctx context.Context, post model.Post) (int, error) {
	postDb := dbModel.Post(post)
	var id int

	err := postRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.post(title, text, user_id) values ($1,$2,$3) RETURNING id`,
		postDb.Title,
		postDb.Text,
		postDb.UserId).Scan(&id)

	return id, err
}

func (postRepository _postRepository) GetPost(ctx context.Context, postId int) (model.Post, error) {
	var post dbModel.Post

	err := postRepository.db.PgConn.QueryRow(ctx,
		`SELECT p.id, p.title, p.text, p.user_id FROM public.post p WHERE p.id=$1`,
		postId).Scan(&post.Id, &post.Title, &post.Text, &post.UserId)

	if err != nil {
		return model.Post{}, fmt.Errorf("get post error: %s", err.Error())
	}

	return *mapper.FromPostDb(&post), nil
}

func (postRepository _postRepository) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	rows, err := postRepository.db.PgConn.Query(ctx,
		`SELECT p.id, p.title, p.text, p.user_id FROM public.post p`)

	if err != nil {
		return []model.Post{}, fmt.Errorf("get all posts get error: %s", err.Error())
	}

	defer rows.Close()

	var posts []dbModel.Post

	for rows.Next() {
		var post dbModel.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Text, &post.UserId); err != nil {
			return []model.Post{}, fmt.Errorf("all posts get error: %s", err.Error())
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return []model.Post{}, fmt.Errorf("all posts get error: %s", err.Error())
	}

	var result []model.Post

	for _, post := range posts {
		result = append(result, *mapper.FromPostDb(&post))
	}

	return result, nil
}
