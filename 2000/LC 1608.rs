impl Solution {
    pub fn special_array(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        let n = nums.len();
        let mut pre = -1;
        for (i, &v) in nums.iter().enumerate() {
            if v >= (n - i) as i32 && (n - i) as i32 > pre {
                return (n - i) as i32
            }
            pre = v
        }

        -1
    }
}