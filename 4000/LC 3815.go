type AuctionSystem struct {
    bids map[key]int
    itemHeaps map[int]*hp
}

type key struct {
    userID, itemID int
}
type bidInfo struct {
    userID, bidAmount int
}

func Constructor() AuctionSystem {
    return AuctionSystem {
        bids:      make(map[key]int),
        itemHeaps: make(map[int]*hp),
    }
}

func (c *AuctionSystem) AddBid(userID int, itemID int, bidAmount int) {
    c.bids[key{userID, itemID}] = bidAmount

    if _, ok := c.itemHeaps[itemID]; !ok {
        c.itemHeaps[itemID] = &hp{}
    }
    heap.Push(c.itemHeaps[itemID], bidInfo{userID, bidAmount})
}

func (c *AuctionSystem) UpdateBid(userID int, itemID int, bidAmount int) {
    c.AddBid(userID, itemID, bidAmount)
}

func (c *AuctionSystem) RemoveBid(userID int, itemID int) {
    delete(c.bids, key{userID, itemID})
}

func (c *AuctionSystem) GetHighestBidder(itemID int) int {
    h := c.itemHeaps[itemID]

    for h != nil && h.Len() > 0 {
        top := (*h)[0]

        currentValidBid, exists := c.bids[key{top.userID, itemID}]
        
        if !exists || currentValidBid != top.bidAmount {
            heap.Pop(h)
        } else {
            return top.userID
        }
    }
    return -1
}


type hp []bidInfo

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool {
    if h[i].bidAmount == h[j].bidAmount {
        return h[i].userID > h[j].userID
    }
    return h[i].bidAmount > h[j].bidAmount
}

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x any) {
    *h = append(*h, x.(bidInfo))
}

func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}