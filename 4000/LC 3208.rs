impl Solution {
    pub fn number_of_alternating_groups(colors: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        let n = colors.len();
        let mut cnt = 1;
        let k = k as usize;

        for i in 1..(n+k-1) {
            if colors[i%n] == colors[(i-1)%n] {
                cnt = 1;
            } else {
                cnt += 1;
                if cnt == k {
                    ans += 1;
                    cnt -= 1;
                }
            }
        }

        ans
    }
}