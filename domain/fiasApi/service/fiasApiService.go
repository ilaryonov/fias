package service

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ilaryonov/fias/domain/fiasApi/entity"
	"log"
	"net/http"
)

const (
	allFiles = "GetAllDownloadFileInfo"
	lastFile = "GetLastDownloadFileInfo"
)

type FiasApiService struct {
	logger logrus.Logger
}

func NewFiasApiService(logger logrus.Logger) *FiasApiService {
	return &FiasApiService{
		logger: logger,
	}
}

func (f *FiasApiService) GetAllDownloadFileInfo() []entity.DownloadFileInfo {
	url := viper.GetString("fiasApi.url") + allFiles
	files := []entity.DownloadFileInfo{}
	result, err := http.Get(url)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
	json.NewDecoder(result.Body).Decode(&files)

	return files
}

func (f *FiasApiService) GetLastDownloadFileInfo() entity.DownloadFileInfo {
	url := viper.GetString("fiasApi.url") + lastFile
	file := entity.DownloadFileInfo{}
	netClient := &http.Client{Transport: &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}}

	result, err := netClient.Get(url)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
	json.NewDecoder(result.Body).Decode(&file)

	return file
}
