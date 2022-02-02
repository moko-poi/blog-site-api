package handler

import (
	"net/http"
	"strconv"
	"time"

	"engineer-jobhunting-api/usecase"

	"github.com/labstack/echo"
)

// BlogHandler blog handlerのinterface
type BlogHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type blogHandler struct {
	blogUsecase usecase.BlogUsecase
}

// NewBlogHandler blog handlerのコンストラクタ
func NewBlogHandler(blogUsecase usecase.BlogUsecase) BlogHandler {
	return &blogHandler{blogUsecase: blogUsecase}
}

type requestBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseBlog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// Post blogを保存するときのハンドラー
func (bh *blogHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestBlog
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdBlog, err := bh.blogUsecase.Create(req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseBlog{
			ID:        createdBlog.ID,
			Title:     createdBlog.Title,
			Content:   createdBlog.Content,
			CreatedAt: createdBlog.CreatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get blogを取得するときのハンドラー
func (bh *blogHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundBlog, err := bh.blogUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseBlog{
			ID:        foundBlog.ID,
			Title:     foundBlog.Title,
			Content:   foundBlog.Content,
			CreatedAt: foundBlog.CreatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put blogを更新するときのハンドラー
func (bh *blogHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestBlog
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedBlog, err := bh.blogUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseBlog{
			ID:        updatedBlog.ID,
			Title:     updatedBlog.Title,
			Content:   updatedBlog.Content,
			CreatedAt: updatedBlog.CreatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete blogを削除するときのハンドラー
func (bh *blogHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = bh.blogUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
