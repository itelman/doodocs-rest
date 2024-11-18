package archiveCreate

import (
	"archive/zip"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ArchiveCreateService struct {
}

func NewArchiveCreateService() *ArchiveCreateService {
	return &ArchiveCreateService{}
}

func (s *ArchiveCreateService) CreateArchive(files []*multipart.FileHeader, validMimeTypes map[string]bool) ([]byte, error) {
	archivesDir := "archives"

	if _, err := os.Stat(archivesDir); os.IsNotExist(err) {
		err := os.Mkdir(archivesDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	archiveFileName := filepath.Join(archivesDir, "archive.zip")
	archiveFile, err := os.Create(archiveFileName)
	if err != nil {
		return nil, err
	}
	defer archiveFile.Close()

	zipWriter := zip.NewWriter(archiveFile)
	defer zipWriter.Close()

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		mime := fileHeader.Header.Get("Content-Type")
		if !validMimeTypes[mime] {
			return nil, errors.New("invalid file format")
		}

		fileName := filepath.Base(fileHeader.Filename)
		fileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, err
		}
	}

	return os.ReadFile(archiveFileName)
}
