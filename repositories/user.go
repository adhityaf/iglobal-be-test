package repositories

import (
	"github.com/adhityaf/iglobal-be-test/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepository) Update(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return user, err
}