package repositories

import (
	"time"

	"github.com/MarcelArt/ModelCraft/models"
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
	Verify(id string) error
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
	err := r.db.Where("(username = ? OR email = ?) and verified_at is not null", username, username).First(&user).Error
	return user, err
}

func (r *UserRepo) Verify(id string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("verified_at", time.Now()).Error
}
