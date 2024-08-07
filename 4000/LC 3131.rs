impl Solution {
    pub fn added_integer(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let min1: i32 = *nums1.iter().min().expect("nums1 should not be empty");
        let min2: i32 = *nums2.iter().min().expect("nums2 should not be empty");

        min2 - min1
    }
}