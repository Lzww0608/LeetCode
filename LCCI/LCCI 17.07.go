
func find(id map[string]string, name string) string {
	for id[name] != name {
		name = id[name]
	}
	return name
}

func union(id map[string]string, freq map[string]int, a, b string) {
	rootA := find(id, a)
	rootB := find(id, b)
	if rootA != rootB {
		if rootA < rootB {
			id[rootB] = rootA
			freq[rootA] += freq[rootB]
			delete(freq, rootB)
		} else {
			id[rootA] = rootB
			freq[rootB] += freq[rootA]
			delete(freq, rootA)
		}
	}
}

func trulyMostPopular(names []string, synonyms []string) (ans []string) {
	m := map[string]int{}
	id := map[string]string{}

	for _, name := range names {
		s := strings.Split(name, "(")
		x := 0
		for i := 0; i < len(s[1])-1; i++ {
			x = x*10 + int(s[1][i]-'0')
		}
		m[s[0]] = x
		id[s[0]] = s[0]
	}

	for _, synonym := range synonyms {
		s := strings.Split(synonym, ",")
		a := s[0][1:]
		b := s[1][:len(s[1])-1]
		if _, ok := id[a]; !ok {
			id[a] = a
		}
		if _, ok := id[b]; !ok {
			id[b] = b
		}
		union(id, m, a, b)
	}

	for name, count := range m {
		if id[name] == name {
			ans = append(ans, name+"("+strconv.Itoa(count)+")")
		}
	}

	return ans
}

