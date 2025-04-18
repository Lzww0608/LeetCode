
import "sync"
type FooBar struct {
    id int
	n int
    mtx sync.Mutex
    cv *sync.Cond
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{n: n}
    fb.cv = sync.NewCond(&fb.mtx)
    return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
        fb.mtx.Lock()
        for fb.id & 1 != 0 {
            fb.cv.Wait()
        }
		// printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
        fb.id++
        fb.cv.Signal()
        fb.mtx.Unlock()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
        fb.mtx.Lock()
        for fb.id & 1 != 1 {
            fb.cv.Wait()
        }
		// printBar() outputs "bar". Do not change or remove this line.
        printBar()
        fb.id++
        fb.cv.Signal()
        fb.mtx.Unlock()
	}
}