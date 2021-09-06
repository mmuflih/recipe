package helper

import (
	"strconv"
	"strings"

	"github.com/gosimple/slug"
	"github.com/mmuflih/recipe/api/config"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Slug interface {
	Make(table, name string) string
}

type slugG struct {
	db *gorm.DB
}

func NewSlug(mySql *config.PgSQL) Slug {
	return slugG{mySql.MainDB}
}

func (s slugG) Make(table, name string) string {
	sl := slug.Make(name)
	slugDB := new(SlugData)
	qSelect := "SELECT slug FROM " + table
	err := s.db.Raw(qSelect+" WHERE slug = ?", sl).Scan(slugDB).Error
	if err != nil {
		return sl
	}

	if slugDB.Slug == "" {
		return sl
	}

	slugs := []*SlugData{}
	err = s.db.Raw(qSelect+" WHERE slug like ?", sl+"%").Scan(&slugs).Error
	if err != nil {
		return sl
	}
	count := 1
	for _, slu := range slugs {
		dbCountS := strings.Replace(slu.Slug, sl+"-", "", -1)
		dbCountI, err := strconv.Atoi(dbCountS)
		if err != nil {
			continue
		}
		if count == dbCountI {
			count++
			continue
		}
	}
	return sl + "-" + strconv.Itoa(count)
}

type SlugData struct {
	Slug string `json:"slug"`
}
