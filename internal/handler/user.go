package handler

import (
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
	var input GetOneUserInfoRequest // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.us.GetByID(ctx, input.ID)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":           user.ID,
		"full_name":    user.FullName,
		"description":  user.Description,
		"email":        user.Email,
		"avatar_url":   user.AvatarUrl,
		"phone_number": user.PhoneNumber,
	})

}

type HasUserTriedInstructorRequest struct {
	ID int32 `json:"id"`
}

func (h *Handler) hasUserTriedInstructor(ctx *gin.Context) {
	var input HasUserTriedInstructorRequest // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hasUsed, err := h.us.GetHasUserTriedInstructor(ctx, input.ID)
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
	var input SetHasUserTriedInstructorToTrueRequest // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.us.SetHasUserTriedInstructorToTrue(ctx, input.ID)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

type UpdateAvatarStruct struct {
	ID int32 `json:"id"`
}

func (h *Handler) updateAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
		return
	}

	userIDStr := ctx.PostForm("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, httpResponses.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "User ID is required",
		})
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
