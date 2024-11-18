package service

import (
	"github.com/itelman/doodocs-rest/config"
	sendEmail "github.com/itelman/doodocs-rest/internal/service/email"
	archiveCreate "github.com/itelman/doodocs-rest/internal/service/zip_create"
	archiveInfo "github.com/itelman/doodocs-rest/internal/service/zip_info"
)

type Service struct {
	Info   *archiveInfo.ArchiveInfoService
	Create *archiveCreate.ArchiveCreateService
	Send   *sendEmail.SendEmailService
	Config config.Config
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Info:   archiveInfo.NewArchiveInfoService(),
		Create: archiveCreate.NewArchiveCreateService(),
		Send:   sendEmail.NewSendEmailService(cfg),
		Config: cfg,
	}
}
