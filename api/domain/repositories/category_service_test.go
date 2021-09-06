package repositories

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mmuflih/recipe/api/domain/models"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

var c = &models.Category{
	ID:        0,
	Name:      "Main Course",
	Slug:      "main-course",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
}

func NewDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return gdb, mock
}

func TestSave(t *testing.T) {
	db, mock := NewDBMock()

	repo := categoryRepo{db}

	// Action
	repo.Save(data)

	// Assertion
}
