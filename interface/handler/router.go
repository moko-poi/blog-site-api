package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, blogHandler BlogHandler) {

	e.POST("/blog", blogHandler.Post())
	e.GET("/blog/:id", blogHandler.Get())
	e.PUT("/blog/:id", blogHandler.Put())
	e.DELETE("/blog/:id", blogHandler.Delete())

}
