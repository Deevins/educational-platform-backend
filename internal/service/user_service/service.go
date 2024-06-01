package user_service

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users_repo"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

var _ handler.UserService = &Service{}

type Service struct {
	repo users_repo.Querier
}

func NewService(repo users_repo.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SetHasUserTriedInstructorToTrue(ctx context.Context, ID int32) error {
	hasUsed, err := s.repo.GetHasUserTriedInstructor(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "failed to get has user tried instructor")
	}

	if hasUsed != nil && !*hasUsed {
		err = s.repo.UpdateHasUserTriedInstructor(ctx, ID)
		if err != nil {
			return errors.Wrap(err, "failed to update has user tried instructor")
		}

		return nil
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, ID int32) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by id")

	}

	return &model.User{
		ID:          user.ID,
		FullName:    user.FullName,
		Description: lo.FromPtrOr(user.Description, ""),
		Email:       user.Email,
		Avatar:      user.Avatar,
		PhoneNumber: user.PhoneNumber,
	}, nil

}

func (s *Service) GetHasUserTriedInstructor(ctx context.Context, ID int32) (bool, error) {
	hasUsed, err := s.repo.GetHasUserTriedInstructor(ctx, ID)
	if err != nil && hasUsed == nil {
		return false, errors.Wrap(err, "failed to get has user tried instructor")
	}

	return *hasUsed, nil
}

func (s *Service) UpdateAvatar(ctx context.Context, ID int32, avatar []byte) error {
	err := s.repo.UpdateAvatar(ctx, &users_repo.UpdateAvatarParams{
		Avatar: avatar,
		ID:     ID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update avatar")
	}

	return nil
}

func (s *Service) AddUserTeachingExperience(ctx context.Context, exp *model.UserUpdateTeachingExperience) error {
	err := s.repo.AddTeachingExperience(ctx, &users_repo.AddTeachingExperienceParams{
		HasVideoKnowledge:     exp.HasVideoKnowledge,
		HasPreviousExperience: exp.HasPreviousExperience,
		CurrentAudienceCount:  exp.CurrentAudienceCount,
		UserID:                exp.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user teaching experience")
	}

	return nil
}

func (s *Service) UpdateUserInfo(ctx context.Context, user model.UserUpdate) error {
	err := s.repo.UpdateUser(ctx, &users_repo.UpdateUserParams{
		FullName:    user.FullName,
		Email:       user.Email,
		Description: &user.Description,
		PhoneNumber: user.PhoneNumber,
		ID:          user.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (s *Service) GetSelfInfo(ctx context.Context, ID int32) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by id")
	}

	return &model.User{
		ID:          user.ID,
		FullName:    user.FullName,
		Description: lo.FromPtrOr(user.Description, ""),
		Email:       user.Email,
		Avatar:      user.Avatar,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
