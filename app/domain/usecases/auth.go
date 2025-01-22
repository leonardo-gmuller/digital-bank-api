package usecases

import (
	"context"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
)

type AuthInput struct {
	Cpf    string
	Secret string
}

type AuthOutput struct {
	AccessToken string
}

func (u *UseCase) Auth(ctx context.Context, input AuthInput, jwt *jwtauth.JWTAuth, jwtExpiresIn int) (AuthOutput, error) {
	repo := u.getAccountRepository(ctx)

	account, err := repo.GetAccountByCpf(ctx, input.Cpf)
	if err != nil {
		return AuthOutput{}, erring.ErrLoginUserNotFound
	}

	if !account.ValidatePassword(input.Secret) {
		return AuthOutput{}, erring.ErrLoginUnauthorized
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": strconv.FormatUint(uint64(account.ID), 10),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		return AuthOutput{}, erring.ErrLoginTokenNotCreated
	}

	return AuthOutput{
		AccessToken: tokenString,
	}, nil
}
