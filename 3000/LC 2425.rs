impl Solution {
    pub fn xor_all_nums(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let m = nums1.len();
        let n = nums2.len();
        let mut ans = 0;
        if m & 1 == 1 {
            for &x in &nums2 {
                ans ^= x;
            }
        }
        if n & 1 == 1 {
            for &x in &nums1 {
                ans ^= x;
            }
        }

        ans
    }
}