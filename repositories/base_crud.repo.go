package repositories

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBaseCrudRepo[TModel any, TDto models.IDTO, TPage any] interface {
	Create(input TDto) (uint, error)
	Read(c *fiber.Ctx, dest []TPage) paginate.Page
	Update(id string, input *TDto) error
	Delete(id string) (TModel, error)
	GetByID(id string) (TModel, error)
}

type BaseCrudRepo[TModel any, TDto models.IDTO, TPage any] struct {
	db        *gorm.DB
	pageQuery string
}

func NewBaseCrudRepo[TModel any, TDto models.IDTO, TPage any](db *gorm.DB) *BaseCrudRepo[TModel, TDto, TPage] {
	return &BaseCrudRepo[TModel, TDto, TPage]{
		db: db,
	}
}

func (r *BaseCrudRepo[TModel, TDto, TPage]) Create(input TDto) (uint, error) {
	err := r.db.Create(&input).Error

	return input.GetID(), err
}

func (r *BaseCrudRepo[TModel, TDto, TPage]) Read(c *fiber.Ctx, dest []TPage) paginate.Page {
	pg := paginate.New()

	query := r.db.Raw(r.pageQuery)
	page := pg.With(query).Request(c.Request()).Response(&dest)

	return page
}

func (r *BaseCrudRepo[TModel, TDto, TPage]) Update(id string, input *TDto) error {
	return r.db.Model(input).Where("id = ?", id).Updates(input).Error
}

func (r *BaseCrudRepo[TModel, TDto, TPage]) Delete(id string) (TModel, error) {
	var model TModel
	err := r.db.Clauses(clause.Returning{}).Delete(&model, id).Error

	return model, err
}

func (r *BaseCrudRepo[TModel, TDto, TPage]) GetByID(id string) (TModel, error) {
	var model TModel
	err := r.db.First(&model, id).Error
	return model, err
}
