package handler

import (
	"context"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/deevins/educational-platform-backend/pkg/httpResponses"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type CourseService interface {
	GetUserCoursesByUserID(ctx context.Context, ID int32) ([]*model.ShortCourse, error)
	GetFullCoursePageInfoByCourseID(ctx context.Context, ID int32) (*model.Course, error)
	GetAllReadyCourses(ctx context.Context) ([]*model.ShortCourse, error)
	GetAllPendingCourses(ctx context.Context) ([]*model.ShortCourse, error)
	GetAllDraftCourses(ctx context.Context) ([]*model.ShortCourse, error)
	ApproveCourse(ctx context.Context, ID int32) (int32, error)
	RejectCourse(ctx context.Context, ID int32) (int32, error)
	SendCourseToCheck(ctx context.Context, ID int32) (int32, error)
	CreateCourseBase(ctx context.Context, base *model.CourseBase) (int32, error)
	SearchCoursesByTitle(ctx context.Context, title string) ([]*model.ShortCourse, error)

	CreateSection(ctx context.Context, courseID int32, input *model.SectionCreation) (int32, error)
	CreateCourseTest(ctx context.Context, sectionID int32, test *model.CreateTestBase) (int32, error)
	CreateCourseLecture(ctx context.Context, sectionID int32, lecture *model.CreateLectureBase) (int32, error)
	AddQuestionsToTest(ctx context.Context, testID int32, question []*model.Question) (int32, error)

	UploadCourseAvatar(ctx context.Context, courseID int32, avatar S3.FileDataType) (string, error)
	UploadCoursePreviewVideo(ctx context.Context, courseID int32, video S3.FileDataType) (string, error)
	UploadCourseLecture(ctx context.Context, lectureID int32, lecture S3.FileDataType) (string, error)
	//UploadCourseTest(ctx context.Context, sectionID int32, test *model.Test) (string, error)

	GetCourseAvatarByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error)
	GetCoursePreviewVideoByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error)
	GetCoursesAvatarsByCourseIDs(ctx context.Context, courseIDs []int32) ([]*model.CourseIDWithResourceLink, error)
	GetInstructorCourses(ctx context.Context, instructorID int32) ([]*model.InstructorCourse, error)
	SearchInstructionCoursesByTitle(ctx context.Context, instructorID int32, title string) ([]*model.InstructorCourse, error)
	RemoveCourseByID(ctx context.Context, courseID int32) error
	UpdateCourseGoals(ctx context.Context, courseID int32, goals *model.UpdateCourseGoals) error
	GetCourseSectionsWithDataByCourseID(ctx context.Context, courseID int32) ([]*model.CourseSection, error)

	UpdateTestTitle(ctx context.Context, testID int32, title string) (string, error)
	UpdateLectureTitle(ctx context.Context, lectureID int32, title string) (string, error)
	UpdateSectionTitle(ctx context.Context, sectionID int32, title string) (string, error)
	RemoveLectureByID(ctx context.Context, lectureID int32) error
	RemoveTestByID(ctx context.Context, testID int32) error
	RemoveSectionByID(ctx context.Context, sectionID int32) error
}

func (h *Handler) getFullCoursePage(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	course, err := h.cs.GetFullCoursePageInfoByCourseID(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (h *Handler) getAllCourses(ctx *gin.Context) {
	courses, err := h.cs.GetAllReadyCourses(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) getLatestEightCourses(ctx *gin.Context) {
	courses, err := h.cs.GetAllReadyCourses(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses[:8])
}

func (h *Handler) createCourseBase(ctx *gin.Context) {
	var input *model.CourseBase

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courseID, err := h.cs.CreateCourseBase(ctx, input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"course_id": courseID})
}

func (h *Handler) deleteCourse(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	err = h.cs.RemoveCourseByID(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "course deleted"})
}

func (h *Handler) searchCoursesByTitle(ctx *gin.Context) {
	if ctx.Param("query") == "" {
		ctx.JSON(400, gin.H{"error": "query is empty"})
		return
	}

	query := ctx.Param("query")

	courses, err := h.cs.SearchCoursesByTitle(ctx, query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

// here we must fetch courses ne smotrya na status
func (h *Handler) getAllCoursesByInstructorID(ctx *gin.Context) {
	if ctx.Param("instructorID") == "" {
		ctx.JSON(400, gin.H{"error": "instructorID is empty"})
		return
	}

	instructorID, err := strconv.ParseInt(ctx.Param("instructorID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "instructorID is not a number"})
		return
	}

	courses, err := h.cs.GetInstructorCourses(ctx, int32(instructorID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) updateCourseGoals(ctx *gin.Context) {
	var goals *model.UpdateCourseGoals

	if ctx.ShouldBindJSON(goals) != nil {
		ctx.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	err = h.cs.UpdateCourseGoals(ctx, int32(courseID), goals)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "goals updated"})
}

func (h *Handler) sendToCheck(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	_, err = h.cs.SendCourseToCheck(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "course sent to check"})
}

func (h *Handler) approveCourse(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	_, err = h.cs.ApproveCourse(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "course approved"})
}

func (h *Handler) rejectCourse(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	_, err = h.cs.RejectCourse(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "course rejected"})
}

func (h *Handler) uploadCourseMaterials(ctx *gin.Context) {

}

func (h *Handler) getCoursesWaitingForApproval(ctx *gin.Context) {
	courses, err := h.cs.GetAllPendingCourses(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) getCoursesApproved(ctx *gin.Context) {
	courses, err := h.cs.GetAllReadyCourses(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) getCoursesByUserID(ctx *gin.Context) {
	fmt.Println(ctx.Param("userID"))
	if ctx.Param("userID") == "" {
		ctx.JSON(400, gin.H{"error": "userID is empty"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "userID is not a number"})
		return
	}

	courses, err := h.cs.GetUserCoursesByUserID(ctx, int32(userID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userID,
		"courses": courses,
	})
}

func (h *Handler) uploadCourseAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
		return
	}

	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		ctx.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to close file: %v", err)
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Чтение содержимого файла
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to read file: %v", err)
		return
	}

	url, err := h.cs.UploadCourseAvatar(ctx, int32(courseID), S3.FileDataType{
		FileName: file.Filename,
		Data:     fileBytes,
	})
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	ctx.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    url, // URL-адрес загруженного файла
	})
}

func (h *Handler) uploadCoursePreviewVideo(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
		return
	}

	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		ctx.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to close file: %v", err)
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Чтение содержимого файла
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to read file: %v", err)
		return
	}

	url, err := h.cs.UploadCourseAvatar(ctx, int32(courseID), S3.FileDataType{
		FileName: file.Filename,
		Data:     fileBytes,
	})
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	ctx.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    url, // URL-адрес загруженного файла
	})
}

func (h *Handler) getCourseMaterials(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}
	sections, err := h.cs.GetCourseSectionsWithDataByCourseID(ctx, int32(courseID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sections)
}

func (h *Handler) createSection(ctx *gin.Context) {
	if ctx.Param("courseID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	var input *model.SectionCreation
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.cs.CreateSection(ctx, int32(courseID), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "section created"})
}

func (h *Handler) createLecture(ctx *gin.Context) {
	if ctx.Param("sectionID") == "" {
		ctx.JSON(400, gin.H{"error": "sectionID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("sectionID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "sectionID is not a number"})
		return
	}
	var input *model.CreateLectureBase
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.cs.CreateCourseLecture(ctx, int32(courseID), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "lecture created"})
}

func (h *Handler) createTest(ctx *gin.Context) {
	if ctx.Param("sectionID") == "" {
		ctx.JSON(400, gin.H{"error": "sectionID is empty"})
		return
	}

	courseID, err := strconv.ParseInt(ctx.Param("sectionID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "sectionID is not a number"})
		return
	}
	var input *model.CreateTestBase
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.cs.CreateCourseTest(ctx, int32(courseID), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "test created"})
}

func (h *Handler) removeSection(ctx *gin.Context) {
	if ctx.Param("sectionID") == "" {
		ctx.JSON(400, gin.H{"error": "sectionID is empty"})
		return
	}

	sectionID, err := strconv.ParseInt(ctx.Param("sectionID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "sectionID is not a number"})
		return
	}

	err = h.cs.RemoveSectionByID(ctx, int32(sectionID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "section deleted"})
}

func (h *Handler) removeLecture(ctx *gin.Context) {
	if ctx.Param("lectureID") == "" {
		ctx.JSON(400, gin.H{"error": "lectureID is empty"})
		return
	}

	lectureID, err := strconv.ParseInt(ctx.Param("lectureID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "lectureID is not a number"})
		return
	}

	err = h.cs.RemoveLectureByID(ctx, int32(lectureID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "lecture deleted"})
}

func (h *Handler) removeTest(ctx *gin.Context) {
	if ctx.Param("testID") == "" {
		ctx.JSON(400, gin.H{"error": "testID is empty"})
		return
	}

	testID, err := strconv.ParseInt(ctx.Param("testID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "testID is not a number"})
		return
	}

	err = h.cs.RemoveTestByID(ctx, int32(testID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "lecture deleted"})

}

type UpdateSectionTitle struct {
	Title string `json:"title"`
}

func (h *Handler) updateSectionTitle(ctx *gin.Context) {
	if ctx.Param("sectionID") == "" {
		ctx.JSON(400, gin.H{"error": "sectionID is empty"})
		return
	}

	sectionID, err := strconv.ParseInt(ctx.Param("sectionID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "sectionID is not a number"})
		return
	}

	var input *UpdateSectionTitle
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title, err := h.cs.UpdateSectionTitle(ctx, int32(sectionID), input.Title)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": title})
}

type UpdateLectureTitle struct {
	Title string `json:"title"`
}

func (h *Handler) updateLectureTitle(ctx *gin.Context) {
	if ctx.Param("lectureID") == "" {
		ctx.JSON(400, gin.H{"error": "lectureID is empty"})
		return
	}

	lectureID, err := strconv.ParseInt(ctx.Param("lectureID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "lectureID is not a number"})
		return
	}
	var input *UpdateLectureTitle
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title, err := h.cs.UpdateLectureTitle(ctx, int32(lectureID), input.Title)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": title})
}

type UpdateTestTitle struct {
	Title string `json:"title"`
}

func (h *Handler) updateTestTitle(ctx *gin.Context) {
	if ctx.Param("testID") == "" {
		ctx.JSON(400, gin.H{"error": "testID is empty"})
		return
	}

	testID, err := strconv.ParseInt(ctx.Param("testID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "testID is not a number"})
		return
	}

	var input *UpdateTestTitle
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title, err := h.cs.UpdateTestTitle(ctx, int32(testID), input.Title)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": title})

}

func (h *Handler) addQuestionsToTest(ctx *gin.Context) {
	if ctx.Param("testID") == "" {
		ctx.JSON(400, gin.H{"error": "testID is empty"})
		return
	}

	testID, err := strconv.ParseInt(ctx.Param("testID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "testID is not a number"})
		return
	}

	var input []*model.Question
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	_, err = h.cs.AddQuestionsToTest(ctx, int32(testID), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "questions added"})
}

func (h *Handler) uploadLectureVideo(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
		return
	}

	if ctx.Param("lectureID") == "" {
		ctx.JSON(400, gin.H{"error": "courseID is empty"})
		return
	}

	lectureID, err := strconv.ParseInt(ctx.Param("lectureID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "lectureID is not a number"})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		ctx.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to close file: %v", err)
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Чтение содержимого файла
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to read file: %v", err)
		return
	}

	url, err := h.cs.UploadCourseLecture(ctx, int32(lectureID), S3.FileDataType{
		FileName: file.Filename,
		Data:     fileBytes,
	})
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	ctx.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    url, // URL-адрес загруженного файла
	})
}
