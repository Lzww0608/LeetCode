const N int = 26

func countOddLetters(n int) (ans int) {
    m := map[int]string {
        0: "zero",
        1: "one",
        2: "two",
        3: "three",
        4: "four",
        5: "five",
        6: "six",
        7: "seven",
        8: "eight",
        9: "nine",
    }

    cnt := [N]int{}
    for n > 0 {
        x := n % 10
        n /= 10
        for _, c := range m[x] {
            y := int(c - 'a')
            cnt[y]++
        }
    }

    for _, v := range cnt {
        if v & 1 == 1 {
            ans++
        }
    }

    return 
}