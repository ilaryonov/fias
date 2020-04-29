package service

import "github.com/sirupsen/logrus"

type FileService struct {
	logger logrus.Logger
}

func NewFileService(logger logrus.Logger) *FileService {
	return &FileService{
		logger: logger,
	}
}