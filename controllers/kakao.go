package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"kakao/models"
)
func KakaoKeyboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.KakaoKeyboardModel())
	}
}
func KakaoMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(models.KakaoRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.KakaoMessageModel(*req))
	}
}