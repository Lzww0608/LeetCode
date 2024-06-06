func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    s1 := strings.Split(sentence1, " ")
    s2 := strings.Split(sentence2, " ")
    n, m := len(s1), len(s2)
    l1, l2, r1, r2 := 0, 0, n - 1, m - 1
    for l1 <= r1 && l2 <= r2 {
        if s1[l1] == s2[l2] {
            l1++
            l2++
        } else {
            break
        }
    }

    for l1 <= r1 && l2 <= r2 {
        if s1[r1] == s2[r2] {
            r1--
            r2--
        } else {
            break
        }
    }

    if n > m {
        return l2 > r2
    }

    return l1 > r1
}
