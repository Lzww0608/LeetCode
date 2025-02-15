impl Solution {
    pub fn number_of_alternating_groups(colors: Vec<i32>) -> i32 {
        let n = colors.len();
        let mut ans = 0;
        for i in 0..n {
            if colors[i] != colors[(i+1)%n] && colors[(i+1)%n] != colors[(i+2)%n] {
                ans += 1;
            }
        }

        ans
    }
}