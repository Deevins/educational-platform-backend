package course

import (
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/courses"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/samber/lo"
	"time"
)

func repackSearchResultsToModel(courses []*courses.SearchCoursesByTitleRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:              course.ID,
			Title:           course.Title,
			CourseAvatarURL: lo.FromPtrOr(course.AvatarUrl, ""),
			Rating:          lo.FromPtrOr(course.Rating, 0.0),
			StudentsCount:   lo.FromPtrOr(course.StudentsCount, 0),
		})
	}

	return modelCourses
}

func repackDBPendingCoursesToModel(courses []*courses.GetAllPendingCoursesRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:              course.CourseID,
			Title:           course.Title,
			AuthorFullName:  course.InstructorName,
			CourseAvatarURL: lo.FromPtrOr(course.CourseAvatarUrl, ""),
			Subtitle:        lo.FromPtrOr(course.Subtitle, ""),
			Rating:          lo.FromPtrOr(course.Rating, 0.0),
			Level:           lo.FromPtrOr(course.Level, ""),
			ReviewsCount:    lo.FromPtrOr(course.RatingsCount, 0),
			LecturesCount:   lo.FromPtrOr(course.LecturesCount, 0),
			StudentsCount:   lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength:  time.Duration(course.LecturesLengthInterval.Microseconds / 1000000 / 60),
			Description:     course.Description,
		})
	}

	return modelCourses
}
func repackDBReadyCoursesToModel(courses []*courses.GetAllReadyCoursesRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:              course.CourseID,
			Title:           course.Title,
			AuthorFullName:  course.InstructorName,
			CourseAvatarURL: lo.FromPtrOr(course.CourseAvatarUrl, ""),
			Subtitle:        lo.FromPtrOr(course.Subtitle, ""),
			Rating:          lo.FromPtrOr(course.Rating, 0.0),
			Level:           lo.FromPtrOr(course.Level, ""),
			ReviewsCount:    lo.FromPtrOr(course.RatingsCount, 0),
			LecturesCount:   lo.FromPtrOr(course.LecturesCount, 0),
			StudentsCount:   lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength:  time.Duration(course.LecturesLengthInterval.Microseconds / 1000000 / 60),
			Description:     course.Description,
		})
	}

	return modelCourses
}
func repackDBDraftCoursesToModel(courses []*courses.GetAllDraftCoursesRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:              course.CourseID,
			Title:           course.Title,
			AuthorFullName:  course.InstructorName,
			CourseAvatarURL: lo.FromPtrOr(course.CourseAvatarUrl, ""),
			Subtitle:        lo.FromPtrOr(course.Subtitle, ""),
			Rating:          lo.FromPtrOr(course.Rating, 0.0),
			Level:           lo.FromPtrOr(course.Level, ""),
			ReviewsCount:    lo.FromPtrOr(course.RatingsCount, 0),
			LecturesCount:   lo.FromPtrOr(course.LecturesCount, 0),
			StudentsCount:   lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength:  time.Duration(course.LecturesLengthInterval.Microseconds / 1000000 / 60),
			Description:     course.Description,
		})
	}

	return modelCourses
}

func repackDBCoursesToShortModel(courses []*courses.GetUserCoursesRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:              course.ID,
			Title:           course.Title,
			CourseAvatarURL: lo.FromPtrOr(course.AvatarUrl, ""),
			Subtitle:        lo.FromPtrOr(course.Subtitle, ""),
			Rating:          lo.FromPtrOr(course.Rating, 0.0),

			StudentsCount:  lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength: time.Duration(course.LecturesLengthInterval.Microseconds / 1000000 / 60),
			Description:    course.Description,
		})
	}

	return modelCourses
}

func mapTypeToDBType(courseType string) courses.HumanResourcesCourseTypes {
	switch courseType {
	case "course":

		return courses.HumanResourcesCourseTypesCourse
	case "practice":
		return courses.HumanResourcesCourseTypesPracticeCourse
	default:
		return courses.HumanResourcesCourseTypesCourse
	}

}

func repackInstructorCoursesToModel(instructorCourses []*courses.GetInstructorCoursesRow) ([]*model.InstructorCourse, error) {
	var coursesList []*model.InstructorCourse
	for _, c := range instructorCourses {

		coursesList = append(coursesList, &model.InstructorCourse{
			ID:        c.ID,
			Title:     c.Title,
			AvatarURL: *c.AvatarUrl,
			Status:    string(c.Status),
		})
	}

	return coursesList, nil
}

func repackSearchInstructorCoursesByTitleToModel(courses []*courses.SearchInstructorCoursesByTitleRow) ([]*model.InstructorCourse, error) {
	var coursesList []*model.InstructorCourse
	for _, c := range courses {
		coursesList = append(coursesList, &model.InstructorCourse{
			ID:        c.ID,
			Title:     c.Title,
			AvatarURL: *c.AvatarUrl,
			Status:    string(c.Status),
		})
	}

	return coursesList, nil
}
