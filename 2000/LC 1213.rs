impl Solution {
    pub fn arrays_intersection(arr1: Vec<i32>, arr2: Vec<i32>, arr3: Vec<i32>) -> Vec<i32> {
        let (mut i, mut j, mut k) = (0, 0, 0);
        let mut ans = Vec::new();

        while i < arr1.len() && j < arr2.len() && k < arr3.len() {
            if arr1[i] < arr2[j] {
                i += 1;
            } else if arr2[j] < arr3[k] {
                j += 1;
            } else {
                if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
                    ans.push(arr1[i]);
                }
                k += 1;
            }
        }

        ans
    }
}