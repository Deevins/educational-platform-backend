package handler

import (
	"context"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/gin-gonic/gin"
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

	UploadCourseAvatar(ctx context.Context, courseID int32, avatar S3.FileDataType) (string, error)
	UploadCoursePreviewVideo(ctx context.Context, courseID int32, video S3.FileDataType) (string, error)
	UploadCourseLecture(ctx context.Context, lectureID int32, lecture S3.FileDataType) (string, error)

	GetCourseAvatarByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error)
	GetCoursePreviewVideoByCourseID(ctx context.Context, courseID int32) (*model.CourseIDWithResourceLink, error)
	GetCoursesAvatarsByCourseIDs(ctx context.Context, courseIDs []int32) ([]*model.CourseIDWithResourceLink, error)
	GetInstructorCourses(ctx context.Context, instructorID int32) ([]*model.InstructorCourse, error)
	SearchInstructionCoursesByTitle(ctx context.Context, instructorID int32, title string) ([]*model.InstructorCourse, error)
	RemoveCourseByID(ctx context.Context, courseID int32) error
	UpdateCourseGoals(ctx context.Context, courseID int32, goals *model.UpdateCourseGoals) error
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

type SearchBase struct {
	Query string `json:"query"`
}

func (h *Handler) searchCoursesByTitle(ctx *gin.Context) {
	var input *SearchBase

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courses, err := h.cs.SearchCoursesByTitle(ctx, input.Query)
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

// "id":          course.ID,
// "title":       course.Title,
// "description": course.Description,
// "price":       course.Price,
// "language":    course.Language,
// "category":    course.Category,
// "level":       course.Level,
// "rating":      course.Rating,
// "students":    course.Students,
// "image":       course.Image,
// "video":       course.Video,
// "created_at":  course.CreatedAt,
// "updated_at":  course.UpdatedAt,
// "instructor":  course.Instructor,
// "reviews":     course.Reviews,
// "curriculum":  course.Curriculum,
// "materials":   course.Materials,
// "goals":       course.Goals,
// "requirements": course.Requirements,
// "audience":    course.Audience,
// "learn":       course.Learn,
// "lectures":    course.Lectures,
// "duration":    course.Duration,
// "status":      course.Status,
// "approved_at": course.ApprovedAt,
// "rejected_at": course.RejectedAt,
// "checked_at":  course.CheckedAt,
// "approved_by": course.ApprovedBy,
// "rejected_by": course.RejectedBy,
// "checked_by":  course.CheckedBy,
// "price":       course.Price,
// "discount":    course.Discount,
// "discount_end": course.DiscountEnd,
// "discount_price": course.DiscountPrice,
// "discount_price_end": course.DiscountPriceEnd,
