package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"auction-system/models.go"
	jsoniter "github.com/json-iterator/go"
)

const (
	minTCPPort = 0
	maxTCPPort = 65535
)

var (
	min, max           int
	minFloat, maxFloat float32
)

func init() {
	rand.Seed(time.Now().UnixNano())
	min = 1
	max = 1000
	minFloat = 1.00
	maxFloat = 1000.00
}

type ResponseJson struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Status     int         `json:"status"`
	Message    string      `json:"message"`
}

func MarshalJson(v interface{}) string {
	var json = jsoniter.ConfigFastest
	buf, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Json Marshall Error. : " + err.Error())
		panic(err)
	}

	return string(buf)
}

func GetRandomInt(min int, max int) int {
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

func IsTCPPortAvailable(port int) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

type bidderStateHandler func(time.Duration, int) models.RequestHandlerFunction

func StartBidderServer(bidder models.BidderStruct, handlerFunc bidderStateHandler) error {
	m := http.NewServeMux()
	s := http.Server{Addr: ":" + strconv.Itoa(bidder.Port), Handler: m}

	fmt.Println("Starting server at " + s.Addr)
	m.HandleFunc("/auction-notification", handlerFunc(bidder.Delay, bidder.Id))
	if err := s.ListenAndServe(); err != nil {
		log.Print("ListenAndServe: ", err)
		return err
	}
	return nil
}

func GetRandomFloat() float32 {
	randomNum := minFloat + rand.Float32()*(maxFloat-minFloat)
	return randomNum
}
