package repositories

import (
	"strings"
	"time"

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

type RecipeRepository interface {
	DBConn() *gorm.DB
	Save(d *models.Recipe, tx *gorm.DB) error
	Update(d *models.Recipe, tx *gorm.DB) error
	Find(id uint) (*models.Recipe, error)
	FindBySlug(slug string) (*models.Recipe, error)
	List(q string, page, size int) *paginator.Paginator
}

type recipeRepo struct {
	db *gorm.DB
}

func NewRecipeRepo(conn *config.PgSQL) RecipeRepository {
	return &recipeRepo{conn.MainDB}
}

func (s recipeRepo) DBConn() *gorm.DB { return s.db }
func (s recipeRepo) Save(d *models.Recipe, tx *gorm.DB) error {
	return s.db.Save(d).Error
}

func (s recipeRepo) Update(d *models.Recipe, tx *gorm.DB) error {
	return s.db.Updates(&d).Error
}

func (s recipeRepo) Find(id uint) (*models.Recipe, error) {
	var item *models.Recipe
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s recipeRepo) FindBySlug(slug string) (*models.Recipe, error) {
	var item *models.Recipe
	err := s.db.First(&item, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s recipeRepo) List(q string, page, size int) *paginator.Paginator {
	var items []struct {
		ID           uint      `json:"id"`
		Name         string    `json:"name"`
		Slug         string    `json:"slug"`
		CategoryName string    `json:"category_name"`
		CategoryID   uint16    `json:"category_id"`
		CreatedAt    time.Time `json:"created_at"`
	}

	var filters []paginator.Filter
	q = strings.ToLower(q)
	filters = append(filters, paginator.Filter{"", "raw", "r.deleted_at is null"})
	filters = append(filters, paginator.Filter{"", "raw",
		"lower(c.name) like '%" + q + "%'" +
			" OR lower(r.name) like '%" + q + "%'" +
			" OR lower(i.name) like '%" + q + "%'"})

	query := `SELECT distinct r.id, r.name, r.slug, c.name as category_name, 
			c.id as category_id, r.created_at
		FROM recipes r
		JOIN categories c
			ON c.id = r.category_id
		JOIN recipe_details rd
			ON rd.recipe_id = r.id
		JOIN ingredients i
			ON i.id = rd.ingredient_id`
	return paginator.MakeRaw(query, &paginator.Config{
		DB:      s.db,
		Page:    page,
		Size:    size,
		OrderBy: []string{"r.created_at desc"},
		Filters: filters,
		ShowSQL: true,
	}, &items)
}
