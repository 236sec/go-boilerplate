package database

import (
	"context"

	"gorm.io/gorm"
)

type Database interface {
	WithContext(ctx context.Context) Database
	Create(value interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Where(query interface{}, args ...interface{}) Database
}

func GetDatabase() Database {
	dbOnce.Do(func() {
		dbInstance = initDatabase()
	})
	return dbInstance
}

var ErrRecordNotFound = gorm.ErrRecordNotFound