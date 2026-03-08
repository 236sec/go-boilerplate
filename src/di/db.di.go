package di

import (
	"sync"

	"goboilerplate.com/src/pkg/database"
)

var GetDB = sync.OnceValue(func() database.Database {
	return database.GetDatabase()
})