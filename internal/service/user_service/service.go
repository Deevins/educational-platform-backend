package user_service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users_repo"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

var jwtKey = []byte(os.Getenv("SIGNING_KEY"))
var passwordSalt = []byte(os.Getenv("SALT"))

// TokenClaimsWithId Structure of token with user id
type TokenClaimsWithId struct {
	jwt.RegisteredClaims
	UserID int32 `json:"user_id"`
}

type Service struct {
	repo users_repo.Querier
}

func NewService(repo users_repo.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, user model.UserCreate) (int32, error) {
	user.Password = generatePasswordHash(user.Password)

	userCreate := &users_repo.CreateUserParams{
		FullName:       user.FullName,
		Email:          user.Email,
		Description:    &user.Description,
		PasswordHashed: user.Password,
		PhoneNumber:    user.PhoneNumber,
	}

	id, err := s.repo.CreateUser(ctx, userCreate)
	return id, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(passwordSalt))
}

func (s *Service) generateToken(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmailAndHashedPassword(ctx, &users_repo.GetUserByEmailAndHashedPasswordParams{
		Email:          email,
		PasswordHashed: generatePasswordHash(password),
	})

	if err != nil {
		return "", errors.Wrap(err, "failed to get user by email and password")
	}

	// token create new JWT token and sign it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithId{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},

		UserID: user.ID,
	})

	return token.SignedString(jwtKey)
}

func (s *Service) parseToken(accessToken string) (int32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaimsWithId{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaimsWithId)

	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaimsWithId")
	}

	return claims.UserID, nil

}
