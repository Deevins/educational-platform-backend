package user

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

var _ handler.UserService = &Service{}

type Service struct {
	repo users.Querier
	s3   S3.Client
}

func NewService(repo users.Querier, s3Client S3.Client) *Service {
	return &Service{
		repo: repo,
		s3:   s3Client,
	}
}

func (s *Service) CheckIfUserRegisteredToCourse(ctx context.Context, userID, courseID int32) (bool, error) {
	_, err := s.repo.CheckIfUserRegisteredToCourse(ctx, &users.CheckIfUserRegisteredToCourseParams{
		UserID:   userID,
		CourseID: courseID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s *Service) RegisterToCourse(ctx context.Context, userID, courseID int32) error {
	_, err := s.repo.RegisterToCourse(ctx, &users.RegisterToCourseParams{
		UserID:   userID,
		CourseID: courseID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SetHasUserTriedInstructorToTrue(ctx context.Context, ID int32) error {
	hasUsed, err := s.repo.GetHasUserTriedInstructor(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "failed to get has user tried instructor")
	}

	if hasUsed != nil && !*hasUsed {
		_, err = s.repo.UpdateHasUserTriedInstructor(ctx, ID)
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
		AvatarUrl:   *user.AvatarUrl,
		PhoneNumber: user.PhoneNumber,
	}, nil

}

func (s *Service) GetHasUserTriedInstructor(ctx context.Context, ID int32) (bool, error) {
	hasUsed, err := s.repo.GetHasUserTriedInstructor(ctx, ID)
	if err != nil || hasUsed == nil {
		return false, errors.Wrap(err, "failed to get has user tried instructor")
	}

	return *hasUsed, nil
}

func (s *Service) UpdateAvatar(ctx context.Context, ID int32, avatar S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(avatar)
	if err != nil {
		return "", err
	}

	_, err = s.repo.UpdateAvatar(ctx, &users.UpdateAvatarParams{
		ID:        ID,
		AvatarUrl: &url,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) AddUserTeachingExperience(ctx context.Context, exp *model.UserUpdateTeachingExperience) error {
	_, err := s.repo.AddTeachingExperience(ctx, &users.AddTeachingExperienceParams{
		VideoKnowledge:     exp.VideoKnowledge,
		PreviousExperience: exp.PreviousExperience,
		CurrentAudience:    exp.CurrentAudience,
		UserID:             exp.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user teaching experience")
	}

	return nil
}

func (s *Service) UpdateUserInfo(ctx context.Context, user *model.UserUpdate) error {
	// TODO: add transaction here to all methods
	if user.Email != "" {
		if _, err := s.repo.UpdateUserEmail(ctx, &users.UpdateUserEmailParams{
			Email: user.Email,
			ID:    user.UserID,
		}); err != nil {
			return errors.Wrap(err, "failed to update user email")
		}
	}

	if user.PhoneNumber != "" {
		if _, err := s.repo.UpdateUserPhone(ctx, &users.UpdateUserPhoneParams{
			PhoneNumber: user.PhoneNumber,
			ID:          user.UserID,
		}); err != nil {
			return errors.Wrap(err, "failed to update user phone")
		}
	}

	if user.Description != "" {
		if _, err := s.repo.UpdateUserDescription(ctx, &users.UpdateUserDescriptionParams{
			Description: &user.Description,
			ID:          user.UserID,
		}); err != nil {
			return errors.Wrap(err, "failed to update user description")
		}
	}

	if user.FullName != "" {
		if _, err := s.repo.UpdateUserFullName(ctx, &users.UpdateUserFullNameParams{
			FullName: user.FullName,
			ID:       user.UserID,
		}); err != nil {
			return errors.Wrap(err, "failed to update user fullname")
		}
	}

	return nil
}
