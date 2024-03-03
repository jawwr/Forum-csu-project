package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"time"
	"user-service/internal/core/dto"
	"user-service/internal/core/helper"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/interface/service"
	"user-service/internal/core/repository/model"
	modelDb "user-service/internal/core/repository/model"
)

type _tokenService struct {
	repository repository.TokenRepository
}

func NewTokenService(repo repository.TokenRepository) service.TokenService {
	return _tokenService{repository: repo}
}

func (service _tokenService) GetToken(ctx context.Context, token string) (*modelDb.Token, error) {
	return service.repository.GetToken(ctx, token)
}

func (service _tokenService) SaveToken(ctx context.Context, token modelDb.Token) error {
	return service.repository.SaveToken(ctx, token)
}

func (service _tokenService) GenerateToken(ctx context.Context, user modelDb.User) (string, error) {
	expire := time.Now().Add(helper.TokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: user.Login,
	})
	jwtToken, err := token.SignedString([]byte(helper.SignInKey))
	if err != nil {
		return "", err
	}
	tokenModel := model.Token{
		Id:      0,
		Token:   jwtToken,
		Expire:  expire,
		Revoked: false,
		UserId:  user.Id,
	}
	if err := service.revokeToken(ctx, user.Id); err != nil {
		return "", err
	}
	if err := service.SaveToken(ctx, tokenModel); err != nil {
		return "", err
	}
	return tokenModel.Token, nil
}

func (service _tokenService) revokeToken(ctx context.Context, userId int) error {
	tokens, err := service.repository.GetValidTokenByUserId(ctx, userId)
	if err != nil {
		return err
	}
	for _, token := range tokens {
		token.Revoked = true

		if err := service.SaveToken(ctx, token); err != nil {
			return err
		}
	}
	return nil
}
