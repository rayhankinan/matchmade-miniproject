package migrations

import (
	"fmt"
	"service/internal/infrastructure"
	"service/internal/models"
)

func MigrateUp() error {
	infrastructure.Init()

	err := infrastructure.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		return fmt.Errorf("failed to create extension: %v", err)
	}

	err = infrastructure.DB.AutoMigrate(&models.User{}, &models.Movie{})
	if err != nil {
		return fmt.Errorf("failed to migrate: %v", err)
	}

	fmt.Println("Migration successful")

	return nil
}

func MigrateDown() error {
	infrastructure.Init()

	err := infrastructure.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		return fmt.Errorf("failed to create extension: %v", err)
	}

	err = infrastructure.DB.Migrator().DropTable(&models.User{}, &models.Movie{})
	if err != nil {
		return fmt.Errorf("failed to drop table: %v", err)
	}

	fmt.Println("Rollback successful")

	return nil
}
