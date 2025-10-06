package usecases

type HealthUseCase interface {
	Apply() error
}

type HealthService struct{}

func (u *HealthService) Apply() error {
	return nil
}

func NewHealthService() HealthUseCase {
	return &HealthService{}
}
