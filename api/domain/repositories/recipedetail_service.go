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

type RecipeDetailRepository interface {
	DBConn() *gorm.DB
	Save(d *models.RecipeDetail, tx *gorm.DB) error
	Update(d *models.RecipeDetail, tx *gorm.DB) error
	Find(id uint) (*models.RecipeDetail, error)
	FindByRecipe(recipeID uint) ([]*models.RecipeDetail, error)
	FindByIngredient(recipeID, ingredientID uint) (*models.RecipeDetail, error)
	List(q string, page, size int) *paginator.Paginator
}

type recipeDetailRepo struct {
	db *gorm.DB
}

func NewRecipeDetailRepo(conn *config.PgSQL) RecipeDetailRepository {
	return &recipeDetailRepo{conn.MainDB}
}
func (s recipeDetailRepo) DBConn() *gorm.DB { return s.db }
func (s recipeDetailRepo) Save(d *models.RecipeDetail, tx *gorm.DB) error {
	return s.db.Save(d).Error
}

func (s recipeDetailRepo) Update(d *models.RecipeDetail, tx *gorm.DB) error {
	return s.db.Updates(&d).Error
}

func (s recipeDetailRepo) Find(id uint) (*models.RecipeDetail, error) {
	var item *models.RecipeDetail
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s recipeDetailRepo) FindByRecipe(recipeID uint) ([]*models.RecipeDetail, error) {
	var items []*models.RecipeDetail
	err := s.db.Find(&items, "deleted_at is null and recipe_id = ?", recipeID).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s recipeDetailRepo) FindByIngredient(recipeID, ingredientID uint) (*models.RecipeDetail, error) {
	var item *models.RecipeDetail
	err := s.db.First(&item, "recipe_id = ? and ingredient_id = ?",
		recipeID, ingredientID).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s recipeDetailRepo) List(q string, page, size int) *paginator.Paginator {
	var items []*models.RecipeDetail
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
