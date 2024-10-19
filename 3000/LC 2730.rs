impl Solution {
    pub fn longest_semi_repetitive_substring(s: String) -> i32 {
        let mut ans = 1;
        let mut cnt = 0;
        let n = s.len();
        let (mut l, mut r) = (0, 1);
        let s = s.as_bytes();
        while r < n {
            if s[r] == s[r-1] {
                cnt += 1;
                if cnt > 1 {
                    while l < r && s[l] != s[l+1] {
                        l += 1;
                    }
                    cnt -= 1;
                    l += 1;
                }
            }
            //println!("{}, {}",l, r);
            ans = std::cmp::max(ans, r - l + 1);
            r += 1;
        }

        ans as i32
    }
}