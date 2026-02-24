func totalDistance(s string) (ans int) {
    m := map[byte]pair {
        'q': pair{0, 0},
        'w': pair{0, 1},
        'e': pair{0, 2},
        'r': pair{0, 3},
        't': pair{0, 4},
        'y': pair{0, 5},
        'u': pair{0, 6},
        'i': pair{0, 7},
        'o': pair{0, 8},
        'p': pair{0, 9},
        'a': pair{1, 0},
        's': pair{1, 1},
        'd': pair{1, 2},
        'f': pair{1, 3},
        'g': pair{1, 4},
        'h': pair{1, 5},
        'j': pair{1, 6},
        'k': pair{1, 7},
        'l': pair{1, 8},
        'z': pair{2, 0},
        'x': pair{2, 1},
        'c': pair{2, 2},
        'v': pair{2, 3},
        'b': pair{2, 4},
        'n': pair{2, 5},
        'm': pair{2, 6},
    }

    x, y := 1, 0
    for i := range s {
        v := m[s[i]]
        ans += abs(v.x - x) + abs(v.y - y)
        x, y = v.x, v.y
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

type pair struct {
    x, y int
}