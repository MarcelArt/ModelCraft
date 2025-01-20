package repositories

import (
	"github.com/MarcelArt/app_standard/models"
	"gorm.io/gorm"
)

const authorizedDevicePageQuery = `
	select 
		user_agent,
		ip,
		user_id
	from authorized_devices
`

type IAuthorizedDeviceRepo interface {
	IBaseCrudRepo[models.AuthorizedDevice, models.AuthorizedDeviceDTO, models.AuthorizedDevicePage]
	GetByRefreshToken(token string) (models.AuthorizedDeviceDTO, error)
}

type AuthorizedDeviceRepo struct {
	BaseCrudRepo[models.AuthorizedDevice, models.AuthorizedDeviceDTO, models.AuthorizedDevicePage]
}

func NewAuthorizedDeviceRepo(db *gorm.DB) *AuthorizedDeviceRepo {
	return &AuthorizedDeviceRepo{
		BaseCrudRepo: BaseCrudRepo[models.AuthorizedDevice, models.AuthorizedDeviceDTO, models.AuthorizedDevicePage]{
			db:        db,
			pageQuery: authorizedDevicePageQuery,
		},
	}
}

func (r *AuthorizedDeviceRepo) GetByRefreshToken(token string) (models.AuthorizedDeviceDTO, error) {
	var device models.AuthorizedDeviceDTO
	err := r.db.Where("refresh_token = ?", token).First(&device).Error
	return device, err
}
