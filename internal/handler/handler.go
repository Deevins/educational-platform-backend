package handler

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/deevins/educational-platform-backend/internal/service/auth"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	CreateUser(ctx context.Context, user model.UserCreate) (auth.RegisterUserResponse, error)
	Authorize(ctx context.Context, email, password string) (auth.LoginUserResponse, error)
}

type UserService interface {
	GetByID(ctx context.Context, ID int32) (*model.User, error)
	GetHasUserTriedInstructor(ctx context.Context, ID int32) (bool, error)
	SetHasUserTriedInstructorToTrue(ctx context.Context, ID int32) error
	GetSelfInfo(ctx context.Context, ID int32) (*model.User, error)
	UpdateAvatar(ctx context.Context, ID int32, avatar S3.FileDataType) error
	AddUserTeachingExperience(ctx context.Context, exp *model.UserUpdateTeachingExperience) error
	UpdateUserInfo(ctx context.Context, user *model.UserUpdate) error
	GetUserAvatarByFileID(ctx context.Context, fileID string) (*model.UserIDWithResourceLink, error)
}

type ThreadService interface{}

type DirectoryService interface {
	GetCategoriesWithSubCategories(ctx context.Context) ([]*model.Category, error)
	GetLanguages(ctx context.Context) ([]*model.Language, error)
}

type Handler struct {
	as AuthService
	us UserService
	cs CourseService
	ts ThreadService
	ds DirectoryService
}

func NewHandler(authSvc AuthService,
	userSvc UserService,
	courseSvc CourseService,
	threadSvc ThreadService,
	directorySvc DirectoryService) *Handler {
	return &Handler{as: authSvc, us: userSvc, cs: courseSvc, ts: threadSvc, ds: directorySvc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(corsMiddleware())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-up", h.signUp) // ok
		authGroup.POST("/sign-in", h.signIn) // ok
	}

	user := router.Group("/users")
	{
		user.GET("/get-one", h.getOneUser)
		user.GET("/has-user-tried-instructor", h.hasUserTriedInstructor)                       // ok
		user.POST("/set-has-user-tried-instructor-to-true", h.setHasUserTriedInstructorToTrue) // ok
		user.PUT("/update-avatar", h.updateAvatar)                                             // poxyi pokA
		user.POST("/add-user-teaching-experience", h.updateUserTeachingExperience)             // ok
		user.PUT("/update-user-info", h.updateUserInfo)                                        // ok
	}
	course := router.Group("/courses")
	{
		course.GET("/get-one/:courseID", h.getOneCourse)                    // temp OK
		course.GET("/get-courses-by-user-id/:userID", h.getCoursesByUserID) // для вывода курсов по id пользователя у него на странице или еще где ok
		course.GET("/get-all", h.getAllCourses)                             // OK courses with READY status for all courses page
		course.GET("/get-latest-eight", h.getLatestEightCourses)            // OK
		course.POST("/create-base", h.createCourseBase)                     // OK type must be 'course' or 'practice'
		//course.PUT("/update", h.updateCourse)
		//course.DELETE("/delete", h.deleteCourse) // poka ne nado
		course.GET("/search-courses-by-title", h.searchCoursesByTitle)
		course.GET("/get-all-courses-by-instructor-id", h.getAllCoursesByInstructorID) // для вывода курсов по id инструктора которые он создал
		course.PUT("/update-course-goals", h.updateCourseGoals)
		course.PUT("/update-course-curriculum", h.updateCourseCurriculum)
		course.PUT("/update/course/basics", h.updateCourseBasics)
		course.POST("/send-to-check", h.sendToCheck)
		course.POST("/approve-course", h.approveCourse)
		course.POST("/reject-course", h.rejectCourse)
		course.POST("/upload-course-materials", h.uploadCourseMaterials)
		course.GET("/get-full-course", h.getFullCourse)          // получаем всю инфу о курсе,который находится в статусе READY это где страница на которой его можно проходить
		course.GET("/get-full-course-to-check", h.getFullCourse) // получаем всю инфу о курсе, который находится в статусе PENDING, это где страница на которой его можно проходить
		course.GET("/get-courses-waiting-for-approval", h.getCoursesWaitingForApproval)
	}
	directories := router.Group("/directories")
	{
		directories.GET("/categories", h.getCategories)
		directories.GET("/languages", h.getLanguages)
		//directories.GET("/filter-by-category-and-subcategory", h.filterByCategoryAndSubcategory)
	}

	threads := router.Group("/threads")
	{
		threads.GET("/get-all-threads", h.getAllThreads)
		threads.POST("/create-thread", h.createThread)
		threads.POST("/answer-in-thread", h.answerInThread)
		threads.GET("/search-threads-by-title", h.searchThreadsByTitle)
		threads.GET("/get-one-thread", h.getOneThread) // here we must get all threads messages
		threads.POST("/add-tag-to-thread", h.addTagToThread)
	}

	files := router.Group("/files")
	{
		files.POST("/upload-user-avatar", h.uploadUserAvatar)
		files.POST("/upload-course-avatar", h.uploadCourseAvatar)
		files.POST("/upload-course-preview-video", h.uploadCoursePreviewVideo)
		files.POST("/upload-course-lecture", h.uploadCourseLecture)
		files.GET("/get-user-avatar/:objectID", h.getUserAvatarByFileID)
		files.GET("/get-course-avatar/:objectID", h.getCourseAvatarByFileID)
		files.GET("/get-course-preview-video/:objectID", h.getCoursePreviewVideoByFileID)
		files.GET("/get-course-preview-video/:objectID", h.getCourseLecturesByFileIDs)
		files.GET("/get-courses-avatars", h.getCoursesAvatarsByFileIDs) // for all courses page

	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
