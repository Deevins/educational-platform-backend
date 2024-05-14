package handler

import (
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (h *Handler) getOneUser(ctx *gin.Context) {
	var input int32 // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) hasUserTriedInstructor(ctx *gin.Context) {
	var input int32 // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hasUsed, err := h.us.GetHasUserTriedInstructor(ctx, input)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"has_used": hasUsed,
	})
}

func (h *Handler) setHasUserTriedInstructorToTrue(ctx *gin.Context) {
	var input int32 // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.us.SetHasUserTriedInstructorToTrue(ctx, input)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getSelfInfo(ctx *gin.Context) {
	var input int32 // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.us.GetSelfInfo(ctx, input)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":           user.ID,
		"full_name":    user.FullName,
		"description":  user.Description,
		"email":        user.Email,
		"avatar":       user.Avatar,
		"phone_number": user.PhoneNumber,
		"linkedin":     user.Linkedin,
		"vk":           user.VK,
		"github":       user.Github,
	})

}

func (h *Handler) updateAvatar(ctx *gin.Context) {
	var input int32 // id
	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to retrieve file: %v", err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to close file: %v", err)
			return
		}
	}(file)

	// Чтение содержимого файла
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to read file: %v", err)
		return
	}

	err = h.us.UpdateAvatar(ctx, input, fileBytes)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	ctx.String(http.StatusOK, "File uploaded successfully")
}

func (h *Handler) updateUserTeachingExperience(ctx *gin.Context) {
	var input model.UserUpdateTeachingExperience

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err := h.us.UpdateUserTeachingExperience(ctx, &input)
	if err != nil {
		return
	}
}

func (h *Handler) updateUserInfo(ctx *gin.Context) {
	var input model.UserUpdate

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
