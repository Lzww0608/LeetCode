impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let n = nums.len() as i32;
        let k = (k % n) as usize;
        nums.reverse();
        nums[..k].reverse();
        nums[k..].reverse();
    }
}