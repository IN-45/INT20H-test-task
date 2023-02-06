package handler

import (
	"github.com/IN-45/INT20H-test-task/pkg/cloudstorage"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/gofiber/fiber/v2"
)

type UploaderHandler struct {
	fileUploader *cloudstorage.FileUploader
}

func NewUploaderHandler(fileUploader *cloudstorage.FileUploader) *UploaderHandler {
	return &UploaderHandler{
		fileUploader: fileUploader,
	}
}

func RegisterUploaderHandler(app *fiber.App, h *UploaderHandler) {
	app.Post("/upload", h.Upload)
}

// Upload
//
//	@Summary	Upload file to cloud storage
//	@Tags		Files Uploading
//	@Param		file	formData	file	true	"File to upload"
//	@Success	200		{object}	dtoUploadedFile
//	@Failure	400		{object}	fiber.Error
//	@Failure	401		{object}	customerrors.UnauthorizedError
//	@Failure	500		{object}	fiber.Error
//	@Router		/upload [post]
func (h *UploaderHandler) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fileURL, err := h.fileUploader.Upload(ctx.Context(), file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(dtoUploadedFile{
		FileURL: fileURL,
	})
}

type dtoUploadedFile struct {
	FileURL string `json:"file_url"`
}
