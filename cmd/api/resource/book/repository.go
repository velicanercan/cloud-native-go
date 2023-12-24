package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) List() (Books, error) {
	books := make([]*Book, 0)
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *Repository) Create(book *Book) (*Book, error) {
	if err := r.DB.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *Repository) Read(id uuid.UUID) (*Book, error) {
	book := &Book{}
	if err := r.DB.Where("id = ?", id).First(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *Repository) Update(book *Book) (int64, error) {
	result := r.DB.Model(&Book{}).
		Select("Title", "Author", "PublishedDate", "ImageURL", "Description", "UpdatedAt").
		Where("id = ?", book.ID).
		Updates(book)

	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.DB.Where("id = ?", id).Delete(&Book{})
	return result.RowsAffected, result.Error
}
