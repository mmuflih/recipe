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

type IngredientRepository interface {
	Save(d *models.Ingredient) error
	Update(d *models.Ingredient) error
	Find(id uint) (*models.Ingredient, error)
	FindBySlug(slug string) (*models.Ingredient, error)
	List(q string, page, size int) *paginator.Paginator
}

type ingredientRepo struct {
	db *gorm.DB
}

func NewIngredientRepo(conn *config.PgSQL) IngredientRepository {
	return &ingredientRepo{conn.MainDB}
}

func (s ingredientRepo) Save(d *models.Ingredient) error {
	return s.db.Save(d).Error
}

func (s ingredientRepo) Update(d *models.Ingredient) error {
	return s.db.Updates(&d).Error
}

func (s ingredientRepo) Find(id uint) (*models.Ingredient, error) {
	var item *models.Ingredient
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s ingredientRepo) FindBySlug(slug string) (*models.Ingredient, error) {
	var item *models.Ingredient
	err := s.db.First(&item, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s ingredientRepo) List(q string, page, size int) *paginator.Paginator {
	var items []*models.Ingredient
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
