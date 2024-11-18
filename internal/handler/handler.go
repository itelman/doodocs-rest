package handler

import (
	s "github.com/itelman/doodocs-rest/internal/service"
	sendEmail "github.com/itelman/doodocs-rest/internal/service/email"
	archiveCreate "github.com/itelman/doodocs-rest/internal/service/zip_create"
	archiveInfo "github.com/itelman/doodocs-rest/internal/service/zip_info"
)

type Handler struct {
	InfoService   *archiveInfo.ArchiveInfoService
	CreateService *archiveCreate.ArchiveCreateService
	SendService   *sendEmail.SendEmailService
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{
		InfoService:   service.Info,
		CreateService: service.Create,
		SendService:   service.Send,
	}
}
