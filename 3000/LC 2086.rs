impl Solution {
    pub fn minimum_buckets(hamsters: String) -> i32 {
        let n = hamsters.len();
        let mut ans = 0;
        let hamsters = hamsters.as_bytes();
        let mut i = 0;
        while i < n {
            if hamsters[i] == b'H' {
                if i + 1 < n && hamsters[i+1] == b'.' {
                    ans += 1;
                    i += 2;
                } else if i >= 1 && hamsters[i-1] == b'.' {
                    ans += 1;
                } else {
                    return -1;
                }
            }
            i += 1;
        }

        ans
    }
}