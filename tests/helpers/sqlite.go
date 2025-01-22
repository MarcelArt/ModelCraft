package helpers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}

func CleanUp(db *gorm.DB, models ...interface{}) error {
	if err := db.Migrator().DropTable(models...); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
