package controllers

import (
	"auction-system/services"
	"auction-system/utils"
	"net/http"

	"auction-system/models.go"

	"github.com/labstack/echo"
)

func ListBiddersHandler(c echo.Context) error {

	bidderList := services.BidderService().GetBiddersList()
	response := utils.MarshalJson(utils.ResponseJson{1, bidderList, 1, "Bidder List"})
	return c.String(http.StatusOK, response)
}

func CreateAndRegisterBidderHandler(c echo.Context) error {

	bidder := new(models.BidderStruct)
	err := c.Bind(bidder)
	if err != nil {
		return err
	}

	_, response := services.BidderService().CreateAndRegisterBidderHandler(bidder)
	return c.String(http.StatusOK, response)
}

func StudentDetails(c echo.Context) error {

	studentId := c.Param("studentId")

	_, response := services.BidderService().CreateAndRegisterBidderHandler(bidder)
	return c.String(http.StatusOK, response)
}
