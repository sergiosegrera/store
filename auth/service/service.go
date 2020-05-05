package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sergiosegrera/store/auth/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, password string) (string, time.Time, error)
	Check(ctx context.Context, token string) error
	Refresh(ctx context.Context, token string) (string, time.Time, error)
}

type Service struct {
	conf *config.Config
}

func NewService(c *config.Config) AuthService {
	return Service{conf: c}
}

func (s Service) Login(ctx context.Context, password string) (string, time.Time, error) {
	err := bcrypt.CompareHashAndPassword(s.conf.Password, []byte(password))
	if err != nil {
		return "", time.Time{}, ErrWrongPassword
	}

	// TODO: Env vars for jwt config?
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.conf.JwtSecretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, err
}

func (s Service) Check(ctx context.Context, token string) error {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.conf.JwtSecretKey, nil
	})
	if err != nil {
		switch err {
		case jwt.ErrSignatureInvalid:
			return ErrInvalidToken
		default:
			return ErrBadToken
		}
	}

	if !tkn.Valid {
		return ErrInvalidToken
	}

	return err
}

func (s Service) Refresh(ctx context.Context, token string) (string, time.Time, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.conf.JwtSecretKey, nil
	})
	if err != nil {
		switch err {
		case jwt.ErrSignatureInvalid:
			return "", time.Time{}, ErrInvalidToken
		default:
			return "", time.Time{}, ErrBadToken
		}
	}

	if !tkn.Valid {
		return "", time.Time{}, ErrInvalidToken
	}

	// TODO: Maybe only refresh at a certain expiration time

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(s.conf.JwtSecretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, err
}

var (
	ErrWrongPassword = errors.New("Wrong password")
	ErrInvalidToken  = errors.New("Invalid token")
	ErrBadToken      = errors.New("Bad token")
)
