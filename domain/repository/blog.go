package repository

import (
	"engineer-jobhunting-api/domain/model"
)

// BlogRepository : blog repository の interface
type BlogRepository interface {
	Create(blog *model.Blog) (*model.Blog, error)
	FindByID(id int) (*model.Blog, error)
	Update(blog *model.Blog) (*model.Blog, error)
	Delete(blog *model.Blog) error
}
