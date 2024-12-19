impl Solution {
    pub fn stable_mountains(height: Vec<i32>, threshold: i32) -> Vec<i32> {
        let n = height.len();
        let mut ans = Vec::new();
        for i in 1..n {
            if height[i-1] > threshold {
                ans.push(i as i32);
            }
        }

        ans
    }
}