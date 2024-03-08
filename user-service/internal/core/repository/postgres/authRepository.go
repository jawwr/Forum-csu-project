package postgres

import (
	"context"
	"fmt"
	"user-service/internal/core/interface/repository"
	dbModel "user-service/internal/core/repository/model"
	"user-service/internal/lib/db"
	"user-service/internal/transport/model"
)

type _userRepo struct {
	*db.Db
}

func NewUserRepository(db *db.Db) repository.UserRepository {
	return _userRepo{db}
}

func (repo _userRepo) GetUserByCredentials(ctx context.Context, userDto model.UserCredentials) (*dbModel.User, error) {
	var user dbModel.User
	row := repo.PgConn.QueryRow(
		ctx,
		`SELECT id, login, password 
			 FROM users
			 WHERE login=$1 AND password=$2`,
		userDto.Login, userDto.Password,
	)

	if err := row.Scan(&user.Id, &user.Login, &user.Password); err != nil {
		return nil, fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return &user, nil
}

func (repo _userRepo) GetUserById(ctx context.Context, userId int) (*dbModel.User, error) {
	var user dbModel.User
	row := repo.PgConn.QueryRow(
		ctx,
		`SELECT id, login, password 
			 FROM users
			 WHERE id = $1`,
		userId,
	)

	if err := row.Scan(&user.Id, &user.Login, &user.Password); err != nil {
		return nil, fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return &user, nil
}

func (repo _userRepo) CreateUser(ctx context.Context, userDto model.UserCredentials) (*dbModel.User, error) {
	var user dbModel.User
	row := repo.PgConn.QueryRow(
		ctx,
		`INSERT INTO users(login, password)
			 VALUES ($1, $2)
			 RETURNING id, login, password`,
		userDto.Login, userDto.Password,
	)

	if err := row.Scan(&user.Id, &user.Login, &user.Password); err != nil {
		return nil, fmt.Errorf("не смогли создать: %x", err)
	}

	return &user, nil
}

func (repo _userRepo) GetAllUsers(ctx context.Context) ([]*dbModel.User, error) {
	var users []*dbModel.User
	rows, err := repo.PgConn.Query(
		ctx,
		`SELECT id, login, password 
			 FROM users`,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user dbModel.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
