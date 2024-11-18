package archiveInfo

import (
	"archive/zip"
	"bytes"
	"errors"
	"mime"
	"path/filepath"

	"github.com/itelman/doodocs-rest/internal/models"
	"github.com/itelman/doodocs-rest/pkg/helpers"
)

type ArchiveInfoService struct {
}

func NewArchiveInfoService() *ArchiveInfoService {
	return &ArchiveInfoService{}
}

func (s *ArchiveInfoService) GetInfoByArchive(fileData []byte, fileName string) (*models.ArchiveInfo, error) {

	if !helpers.IsZipFile(fileName) {
		return nil, errors.New("not a valid ZIP file")
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileData), int64(len(fileData)))
	if err != nil {
		return nil, err
	}

	var totalSize float64
	var files []models.FileInfo

	for _, file := range zipReader.File {
		info := models.FileInfo{
			FilePath: file.Name,
			Size:     float64(file.UncompressedSize64),
			MimeType: mime.TypeByExtension(filepath.Ext(file.Name)),
		}

		totalSize += info.Size
		files = append(files, info)
	}

	archiveInfo := &models.ArchiveInfo{
		FileName:    fileName,
		ArchiveSize: float64(len(fileData)),
		TotalSize:   totalSize,
		TotalFiles:  float64(len(files)),
		Files:       files,
	}

	return archiveInfo, nil
}
