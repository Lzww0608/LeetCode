func accountBalanceAfterPurchase(purchaseAmount int) (ans int) {
    t := purchaseAmount % 10
    if t < 5 {
        ans = 100 - (purchaseAmount - t)
    } else {
        ans = 100 - (purchaseAmount + 10 - t) 
    }

    return
}