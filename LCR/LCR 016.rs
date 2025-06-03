impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let s = s.as_bytes();
        let mut ans = 0;
        let mut cnt = vec![0; 256];
        let mut l = 0;
        for r in 0..s.len() {
            let x = s[r] as usize;
            cnt[x] += 1;
            while cnt[x] > 1 { 
                let y = s[l] as usize;
                cnt[y] -= 1;
                l += 1;
            }
            ans = ans.max(r - l + 1);
        }

        ans as _
    }
}