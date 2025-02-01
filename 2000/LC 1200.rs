impl Solution {
    pub fn minimum_abs_difference(mut arr: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans = Vec::new();
        arr.sort_unstable();
        let mut mn = std::i32::MAX;
        for i in 1..arr.len() {
            let dif = arr[i] - arr[i-1];
            if dif < mn {
                mn = dif;
                ans = vec![[arr[i-1], arr[i]].to_vec()];
            } else if dif == mn {
                ans.push([arr[i-1], arr[i]].to_vec());
            }
        }

        ans
    }
}