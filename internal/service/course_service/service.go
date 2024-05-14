package course_service

import (
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/courses_repo"
)

//var _ handler.UserService = &Service{}

type Service struct {
	repo courses_repo.Querier
}

func NewService(repo courses_repo.Querier) *Service {
	return &Service{
		repo: repo,
	}
}
