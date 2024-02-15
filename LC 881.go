func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    ans := 0
    l, r := 0, len(people) - 1
    for l <= r {
        if people[r] + people[l] <= limit {
            l, r = l + 1, r - 1
        } else {
            r--
        }
        ans++
    }
    return ans
}