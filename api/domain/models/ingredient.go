package models

import "time"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Ingredient struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewIngredient(name, slug string) *Ingredient {
	return &Ingredient{
		Name: name,
		Slug: slug,
	}
}

func (i Ingredient) PublicResponse() interface{} {
	return struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}{
		ID:   i.ID,
		Name: i.Name,
		Slug: i.Slug,
	}
}
