func suggestedProducts(products []string, searchWord string) (ans [][]string) {
    sort.Slice(products, func(i, j int) bool {
        return products[i] < products[j]
    })
    for i := range searchWord {
        s := searchWord[0:i+1]
        pre := []string{}
        for _, t := range products {
            if strings.HasPrefix(t, s) {
                pre = append(pre, t)
            }
            if len(pre) == 3 {
                break
            }
        }
        ans = append(ans, pre)
    }

    return
}

