package course

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/courses"
	"github.com/deevins/educational-platform-backend/internal/model"
)

var _ handler.CourseService = &Service{}

type Service struct {
	repo courses.Querier
	s3   S3.Client
}

func NewService(repo courses.Querier, s3Client S3.Client) *Service {
	return &Service{
		repo: repo,
		s3:   s3Client,
	}
}

func (s *Service) CreateCourseBase(ctx context.Context, base *model.CourseBase) (int32, error) {
	courseID, err := s.repo.CreateCourseBase(ctx, &courses.CreateCourseBaseParams{
		Title:       base.Title,
		Type:        mapTypeToDBType(base.Type),
		AuthorID:    base.AuthorID,
		CategoryID:  &base.CategoryID,
		TimePlanned: &base.TimePlanned,
	})
	if err != nil {
		return 0, err
	}

	return courseID, nil
}

func (s *Service) GetAllPendingCourses(ctx context.Context) ([]*model.ShortCourse, error) {
	coursesList, err := s.repo.GetAllPendingCourses(ctx)
	if err != nil {
		return nil, err
	}

	return repackDBCoursesToModel(coursesList), nil
}

func (s *Service) GetAllDraftCourses(ctx context.Context) ([]*model.ShortCourse, error) {
	coursesList, err := s.repo.GetAllDraftCourses(ctx)
	if err != nil {
		return nil, err
	}

	return repackDBCoursesToModel(coursesList), nil
}

func (s *Service) GetAllReadyCourses(ctx context.Context) ([]*model.ShortCourse, error) {
	coursesList, err := s.repo.GetAllReadyCourses(ctx)
	if err != nil {
		return nil, err
	}

	return repackDBCoursesToModel(coursesList), nil
}

func (s *Service) ApproveCourse(ctx context.Context, ID int32) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SentCourseToCheck(ctx context.Context, ID int32) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RejectCourse(ctx context.Context, ID int32) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetFullCoursePageInfoByCourseID(ctx context.Context, ID int32) (*model.Course, error) {
	course := &model.Course{
		ID:               ID,
		Title:            "title",
		Subtitle:         "subtitle",
		Description:      "desc",
		Language:         "lang",
		AvatarURL:        "dick",
		Requirements:     []string{"req1", "req2"},
		Level:            "based",
		LecturesLength:   300,
		LecturesCount:    15,
		StudentsCount:    2133,
		ReviewsCount:     123,
		Rating:           4.1,
		PreviewVideoURL:  "https://youtu.be/n5OgakKNE_0",
		TargetAudience:   []string{"aud1", "aud2"},
		WhatYouWillLearn: []string{"learn1", "learn2"},
		CourseGoals:      []string{"goal1", "goal2"},
		Instructor: model.CourseInstructor{
			ID:            2,
			FullName:      "full name",
			Avatar:        []byte("avatar"),
			Description:   "desc",
			StudentsCount: 123,
			CoursesCount:  20,
			Rating:        4.5,
		},
	}
	return course, nil
}

func (s *Service) GetUserCoursesByUserID(ctx context.Context, ID int32) ([]*model.ShortCourse, error) {
	userCourses, err := s.repo.GetUserCourses(ctx, ID)
	if err != nil {
		return nil, err

	}

	return repackDBCoursesToShortModel(userCourses), nil
}

func (s *Service) SearchCoursesByTitle(ctx context.Context, title string) ([]*model.ShortCourse, error) {
	userCourses, err := s.repo.SearchCoursesByTitle(ctx, title)
	if err != nil {
		return nil, err
	}

	return repackSearchResultsToModel(userCourses), nil
}

func (s *Service) UploadUserAvatar(ctx context.Context, userID int32, avatar S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(avatar)
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateUserAvatar(ctx, userID, url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) UploadCourseAvatar(ctx context.Context, courseID int32, avatar S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(avatar)
	if err != nil {
		return "", err
	}

	_, err = s.repo.UpdateCourseAvatar(ctx, &courses.UpdateCourseAvatarParams{
		AvatarUrl: &url,
		ID:        courseID,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) UploadCoursePreviewVideo(ctx context.Context, courseID int32, video S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(video)
	if err != nil {
		return "", err
	}

	_, err = s.repo.UpdateCoursePreviewVideo(ctx, &courses.UpdateCoursePreviewVideoParams{
		PreviewVideoUrl: &url,
		ID:              courseID,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) UploadCourseLecture(ctx context.Context, courseID int32, lecture S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(lecture)
	if err != nil {
		return "", err
	}

	_, err = s.repo.InsertLectureAndCourseLecture(ctx, &courses.InsertLectureAndCourseLectureParams{
		CourseID: courseID,
		VideoUrl: url,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) GetCourseAvatarByFileID(ctx context.Context, fileID string) (*model.CourseIDWithResourceLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetCoursePreviewVideoByFileID(ctx context.Context, fileID string) (*model.CourseIDWithResourceLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetCourseLecturesByFileIDs(ctx context.Context, fileIDs []string) ([]*model.CourseIDWithResourceLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetCoursesAvatarsByFileIDs(ctx context.Context, fileIDs []string) ([]*model.CourseIDWithResourceLink, error) {
	//TODO implement me
	panic("implement me")
}
