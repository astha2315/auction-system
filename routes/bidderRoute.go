package route

import (
	"auction-system/controllers"

	"github.com/labstack/echo"
)

func BidderRouteService(e *echo.Echo) {

	e.GET("/list-bidders", controllers.ListBiddersHandler)
	e.POST("/create-bidder", controllers.CreateAndRegisterBidderHandler)

}
