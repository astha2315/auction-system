package route

import (
	"auction-system/controllers"

	"github.com/labstack/echo"
)

func AuctionRouteService(e *echo.Echo) {

	e.GET("/auction/bidRound/start", controllers.BidRoundHandler)
	e.GET("/list-auctions", controllers.ListEndpointHandler)
	e.POST("/register-auction", controllers.RegisterAuctionHandler)

}
