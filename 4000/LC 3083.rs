impl Solution {
    pub fn is_substring_present(s: String) -> bool {
        let s = s.as_bytes();
        let n = s.len();
        let mut cnt = vec![0; 26];

        for i in 0..n-1 {
            let x = (s[i] - b'a') as usize;
            let y = (s[i+1] - b'a') as usize;
            cnt[x] |= 1 << (y as i32);
            if (cnt[y] >> (x as i32)) & 1 == 1 {
                return true;
            }
        }

        false
    }
}