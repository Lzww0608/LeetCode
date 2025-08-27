func findRepeatDocument(documents []int) int {
    for i, x := range documents {
        if x == i {
            continue
        }
        if x == documents[x] {
            return x
        }
        documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
        
    }
    return -1
}
