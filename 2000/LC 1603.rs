struct ParkingSystem {
    big: i32,
    medium: i32,
    small: i32,
}

impl ParkingSystem {
    fn new(big: i32, medium: i32, small: i32) -> Self {
        ParkingSystem { big, medium, small }
    }
    
    fn add_car(&mut self, car_type: i32) -> bool {
        if car_type == 1 {
            if self.big > 0 {
                self.big -= 1;
                true
            } else {
                false
            }
        } else if car_type == 2 {
            if self.medium > 0 {
                self.medium -= 1;
                true
            } else {
                false
            }
        } else {
            if self.small > 0 {
                self.small -= 1;
                true
            } else {
                false
            }
        }
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * let obj = ParkingSystem::new(big, medium, small);
 * let ret_1: bool = obj.add_car(carType);
 */
