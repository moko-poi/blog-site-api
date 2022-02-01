package model

import (
	"errors"
	"time"
)

// Blog : Blog記事の構造体
type Blog struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

// NewBlog : blog の コンストラクタ
func NewBlog(title string, content string) (*Blog, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	//TODO : CreatedAt を 人間の読みやすい形に変形する
	blog := &Blog{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return blog, nil
}

// Set : blog の セッター
func (b *Blog) Set(title string, content string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	b.Title = title
	b.Content = content
	b.CreatedAt = time.Now()

	return nil
}
