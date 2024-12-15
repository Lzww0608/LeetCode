impl Solution {
    pub fn min_subsequence(mut nums: Vec<i32>) -> Vec<i32> {
        nums.sort_unstable_by(|a, b| b.cmp(a));
        let mut ans: Vec<i32> = Vec::new();
        let mut sum = 0;
        for &x in &nums {
            sum += x;
        }

        let mut cur = 0;
        for &x in &nums {
            cur += x;
            ans.push(x);
            if cur * 2 > sum {
                break;
            }
        }

        ans
    }
}