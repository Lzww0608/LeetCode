
import "strings"
type FileSystem struct {
    content string
    dir map[string]*FileSystem
}

func (c *FileSystem) f(path string) *FileSystem {
    s := strings.Split(path, "/")
	p := c
	for _, t := range s {
		if t == "" {
			continue
		}
		if p.dir == nil {
			p.dir = make(map[string]*FileSystem)
		}
		if _, ok := p.dir[t]; !ok {
			p.dir[t] = &FileSystem{}
		}
		p = p.dir[t]
	}
	return p
}

func Constructor() FileSystem {
    return FileSystem{}
}


func (c *FileSystem) Ls(path string) (ans []string) {
    p := c.f(path)
    if len(p.content) > 0 {
        return []string{path[strings.LastIndex(path, "/") + 1:]}
    }
    for k := range p.dir {
        ans = append(ans, k)
    }
    sort.Slice(ans, func(i, j int) bool {
        return ans[i] < ans[j]
    })
    return 
}


func (c *FileSystem) Mkdir(path string)  {
    c.f(path)
}


func (c *FileSystem) AddContentToFile(filePath string, content string)  {
    p := c.f(filePath)
    p.content += content
}


func (c *FileSystem) ReadContentFromFile(filePath string) string {
    p := c.f(filePath)
    return p.content
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */