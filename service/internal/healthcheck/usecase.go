package healthcheck

import (
	"context"
)

type IHealthcheckUseCase interface {
	Ping(context.Context) error
}

type HealthcheckUseCase struct {
	repo IHealthcheckRepository
}

// Manipulate data from repository
func (u *HealthcheckUseCase) Ping(ctx context.Context) (err error) {
	err = u.repo.Ping(ctx)
	return
}

func NewUseCase(repo IHealthcheckRepository) (usecase IHealthcheckUseCase) {
	usecase = &HealthcheckUseCase{
		repo,
	}
	return
}
