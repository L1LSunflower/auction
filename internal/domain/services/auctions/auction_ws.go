package auctions

var eventMap = map[int]float64{}

var stack = map[int]map[int]bool{}

func RegisterNew(auctionID, wsID int) {
	if st, ok := stack[auctionID]; !ok {
		stack[wsID] = map[int]bool{
			wsID: true,
		}
	} else {
		st[wsID] = true
	}
}

func LengthStackOfAuction(auctionID int) int {
	if st, ok := stack[auctionID]; ok {
		return len(st)
	}
	return 0
}

func DataSent(auctionID int) bool {
	for _, st := range stack[auctionID] {
		if !st {
			return false
		}
	}
	return true
}

func SetActual(auctionID, wsID int) {
	stack[auctionID][wsID] = true
}

func CheckActual(auctionID, wsID int) bool {
	return stack[auctionID][wsID]
}

func DeleteStack(auctionID int) {
	if _, ok := stack[auctionID]; ok {
		delete(stack, auctionID)
	}
}

func CheckAuction(auctionID int) bool {
	if _, ok := stack[auctionID]; ok {
		return true
	}
	return false
}

func RegisterNewEvent(auctionID int, price float64) {
	if _, ok := eventMap[auctionID]; !ok {
		eventMap[auctionID] = price
	}

	for wsi := range stack[auctionID] {
		stack[auctionID][wsi] = false
	}
}

func AuctionPrice(auctionID int) float64 {
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
