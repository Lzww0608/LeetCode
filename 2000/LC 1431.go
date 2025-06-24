
import "slices"
func kidsWithCandies(candies []int, extraCandies int) []bool {
    mx := slices.Max(candies)
    ans := make([]bool, len(candies))
    for i, x := range candies {
        ans[i] = x + extraCandies >= mx
    }

    return ans
}