package handler

import (
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/pkg/httpResponses"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
)

func (h *Handler) uploadUserAvatar(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No file is received",
			Details: err,
		})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			// Если файл не удается закрыть, возвращаем ошибку с соответствующим статусом и сообщением
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Читаем содержимое файла в байтовый срез
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		// Если содержимое файла не удается прочитать, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to read the file",
			Details: err,
		})
		return
	}

	// Создаем структуру FileDataType для хранения данных файла
	fileData := S3.FileDataType{
		FileName: file.Filename, // Имя файла
		Data:     fileBytes,     // Содержимое файла в виде байтового среза
	}

	// Сохраняем файл в MinIO с помощью метода CreateOne
	link, err := h.s3.CreateOne(fileData)
	if err != nil {
		// Если не удается сохранить файл, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the file",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом сохраненного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    link, // URL-адрес загруженного файла
	})
}

func (h *Handler) uploadCourseAvatar(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No file is received",
			Details: err,
		})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			// Если файл не удается закрыть, возвращаем ошибку с соответствующим статусом и сообщением
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Читаем содержимое файла в байтовый срез
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		// Если содержимое файла не удается прочитать, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to read the file",
			Details: err,
		})
		return
	}

	// Создаем структуру FileDataType для хранения данных файла
	fileData := S3.FileDataType{
		FileName: file.Filename, // Имя файла
		Data:     fileBytes,     // Содержимое файла в виде байтового среза
	}

	// Сохраняем файл в MinIO с помощью метода CreateOne
	link, err := h.s3.CreateOne(fileData)
	if err != nil {
		// Если не удается сохранить файл, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the file",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом сохраненного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    link, // URL-адрес загруженного файла
	})
}

func (h *Handler) uploadCoursePreviewVideo(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No file is received",
			Details: err,
		})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			// Если файл не удается закрыть, возвращаем ошибку с соответствующим статусом и сообщением
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Читаем содержимое файла в байтовый срез
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		// Если содержимое файла не удается прочитать, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to read the file",
			Details: err,
		})
		return
	}

	// Создаем структуру FileDataType для хранения данных файла
	fileData := S3.FileDataType{
		FileName: file.Filename, // Имя файла
		Data:     fileBytes,     // Содержимое файла в виде байтового среза
	}

	// Сохраняем файл в MinIO с помощью метода CreateOne
	link, err := h.s3.CreateOne(fileData)
	if err != nil {
		// Если не удается сохранить файл, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the file",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом сохраненного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    link, // URL-адрес загруженного файла
	})
}

func (h *Handler) uploadCourseLecture(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No file is received",
			Details: err,
		})
		return
	}

	// Открываем файл для чтения
	f, err := file.Open()
	if err != nil {
		// Если файл не удается открыть, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			// Если файл не удается закрыть, возвращаем ошибку с соответствующим статусом и сообщением
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}(f) // Закрываем файл после завершения работы с ним

	// Читаем содержимое файла в байтовый срез
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		// Если содержимое файла не удается прочитать, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to read the file",
			Details: err,
		})
		return
	}

	// Создаем структуру FileDataType для хранения данных файла
	fileData := S3.FileDataType{
		FileName: file.Filename, // Имя файла
		Data:     fileBytes,     // Содержимое файла в виде байтового среза
	}

	// Сохраняем файл в MinIO с помощью метода CreateOne
	link, err := h.s3.CreateOne(fileData)
	if err != nil {
		// Если не удается сохранить файл, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the file",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом сохраненного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    link, // URL-адрес загруженного файла
	})
}

func (h *Handler) getCoursePreviewVideoByFileID(c *gin.Context) {
	// Получаем идентификатор объекта из параметров URL
	objectID := c.Param("objectID")

	// Используем сервис MinIO для получения ссылки на объект
	link, err := h.s3.GetOne(objectID)
	if err != nil {
		// Если произошла ошибка при получении объекта, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом полученного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File retrieved successfully",
		Data:    link, // URL-адрес полученного файла
	})
}

func (h *Handler) getCourseAvatarByFileID(c *gin.Context) {
	// Получаем идентификатор объекта из параметров URL
	objectID := c.Param("objectID")

	// Используем сервис MinIO для получения ссылки на объект
	link, err := h.s3.GetOne(objectID)
	if err != nil {
		// Если произошла ошибка при получении объекта, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом полученного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File retrieved successfully",
		Data:    link, // URL-адрес полученного файла
	})
}

func (h *Handler) getUserAvatarByFileID(c *gin.Context) {
	// Получаем идентификатор объекта из параметров URL
	objectID := c.Param("objectID")

	// Используем сервис MinIO для получения ссылки на объект
	link, err := h.s3.GetOne(objectID)
	if err != nil {
		// Если произошла ошибка при получении объекта, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом полученного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File retrieved successfully",
		Data:    link, // URL-адрес полученного файла
	})
}

func (h *Handler) getCourseLecturesByFileIDs(c *gin.Context) {
	// Получаем идентификатор объекта из параметров URL
	objectID := c.Param("objectID")

	// Используем сервис MinIO для получения ссылки на объект
	link, err := h.s3.GetOne(objectID)
	if err != nil {
		// Если произошла ошибка при получении объекта, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом полученного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File retrieved successfully",
		Data:    link, // URL-адрес полученного файла
	})
}

func (h *Handler) getCoursesAvatarsByFileIDs(c *gin.Context) {
	// Получаем идентификатор объекта из параметров URL
	objectID := c.Param("objectID")

	// Используем сервис MinIO для получения ссылки на объект
	link, err := h.s3.GetOne(objectID)
	if err != nil {
		// Если произошла ошибка при получении объекта, возвращаем ошибку с соответствующим статусом и сообщением
		c.JSON(http.StatusInternalServerError, httpResponses.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ с URL-адресом полученного файла
	c.JSON(http.StatusOK, httpResponses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File retrieved successfully",
		Data:    link, // URL-адрес полученного файла
	})
}
