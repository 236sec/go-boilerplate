package database

import (
	"gorm.io/gorm"
)

type Database interface {
	Create(value any) error
	Find(dest any, conds ...any) error
	First(dest any, conds ...any) error
	Where(query any, args ...any) Database
}

func GetDatabase() Database {
	dbOnce.Do(func() {
		dbInstance = initDatabase()
	})
	return dbInstance
}

var ErrRecordNotFound = gorm.ErrRecordNotFound