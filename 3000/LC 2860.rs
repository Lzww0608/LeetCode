impl Solution {
    pub fn count_ways(mut nums: Vec<i32>) -> i32 {
        nums.sort();
        let mut ans = 1;
        if nums[0] > 0 {
            ans += 1;
        }

        for i in 1..nums.len() {
            if (i as i32) > nums[i-1] && (i as i32) < nums[i] {
                ans += 1;
            }
        }

        ans
    }
}