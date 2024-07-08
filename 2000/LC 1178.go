func findNumOfValidWords(words []string, puzzles []string) []int {
    m := map[int]int{}
    for _, word := range words {
        mask := 0
        for _, c := range word {
            x := int(c - 'a')
            mask |= 1 << x
        }
        m[mask]++
    }

    ans := make([]int, len(puzzles))
    for i, puzzle := range puzzles {
        cnt, mask := 0, 0
        for _, c := range puzzle[1:] {
            x := int(c - 'a')
            mask |= 1 << x
        }

        for k := mask; k != 0; k = (k - 1) & mask {
            t := k | (1 << int(puzzle[0] - 'a'))
            cnt += m[t]
        }
        
        cnt += m[1 << int(puzzle[0] - 'a')]
        ans[i] = cnt
    }

    return ans
}