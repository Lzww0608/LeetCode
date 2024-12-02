impl Solution {
    pub fn min_swaps(s: String) -> i32 {
        let mut ans = 0;
        let mut cnt = 0;
        let s = s.as_bytes();
        for &c in s {
            if c == b'[' {
                cnt += 1;
            } else {
                cnt -= 1;
                if cnt < 0 {
                    cnt += 2;
                    ans += 1;
                }
            }
        }

        ans
    }
}