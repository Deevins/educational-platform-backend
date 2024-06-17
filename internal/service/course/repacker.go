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
			StudentsCount:   lo.FromPtrOr(course.StudentsCount, 0),
			LecturesLength:  time.Duration(course.LecturesLengthInterval.Microseconds / 1000000 / 60),
			Description:     course.Description,
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

func repackInstructorCoursesToModel(instructorCourses []*courses.GetInstructorCoursesRow) []*model.InstructorCourse {
	var coursesList []*model.InstructorCourse
	for _, c := range instructorCourses {
		coursesList = append(coursesList, &model.InstructorCourse{
			ID:           c.ID,
			Title:        c.Title,
			AvatarURL:    lo.FromPtrOr(c.AvatarUrl, ""),
			Rating:       lo.FromPtrOr(c.Rating, 0.0),
			ReviewsCount: lo.FromPtrOr(c.RatingsCount, 0),
		})
	}

	return coursesList
}

func repackSearchInstructorCoursesByTitleToModel(courses []*courses.SearchInstructorCoursesByTitleRow) []*model.InstructorCourse {
	var coursesList []*model.InstructorCourse
	for _, c := range courses {
		coursesList = append(coursesList, &model.InstructorCourse{
			ID:           c.ID,
			Title:        c.Title,
			AvatarURL:    lo.FromPtrOr(c.AvatarUrl, ""),
			Rating:       lo.FromPtrOr(c.Rating, 0.0),
			ReviewsCount: lo.FromPtrOr(c.RatingsCount, 0),
		})
	}

	return coursesList
}

func repackData(
	sections []*courses.GetSectionsByCourseIDRow,
	tests []*courses.GetTestsByCourseIDRow,
	lectures []*courses.GetLecturesByCourseIDRow,
) []*model.CourseSection {
	sectionMap := make(map[int32]*model.CourseSection)

	for _, sec := range sections {
		section := &model.CourseSection{
			SectionID:          sec.SectionID,
			SectionTitle:       sec.SectionTitle,
			SectionDescription: sec.SectionDescription,
			SerialNumber:       sec.SectionSerialNumber,
			Lectures:           []*model.Lecture{},
			Tests:              []*model.Test{},
		}
		sectionMap[sec.SectionID] = section
	}

	for _, lec := range lectures {
		section, exists := sectionMap[lec.SectionID]
		if exists {
			lecture := &model.Lecture{
				ID:           lo.FromPtr(lec.LectureID),
				SerialNumber: lo.FromPtrOr(lec.LectureSerialNumber, 0),
				Title:        lo.FromPtrOr(lec.LectureTitle, ""),
				Type:         "lecture",
				Description:  lo.FromPtrOr(lec.LectureDescription, ""),
				VideoURL:     lo.FromPtrOr(lec.LectureVideoUrl, ""),
			}
			section.Lectures = append(section.Lectures, lecture)
		}
	}

	for _, tst := range tests {
		section, exists := sectionMap[tst.SectionID]
		if exists {
			var test *model.Test
			for _, t := range section.Tests {
				if t.TestID == lo.FromPtr(tst.TestID) {
					test = t
					break
				}
			}

			if test == nil {
				test = &model.Test{
					TestID:       lo.FromPtrOr(tst.TestID, 0),
					TestName:     lo.FromPtrOr(tst.TestName, ""),
					Description:  lo.FromPtrOr(tst.TestDescription, ""),
					Type:         "test",
					SerialNumber: lo.FromPtrOr(tst.TestSerialNumber, 0),
					Questions:    []model.Question{},
				}
				section.Tests = append(section.Tests, test)
			}

			if tst.QuestionID != nil {
				var question *model.Question
				for i := range test.Questions {
					if test.Questions[i].QuestionBody == lo.FromPtrOr(tst.QuestionBody, "") {
						question = &test.Questions[i]
						break
					}
				}

				if question == nil {
					newQuestion := model.Question{
						ID:           lo.FromPtr(tst.QuestionID),
						QuestionBody: lo.FromPtrOr(tst.QuestionBody, ""),
						Answers:      []model.Response{},
					}
					test.Questions = append(test.Questions, newQuestion)
					question = &test.Questions[len(test.Questions)-1]
				}

				if tst.AnswerID != nil {
					answer := model.Response{
						ResponseText: lo.FromPtrOr(tst.AnswerBody, ""),
						Description:  lo.FromPtrOr(tst.AnswerDescription, ""),
						IsCorrect:    lo.FromPtrOr(tst.AnswerIsCorrect, false),
					}
					question.Answers = append(question.Answers, answer)
				}
			}
		}
	}

	var courseSections []*model.CourseSection
	for _, section := range sectionMap {
		courseSections = append(courseSections, section)
	}

	return courseSections
}
