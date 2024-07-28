func findMinStep(board string, hand string) int {
    type pair struct {
        c byte
        x int
    }
    ans, used := math.MaxInt32, 0
    rem := map[byte]int{}
    st := []pair{}
    n := len(board)

    var dfs func(int, bool)
    dfs = func(pos int, inserted bool) {
        if used >= ans {
            return
        }

        if pos == n {
            if len(st) == 0 {
                ans = used
            }
            return
        }

        // pos位置小球加入
        if len(st) > 0 && st[len(st)-1].c == board[pos] {
            st[len(st)-1].x++
        } else {
            st = append(st, pair{board[pos], 1})
        }

        //1.不额外插入小球
        //1.1 当前颜色满足三个，且后面没有相同颜色的球
        if (st[len(st)-1].x >= 3 && (pos + 1 == n || board[pos+1] != board[pos])) {
            tmp := st[len(st)-1]
            st = st[:len(st)-1]
            dfs(pos + 1, false)
            st = append(st, tmp)
        }

        //1.2当前小球颜色不满三个
        //1.3当前颜色小球恰好满足三个，但当前小球与上一小球不同色或者之前插入过非同色球，入XX....X
        if (st[len(st)-1].x < 3 || 
            (st[len(st)-1].x == 3 && (board[pos] != board[pos-1] || inserted))) {
            dfs(pos + 1, false);
        }

        //2.插入与当前位置同色小球
        if (rem[board[pos]] > 0 &&
            (pos + 1 == n || board[pos+1] != board[pos])) {
            limit := rem[board[pos]]
            for i := 1; i <= min(2, limit); i++ {
                rem[board[pos]] -= i
                used += i
                st[len(st)-1].x += i
                if (st[len(st)-1].x >= 3) {
                    tmp := st[len(st)-1]
                    st = st[:len(st)-1]
                    dfs(pos + 1, false);
                    st = append(st, tmp)
                } else {
                    dfs(pos + 1, false);
                }

                st[len(st)-1].x -= i
                used -= i
                rem[board[pos]] += i
            }
        }

        //3.后面与相同颜色的球，尝试插入与当前位置颜色不同的小球
        if (pos + 1 < n && board[pos+1] == board[pos]) {
            tmp := st[len(st)-1]
            if tmp.x >= 3 {
                st = st[:len(st)-1]
            }

            for c, x := range rem {
                if c == board[pos] || x == 0 {
                    continue
                }

                for j := 1; j <= min(3, x); j++ {
                    rem[c] -= j
                    used += j

                    if (len(st) > 0 && st[len(st)-1].c == c) {
                        st[len(st)-1].x += j
                    } else {
                        st = append(st, pair{c, j})
                    }

                    if (st[len(st)-1].x >= 3) {
                        tmp2 := st[len(st)-1]
                        st = st[:len(st)-1]
                        dfs(pos + 1, true)
                        st = append(st, tmp2)
                    } else {
                        dfs(pos + 1, true)
                    }

                    if st[len(st)-1].x > j {
                        st[len(st)-1].x -= j
                    } else {
                        st = st[:len(st)-1]
                    }

                    used -= j
                    rem[c] += j
                }
            }
            if tmp.x >= 3 {
                st = append(st, tmp)
            }
        }

        if st[len(st)-1].x == 1 {
            st = st[:len(st)-1]
        } else {
            st[len(st)-1].x--
        }
    }


    for i := range hand {
        rem[hand[i]]++
    }

    dfs(0, false)
    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}