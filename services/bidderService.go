package services

import (
	"auction-system/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"auction-system/models.go"
)

var (
	allBidders        models.BidderAppState
	allottedBidderIds map[int]struct{}
)

func init() {
	allBidders = models.BidderAppState{}
	allottedBidderIds = map[int]struct{}{}
}

type bidderService struct{}

type BidderServiceIF interface {
	GetBiddersList() models.BidderAppState
	CreateAndRegisterBidderHandler(bidder *models.BidderStruct) (int, string)
}

func BidderService() BidderServiceIF {
	return &bidderService{}
}

func (self *bidderService) GetBiddersList() models.BidderAppState {
	return allBidders
}

func (self *bidderService) CreateAndRegisterBidderHandler(bidder *models.BidderStruct) (int, string) {

	if len(allBidders.BidderList) >= 1000 {

		return 0, utils.MarshalJson(utils.ResponseJson{0, nil, 0, "BidderList at max capacity"})

	}

	// Check against allotted id's if any repeat ID is given.
	// Random Id creation would run 100 times, and if still no unique Id, then drop the creation
	i := 0
	for i <= 100 {
		newId := utils.GetRandomInt(1, 1000)
		if _, found := allottedBidderIds[newId]; !found {
			bidder.Id = newId
			allottedBidderIds[newId] = struct{}{}
			break
		} else {

			return 0, utils.MarshalJson(utils.ResponseJson{0, nil, 0, "Could not generate unique Bidder"})

		}

	}
	if !utils.IsTCPPortAvailable(bidder.Port) {

		return 0, utils.MarshalJson(utils.ResponseJson{0, nil, 0, "Port in use."})

	}

	go utils.StartBidderServer(*bidder, BidderNotificationHandler)

	allBidders.Lock()
	defer allBidders.Unlock()

	allBidders.BidderList = append(allBidders.BidderList, *bidder)

	return 1, utils.MarshalJson(utils.ResponseJson{http.StatusOK, map[string]string{"success": "true", "bidder_id": strconv.Itoa(bidder.Id)}, 0, "List of all Bids "})

}

func BidderNotificationHandler(t time.Duration, id int) models.RequestHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t * time.Millisecond)

		var bidResp models.BidResponse
		bidResp.BidderId = id
		bidResp.Price = utils.GetRandomFloat()
		fmt.Println("BidCreated")
		fmt.Printf("%+v", bidResp)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bidResp)

	}
}
