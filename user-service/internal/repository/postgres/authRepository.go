package postgres

import (
	"context"
	"fmt"
	"user-service/internal/core/interface/repository"
	"user-service/internal/lib/db"
)

type userDB struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type _userRepo struct {
	*db.Db
}

func NewRepo(db *db.Db) repository.UserRepository {
	return _userRepo{db}
}

func (repo _userRepo) GetUser(ctx context.Context, login, hashPassword string) (string, error) {
	var user userDB

	row := repo.PgConn.QueryRow(nil, `SELECT * FROM users WHERE login=$1 and password=$2`, login, hashPassword)

	if err := row.Scan(&user); err != nil {
		return "", fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return login, nil

}

func (repo _userRepo) CreateUser(ctx context.Context, login, hashPassword string) (string, error) {
	_, err := repo.PgConn.Exec(
		ctx,
		`INSERT INTO users(login, pass) values ($1, $2)`,
		login, hashPassword,
	)

	if err != nil {
		return "", fmt.Errorf("не смогли создать: %x", err)
	}

	return login, nil
}
