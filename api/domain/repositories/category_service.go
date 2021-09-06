package repositories

import (
	paginator "github.com/mmuflih/gorm-paginator"
	"github.com/mmuflih/recipe/api/config"
	"github.com/mmuflih/recipe/api/domain/models"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type CategoryRepository interface {
	Save(d *models.Category) error
	Update(d *models.Category) error
	Find(id uint16) (*models.Category, error)
	FindBySlug(slug string) (*models.Category, error)
	List(q string, page, size int) *paginator.Paginator
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(conn *config.PgSQL) CategoryRepository {
	return &categoryRepo{conn.MainDB}
}

func (s categoryRepo) Save(d *models.Category) error {
	return s.db.Save(d).Error
}

func (s categoryRepo) Update(d *models.Category) error {
	return s.db.Updates(&d).Error
}

func (s categoryRepo) Find(id uint16) (*models.Category, error) {
	var item *models.Category
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s categoryRepo) FindBySlug(slug string) (*models.Category, error) {
	var item *models.Category
	err := s.db.First(&item, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s categoryRepo) List(q string, page, size int) *paginator.Paginator {
	var items []*models.Category
	var filters []paginator.Filter
	filters = append(filters, paginator.Filter{"", "raw", "deleted_at is null"})
	return paginator.Make(&paginator.Config{
		DB:      s.db,
		Page:    page,
		Size:    size,
		OrderBy: []string{"created_at desc"},
		Filters: filters,
		ShowSQL: true,
	}, &items)
}
