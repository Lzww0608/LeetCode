const MOD int = 1_000_000_007

func rangeSum(x, y int) int {
    return (x + y) * (x - y + 1) / 2
}

func maxProfit(inventory []int, orders int) (ans int) {
    n := len(inventory)
    sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
    inventory = append(inventory, 0)
    for i := 0; i < n && orders > 0; i++ {
        cnt := i + 1
        if orders >= cnt * (inventory[i] - inventory[i+1]) {
            ans = (ans + cnt * rangeSum(inventory[i], inventory[i+1] + 1)) % MOD
            orders -= cnt * (inventory[i] - inventory[i+1])
        } else {
            q := orders / cnt
            r := orders % cnt
            ans = (ans + cnt * rangeSum(inventory[i], inventory[i] - q + 1)) % MOD
            ans = (ans + r * (inventory[i] - q)) % MOD
            orders = 0
        }
    }

    return
}