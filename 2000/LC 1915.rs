impl Solution {
    pub fn wonderful_substrings(word: String) -> i64 {
        let mut ans = 0;
        let mut cnt = vec![0; 1024];
        cnt[0] = 1;
        let mut mask = 0;
        
        let s = word.as_bytes();
        for i in 0..s.len() {
            let x = 1 << (s[i] - b'a');
            mask ^= x;
            ans += cnt[mask as usize] as i64;
            let mut i = 1;
            while i < 1024 {
                ans += cnt[(mask ^ i) as usize] as i64;
                i <<= 1;
            }
            cnt[mask as usize] += 1;
        }

        ans
    }
}