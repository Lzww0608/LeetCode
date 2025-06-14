
import "container/heap"
type VideoSharingPlatform struct {
    name2Id map[string]int
    id2Name map[int]string 
    likes map[int]int
    dislikes map[int]int
    views map[int]int
    h *hp
}

const N int = 100_001

func Constructor() VideoSharingPlatform {
    h := &hp{}
    for i := range N {
        heap.Push(h, i)
    }
    return VideoSharingPlatform{map[string]int{}, map[int]string{}, map[int]int{}, map[int]int{}, map[int]int{}, h}
}


func (c *VideoSharingPlatform) Upload(video string) int {
    id := heap.Pop(c.h).(int)
    c.name2Id[video] = id
    c.id2Name[id] = video
    return id
}


func (c *VideoSharingPlatform) Remove(videoId int)  {
    if name, ok := c.id2Name[videoId]; ok {
        delete(c.name2Id, name)
        delete(c.id2Name, videoId)
        delete(c.dislikes, videoId)
        delete(c.likes, videoId)
        delete(c.views, videoId)
        heap.Push(c.h, videoId)
    }
}


func (c *VideoSharingPlatform) Watch(videoId int, startMinute int, endMinute int) string {
    if _, ok := c.id2Name[videoId]; !ok {
        return "-1"
    }
    c.views[videoId]++
    name := c.id2Name[videoId]

    return name[startMinute: min(endMinute + 1, len(name))]
}


func (c *VideoSharingPlatform) Like(videoId int)  {
    if _, ok := c.id2Name[videoId]; ok {
        c.likes[videoId]++
    }
}


func (c *VideoSharingPlatform) Dislike(videoId int)  {
    if _, ok := c.id2Name[videoId]; ok {
        c.dislikes[videoId]++
    }
}


func (c *VideoSharingPlatform) GetLikesAndDislikes(videoId int) []int {
    if _, ok := c.id2Name[videoId]; !ok {
        return []int{-1}
    }

    return []int{c.likes[videoId], c.dislikes[videoId]}
}


func (c *VideoSharingPlatform) GetViews(videoId int) int {
    if _, ok := c.id2Name[videoId]; !ok {
        return -1
    }

    return c.views[videoId]
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    n := len(*h)
    old := *h
    x = old[n-1]
    *h = old[:n-1]
    return 
}


/**
 * Your VideoSharingPlatform object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Upload(video);
 * obj.Remove(videoId);
 * param_3 := obj.Watch(videoId,startMinute,endMinute);
 * obj.Like(videoId);
 * obj.Dislike(videoId);
 * param_6 := obj.GetLikesAndDislikes(videoId);
 * param_7 := obj.GetViews(videoId);
 */