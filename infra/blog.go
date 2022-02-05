package infra

import (
	"engineer-jobhunting-api/domain/model"
	"engineer-jobhunting-api/domain/repository"
	"github.com/jinzhu/gorm"
)

// BlogRepository : blog repository の 構造体
type BlogRepository struct {
	Conn *gorm.DB
}

// NewBlogRepository : blog repositoryのコンストラクタ
func NewBlogRepository(conn *gorm.DB) repository.BlogRepository {
	return &BlogRepository{Conn: conn}
}

// Create blogの保存
func (br *BlogRepository) Create(blog *model.Blog) (*model.Blog, error) {
	if err := br.Conn.Create(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

// FindByID : blogをIDで取得
func (br *BlogRepository) FindByID(id int) (*model.Blog, error) {
	blog := &model.Blog{ID: id}

	if err := br.Conn.First(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

// Update : blogの更新
func (br *BlogRepository) Update(blog *model.Blog) (*model.Blog, error) {
	if err := br.Conn.Model(&blog).Updates(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

// Delete : blogの削除
func (br *BlogRepository) Delete(blog *model.Blog) error {
	if err := br.Conn.Delete(&blog).Error; err != nil {
		return err
	}

	return nil
}
