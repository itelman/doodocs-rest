package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itelman/doodocs-rest/internal/models"
)

func (h *Handler) CreateArchiveByFiles(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := c.Request.MultipartForm.File["files[]"]

	validMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
	}

	archiveFiles, err := h.CreateService.CreateArchive(files, validMimeTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=archives.zip")
	c.Data(http.StatusOK, "application/zip", archiveFiles)
	c.JSON(http.StatusOK, models.MessageOK{Message: "OK"})
}
