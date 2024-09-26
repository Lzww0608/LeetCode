impl Solution {
    pub fn add_rungs(rungs: Vec<i32>, dist: i32) -> i32 {
        let mut ans = (rungs[0] - 1) / dist;
        for i in 1..rungs.len() {
            ans += (rungs[i] - rungs[i-1] - 1) / dist;
        }

        ans
    }
}