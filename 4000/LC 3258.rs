impl Solution {
    pub fn count_k_constraint_substrings(s: String, k: i32) -> i32 {
        let s = s.as_bytes();
        let mut ans = 0;
        let n = s.len();

        for i in 0..n {
            let (mut a, mut b) = (0, 0);
            for j in i..n {
                if s[j] == b'0' {
                    a += 1;
                } else {
                    b += 1;
                }

                if a <= k || b <= k {
                    ans += 1;
                } else {
                    break;
                }
            }
        }

        ans
    }
}