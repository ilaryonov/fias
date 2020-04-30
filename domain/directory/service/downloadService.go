package service

import (
	"archive/zip"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	addressEntity "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/entity"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type DownloadService struct {
	logger logrus.Logger
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func NewDownloadService(logger logrus.Logger) *DownloadService {
	return &DownloadService{
		logger: logger,
	}
}

func (d *DownloadService) DownloadFile(url string) (*entity.File, error) {
	fileName := path.Base(url)
	filepath := viper.GetString("directory.filePath") + fileName
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return nil, err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return nil, err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return nil, err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return nil, err
	}
	return &entity.File{Path: filepath}, nil
}

func (d *DownloadService) Unzip(file *entity.File) ([]entity.File, error) {
	d.logger.Info(file)
	dest := viper.GetString("directory.filePath")
	var filenames []entity.File

	r, err := zip.OpenReader(file.Path)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}
		fileExt := path.Ext(fpath)
		isAddr, _ := regexp.MatchString(addressEntity.GetAddressXmlFile(), f.Name)
		isHouse, _ := regexp.MatchString(addressEntity.GetHouseXmlFile(), f.Name)
		if (!isHouse && !isAddr) || fileExt != ".XML"{
			continue
		}
		filenames = append(filenames, entity.File{Path: fpath})

		if f.FileInfo().IsDir() {
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
