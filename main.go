package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORSのミドルウェアを全許可の設定で追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// ルートを設定：すべてのリクエストに対して200を返す
	e.Any("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "HTTP Status Code 200 Returned")
	})

	// サーバーをポート8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
