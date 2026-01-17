package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/user"
)

type Service interface {
	GenAccessToken(userID uuid.UUID) (string, error)
	GenRefreshToken(userID uuid.UUID) (string, error)
	ValidateAccessToken(tokenStr string) (*Claims, error)
	ValidateRefreshToken(tokenStr string) (*Claims, error)
	RegisterUser(ctx context.Context, username string) (uuid.UUID, error)
	LoginUser(ctx context.Context, username string) (string, string, error)
}

type authSvc struct {
	AccessSecret string
	RefreshSecret string
	UserSvc user.Service
}

func (s *authSvc) genToken(userID string, duration time.Duration, key []byte) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: "CVWO-GWG",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
				IssuedAt: jwt.NewNumericDate(time.Now()),
			},
		})
	return token.SignedString(key)
}

func (s *authSvc) validateToken(tokenStr string, key []byte) (*Claims, error){
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
									  func (token *jwt.Token) (any, error){
										  return key, nil
									  },
									  jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*Claims); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}

func (s *authSvc) GenAccessToken(userID uuid.UUID) (string, error) {
	//15 minute magic number atm... parameterise this soon
	return s.genToken(userID.String(), 15 * time.Minute, []byte(s.AccessSecret))
}

func (s *authSvc) GenRefreshToken(userID uuid.UUID) (string, error) {
	return s.genToken(userID.String(), 12 * time.Hour, []byte(s.RefreshSecret))
}

func (s *authSvc) ValidateAccessToken(tokenStr string) (*Claims, error){
	return s.validateToken(tokenStr, []byte(s.AccessSecret))
}

func (s *authSvc) ValidateRefreshToken(tokenStr string) (*Claims, error){
	return s.validateToken(tokenStr, []byte(s.RefreshSecret))
}

func (s *authSvc) RegisterUser(ctx context.Context, username string) (uuid.UUID, error){
	return s.UserSvc.RegisterUser(ctx, username)
}

func (s *authSvc) LoginUser(ctx context.Context, username string) (string, string, error){
	user, err := s.UserSvc.FindByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}
	accessToken, Aerr := s.GenAccessToken(user.ID)
	if Aerr != nil {
		return "", "", errors.Join(ErrAccessTokenGen, Aerr)
	}
	refreshToken, Rerr := s.GenRefreshToken(user.ID)
	if Rerr != nil {
		return "", "", errors.Join(ErrRefreshTokenGen, Rerr)
	}
	return accessToken, refreshToken, nil
}

func NewAuthSvc(us user.Service, accessSecret string, refreshSecret string) *authSvc{
	return &authSvc{
		UserSvc: us, 
		AccessSecret: accessSecret, 
		RefreshSecret: refreshSecret,
	}
}