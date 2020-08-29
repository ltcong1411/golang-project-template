package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func createAd(c echo.Context) error {
	return c.JSON(http.StatusOK, "createAd")
}

func listAd(c echo.Context) error {
	return c.JSON(http.StatusOK, "listAd")
}

func getAdDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, "getAdDetail")
}

func updateAd(c echo.Context) error {
	return c.JSON(http.StatusOK, "updateAd")
}

func deleteAd(c echo.Context) error {
	return c.JSON(http.StatusOK, "deleteAd")
}
