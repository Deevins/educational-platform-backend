package handler

import (
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

type GetOneUserInfoRequest struct {
	ID int32 `json:"id"`
}

func (h *Handler) getOneUser(ctx *gin.Context) {
	if ctx.Param("userID") == "" {
		ctx.JSON(400, gin.H{"error": "userID is empty"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "userID is not a number"})
		return
	}

	user, err := h.us.GetByID(ctx, int32(userID))
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type HasUserTriedInstructorRequest struct {
	ID int32 `json:"id"`
}

func (h *Handler) hasUserTriedInstructor(ctx *gin.Context) {
	if ctx.Param("userID") == "" {
		ctx.JSON(400, gin.H{"error": "userID is empty"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	hasUsed, err := h.us.GetHasUserTriedInstructor(ctx, int32(userID))
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"has_used": hasUsed,
	})
}

type SetHasUserTriedInstructorToTrueRequest struct {
	ID int32 `json:"id"`
}

func (h *Handler) setHasUserTriedInstructorToTrue(ctx *gin.Context) {
	if ctx.Param("userID") == "" {
		ctx.JSON(400, gin.H{"error": "userID is empty"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	err = h.us.SetHasUserTriedInstructorToTrue(ctx, int32(userID))
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

type UpdateAvatarStruct struct {
	ID int32 `json:"id"`
}

func (h *Handler) updateAvatar(ctx *gin.Context) {
	if ctx.Param("userID") == "" {
		ctx.JSON(400, gin.H{"error": "userID is empty"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "courseID is not a number"})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
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

	url, err := h.us.UpdateAvatar(ctx, int32(userID), S3.FileDataType{
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

func (h *Handler) updateUserTeachingExperience(ctx *gin.Context) {
	var input model.UserUpdateTeachingExperience

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err := h.us.AddUserTeachingExperience(ctx, &input)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) updateUserInfo(ctx *gin.Context) {
	var input *model.UserUpdate

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.us.UpdateUserInfo(ctx, input)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

type RegisterToCourseRequest struct {
	UserID   int32  `json:"userID"`
	CourseID string `json:"courseID"`
}

func (h *Handler) registerOnCourse(ctx *gin.Context) {
	var input *RegisterToCourseRequest

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	courseIDI, _ := strconv.Atoi(input.CourseID)

	err := h.us.RegisterToCourse(ctx, input.UserID, int32(courseIDI))
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered to course",
	})
}

func (h *Handler) checkIfUserRegisteredToCourse(ctx *gin.Context) {
	userID := ctx.Query("userID")
	courseID := ctx.Query("courseID")

	if userID == "" || courseID == "" {
		model.NewErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("userID and courseID are required"))
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("userID must be a number"))
		return
	}

	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("courseID must be a number"))
		return
	}

	isRegistered, err := h.us.CheckIfUserRegisteredToCourse(ctx, int32(userIDInt), int32(courseIDInt))
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "User registered to course",
		"is_registered": isRegistered,
	})

}
