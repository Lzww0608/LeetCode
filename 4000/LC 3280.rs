impl Solution {
    pub fn convert_date_to_binary(date: String) -> String {
        date.split('-')
            .map(|s| {
                let num: u32 = s.parse().unwrap(); 
                format!("{:b}", num)
            })
            .collect::<Vec<String>>()
            .join("-")
    }
}