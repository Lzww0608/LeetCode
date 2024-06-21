func numberOfWeakCharacters(properties [][]int) (ans int) {
    sort.Slice(properties, func(i, j int) bool {
        if properties[i][0] == properties[j][0] {
            return properties[i][1] < properties[j][1]
        }
        return properties[i][0] > properties[j][0]
    })

    maxD := 0
    for _, v := range properties {
        if v[1] < maxD {
            ans++
        } else {
            maxD = v[1]
        }
    }
    
    return 
}
