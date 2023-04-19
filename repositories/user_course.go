package repositories

import (
	"github.com/adhityaf/iglobal-be-test/models"
	"gorm.io/gorm"
)

type UserCourseRepository interface {
	Create(userCourse *models.UserCourse) (*models.UserCourse, error)
	FindById(courseId int) (*models.UserCourse, error)
}

type userCourseRepository struct {
	db *gorm.DB
}

func NewUserCourseRepository(db *gorm.DB) UserCourseRepository {
	return &userCourseRepository{
		db: db,
	}
}

func (u *userCourseRepository) Create(userCourse *models.UserCourse) (*models.UserCourse, error) {
	err := u.db.Create(&userCourse).Error
	return userCourse, err
}

func (u *userCourseRepository) FindById(courseId int) (*models.UserCourse, error) {
	var userCourse *models.UserCourse
	err := u.db.Where("course_id = ?", courseId).First(&userCourse).Error
	return userCourse, err
}
