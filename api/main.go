package main

import (
	"engineer-jobhunting-api/config"
	"engineer-jobhunting-api/infra"
	"engineer-jobhunting-api/interface/handler"
	"engineer-jobhunting-api/usecase"
	"github.com/labstack/echo"
)

func main() {
	blogRepository := infra.NewBlogRepository(config.NewDB())
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	e := echo.New()
	handler.InitRouting(e, blogHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
