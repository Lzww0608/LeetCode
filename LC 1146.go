type SnapshotArray struct {
    m []map[int]int
}

var cnt int = 0

func Constructor(length int) SnapshotArray {
    m := make([]map[int]int, length)
    cnt = 0
    return SnapshotArray{m}
}


func (this *SnapshotArray) Set(index int, val int)  {
    if len(this.m[index]) == 0 {
        this.m[index] = make(map[int]int)
    }
    this.m[index][cnt] = val
}


func (this *SnapshotArray) Snap() int {
    cnt++
    return cnt - 1
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    snapIds := this.m[index] 
    left, right := 0, len(snapIds)-1
    result := -1
    for left <= right {
        mid := left + (right - left) / 2
        if snapIds[mid] <= snap_id {
            result = snapIds[mid]
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return result
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */