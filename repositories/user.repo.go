package repositories

import (
	"github.com/MarcelArt/app_standard/models"
	"gorm.io/gorm"
)

const userPageQuery = `
	select 
		username, 
		email
	from users
`

type IUserRepo interface {
	IBaseCrudRepo[models.User, models.UserDTO, models.UserPage]
	GetByUsernameOrEmail(username string) (models.UserDTO, error)
}

type UserRepo struct {
	BaseCrudRepo[models.User, models.UserDTO, models.UserPage]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseCrudRepo: BaseCrudRepo[models.User, models.UserDTO, models.UserPage]{
			db:        db,
			pageQuery: userPageQuery,
		},
	}
}

func (r *UserRepo) GetByUsernameOrEmail(username string) (models.UserDTO, error) {
	var user models.UserDTO
	err := r.db.Where("username = ? OR email = ?", username, username).First(&user).Error
	return user, err
}
