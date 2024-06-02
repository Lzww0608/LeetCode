func distributeCandies(candies int, num_people int) []int {
    ans := make([]int, num_people)
    sum := 0
    for x, i := 1, 0; sum < candies; x, i = x + 1, (i + 1) % num_people {
        if sum + x <= candies {
            ans[i] += x
        } else {
            ans[i] += candies - sum
        }
        sum += x
    }

    return ans
}
