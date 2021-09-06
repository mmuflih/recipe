package models

import "time"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Recipe struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Slug       string     `json:"slug"`
	CategoryID uint16     `json:"category_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func NewRecipe(name, slug string, categoryID uint16) *Recipe {
	return &Recipe{
		Name:       name,
		Slug:       slug,
		CategoryID: categoryID,
	}
}
