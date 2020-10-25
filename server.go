package main

import (
	route "auction-system/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	fmt.Println(e)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	route.AuctionRouteService(e)

	e.Logger.Fatal(e.Start(":1323"))
}
