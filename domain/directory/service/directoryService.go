package service

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/file/entity"
	"log"
	"os"
	"strings"
)

type DirectoryService struct {
	logger logrus.Logger
}

func NewDirectoryService(logger logrus.Logger) *DirectoryService {
	return &DirectoryService{
		logger: logger,
	}
}

func (d *DirectoryService) UnzipFiles(file *entity.File) {
	files, err := file.Unzip(app.Config.FilePath+fileName, app.Config.FilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Unzipped:\n" + strings.Join(files, "\n"))
	err = os.Remove(app.Config.FilePath + fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
}