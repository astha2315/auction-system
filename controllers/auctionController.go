package controllers

import (
	"net/http"

	"auction-system/services"

	"auction-system/models.go"

	"github.com/labstack/echo"
)

func BidRoundHandler(c echo.Context) error {

	_, response := services.AuctionService().BidRoundHandler()
	return c.String(http.StatusOK, response)
}

func ListEndpointHandler(c echo.Context) error {

	_, response := services.AuctionService().ListEndpointHandler()
	return c.String(http.StatusOK, response)
}

func RegisterAuctionHandler(c echo.Context) error {

	auction := new(models.AuctionStruct)
	err := c.Bind(auction)
	if err != nil {
		return err
	}

	_, response := services.AuctionService().RegisterAuctionHandler(auction)
	return c.String(http.StatusOK, response)
}
