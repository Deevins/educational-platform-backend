package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users"
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
	UserID   int32  `json:"user_id"`
	UserRole string `json:"user_role"`
}

type Service struct {
	repo users.Querier
}

func NewService(repo users.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

type RegisterUserResponse struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

func (s *Service) CreateUser(ctx context.Context, user model.UserCreate) (RegisterUserResponse, error) {
	user.Password = generatePasswordHash(user.Password)

	userCreate := &users.CreateUserParams{
		FullName:       user.FullName,
		Email:          user.Email,
		Description:    &user.Description,
		PasswordHashed: user.Password,
		PhoneNumber:    user.PhoneNumber,
	}

	id, err := s.repo.CreateUser(ctx, userCreate)
	if err != nil {
		return RegisterUserResponse{}, errors.Wrap(err, "failed to create user")
	}

	tkn, err := s.getUserByIDAndGenerateToken(ctx, id)

	return RegisterUserResponse{
		UserID: id,
		Token:  tkn,
	}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(passwordSalt))
}

type LoginUserResponse struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

func (s *Service) Authorize(ctx context.Context, email, password string) (LoginUserResponse, error) {
	user, err := s.repo.GetUserByEmailAndHashedPassword(ctx, &users.GetUserByEmailAndHashedPasswordParams{
		Email:          email,
		PasswordHashed: generatePasswordHash(password),
	})

	if err != nil {
		return LoginUserResponse{}, errors.Wrap(err, "failed to get user by email and password")
	}

	role, _ := user.Role.Value()

	// token create new JWT token and sign it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithId{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   user.ID,
		UserRole: role.(string),
	})

	jwtt, err := token.SignedString(jwtKey)
	if err != nil {
		return LoginUserResponse{}, errors.Wrap(err, "failed to sign token")
	}

	return LoginUserResponse{
		UserID: user.ID,
		Token:  jwtt,
	}, nil
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

func (s *Service) getUserByIDAndGenerateToken(ctx context.Context, id int32) (string, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return "", errors.Wrap(err, "failed to get user by id")
	}

	role, _ := user.Role.Value()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithId{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   user.ID,
		UserRole: role.(string),
	})

	jwtt, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")

	}
	return jwtt, nil

}
