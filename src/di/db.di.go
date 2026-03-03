package di

import (
	"sync"

	"goboilerplate.com/src/utils"
)

var GetDB = sync.OnceValue(func() *utils.Database {
	return utils.GetDatabase()
})

var GetGormDB = sync.OnceValue(func() utils.GormDB {
	return GetDB()
})