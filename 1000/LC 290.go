func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	m := make(map[interface{}]int)
	for i, word := range words {
		if _, ok := m[pattern[i]]; !ok {
			m[pattern[i]] = i
		}
		if _, ok := m[word]; !ok {
			m[word] = i
		}
		if m[pattern[i]] != m[word] {
			return false
		}
	}

	return true
}
