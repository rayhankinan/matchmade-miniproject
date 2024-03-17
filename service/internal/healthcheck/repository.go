package healthcheck

import (
	"context"

	"gorm.io/gorm"
)

type IHealthcheckRepository interface {
	Ping(context.Context) error
}

type HealthcheckRepository struct {
	db *gorm.DB
}

// Directly access the database
func (r *HealthcheckRepository) Ping(ctx context.Context) (err error) {
	// Get generic database object sql.DB to use its functions
	sqlDB, err := r.db.DB()
	if err != nil {
		return
	}

	// Ping database to check if it's still alive
	err = sqlDB.PingContext(ctx)
	return
}

func NewRepository(db *gorm.DB) (repo IHealthcheckRepository) {
	repo = &HealthcheckRepository{
		db,
	}
	return
}
