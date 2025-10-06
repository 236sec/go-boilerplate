package di

import (
	"sync"

	"goboilerplate.com/src/rest/usecases"
)

// Create a singleton instance of GetHealthService
var (
	healthServiceInstance usecases.HealthUseCase
	healthServiceOnce     sync.Once
)

func GetHealthService() usecases.HealthUseCase {
	healthServiceOnce.Do(func() {
		healthServiceInstance = usecases.NewHealthService()
	})
	return healthServiceInstance
}
