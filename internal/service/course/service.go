package course

import (
	"context"
	"encoding/json"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/courses"
	"github.com/deevins/educational-platform-backend/internal/model"
	"time"
)

var _ handler.CourseService = &Service{}

type Service struct {
	repo courses.Querier
	s3   S3.Client
}

func (s *Service) UpdateTestTitle(ctx context.Context, testID int32, title string) (string, error) {
	_, err := s.repo.UpdateTestTitle(ctx, &courses.UpdateTestTitleParams{
		ID:   testID,
		Name: title,
	})
	if err != nil {
		return "", err
	}

	return title, nil
}

func (s *Service) UpdateLectureTitle(ctx context.Context, lectureID int32, title string) (string, error) {
	_, err := s.repo.UpdateLectureTitle(ctx, &courses.UpdateLectureTitleParams{
		ID:    lectureID,
		Title: title,
	})
	if err != nil {
		return "", err
	}

	return title, nil
}

func (s *Service) UpdateSectionTitle(ctx context.Context, sectionID int32, title string) (string, error) {
	_, err := s.repo.UpdateSectionTitle(ctx, &courses.UpdateSectionTitleParams{
		ID:    sectionID,
		Title: title,
	})
	if err != nil {
		return "", err
	}

	return title, nil
}

func (s *Service) RemoveLectureByID(ctx context.Context, lectureID int32) error {
	_, err := s.repo.RemoveLecture(ctx, &courses.RemoveLectureParams{
		ID: lectureID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveTestByID(ctx context.Context, testID int32) error {
	_, err := s.repo.RemoveTest(ctx, &courses.RemoveTestParams{
		ID: testID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveSectionByID(ctx context.Context, sectionID int32) error {
	_, err := s.repo.RemoveSection(ctx, sectionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateSection(ctx context.Context, courseID int32, input *model.SectionCreation) (int32, error) {
	sectionID, err := s.repo.CreateSection(ctx, &courses.CreateSectionParams{
		CourseID:     courseID,
		Title:        input.Title,
		Description:  input.Description,
		SerialNumber: input.SerialNumber,
	})
	if err != nil {
		return 0, err
	}

	return sectionID, nil
}

func (s *Service) CreateCourseTest(ctx context.Context, sectionID int32, test *model.CreateTestBase) (int32, error) {
	id, err := s.repo.CreateTest(ctx, &courses.CreateTestParams{
		SectionID:    sectionID,
		Name:         test.Title,
		SerialNumber: test.SerialNumber,
		Description:  test.Description,
	})
	if err != nil {
		return 0, err

	}
	return id, nil
}

func (s *Service) CreateCourseLecture(ctx context.Context, sectionID int32, lecture *model.CreateLectureBase) (int32, error) {
	id, err := s.repo.CreateLecture(ctx, &courses.CreateLectureParams{
		SectionID:    sectionID,
		Title:        lecture.Title,
		SerialNumber: lecture.SerialNumber,
		Description:  lecture.Description,
	})
	if err != nil {
		return 0, err

	}
	return id, nil
}

func (s *Service) AddQuestionsToTest(ctx context.Context, testID int32, questions []*model.Question) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo courses.Querier, s3Client S3.Client) *Service {
	return &Service{
		repo: repo,
		s3:   s3Client,
	}
}

func (s *Service) RemoveCourseByID(ctx context.Context, courseID int32) error {
	// TODO: remove connected resources from S3 and DB
	_, err := s.repo.RemoveCourse(ctx, courseID)
	if err != nil {
		return err
	}

	return nil
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

func (s *Service) GetInstructorCourses(ctx context.Context, instructorID int32) ([]*model.InstructorCourse, error) {
	instructorCourses, err := s.repo.GetInstructorCourses(ctx, instructorID)
	if err != nil {
		return nil, err
	}

	return repackInstructorCoursesToModel(instructorCourses)
}

func (s *Service) SearchInstructionCoursesByTitle(ctx context.Context, instructorID int32, title string) ([]*model.InstructorCourse, error) {
	instructorCourses, err := s.repo.SearchInstructorCoursesByTitle(ctx, &courses.SearchInstructorCoursesByTitleParams{
		Title:    title,
		AuthorID: instructorID,
	})

	if err != nil {
		return nil, err

	}

	return repackSearchInstructorCoursesByTitleToModel(instructorCourses)
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
	id, err := s.repo.ApproveCourse(ctx, ID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) SendCourseToCheck(ctx context.Context, ID int32) (int32, error) {
	id, err := s.repo.SendCourseToCheck(ctx, ID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) RejectCourse(ctx context.Context, ID int32) (int32, error) {
	id, err := s.repo.RejectCourse(ctx, ID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) GetFullCoursePageInfoByCourseID(ctx context.Context, ID int32) (*model.Course, error) {
	fc, err := s.repo.GetFullCourseInfoWithInstructorByCourseID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:              fc.CourseID,
		Title:           fc.Title,
		Subtitle:        *fc.Subtitle,
		Description:     fc.Description,
		Language:        *fc.Language,
		AvatarURL:       *fc.CourseAvatarUrl,
		Requirements:    fc.Requirements,
		Level:           *fc.Level,
		LecturesLength:  time.Duration(*fc.LecturesLength),
		LecturesCount:   int(*fc.LecturesCount),
		StudentsCount:   int(*fc.StudentsCount),
		ReviewsCount:    int(*fc.RatingsCount),
		Rating:          *fc.Rating,
		PreviewVideoURL: *fc.PreviewVideoUrl,
		TargetAudience:  fc.TargetAudience,
		CourseGoals:     fc.CourseGoals,
		Instructor: model.CourseInstructor{
			ID:            fc.InstructorID,
			FullName:      fc.InstructorFullName,
			AvatarURL:     *fc.InstructorAvatarUrl,
			Description:   *fc.InstructorDescription,
			StudentsCount: fc.InstructorStudentsCount,
			CoursesCount:  fc.InstructorCoursesCount,
			Rating:        float64(fc.InstructorRating),
		},
		CreatedAt: fc.CourseCreatedAt.Time,
		Status:    string(fc.CourseStatus),
		Category:  fc.CategoryTitle,
	}, nil
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

func (s *Service) UploadCourseLecture(ctx context.Context, lectureID int32, lecture S3.FileDataType) (string, error) {
	url, err := s.s3.CreateOne(lecture)
	if err != nil {
		return "", err
	}

	_, err = s.repo.UpdateLectureVideoUrl(ctx, &courses.UpdateLectureVideoUrlParams{
		VideoUrl: url,
		ID:       lectureID,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) GetCourseAvatarByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error) {
	url, err := s.repo.GetCourseAvatarByID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return &model.CourseIDWithResourceLink{
		CourseID: courseID,
		Link:     *url,
	}, nil
}

func (s *Service) GetCoursePreviewVideoByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error) {
	url, err := s.repo.GetCoursePreviewVideoByID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return &model.CourseIDWithResourceLink{
		CourseID: courseID,
		Link:     *url,
	}, nil
}

func (s *Service) GetCoursesAvatarsByCourseIDs(ctx context.Context, courseIDs []int32) ([]*model.CourseIDWithResourceLink, error) {
	ds, err := s.repo.GetCoursesAvatarsByIDs(ctx, courseIDs)
	if err != nil {
		return nil, err
	}

	return repackCourseAvatarsToModel(ds), nil
}

func repackCourseAvatarsToModel(ds []*courses.GetCoursesAvatarsByIDsRow) []*model.CourseIDWithResourceLink {
	var res []*model.CourseIDWithResourceLink
	for _, d := range ds {
		res = append(res, &model.CourseIDWithResourceLink{
			CourseID: d.ID,
			Link:     *d.AvatarUrl,
		})
	}

	return res
}

func (s *Service) UpdateCourseGoals(ctx context.Context, courseID int32, goals *model.UpdateCourseGoals) error {
	_, err := s.repo.UpdateCourseGoals(ctx, &courses.UpdateCourseGoalsParams{
		CourseGoals:    goals.Goals,
		Requirements:   goals.Requirements,
		TargetAudience: goals.TargetAudience,
		ID:             courseID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetCourseSectionsWithDataByCourseID(ctx context.Context, courseID int32) ([]*model.CourseSection, error) {
	rows, err := s.repo.GetSectionsWithLecturesAndTestsByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	var sections []*model.CourseSection
	for _, row := range rows {
		var lectures []*model.Lecture
		var tests []*model.Test

		err := json.Unmarshal(row.Lectures, &lectures)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(row.Tests, &tests)
		if err != nil {
			return nil, err
		}

		section := &model.CourseSection{
			SectionID:          row.SectionID,
			SectionTitle:       row.SectionTitle,
			SectionDescription: row.SectionDescription,
			Lectures:           lectures,
			Tests:              tests,
		}

		sections = append(sections, section)
	}

	return sections, nil
}
