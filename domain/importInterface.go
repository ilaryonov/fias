package domain

import (
	"github.com/jinzhu/gorm"
	"os"
	"sync"
)

type ImportInterface interface {
	Import(f os.FileInfo, wg *sync.WaitGroup, db *gorm.DB) bool
	BatchInsert()
}