func findRestaurant(list1 []string, list2 []string) (ans []string) {
    m := map[string]int{}
    mn := math.MaxInt32
    for i, s := range list1 {
        m[s] = i
    }
    for i, s := range list2 {
        if i > mn {
            break
        }
        if j, ok := m[s]; ok {
            if i + j == mn {
                ans = append(ans, s)
            } else if i + j < mn {
                mn = i + j
                ans = []string{s}
            }
        }
    }

    return 
}
