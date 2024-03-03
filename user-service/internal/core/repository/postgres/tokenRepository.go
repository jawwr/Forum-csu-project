package postgres

import (
	"context"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/repository/model"
	"user-service/internal/lib/db"
)

type _tokenRepository struct {
	*db.Db
}

func NewTokenRepository(db *db.Db) repository.TokenRepository {
	return _tokenRepository{db}
}

func (repo _tokenRepository) GetToken(ctx context.Context, token string) (*model.Token, error) {
	var savedToken model.Token
	row := repo.PgConn.QueryRow(
		ctx,
		`SELECT id, token, revoked, expire, user_id
			 FROM tokens
			 WHERE token = $1`,
		token,
	)
	if err := row.Scan(&savedToken.Id, &savedToken.Token, &savedToken.Revoked, &savedToken.Expire, &savedToken.UserId); err != nil {
		return nil, err
	}
	return &savedToken, nil
}

func (repo _tokenRepository) SaveToken(ctx context.Context, token model.Token) error {
	if token.Id != 0 {
		_, err := repo.PgConn.Query(
			ctx,
			`UPDATE tokens
			 SET token = $2, revoked = $3
			 WHERE id = $1`,
			token.Id, token.Token, token.Revoked,
		)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := repo.PgConn.Query(
		ctx,
		`INSERT INTO tokens(token, revoked, expire, user_id)
			 VALUES ($1, $2, $3, $4)`,
		token.Token, token.Revoked, token.Expire, token.UserId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo _tokenRepository) GetValidTokenByUserId(ctx context.Context, userId int) ([]model.Token, error) {
	var tokens []model.Token
	row, err := repo.PgConn.Query(
		ctx,
		`SELECT id, token, revoked, expire, user_id
			 FROM tokens
			 WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var token model.Token
		if err := row.Scan(&token.Id, &token.Token, &token.Revoked, &token.Expire, &token.UserId); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
