// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package courses

import (
	"context"
)

type Querier interface {
	AddCourseBasicInfo(ctx context.Context, arg *AddCourseBasicInfoParams) (int32, error)
	ApproveCourse(ctx context.Context, id int32) (int32, error)
	CancelPublishingCourse(ctx context.Context, id int32) (int32, error)
	CreateAnswer(ctx context.Context, arg *CreateAnswerParams) (int32, error)
	CreateCourseBase(ctx context.Context, arg *CreateCourseBaseParams) (int32, error)
	CreateLecture(ctx context.Context, arg *CreateLectureParams) (int32, error)
	CreateQuestion(ctx context.Context, arg *CreateQuestionParams) (int32, error)
	CreateSection(ctx context.Context, arg *CreateSectionParams) (int32, error)
	CreateTest(ctx context.Context, arg *CreateTestParams) (int32, error)
	GetAllDraftCourses(ctx context.Context) ([]*GetAllDraftCoursesRow, error)
	GetAllPendingCourses(ctx context.Context) ([]*GetAllPendingCoursesRow, error)
	GetAllReadyCourses(ctx context.Context) ([]*GetAllReadyCoursesRow, error)
	GetCourseAvatarByID(ctx context.Context, id int32) (*string, error)
	GetCourseBasicInfo(ctx context.Context, id int32) (*GetCourseBasicInfoRow, error)
	GetCourseGoals(ctx context.Context, id int32) (*GetCourseGoalsRow, error)
	GetCoursePreviewVideoByID(ctx context.Context, id int32) (*string, error)
	GetCourseReviewsByCourseID(ctx context.Context, id int32) ([]*GetCourseReviewsByCourseIDRow, error)
	GetCourseSectionSerialNumber(ctx context.Context, id int32) (int32, error)
	GetCoursesAvatarsByIDs(ctx context.Context, dollar_1 []int32) ([]*GetCoursesAvatarsByIDsRow, error)
	GetFullCourseByID(ctx context.Context, id int32) (*GetFullCourseByIDRow, error)
	GetFullCourseInfoWithInstructorByCourseID(ctx context.Context, id int32) (*GetFullCourseInfoWithInstructorByCourseIDRow, error)
	GetInstructorCourses(ctx context.Context, authorID int32) ([]*GetInstructorCoursesRow, error)
	GetLectureByID(ctx context.Context, id int32) (*GetLectureByIDRow, error)
	GetLectureSerialNumber(ctx context.Context, sectionID int32) (int32, error)
	GetLecturesByCourseID(ctx context.Context, courseID int32) ([]*GetLecturesByCourseIDRow, error)
	GetSectionsByCourseID(ctx context.Context, courseID int32) ([]*GetSectionsByCourseIDRow, error)
	GetSectionsWithLecturesAndTestsByCourseID(ctx context.Context, courseID int32) ([]*GetSectionsWithLecturesAndTestsByCourseIDRow, error)
	GetTestSerialNumber(ctx context.Context, sectionID int32) (int32, error)
	GetTestsByCourseID(ctx context.Context, courseID int32) ([]*GetTestsByCourseIDRow, error)
	GetUserCourses(ctx context.Context, userID int32) ([]*GetUserCoursesRow, error)
	RejectCourse(ctx context.Context, id int32) (int32, error)
	RemoveAnswer(ctx context.Context, id int32) (int32, error)
	RemoveCourse(ctx context.Context, id int32) (int32, error)
	RemoveLecture(ctx context.Context, arg *RemoveLectureParams) (int32, error)
	RemoveQuestion(ctx context.Context, id int32) (int32, error)
	RemoveSection(ctx context.Context, id int32) (int32, error)
	RemoveTest(ctx context.Context, arg *RemoveTestParams) (int32, error)
	SearchCoursesByTitle(ctx context.Context, title string) ([]*SearchCoursesByTitleRow, error)
	SearchInstructorCoursesByTitle(ctx context.Context, arg *SearchInstructorCoursesByTitleParams) ([]*SearchInstructorCoursesByTitleRow, error)
	SendCourseToCheck(ctx context.Context, id int32) (int32, error)
	UpdateAnswer(ctx context.Context, arg *UpdateAnswerParams) (int32, error)
	UpdateCourseAvatar(ctx context.Context, arg *UpdateCourseAvatarParams) (*string, error)
	UpdateCourseGoals(ctx context.Context, arg *UpdateCourseGoalsParams) (int32, error)
	UpdateCoursePreviewVideo(ctx context.Context, arg *UpdateCoursePreviewVideoParams) (*string, error)
	UpdateLectureTitle(ctx context.Context, arg *UpdateLectureTitleParams) (int32, error)
	UpdateLectureVideoAddedInfo(ctx context.Context, arg *UpdateLectureVideoAddedInfoParams) (int32, error)
	UpdateLectureVideoUrl(ctx context.Context, arg *UpdateLectureVideoUrlParams) (int32, error)
	UpdateLecturesInfo(ctx context.Context, arg *UpdateLecturesInfoParams) (int32, error)
	UpdateQuestion(ctx context.Context, arg *UpdateQuestionParams) (int32, error)
	UpdateSectionTitle(ctx context.Context, arg *UpdateSectionTitleParams) (int32, error)
	UpdateTestTitle(ctx context.Context, arg *UpdateTestTitleParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
