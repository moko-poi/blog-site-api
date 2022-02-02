package usecase

import (
	"engineer-jobhunting-api/domain/model"
	"engineer-jobhunting-api/domain/repository"
)

// BlogUsecase blog usecaseのinterface
type BlogUsecase interface {
	Create(title, content string) (*model.Blog, error)
	FindByID(id int) (*model.Blog, error)
	Update(id int, title, content string) (*model.Blog, error)
	Delete(id int) error
}

type blogUsecase struct {
	blogRepo repository.BlogRepository
}

// NewBlogUsecase blog usecaseのコンストラクタ
func NewBlogUsecase(blogRepo repository.BlogRepository) BlogUsecase {
	return &blogUsecase{blogRepo: blogRepo}
}

// Create blogを保存するときのユースケース
func (bu *blogUsecase) Create(title, content string) (*model.Blog, error) {
	blog, err := model.NewBlog(title, content)
	if err != nil {
		return nil, err
	}

	createdBlog, err := bu.blogRepo.Create(blog)
	if err != nil {
		return nil, err
	}

	return createdBlog, nil
}

// FindByID blogをIDで取得するときのユースケース
func (bu *blogUsecase) FindByID(id int) (*model.Blog, error) {
	foundBlog, err := bu.blogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundBlog, nil
}

// Update blogを更新するときのユースケース
func (bu *blogUsecase) Update(id int, title, content string) (*model.Blog, error) {
	targetBlog, err := bu.blogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetBlog.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedBlog, err := bu.blogRepo.Update(targetBlog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

// Delete blogを削除するときのユースケース
func (bu *blogUsecase) Delete(id int) error {
	blog, err := bu.blogRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = bu.blogRepo.Delete(blog)
	if err != nil {
		return err
	}

	return nil
}
