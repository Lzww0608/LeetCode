use std::collections::HashMap;

impl Solution {
    pub fn length_of_longest_substring_k_distinct(s: String, k: i32) -> i32 {
        let n = s.len();
        let mut ans = 0;
        let s = s.as_bytes();
        let mut m: HashMap<u8, i32> = HashMap::new();

        let (mut l, mut r) = (0, 0);
        while r < n {
            *m.entry(s[r]).or_insert(0) += 1;

            while m.len() > k as usize {
                if let Some(cnt) = m.get_mut(&s[l]) {
                    *cnt -= 1;
                    if *cnt == 0 {
                        m.remove(&s[l]);
                    }
                }

                l += 1;
            }

            ans = ans.max(r - l + 1);
            r += 1;
        }

        ans as i32
    }
}