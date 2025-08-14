const N int = 605
const D int = 60
func numPairsDivisibleBy60(time []int) (ans int) {
    cnt := [N]int{}
    for _, x := range time {
        ans += cnt[(D - x % D) % D]
        cnt[x % D]++
    }

    return
}