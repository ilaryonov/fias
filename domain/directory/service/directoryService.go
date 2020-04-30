package service

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/entity"
)

type DirectoryService struct {
	logger logrus.Logger
	downloadService *DownloadService
}

func NewDirectoryService(logger logrus.Logger) *DirectoryService {
	return &DirectoryService{
		logger: logger,
		downloadService: NewDownloadService(logger),
	}
}

func (d *DirectoryService) DownloadAndExtractFile(url string) *[] entity.File {
	file, err := d.downloadService.DownloadFile(url)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	extractedFiles, err := d.downloadService.Unzip(file)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return  &extractedFiles
}