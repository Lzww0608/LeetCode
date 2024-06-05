func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) (ans []int) {
    m := map[int]int{}
    for i, e := range restaurants {
        m[e[0]] = i
        if e[2] == 0 && veganFriendly == 1 {
            continue
        }
        if e[3] <= maxPrice && e[4] <= maxDistance {
            ans = append(ans, e[0])
        }
    }


    sort.Slice(ans, func(i, j int) bool {
        if restaurants[m[ans[i]]][1] == restaurants[m[ans[j]]][1] {
            return ans[i] > ans[j]
        }
        return restaurants[m[ans[i]]][1] > restaurants[m[ans[j]]][1]
    })

    return 
}
