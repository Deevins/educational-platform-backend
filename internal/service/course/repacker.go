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
			ID:            course.ID,
			Title:         course.Title,
			AvatarURL:     lo.FromPtrOr(course.AvatarUrl, ""),
			Rating:        lo.FromPtrOr(course.Rating, 0.0),
			StudentsCount: lo.FromPtrOr(course.StudentsCount, 0),
		})
	}

	return modelCourses
}

func repackDBCoursesToModel(courses []*courses.HumanResourcesCourse) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:             course.ID,
			Title:          course.Title,
			AvatarURL:      lo.FromPtrOr(course.AvatarUrl, ""),
			Subtitle:       lo.FromPtrOr(course.Subtitle, ""),
			Rating:         lo.FromPtrOr(course.Rating, 0.0),
			StudentsCount:  lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength: time.Duration(lo.FromPtrOr(course.LecturesLength, 0)),
			Description:    course.Description,
		})
	}

	return modelCourses
}

func repackDBCoursesToShortModel(courses []*courses.GetUserCoursesRow) []*model.ShortCourse {
	var modelCourses []*model.ShortCourse

	for _, course := range courses {
		modelCourses = append(modelCourses, &model.ShortCourse{
			ID:             course.ID,
			Title:          course.Title,
			AvatarURL:      lo.FromPtrOr(course.AvatarUrl, ""),
			Subtitle:       lo.FromPtrOr(course.Subtitle, ""),
			Rating:         lo.FromPtrOr(course.Rating, 0.0),
			StudentsCount:  lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength: time.Duration(lo.FromPtrOr(course.LecturesLength, 0)),
			Description:    course.Description,
		})
	}

	return modelCourses
}

func mapTypeToDBType(courseType string) courses.NullHumanResourcesCourseTypes {
	switch courseType {
	case "course":
		return courses.NullHumanResourcesCourseTypes{
			HumanResourcesCourseTypes: courses.HumanResourcesCourseTypesCourse,
			Valid:                     true,
		}
	case "practice":
		return courses.NullHumanResourcesCourseTypes{
			HumanResourcesCourseTypes: courses.HumanResourcesCourseTypesPracticeCourse,
			Valid:                     true,
		}
	default:
		return courses.NullHumanResourcesCourseTypes{
			HumanResourcesCourseTypes: courses.HumanResourcesCourseTypesCourse,
			Valid:                     true,
		}
	}

}
