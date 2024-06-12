
func isPalindrome(word string) bool {
    i, j := 0, len(word)-1
    for i < j {
        if word[i] != word[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func palindromePairs(words []string) [][]int {
    reverseMap := make(map[string]int)
    for i, word := range words {
        reversed := reverseString(word)
        reverseMap[reversed] = i
    }

    var results [][]int
    for i, word := range words {
        for j := 0; j <= len(word); j++ {
            left := word[:j]
            right := word[j:]

            if isPalindrome(left) {
                if idx, exists := reverseMap[right]; exists && idx != i {
                    results = append(results, []int{idx, i})
                }
            }

            if j != len(word) && isPalindrome(right) {
                if idx, exists := reverseMap[left]; exists && idx != i {
                    results = append(results, []int{i, idx})
                }
            }
        }
    }

    return results
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


