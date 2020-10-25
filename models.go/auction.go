package models

import "sync"

type (
	AuctionStruct struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	AuctionAppState struct {
		sync.Mutex                  // <-- mutex protection
		AuctionList []AuctionStruct `json:"auction_list"`
		LiveAuction AuctionStruct   `json:"live_auction"`
	}
)
