/**
 * Definition for a street.
 * type Street interface {
 *     OpenDoor()
 *     CloseDoor()
 *     IsDoorOpen() bool
 *     MoveRight()
 *     MoveLeft()
 * }
 */
func houseCount(street Street, k int) int {
    for i := 0; i < k; i++ {
        street.CloseDoor()
        street.MoveLeft()
    }

    for i := 0; i < k; i++ {
        if street.IsDoorOpen() {
            return i
        }
        street.OpenDoor()
        street.MoveLeft()
    }

    return k
}