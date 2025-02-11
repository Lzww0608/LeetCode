impl Solution {
    pub fn maximum_units(mut box_types: Vec<Vec<i32>>, mut truck_size: i32) -> i32 {
        let mut ans = 0;
        box_types.sort_by(|a, b| b[1].cmp(&a[1]));

        for v in box_types {
            let t = truck_size.min(v[0]);
            ans += t * v[1];
            truck_size -= t;
        }

        ans
    }
}