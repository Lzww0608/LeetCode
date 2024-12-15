
import "strings"
type FileSystem struct {
    val int
    children map[string]*FileSystem
}

func (c *FileSystem) insert(str []string, val int) bool {
    if len(str) == 0 {
        return true
    } else if len(str) == 1 {
        if _, ok := c.children[str[0]]; ok {
            return false
        }
        c.children[str[0]] = &FileSystem{val: val, children: map[string]*FileSystem{}}
        return true
    }
    v, ok := c.children[str[0]]
    if !ok {
        return false
    }
    return v.insert(str[1:], val)
}

func (c *FileSystem) find(str []string) int {
    if len(str) == 0 {
        return c.val
    } 
    v, ok := c.children[str[0]]
    if !ok {
        return -1
    }
    return v.find(str[1:])
}

func Constructor() FileSystem {
    return FileSystem{0, map[string]*FileSystem{}}
}


func (c *FileSystem) CreatePath(path string, value int) bool {
    s := strings.Split(path, "/")
    s = s[1:]
    return c.insert(s, value)
}


func (c *FileSystem) Get(path string) int {
    s := strings.Split(path, "/")
    s = s[1:]
    return c.find(s)
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */