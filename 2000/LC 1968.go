
import "math/rand"
func rearrangeArray(nums []int) []int {
next:
    for {
        rand.Shuffle(len(nums), func(i, j int) {nums[i], nums[j] = nums[j], nums[i]})
        for i := 1; i < len(nums) - 1; i++ {
            if nums[i] * 2 == nums[i-1] + nums[i+1] {
                continue next
            }
        }
        break
    }

    return nums
}