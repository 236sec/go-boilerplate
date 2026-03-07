package di

import (
	"sync"

	"goboilerplate.com/src/utils/database"
)

var GetDB = sync.OnceValue(func() database.Database {
	return database.GetDatabase()
})