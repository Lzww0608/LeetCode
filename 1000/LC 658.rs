impl Solution {
    pub fn find_closest_elements(arr: Vec<i32>, k: i32, x: i32) -> Vec<i32> {
        let pos = match arr.binary_search(&x) {
            Ok(i) => i,
            Err(i) => i,
        };

        let mut l = pos as i32 - 1;
        let mut r = pos as i32;
        for _ in 0..k {
            if r >= arr.len() as i32 || l >= 0 && x - arr[l as usize] <= arr[r as usize] - x {
                l -= 1;
            } else {
                r += 1;
            }
        }

        arr[(l + 1) as usize..r as usize].to_vec()
    }
}