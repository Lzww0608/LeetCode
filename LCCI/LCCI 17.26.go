func computeSimilarities(docs [][]int) []string {
	mp1 := make(map[int][]int)
	for i, doc := range docs {
		for _, word := range doc {
			mp1[word] = append(mp1[word], i)
		}
	}

	mp2 := make(map[int]map[int]int)
	for _, ids := range mp1 {
		for i := 0; i+1 < len(ids); i++ {
			for j := i + 1; j < len(ids); j++ {
				if mp2[ids[i]] == nil {
					mp2[ids[i]] = make(map[int]int)
				}
				mp2[ids[i]][ids[j]]++
			}
		}
	}

	var result []string
	for id1, subMap := range mp2 {
		for id2, count := range subMap {
			similarity := float64(count) / float64(len(docs[id1])+len(docs[id2])-count)
			formattedStr := fmt.Sprintf("%d,%d: %0.4f", id1, id2, similarity+1e-9)
			result = append(result, formattedStr)
		}
	}

	//sort.Strings(result)
	return result
}
