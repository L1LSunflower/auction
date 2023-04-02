package auctions

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
)

var eventMap = map[int]map[string]any{}

var stack = map[int]map[*websocket.Conn]bool{}

func RegisterNew(auctionID int, ws *websocket.Conn) {
	if st, ok := stack[auctionID]; !ok {
		stack[auctionID] = map[*websocket.Conn]bool{
			ws: true,
		}
	} else {
		st[ws] = true
	}
	fmt.Printf("\nSTACK: %#v\n\n", stack)
}

func LengthStackOfAuction(auctionID int) int {
	if st, ok := stack[auctionID]; ok {
		return len(st) + 1
	}
	return 1
}

func DataSent(auctionID int) bool {
	for _, st := range stack[auctionID] {
		if !st {
			return false
		}
	}
	return true
}

func SetActual(auctionID int, wsID *websocket.Conn) {
	stack[auctionID][wsID] = true
}

func CheckActual(auctionID int, wsID *websocket.Conn) bool {
	return stack[auctionID][wsID]
}

func DeleteStack(auctionID int) {
	if _, ok := stack[auctionID]; ok {
		delete(stack, auctionID)
	}
}

func DeleteConsumer(auctionID int, ws *websocket.Conn) {
	if len(stack[auctionID]) == 1 {
		delete(stack, auctionID)
	} else {
		delete(stack[auctionID], ws)
	}
}

func CheckAuction(auctionID int) bool {
	if _, ok := stack[auctionID]; ok {
		return true
	}
	return false
}

func RegisterNewEvent(auctionID int, userID string, price float64) {
	if _, ok := eventMap[auctionID]; !ok {
		eventMap[auctionID] = map[string]any{
			"user_id": userID,
			"price":   price,
		}
	}

	for wsi := range stack[auctionID] {
		stack[auctionID][wsi] = false
	}
}

func AuctionPrice(auctionID int) map[string]any {
	return eventMap[auctionID]
}

func CheckEvent(auctionID int) bool {
	if _, ok := eventMap[auctionID]; ok {
		return true
	}
	return false
}

func DeleteEvent(auctionID int) {
	if ok := CheckEvent(auctionID); ok {
		delete(eventMap, auctionID)
	}
}
