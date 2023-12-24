package tracker

import (
	"errors"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var ErrInvalidDatabase = errors.New("unable to open database")

func openOrCreateDatabase(path string) error {
	d, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if e != nil {
		return ErrInvalidDatabase
	}
	db = d
	return nil
}
