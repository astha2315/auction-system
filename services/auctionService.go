package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"auction-system/utils"

	"auction-system/models.go"
)

const urlPrepend, auctionNotifURL = "http://127.0.0.1:", "/auction-notification"

var (
	allAuctions        models.AuctionAppState
	allottedAuctionIds map[int]struct{}
)

func init() {
	allAuctions = models.AuctionAppState{} // Initialize Global App State only once
	allottedAuctionIds = map[int]struct{}{}
}

type auctionService struct{}

type AuctionServiceIF interface {
	BidRoundHandler() (int, string)
	ListEndpointHandler() (int, string)
	RegisterAuctionHandler(auction *models.AuctionStruct) (int, string)
}

func AuctionService() AuctionServiceIF {
	return &auctionService{}
}

func (self *auctionService) BidRoundHandler() (int, string) {

	listLen := len(allAuctions.AuctionList)
	var resp interface{}
	// Start an auction Round if list of auctions has entry
	if listLen <= 0 {
		resp = map[string]string{"auction_id": "0", "price": "null", "bidder_id": "null"}
		return 0, utils.MarshalJson(utils.ResponseJson{0, nil, 0, "No active auction for starting the bid"})
	}

	allAuctions.LiveAuction = allAuctions.AuctionList[0]
	allAuctions.AuctionList[0] = allAuctions.AuctionList[len(allAuctions.AuctionList)-1]
	allAuctions.AuctionList = allAuctions.AuctionList[:len(allAuctions.AuctionList)-1]

	// Create a channel to collect bid notification responses.
	bidEntriesChannel, biddersObj := make(chan models.BidResponse, 10), BidderService().GetBiddersList()
	select {
	case bid := <-bidEntriesChannel:
		fmt.Println("received bid", bid)

	default:
		fmt.Println("no bids received")
	}

	// SendAuctionNotification: Takes bidders object and concurrently Notifies all bidders.
	go SendAuctionNotification(biddersObj, bidEntriesChannel)

	// Set 200 millisecond timer.
	timer := time.NewTimer(200 * time.Millisecond)
	<-timer.C
	// Close channel as timer is up.
	close(bidEntriesChannel)

	resp = map[string]string{"auction_id": strconv.Itoa(allAuctions.LiveAuction.Id), "price": "null", "bidder_id": "null"}
	highestBidder := models.BidResponse{}
	// If no bids received return
	if len(bidEntriesChannel) <= 0 {
		return 0, utils.MarshalJson(utils.ResponseJson{0, resp, 0, ""})

	}
	for i := range bidEntriesChannel {
		if i.Price > highestBidder.Price {
			highestBidder.BidderId = i.BidderId
			highestBidder.Price = i.Price
		}
	}

	resp = map[string]string{"auction_id": strconv.Itoa(allAuctions.LiveAuction.Id), "price": fmt.Sprintf("%f", highestBidder.Price),
		"bidder_id": strconv.Itoa(highestBidder.BidderId)}
	return 1, utils.MarshalJson(utils.ResponseJson{http.StatusOK, resp, 0, ""})

}

func SendAuctionNotification(biddersObj models.BidderAppState, bidEntriesChannel chan models.BidResponse) {
	// Create an http client for making requests.
	client := http.Client{Timeout: 200 * time.Millisecond}
	for _, v := range biddersObj.BidderList {
		url := urlPrepend + strconv.Itoa(v.Port) + auctionNotifURL
		go sendRequests(client, url, bidEntriesChannel)
	}
}

func sendRequests(client http.Client, url string, channel chan models.BidResponse) models.BidResponse {
	var bidResp models.BidResponse
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Response error: ", err)
		return bidResp
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&bidResp)
	select {
	case channel <- bidResp:
		fmt.Println("sent message", bidResp)
	default:
		fmt.Println("no message sent")
	}
	return bidResp
}

func (self *auctionService) ListEndpointHandler() (int, string) {

	return 1, utils.MarshalJson(utils.ResponseJson{http.StatusOK, allAuctions, 0, "List of all Auctions "})

}

func (self *auctionService) RegisterAuctionHandler(auction *models.AuctionStruct) (int, string) {

	// Check against allotted id's if any repeat ID is given.
	// Random Id creation would run 100 times, and if still no unique Id, then drop the creation
	i := 0
	for i <= 100 {
		newId := utils.GetRandomInt(1, 1000)
		if _, found := allottedAuctionIds[newId]; !found {
			auction.Id = newId
			break
		} else {
			return 0, utils.MarshalJson(utils.ResponseJson{0, nil, 0, "Could not generate unique Auction"})

		}

	}
	// Start lock for allAuctions i.e. AppState
	allAuctions.Lock()
	defer allAuctions.Unlock()

	allAuctions.AuctionList = append(allAuctions.AuctionList, *auction)
	return 1, utils.MarshalJson(utils.ResponseJson{http.StatusOK, map[string]string{"success": "true", "auction_id": strconv.Itoa(auction.Id)}, 0, "List of all Auctions "})

}
