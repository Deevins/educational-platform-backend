package handler

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/deevins/educational-platform-backend/internal/service/auth_service"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	CreateUser(ctx context.Context, user model.UserCreate) (auth_service.RegisterUserResponse, error)
	GenerateToken(ctx context.Context, email, password string) (auth_service.LoginUserResponse, error)
}

type UserService interface {
	GetByID(ctx context.Context, ID int32) (*model.User, error)
	GetHasUserTriedInstructor(ctx context.Context, ID int32) (bool, error)
	SetHasUserTriedInstructorToTrue(ctx context.Context, ID int32) error
	GetSelfInfo(ctx context.Context, ID int32) (*model.User, error)
	UpdateAvatar(ctx context.Context, ID int32, avatar []byte) error
	AddUserTeachingExperience(ctx context.Context, exp *model.UserUpdateTeachingExperience) error
	UpdateUserInfo(ctx context.Context, user model.UserUpdate) error
}

type CourseService interface{}

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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	user := router.Group("/user")
	{
		user.GET("/get-one", h.getOneUser)
		user.GET("/has-user-tried-instructor", h.hasUserTriedInstructor)
		user.POST("/set-has-user-tried-instructor-to-true", h.setHasUserTriedInstructorToTrue) // dunno
		user.GET("/get-self-info", h.getSelfInfo)
		user.PUT("/update-avatar", h.updateAvatar)
		user.POST("/add-user-teaching-experience", h.updateUserTeachingExperience) //
		user.PUT("/update-user-info", h.updateUserInfo)                            // dunno
	}
	course := router.Group("/course")
	{
		course.GET("/get-one", h.getOneCourse)
		course.GET("/get-courses-by-user-id", h.getAllCoursesByUserID) // для вывода курсов по id пользователя у него на странице или еще где
		course.GET("/get-all", h.getAllCourses)                        // add check for course status
		course.GET("/get-last-eight", h.getLastEightCourses)
		course.POST("/create-base", h.createCourseBase)
		course.PUT("/update", h.updateCourse)
		course.DELETE("/delete", h.deleteCourse)
		course.GET("/search-courses-by-title", h.searchCoursesByTitle)
		course.GET("/get-all-courses-by-instructor-id", h.getAllCoursesByInstructorID) // для вывода курсов по id инструктора которые он создал
		course.PUT("/update-course-goals", h.updateCourseGoals)
		course.PUT("/update-course-curriculum", h.updateCourseCurriculum)
		course.PUT("/update/course/basics", h.updateCourseBasics)
		course.POST("/send-to-check", h.sendToCheck)
		course.POST("/approve-course", h.approveCourse)
		course.POST("/reject-course", h.rejectCourse)
		course.POST("/upload-course-materials", h.uploadCourseMaterials)
		course.GET("/get-course-info", h.getCourseInfo)
		course.GET("/get-courses-waiting-for-approval", h.getCoursesWaitingForApproval)
		course.GET("/get-courses-approved", h.getCoursesApproved)

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
