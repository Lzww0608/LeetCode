impl Solution {
    pub fn find_permutation(s: String) -> Vec<i32> {
        let n = s.len() + 1;
        let mut ans: Vec<i32> = (1..=n as i32).collect();

        let s = s.as_bytes();
        let mut i = 0;
        while i < n - 1 {
            if s[i] == b'D' {
                let mut j = i;
                while j < n - 1 && s[j] == b'D' {
                    j += 1;
                }

                ans[i..=j].reverse();
                i = j;
            } else {
                i += 1;
            }
        }

        ans
    }
}