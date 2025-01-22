package repositories_test

import (
	"testing"
	"time"

	"github.com/MarcelArt/ModelCraft/models"
	"github.com/MarcelArt/ModelCraft/repositories"
	"github.com/MarcelArt/ModelCraft/tests/helpers"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetByUsernameOrEmail(t *testing.T) {
	db, err := helpers.ConnectDB()
	assert.NoError(t, err)
	defer helpers.CleanUp(db, &models.User{})

	err = db.AutoMigrate(
		&models.User{},
	)
	assert.NoError(t, err)

	userTest1 := models.User{
		Username:   "string",
		Email:      "string@yopmail.com",
		Password:   "stringest",
		Salt:       "stringest",
		VerifiedAt: time.Now(),
	}
	err = db.Create(&userTest1).Error
	assert.NoError(t, err)

	userRepo := repositories.NewUserRepo(db)

	t.Run("should be able to retrieve data by username", func(t *testing.T) {
		user, err := userRepo.GetByUsernameOrEmail("string")
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, userTest1.Username, user.Username)
	})

	t.Run("should be able to retrieve data by email", func(t *testing.T) {
		user, err := userRepo.GetByUsernameOrEmail("string@yopmail.com")
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, userTest1.Email, user.Email)
	})

	userTest2 := models.UserDTO{
		Username: "unverified",
		Email:    "unverified@yopmail.com",
		Password: "unverified",
		Salt:     "unverified",
	}
	err = db.Create(&userTest2).Error
	assert.NoError(t, err)

	t.Run("should return error when user not verified", func(t *testing.T) {
		user, err := userRepo.GetByUsernameOrEmail("unverified")
		assert.Error(t, err)
		assert.Empty(t, user)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

}
