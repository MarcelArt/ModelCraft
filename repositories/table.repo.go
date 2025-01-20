package repositories

import "gorm.io/gorm"

type ITableRepo interface {
	GetTables() ([]string, error)
}

type TableRepo struct {
	db *gorm.DB
}

func NewTableRepo(db *gorm.DB) *TableRepo {
	return &TableRepo{
		db: db,
	}
}

func (r *TableRepo) GetTables() ([]string, error) {
	var tables []string
	err := r.db.Table("information_schema.tables").Select("table_name").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error
	return tables, err
}
