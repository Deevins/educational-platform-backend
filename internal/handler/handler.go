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
	UpdateAvatar(ctx context.Context, ID int32, avatar S3.FileDataType) (string, error)
	AddUserTeachingExperience(ctx context.Context, exp *model.UserUpdateTeachingExperience) error
	UpdateUserInfo(ctx context.Context, user *model.UserUpdate) error
	CheckIfUserRegisteredToCourse(ctx context.Context, userID, courseID int32) (bool, error)
	RegisterToCourse(ctx context.Context, userID, courseID int32) error
}

type ThreadService interface{}

type DirectoryService interface {
	GetCategoriesWithSubCategories(ctx context.Context) ([]*model.Category, error)
	GetLanguages(ctx context.Context) ([]*model.Language, error)
	GetMetasCount(ctx context.Context) (*model.MetasCount, error)
	GetLevels(ctx context.Context) ([]*model.Level, error)
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
		authGroup.POST("/sign-up", h.signUp)
		authGroup.POST("/sign-in", h.signIn)
	}

	user := router.Group("/users")
	{
		user.GET("/get-one/:userID", h.getOneUser)
		user.GET("/has-user-tried-instructor/:userID", h.hasUserTriedInstructor)
		user.POST("/set-has-user-tried-instructor-to-true/:userID", h.setHasUserTriedInstructorToTrue)
		user.POST("/add-user-teaching-experience", h.updateUserTeachingExperience)
		user.PUT("/update-user-info", h.updateUserInfo)
		user.POST("/upload-avatar/:userID", h.updateAvatar)
		user.POST("/register-on-course", h.registerOnCourse)
		user.GET("/check-if-user-registered-to-course", h.checkIfUserRegisteredToCourse)
		user.GET("/get-all-courses-by-instructor-id/:instructorID", h.getAllCoursesByInstructorID) // OK для вывода курсов по id инструктора которые он создал
	}
	course := router.Group("/courses")
	{
		course.GET("/get-courses-by-user-id/:userID", h.getCoursesByUserID) // для вывода курсов по id пользователя у него на странице или еще где ok
		course.GET("/get-all", h.getAllCoursesWithFilters)                  // OK courses with READY status for all courses page
		course.GET("/get-latest-eight", h.getLatestEightCourses)
		course.POST("/create-base", h.createCourseBase) // OK type must be 'course' or 'practice'
		course.DELETE("/delete/:courseID", h.deleteCourse)
		course.POST("/cancel-publishing/:courseID", h.cancelPublishing)
		course.GET("/search-courses-by-title/:query", h.searchCoursesByTitle)
		course.POST("/send-for-approval/:courseID", h.sendToCheck)
		course.POST("/approve-course/:courseID", h.approveCourse)
		course.POST("/reject-course/:courseID", h.rejectCourse)
		//course.POST("/upload-course-materials/:courseID", h.uploadCourseMaterials)
		course.GET("/get-full-course/:courseID", h.getFullCoursePage) // получаем всю инфу о курсе,который находится в статусе READY это где страница на которой его можно проходить
		course.GET("/get-courses-waiting-for-approval", h.getCoursesWaitingForApproval)
		course.POST("/upload-course-avatar/:courseID", h.uploadCourseAvatar)
		course.POST("/upload-course-preview-video/:courseID", h.uploadCoursePreviewVideo)

		course.PUT("/update-course-goals/:courseID", h.updateCourseGoals)
		course.GET("/get-course-goals/:courseID", h.getCourseGoals)

		course.GET("/get-course-basic-info/:courseID", h.getCourseBasicInfo)
		course.PUT("/update-course-basic-info/:courseID", h.updateCourseBasicInfo)

		course.GET("/get-course-materials/:courseID", h.getCourseMaterials)
		course.POST("/create-section/:courseID", h.createSection)
		course.POST("/create-lecture/:sectionID", h.createLecture)
		course.POST("/create-test/:sectionID", h.createTest)
		course.POST("/remove-section/:sectionID", h.removeSection)
		course.POST("/remove-lecture/:lectureID", h.removeLecture)
		course.POST("/remove-test/:testID", h.removeTest)
		course.POST("/update-section-title/:sectionID", h.updateSectionTitle)
		course.POST("/update-lecture-title/:lectureID", h.updateLectureTitle)
		course.POST("/update-test-title/:testID", h.updateTestTitle)

		course.POST("/add-questions-to-test/:testID", h.addQuestionsToTest)
		course.POST("/add-question-to-test/:testID", h.addQuestionToTest)
		course.POST("/edit-question/:questionID", h.editTestQuestion)
		course.POST("/remove-question-from-test/:questionID", h.removeQuestionFromTest)

		course.POST("/update-lecture-video-url/:courseID/:lectureID", h.uploadLectureVideo)
		course.POST("/remove-lecture-video/:courseID/:lectureID", h.removeLectureVideo)
	}
	directories := router.Group("/directories")
	{
		directories.GET("/categories", h.getCategories)
		directories.GET("/languages", h.getLanguages)
		directories.GET("/levels", h.getLevels)
		directories.GET("/meta-counts", h.getMetasCount)
		//directories.GET("/filter-by-category-and-subcategory", h.filterByCategoryAndSubcategory)
	}

	threads := router.Group("/threads")
	{
		threads.GET("/get-all", h.getAllThreads)
		threads.POST("/create", h.createThread)
		threads.POST("/answer", h.answerInThread)
		threads.GET("/search-by-title/:query", h.searchThreadsByTitle)
		threads.GET("/get-one", h.getOneThread) // here we must get all threads messages
		threads.POST("/add-tag-to", h.addTagToThread)
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
