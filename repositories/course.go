package repositories

import (
	"github.com/adhityaf/iglobal-be-test/models"
	"github.com/adhityaf/iglobal-be-test/params"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *models.Course) (*models.Course, error)
	Update(course *models.Course) (*models.Course, error)
	Delete(course *models.Course) (*models.Course, error)
	FindById(courseId int) (*models.Course, error)
	FindAllCourses() (*[]models.Course, error)
	FindAllBySearchAndFilter(criteria params.CourseSearchCriteria) (*[]models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (c *courseRepository) Create(course *models.Course) (*models.Course, error) {
	err := c.db.Create(&course).Error
	return course, err
}

func (c *courseRepository) Update(course *models.Course) (*models.Course, error) {
	err := c.db.Save(&course).Error
	return course, err
}

func (c *courseRepository) Delete(course *models.Course) (*models.Course, error) {
	err := c.db.Delete(&course).Error
	return course, err
}

func (c *courseRepository) FindById(courseId int) (*models.Course, error) {
	var course *models.Course
	err := c.db.Where("course_id = ?", courseId).First(&course).Error
	return course, err
}

func (c *courseRepository) FindAllCourses() (*[]models.Course, error) {
	var courses *[]models.Course
	err := c.db.Find(&courses).Error
	return courses, err
}

func (c *courseRepository) FindAllBySearchAndFilter(criteria params.CourseSearchCriteria) (*[]models.Course, error) {
	var courses *[]models.Course
	offset := (criteria.Page - 1) * criteria.Size
	queryBuilder := c.db.Limit(criteria.Size).Offset(offset)

	err := queryBuilder.
		Where(
			c.db.Where("title LIKE ?", "%"+criteria.Query+"%").
			Or("description LIKE ?", "%"+criteria.Query+"%"),
		).Where("level = ?", criteria.Level).
		Where("language = ?", criteria.Language).
		Find(&courses).Error
	return courses, err
}
